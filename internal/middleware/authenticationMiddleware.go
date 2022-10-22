package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/pandukamuditha/simple-blog/internal/common"
)

type TokenClaimsKey struct{}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		inputToken := strings.Split(r.Header.Get("Authorization"), " ")[1]

		if inputToken == "" {
			common.WriteJsonResponse(
				rw,
				http.StatusUnauthorized,
				common.ErrorResponse{Err: "invalid token"},
			)
			return
		}

		processedToken, err := jwt.ParseWithClaims(inputToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return common.JwtSigningSecret, nil
		})

		if err != nil {
			common.WriteJsonResponse(
				rw,
				http.StatusUnauthorized,
				common.ErrorResponse{Err: "token error"},
			)
			return
		}

		var processedClaims jwt.StandardClaims

		if claims, ok := processedToken.Claims.(*jwt.StandardClaims); ok && processedToken.Valid {
			processedClaims = *claims
		} else {
			common.WriteJsonResponse(
				rw,
				http.StatusUnauthorized,
				common.ErrorResponse{Err: "token error"},
			)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), TokenClaimsKey{}, processedClaims))
		next.ServeHTTP(rw, r)
	})
}
