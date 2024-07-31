package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `password@123`

	ePass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ePass))

	err = bcrypt.CompareHashAndPassword(ePass, []byte(`password@1234`))
	if err != nil {
		fmt.Println("Passwords do not match")
	} else {
		fmt.Println("Password match")
	}

}
