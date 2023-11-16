package utils

import (
	"github.com/device-server/domain/constants"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

// GenToken 生成JWT
func MakeTokenString(key []byte, SigningAlgorithm string, username string, password string) string {
	if SigningAlgorithm == "" {
		SigningAlgorithm = "HS256"
	}

	token := jwt.New(jwt.GetSigningMethod(SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)
	claims[constants.KeyId] = username
	claims[constants.KeyExp] = time.Now().Add(time.Hour).Unix()
	claims[constants.KeyIdx] = time.Now().Unix()
	claims[constants.KeyPassword] = password
	var tokenString string
	if SigningAlgorithm == "RS256" {
		keyData, _ := os.ReadFile("testdata/jwtRS256.key")
		signKey, _ := jwt.ParseRSAPrivateKeyFromPEM(keyData)
		tokenString, _ = token.SignedString(signKey)
	} else {
		var err error
		tokenString, err = token.SignedString(key)
		if err != nil {
			logrus.Errorf("make token failed:%s", err.Error())
		}
	}

	return tokenString
}
