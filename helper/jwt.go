package helper

import (
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
}
