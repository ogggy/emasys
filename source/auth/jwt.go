package auth

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
)

type jwtProvider struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

type UserClaim struct {
	*jwt.StandardClaims
	User struct {
		UID string `json:"uid"`
	} `json:"user"`
}

type Claim struct {
	*jwt.StandardClaims
	Data map[string]interface{}
}

func (me *jwtProvider) Init() error {
	
	

	return nil
}
