package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmunongo/linkkeep/api/src/handlers/auth"
)

func GetUserIDFromToken(t string) (uint, error) {
	//t is Bearer + token, so remove the Bearer
	t = t[len("Bearer "):]
	
	token, _, err := new(jwt.Parser).ParseUnverified(t, &auth.JwtCustomClaims{})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*auth.JwtCustomClaims)

	if !ok {
		return 0, err
	}

	return uint(claims.ID), nil
}