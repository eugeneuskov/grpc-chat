package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	uuid "github.com/satori/go.uuid"
	"time"
)

type authService struct {
	repository repositories.Auth
	config     *config.Auth
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

func (a *authService) Login(login, password string) (string, error) {
	user, err := a.repository.Login(login, generatePasswordHash(password, a.config.HashSalt))
	if err != nil {
		return "", err
	}

	return a.generateAuthToken(user)
}

func (a *authService) CheckToken(token string) (*structs.User, error) {
	claims, err := a.parseToken(token)
	if err != nil {
		return nil, err
	}

	id, err := uuid.FromString(claims.UserId)
	if err != nil {
		return nil, errors.New("wrong token")
	}

	return &structs.User{
		ID:       id,
		Username: claims.Username,
	}, nil
}

func (a *authService) generateAuthToken(user *structs.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(a.config.Tll) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:   user.ID.String(),
		Username: user.Username,
	})

	return token.SignedString([]byte(a.config.SingingKey))
}

func (a *authService) parseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(a.config.SingingKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims, nil
}

func newAuthService(repository repositories.Auth, config *config.Auth) *authService {
	return &authService{repository, config}
}
