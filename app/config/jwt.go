package config

/**
 *@Author: ANYARONKE Dare Samuel
 */

import "github.com/golang-jwt/jwt/v4"

var JWT_SECRET = []byte("aqwzsxedcrfvtgbyhnujujikolpmamzlekjhgfdswqazx")

//var COOKIE_TOKEN = "token"

type TWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
