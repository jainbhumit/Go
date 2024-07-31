package main

import "golang.org/x/crypto/bcrypt"

func main() {
	pass := `password@123`

	ePass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

}
