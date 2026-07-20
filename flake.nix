{
  description = "Lightweight system information tool focused on clean output";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      lib = nixpkgs.lib;
      systems = [
        "x86_64-linux"
        "aarch64-linux"
      ];

      forAllSystems = lib.genAttrs systems;

      packageFor = system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        pkgs.buildGoModule {
          pname = "dfetch";
          version = "1.0.0";
          src = self;

          vendorHash = "sha256-WIp9R1wJAi1YIQtm3YfBY83xK66FiSsFXRpnudZh1Xg=";

          subPackages = [ "." ];

          meta = with lib; {
            description = "Lightweight system information tool focused on clean output";
            homepage = "https://github.com/David17c/Dfetch";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };
    in
    {
      packages = forAllSystems (system: {
        default = packageFor system;
      });

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/dfetch";
          meta.description = "Run Dfetch";
        };
      });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              git
            ];
          };
        });

      nixosModules.default = { config, pkgs, lib, ... }:
        let
          cfg = config.programs.dfetch;
        in
        {
          options.programs.dfetch = {
            enable = lib.mkEnableOption "Dfetch";

            package = lib.mkOption {
              type = lib.types.package;
              default = self.packages.${pkgs.stdenv.hostPlatform.system}.default;
              defaultText = lib.literalExpression "inputs.dfetch.packages.${pkgs.stdenv.hostPlatform.system}.default";
              description = "The Dfetch package to install.";
            };
          };

          config = lib.mkIf cfg.enable {
            environment.systemPackages = [ cfg.package ];
          };
        };
    };
}
