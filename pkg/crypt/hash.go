package crypt

import "golang.org/x/crypto/bcrypt"

func Hash(pwd string) (hashed string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	if err != nil {
		return
	}

	return string(hash), nil
}

func CompareHashAndPwd(hash string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
