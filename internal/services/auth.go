package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/kr1ny77/BasketForm-AI/internal/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailTaken         = errors.New("email already taken")
	ErrNicknameTaken      = errors.New("nickname already taken")
	ErrInvalidToken       = errors.New("invalid token")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthService struct {
	storage   *Storage
	jwtSecret []byte
}

func NewAuthService(storage *Storage) *AuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "basketform-ai-default-secret"
	}
	return &AuthService{
		storage:   storage,
		jwtSecret: []byte(secret),
	}
}

func (a *AuthService) Register(email, nickname, password string) (*models.User, error) {
	if _, ok := a.storage.GetUserByEmail(email); ok {
		return nil, ErrEmailTaken
	}
	if _, ok := a.storage.GetUserByNickname(nickname); ok {
		return nil, ErrNicknameTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           genID(),
		Email:        email,
		Nickname:     nickname,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if err := a.storage.SaveUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthService) Login(email, password string) (string, *models.User, error) {
	user, ok := a.storage.GetUserByEmail(email)
	if !ok {
		return "", nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := a.generateToken(user.ID)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (a *AuthService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return a.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return "", ErrInvalidToken
	}
	return userID, nil
}

func (a *AuthService) ChangePassword(userID, oldPassword, newPassword string) error {
	user, ok := a.storage.GetUserByID(userID)
	if !ok {
		return ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return ErrInvalidCredentials
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)
	return a.storage.SaveUser(user)
}

func (a *AuthService) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.jwtSecret)
}
