package token

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "errors"
	"fmt"
	// "log"
	// "net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(user *models.User) (string, error) {
	var secretJWT = []byte(os.Getenv("SECRET_KEY_JWT"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = map[string]string{
		"id":    fmt.Sprintf("%d", user.Id),
		"uuid":  user.Uuid,
		"name":  user.Name,
		"email": user.Email,
	}
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretJWT)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY_JWT")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenMetadata(tokenString string) (*models.AccessDetails, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY_JWT")), nil
	})

	fmt.Println("token")
	fmt.Println(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims")
		fmt.Println(claims)
		fmt.Println("claims user")
		fmt.Println(claims["user"])
		user := claims["user"]
		md, _ := user.(map[string]interface{})
		email := md["email"].(string)
		uuid := md["uuid"].(string)
		fmt.Println(email)
		fmt.Println(uuid)
		return &models.AccessDetails{
			AccessUuid: uuid,
			Email:      email,
		}, nil
	} else {
		fmt.Println("err")
		fmt.Println(err)
		return nil, err
	}
}
