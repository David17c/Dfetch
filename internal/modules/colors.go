package modules

func Color(NoColor bool) string {
	if NoColor {
		return ""
	}
	return "\033[30m███\033[31m███\033[32m███\033[33m███\033[34m███\033[35m███\033[36m███\033[37m███\n\033[90m███\033[91m███\033[92m███\033[93m███\033[94m███\033[95m███\033[96m███\033[97m███"
}
