package libs

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"real-holat/config"
	"real-holat/pkg/constants"

	jwtgo "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// GetUserIDFromToken ...
func GetUserIDFromToken(accessToken string) (uuid.UUID, error) {
	cfg := config.LoadConfig(".")

	claims, err := ExtractClaims(accessToken, []byte(cfg.JwtSecretKey))
	if err != nil {
		claims, err = ExtractClaims(accessToken, []byte(""))
		if err != nil {
			log.Println("could not extract claims:", err)
			return uuid.Nil, err
		}
	}

	if _, ok := claims["user_id"]; !ok {
		return uuid.Nil, fmt.Errorf("user_id is not set: %w", "invalid token")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("could not extract claims: %w", "invalid token")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("'user_id' could not be parsed as uuid: %w", "invalid token")
	}
	return userID, nil
}

func ExtractClaims(tokenString string, signingKey []byte) (jwtgo.MapClaims, error) {
	claims := jwtgo.MapClaims{}
	if tokenString == "" {
		claims["role"] = constants.UnAuthorized
		return claims, nil
	}
	if strings.Contains(tokenString, "Basic") {
		claims["role"] = constants.UnAuthorized
		return claims, nil
	}

	// Remove Bearer prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("invalid jwt token"))
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return claims, nil
}
