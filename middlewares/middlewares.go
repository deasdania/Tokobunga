package middlewares

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	"github.com/gin-gonic/gin"

	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

// func GenerateJWT(user *models.User) (string, error) {
// 	var secretJWT = []byte(os.Getenv(utilities.KEY_JWT))
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["user"] = map[string]string{
// 		"id":    fmt.Sprintf("%d", user.Id),
// 		"uuid":  user.Uuid,
// 		"name":  user.Name,
// 		"email": user.Email,
// 	}
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	tokenString, err := token.SignedString(secretJWT)

// 	if err != nil {
// 		fmt.Errorf("Something went wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func GenerateToken(duration int64) string {
// 	mySigningKey := []byte("SignKey")

// 	claims := &jwt.StandardClaims{
// 		ExpiresAt: time.Now().Unix() + duration,
// 	}

// 	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenKey, err := tokenClaims.SignedString(mySigningKey)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return tokenKey
// }

// func VerifyToken(tokenString string) (bool, error) {
// 	mySigningKey := []byte(os.Getenv(utilities.KEY_JWT))

// 	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return mySigningKey, nil
// 	})

// 	if err != nil {
// 		return false, err
// 	}

// 	if token.Valid {
// 		return token.Valid, err
// 	} else {
// 		log.Println(`error => `, err)
// 	}
// 	return token.Valid, err
// }

// func TokenValidation() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if len(c.Request.Header["Authorization"]) == 0 {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "you are not unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		} else {
// 			token := c.Request.Header["Authorization"][0]
// 			validity, _ := VerifyToken(token)
// 			if validity {
// 				c.Next()
// 			} else {
// 				c.JSON(http.StatusUnauthorized, gin.H{
// 					"error": "you are not unauthorized",
// 				})
// 				c.Abort()
// 				return
// 			}
// 		}
// 	}
// }
