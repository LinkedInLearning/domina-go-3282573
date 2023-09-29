package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	llaveSecreta := []byte("llave_secreta")

	nuevoToken := jwt.New(jwt.SigningMethodHS256)

	claims := nuevoToken.Claims.(jwt.MapClaims)
	claims["usuario"] = "mi_usuario"

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, error := nuevoToken.SignedString(llaveSecreta)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(token)
}
