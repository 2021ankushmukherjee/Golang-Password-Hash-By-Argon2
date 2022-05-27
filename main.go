package main

import (
	"fmt"
	"log"

	"github.com/itrepablik/pwd"
)

func main() {

	// Method 1: To initialize the pwd with custom configs, use the SetArgon2Configs() method
	var pwdConfig = &pwd.Argon2Configs{
		Memory:      128 * 1024,
		Iterations:  1,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}
	pwdConfig.SetArgon2Configs(pwdConfig)

	// Generate Argon2id secured hash password
	argon2Hash, err := pwdConfig.HashAndSalt("ankush")
	if err != nil {
		errInfo := fmt.Sprintf("error hashing password: %s", err.Error())
		log.Fatal(errInfo)
		return
	}
	fmt.Println("argon2Hash: ", argon2Hash)

	// Validate Argon2id secured hash password
	isPwdCorrect, err := pwdConfig.CheckPasswordHash("ankush", argon2Hash)
	if err != nil {
		errInfo := fmt.Sprintf("error checking password: %s", err.Error())
		log.Fatal(errInfo)
		return
	}
	fmt.Println("Password is correct:", isPwdCorrect)

	// Method 2: To initialize the pwd with default configs, use the NewArgon2id() method
	var pwd = pwd.NewArgon2id()

	// Generate Argon2id secured hash password
	argon2Hash1, err := pwd.HashAndSalt("ankush")
	if err != nil {
		errInfo := fmt.Sprintf("error hashing password: %s", err.Error())
		log.Fatal(errInfo)
		return
	}
	fmt.Println("argon2Hash1: ", argon2Hash1)

	// Validate Argon2id secured hash password
	isPwdCorrect1, err := pwd.CheckPasswordHash("ankush", argon2Hash1)
	if err != nil {
		errInfo := fmt.Sprintf("error checking password: %s", err.Error())
		log.Fatal(errInfo)
		return
	}
	fmt.Println("Password is correct:", isPwdCorrect1)
}
