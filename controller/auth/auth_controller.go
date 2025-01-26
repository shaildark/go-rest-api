package auth

import (
	"fmt"

	"example.com/go-api/db"
	"example.com/go-api/models"
	"example.com/go-api/request"
	"example.com/go-api/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		msg := "Failed to connect to database"
		context.JSON(500, response.APIResponse(500, "Something went wrong", msg, nil, nil))
		return
	}

	fmt.Println(loginReq.Email)
	fmt.Println(loginReq.Password)

	var user models.User
	fmt.Println(user.Email)
	result := database.Where("email = ?", loginReq.Email).First(&user)

	if result.Error != nil {
		context.JSON(400, response.APIResponse(400, "", result.Error, nil, nil))
		return
	}

	if user.Email == "" {
		context.JSON(400, response.APIResponse(400, "", "Email already exists", nil, nil))
		return
	}

	fmt.Println(user.Email)
	if result.Error != gorm.ErrRecordNotFound {
		context.JSON(400, response.APIResponse(400, "", "Email already exists", nil, nil))
		return
	}

	data := response.APIResponse(200, "success", nil, nil, nil)
	context.JSON(200, data)
	return
}

func LoginUser(context *gin.Context) {
	data := response.APIResponse(200, "success", nil, nil, nil)
	context.JSON(200, data)
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
