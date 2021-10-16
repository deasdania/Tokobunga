package utilities

import "os"

const (
	// endpoints
	CREATE_ACCOUNT        = "/create/account"
	CREATE_ACCOUNT_PUBLIC = "/createaccount"
	TEST                  = "/test"
	GET_ALL_ACCOUNT       = "/users"
	GET_ACCOUNT           = "/user"
	// LOGIN           = "/login"
	CHANGE_PASSWORD = "/change/account/password"

	GENERATE_UUID = "/generate/uuid"

	LOGIN      = "/login"
	CHECK_AUTH = "/check/authorize"
	LOGOUT     = "/logout"
	REFRESH    = "/refresh"

	GET_ROLE    = "/role"
	CREATE_ROLE = "/create/role"

	CREATE_PRODUCT = "/create/product"
	UPDATE_PRODUCT = "/update/product"

	CREATE_PRODUCT_CATEGORY = "/create/productcategory"
	UPDATE_PRODUCT_CATEGORY = "/update/productcategory"

	MEMBER = "member"
	ADMIN  = "admin"
)

var (
	KEY_JWT      = os.Getenv("SECRET_KEY_JWT")
	ACCOUNT_PORT = os.Getenv("ACCOUNT_PORT")
	REDIS_URL    = os.Getenv("REDIS_URL")
)
