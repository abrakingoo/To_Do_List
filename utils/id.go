package utils

import(
	"crypto/rand"
	"log"
	"encoding/hex"

)

func GetId() string {
	id := make([]byte, 8)

	_, err := rand.Read(id)

	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(id)
}