package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HandleError(message string, err interface{}) {
	log.Println("========== Start Error Message ==========")
	log.Println("Message => " + message + ".")
	if err != nil {
		log.Println("Error => ", err)
	}
	log.Println("========== End Of Error Message ==========")
	log.Println()
}

type JSONResponse struct {
	Status        int         `json:"status"`
	Message 	  string      `json:"message"`
	Data          interface{} `json:"data"`
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}