package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/alexspx/gocms/src/initializers"
	"github.com/alexspx/gocms/src/models"
	utils "github.com/alexspx/gocms/src/utilities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.UserSignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Username:  strings.ToLower(payload.Username),
		Email:     payload.Email,
		Password:  hashedPassword,
		Approved:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res := ac.DB.Create(&newUser)

	if res.Error != nil && strings.Contains(res.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if res.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something went wrong"})
		return
	}

	userResponse := models.UserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Name:      newUser.Name,
		Email:     newUser.Username,
		Approved:  newUser.Approved,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created", "user": userResponse})
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var potentialUser models.User
	res := ac.DB.First(&potentialUser, "username = ?", strings.ToLower(payload.Username))
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid username"}) // a bit of a security improvement would be to return the same message on both invalid password and username
		return
	}

	if err := utils.VerifyPassword(potentialUser.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Wrong password"})
		return
	}

	config, _ := initializers.LoadConfig(".") // TO-DO: Pass it as a parameter

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, potentialUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, potentialUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true) // secure for prod??
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}
