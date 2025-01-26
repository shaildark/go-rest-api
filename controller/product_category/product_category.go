package productcategory

import (
	"example.com/go-api/response"
	"github.com/gin-gonic/gin"
)

func ListAllProductCategory(context *gin.Context) {

	data := response.APIResponse(200, "success", nil, nil, nil)
	context.JSON(200, data)
	return
}
