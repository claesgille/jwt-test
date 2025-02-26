package user

var users = map[string]string{}

func GetRole(username string) string {
	if username == "senior@mail.com" {
		return "senior"
	}
	return "employee"
}

func ValidateUser(username, password string) bool {
	return username == "employee@mail.com" && password == "password" || username == "senior@mail.com" && password == "password"
}

func LogOut(username string, tokenString string) {
	users[username] = tokenString
}

func VerifyUserToken(username string, tokenString string) bool {
	return false
}
