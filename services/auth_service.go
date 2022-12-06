package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/oxyrinchus/goilerplate/common"
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/repositories"
)

type AccessTokenClaims struct {
	UserUUID string `json:"uid"`
	Role     string `json:"role"`
	// Role     []string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	// UserUUID string `json:"uid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	logger         lib.Logger
	userService    UserService
	authRepository repositories.AuthRepository
	env            lib.Env
}

func NewAuthService(logger lib.Logger, userService UserService, authRepository repositories.AuthRepository, env lib.Env) AuthService {
	return AuthService{
		logger:         logger,
		userService:    userService,
		authRepository: authRepository,
		env:            env,
	}
}

func (as AuthService) SignUp(email, password, name, role string) (result bool, err error) {
	foundUser, err := as.userService.FindUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		as.logger.Error(err)
		return false, err
	}

	if foundUser.ID != "" {
		as.logger.Info("duplicated email")
		return false, nil
	}

	newUser := models.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
		Role:     role,
	}

	if err := as.userService.CreateUser(newUser); err != nil {
		as.logger.Error(err)
		return false, err
	}

	return true, nil
}

func (as AuthService) SignIn(email, password string) (accessToken string, refreshToken string, err error) {
	user, err := as.userService.FindUserByEmail(email)
	if err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	if user.ID == "" {
		as.logger.Info("Not found user")
		return "", "", nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	accessToken, err = as.generateAccessToken(user)
	if err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	newSessionID := uuid.New().String()
	refreshToken, err = as.generateRefreshToken(newSessionID)
	if err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	if err = as.registerTokenSession(newSessionID, user.ID); err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (as AuthService) Authorize(accessToken, refreshToken string) (userID string, newAccessToken string, err error) {
	if accessToken != "" {
		accessTokenClaims, err := as.verifyAccessToken(accessToken)
		if err == nil {
			return accessTokenClaims.UserUUID, "", nil
		}
	}

	if refreshToken == "" {
		return "", "", errors.New("refreshToken is empty")
	}

	refreshTokenClaims, err := as.verifyRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	userID, err = as.findTokenSession(refreshTokenClaims.RegisteredClaims.ID)
	if err != nil {
		return "", "", err
	}

	user, err := as.userService.FindUserByID(userID)
	if err != nil {
		return "", "", err
	}

	accessToken, err = as.generateAccessToken(user)
	if err != nil {
		as.logger.Error(err)
		return "", "", err
	}

	return userID, accessToken, nil
}

func (as AuthService) generateAccessToken(user models.User) (token string, err error) {
	claims := AccessTokenClaims{
		UserUUID: user.ID,
		Role:     user.Role,
		// Email:    user.Email,
		// Name:     user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * common.ACCESS_TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "goilerplate",
		},
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(as.env.JWTSecret)); err != nil {
		return "", err
	}

	return token, nil
}

func (as AuthService) generateRefreshToken(tokenID string) (token string, err error) {
	claims := RefreshTokenClaims{
		//UserUUID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * common.REFRESH_TOKEN_TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "goilerplate",
			ID:        tokenID,
		},
	}

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(as.env.JWTSecret)); err != nil {
		return "", err
	}

	return token, nil
}

func (as AuthService) verifyAccessToken(tokenString string) (*AccessTokenClaims, error) {
	claims := AccessTokenClaims{}

	if err := common.VerifyToken(tokenString, &claims, as.env.JWTSecret); err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return &claims, nil
}

func (as AuthService) verifyRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	claims := RefreshTokenClaims{}

	if err := common.VerifyToken(tokenString, &claims, as.env.JWTSecret); err != nil {
		as.logger.Error(err)
		return nil, err
	}

	return &claims, nil
}

func (as AuthService) registerTokenSession(tokenID, userID string) error {
	return as.authRepository.Set(tokenID, userID, time.Second*common.REFRESH_TOKEN_TTL)
}

func (as AuthService) findTokenSession(tokenID string) (userID string, err error) {
	return as.authRepository.Get(tokenID)
}
