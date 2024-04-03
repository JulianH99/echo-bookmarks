package help

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/JulianH99/gomarks/storage/models"
)

func GenerateSessionToken(user models.User) string {

	// define in an env var later
	secret := "thisisaverybigst"

	return encrypt(user.Nick, secret)
}

func encrypt(str, secret string) string {
	block, err := aes.NewCipher([]byte(secret))

	if err != nil {
		return ""
	}

	// this is not necerarily secure: using the same secret as iv
	// but it will make things easier by now
	cfb := cipher.NewCFBEncrypter(block, []byte(secret))
	cipherText := make([]byte, len(str))

	cfb.XORKeyStream(cipherText, []byte(str))

	return base64Encode(string(cipherText))

}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
