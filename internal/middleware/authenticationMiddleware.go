package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	"github.com/pandukamuditha/simple-blog/internal/common"
	"github.com/pandukamuditha/simple-blog/internal/repositories"
)

type TokenClaimsKey struct{}
type UserId struct{}

func AuthenticationMiddleware(logger *common.Logger, db *pgx.Conn) func(next http.Handler) http.Handler {
	// get access to query interface
	var queries = repositories.New(db)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			// except .. endpoints
			openRoutes := []string{"auth/login", "auth/signup", "health"}

			for _, openRoute := range openRoutes {
				if strings.Contains(r.URL.Path, openRoute) {
					next.ServeHTTP(rw, r)
					return
				}
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				common.WriteJsonResponse(
					rw,
					http.StatusUnauthorized,
					common.ErrorResponse{Err: "no auth token"},
				)
				return
			}

			inputToken := strings.Split(authHeader, " ")[1]

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
					common.ErrorResponse{Err: "invalid token"},
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
					common.ErrorResponse{Err: "invalid token"},
				)
				return
			}

			// get user id
			authUserId, err := queries.GetUserIdByUsername(r.Context(), processedClaims.Subject)

			if err != nil {
				common.WriteJsonResponse(
					rw,
					http.StatusUnauthorized,
					common.ErrorResponse{Err: "user not found"},
				)
				return
			}

			r = r.WithContext(
				context.WithValue(
					context.WithValue(
						r.Context(),
						TokenClaimsKey{},
						processedClaims,
					),
					UserId{},
					authUserId,
				),
			)

			next.ServeHTTP(rw, r)
		})
	}
}
