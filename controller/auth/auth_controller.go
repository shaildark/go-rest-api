package auth

import (
	"example.com/go-api/db"
	"example.com/go-api/models"
	"example.com/go-api/request"
	"example.com/go-api/response"
	"example.com/go-api/utility/jwt"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var loginReq request.LoginRequest

	if err := context.ShouldBindJSON(&loginReq); err != nil {
		msg := request.ValidationErrorResponse(err)
		context.JSON(400, response.APIResponse(400, "Something went wrong", msg, nil, nil))
		return
	}

	database, err := db.Connection()

	if err != nil {
		context.JSON(500, response.APIResponse(500, "Something went wrong", "Failed to connect to database", nil, nil))
		return
	}

	var existingUser models.User
	database.Where("email = ?", loginReq.Email).First(&existingUser)
	if existingUser.Email != "" {
		context.JSON(400, response.APIResponse(400, "", "Email Already Exist", nil, nil))
		return
	}

	newUser := models.User{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}
	newUser.SetPassword()

	userResult := newUser.Create()
	if userResult != nil {
		context.JSON(400, response.APIResponse(400, "Something went wrong.", userResult, nil, nil))
		return
	}

	context.JSON(200, response.APIResponse(200, "User created successfully", nil, nil, nil))
	return
}

func LoginUser(context *gin.Context) {
	var loginReq request.LoginRequest

	if err := context.ShouldBindJSON(&loginReq); err != nil {
		msg := request.ValidationErrorResponse(err)
		context.JSON(400, response.APIResponse(400, "Something went wrong", msg, nil, nil))
		return
	}

	database, err := db.Connection()

	if err != nil {
		context.JSON(500, response.APIResponse(500, "Something went wrong", "Failed to connect to database", nil, nil))
		return
	}

	var user models.User
	database.Where("email = ?", loginReq.Email).First(&user)

	if user.Email == "" {
		context.JSON(400, response.APIResponse(400, "", "Email not found", nil, nil))
		return
	}

	passwordCheck := user.CheckPassword(loginReq.Password)

	if !passwordCheck {
		context.JSON(400, response.APIResponse(400, "", "Invalid Password", nil, nil))
		return
	}

	token, err := jwt.GenerateToken(user.Email, int64(user.ID))

	if err != nil {
		context.JSON(500, response.APIResponse(500, "Something went wrong", "Failed to generate token", nil, nil))
		return
	}

	context.JSON(200, response.APIResponse(200, "Login successful", map[string]string{"token": token}, nil, nil))
	return
}

func ForgetPassword(context *gin.Context) {
	data := response.APIResponse(200, "success", nil, nil, nil)
	context.JSON(200, data)
	return
}

func SetNewPassword(context *gin.Context) {
	data := response.APIResponse(200, "success", nil, nil, nil)
	context.JSON(200, data)
	return
}
