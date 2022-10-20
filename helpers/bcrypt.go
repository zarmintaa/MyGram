package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePass(h, p string) bool {
	pass, hash := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return nil == err
}
