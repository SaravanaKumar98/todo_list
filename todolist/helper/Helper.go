package helper

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UuidGenerate() (string, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	newUUID_str := strings.Trim(string(newUUID), "\n")

	return newUUID_str, nil
}

func MakeHass(pass string) string {
	hass, _ := bcrypt.GenerateFromPassword([]byte(pass), 15)
	return string(hass)
}

func CompareHash(hass string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hass), []byte(password))
	return err == nil
}

func SetAuthCookie(c *gin.Context, uuid string) error {
	str, err := Crypt(uuid, "encrypt")
	if err != nil {
		return err
	}
	c.SetCookie("auth", string(str), 3600, "/", "localhost", false, true)
	return nil
}
