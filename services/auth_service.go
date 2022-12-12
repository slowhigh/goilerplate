package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/repositories"
)

type AccessTokenClaims struct {
	UserUUID string `json:"uid"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

type AuthService struct {
	logger         lib.Logger
	userService    UserService
	authRepository repositories.AuthRepository
	env            lib.Env
}

// NewAuthService initialize auth service
func NewAuthService(logger lib.Logger, userService UserService, authRepository repositories.AuthRepository, env lib.Env) AuthService {
	return AuthService{
		logger:         logger,
		userService:    userService,
		authRepository: authRepository,
		env:            env,
	}
}

// SignUp signs up user.
func (as AuthService) SignUp(email, password, name, role string) error {
	newUser := models.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
		Role:     role,
	}

	err := as.userService.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

// SignIn signs in user.
func (as AuthService) SignIn(email, password string, accessToken *string, refreshToken *string) (bool, error) {
	user, err := as.userService.GetUserInfoByEmail(email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		as.logger.Info("Not same password")
		return false, nil
	} else if err != nil {
		as.logger.Error(err)
		return false, err
	}

	*accessToken, err = as.generateAccessToken(user)
	if err != nil {
		return false, err
	}

	newSessionID := uuid.New().String()
	*refreshToken, err = as.generateRefreshToken(newSessionID)
	if err != nil {
		return false, err
	}

	err = as.registerTokenSession(newSessionID, user.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Authorize authorizes the generated token
func (as AuthService) Authorize(accessToken, refreshToken string, ) (userID string, newAccessToken string, err error) {
	if accessToken != "" {
		accessTokenClaims := as.verifyAccessToken(accessToken)
		if accessTokenClaims != nil {
			return accessTokenClaims.UserUUID, "", nil
		}
	}

	if refreshToken == "" {
		return "", "", nil
	}

	refreshTokenClaims := as.verifyRefreshToken(refreshToken)
	if refreshTokenClaims == nil {
		return "", "", nil
	}

	userID, err = as.findTokenSession(refreshTokenClaims.RegisteredClaims.ID)
	if err != nil {
		return "", "", err
	}
	if userID == "" {
		return "", "", nil
	}

	user, err := as.userService.GetUserInfoByID(userID)
	if err != nil {
		return "", "", err
	}

	accessToken, err = as.generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	return userID, accessToken, nil
}

// generateAccessToken generates the access token with given user.
func (as AuthService) generateAccessToken(user models.User) (string, error) {
	claims := AccessTokenClaims{
		UserUUID: user.ID,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * common.ACCESS_TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "goilerplate",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(as.env.JWTSecret))
	if err != nil {
		as.logger.Error(err)
		return "", err
	}

	as.logger.Debugf("[generateAccessToken] {token:%s}", token)
	return token, nil
}

// generateRefreshToken generates the refresh token with given token's ID.
func (as AuthService) generateRefreshToken(tokenID string) (string, error) {
	claims := RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * common.REFRESH_TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "goilerplate",
			ID:        tokenID,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(as.env.JWTSecret))
	if err != nil {
		as.logger.Error(err)
		return "", err
	}

	as.logger.Debugf("[generateRefreshToken] {token:%s}", token)
	return token, nil
}

// verifyAccessToken verifies the signatured access-token and return parsed access-claims.
func (as AuthService) verifyAccessToken(tokenString string) *AccessTokenClaims {
	claims := AccessTokenClaims{}

	if !common.VerifyToken(tokenString, &claims, as.env.JWTSecret) {
		return nil
	}

	return &claims
}

// verifyRefreshToken verifies the signatured refresh-token and return parsed refresh-claims.
func (as AuthService) verifyRefreshToken(tokenString string) *RefreshTokenClaims {
	claims := RefreshTokenClaims{}

	if !common.VerifyToken(tokenString, &claims, as.env.JWTSecret) {
		return nil
	}

	return &claims
}

// registerTokenSession registers tokenID and userID in the session.
func (as AuthService) registerTokenSession(tokenID, userID string) error {
	return as.authRepository.Set(tokenID, userID, time.Second*common.REFRESH_TOKEN_TTL)
}

// findTokenSession finds the userID matching the given tokenID in the session.
func (as AuthService) findTokenSession(tokenID string) (userID string, err error) {
	return as.authRepository.Get(tokenID)
}
