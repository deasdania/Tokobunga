package authjwt

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type JWTService interface {
	CreateToken(email string) (*models.TokenDetails, error)
	CreateAuth(email string, td *models.TokenDetails) error
	DeleteTokens(authD *models.AccessDetails) error
	FetchAuth(authD *models.AccessDetails) (string, error)
	DeleteAuth(givenUuid string) (int64, error)
}
