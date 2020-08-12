package services

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const passwordSalt = "sjflksjsaljflsafjsalfjçhçljgçhlkgfçbgkldsfuoi3498543rhyfewkjhr3k4h53kjb534b"

// HashPassword generates an hashed and salted string
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+passwordSalt), 10)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// CheckPasswordHash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passwordSalt))
	return err == nil
}

// GenerateRandomToken
func GenerateRandomToken() string {
	return tokenGenerator(64)
}

func tokenGenerator(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
