package encrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Helper interface {
	HashAndSalt(pwd []byte) string
	ValidateHash(secret, hash string) bool
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
func ValidateHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}

// func (db *userConnection) InsertUser(user entity.User) entity.User {
// 	user.Password = hashAndSalt([]byte(user.Password))
// 	db.connection.Save(&user)
// 	return user
// }
