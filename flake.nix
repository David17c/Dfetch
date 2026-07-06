{
  description = "A lightweight system information tool focused on clean output";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    {
      packages.x86_64-linux = {
        default = nixpkgs.legacyPackages.x86_64-linux.buildGoModule {
          pname = "dfetch";
          version = "1.0.0";
          src = self;

          vendorHash = null;

          subPackages = [ "." ];

          meta = with nixpkgs.lib; {
            description = "A lightweight system information tool focused on clean output";
            homepage = "https://github.com/crispdark/Dfetch";
            license = licenses.mit;
            maintainers = [];
            platforms = platforms.linux;
          };
        };
      };

      packages.aarch64-linux = {
        default = nixpkgs.legacyPackages.aarch64-linux.buildGoModule {
          pname = "dfetch";
          version = "1.0.0";
          src = self;

          vendorHash = null;

          subPackages = [ "." ];

          meta = with nixpkgs.lib; {
            description = "A lightweight system information tool focused on clean output";
            homepage = "https://github.com/crispdark/Dfetch";
            license = licenses.mit;
            maintainers = [];
            platforms = platforms.linux;
          };
        };
      };

      devShells.x86_64-linux.default = nixpkgs.legacyPackages.x86_64-linux.mkShell {
        buildInputs = with nixpkgs.legacyPackages.x86_64-linux; [
          go
          git
        ];
      };

      devShells.aarch64-linux.default = nixpkgs.legacyPackages.aarch64-linux.mkShell {
        buildInputs = with nixpkgs.legacyPackages.aarch64-linux; [
          go
          git
        ];
      };
    };
}
