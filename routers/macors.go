package routers

func UserNameMacro(username string) func(string) bool {
	return func(s string) bool {
		return len(s) > 0 && len(s) < 12
	}
}
