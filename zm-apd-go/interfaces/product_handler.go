package interfaces

import (
	"net/http"
	"strconv"
	"zm-apd-go/application/param"
	"zm-apd-go/application/service"

	"github.com/gin-gonic/gin"
)

//ProductHandler struct defines the dependencies that will be used
type ProductHandler struct {
	app service.ProductApp
}

//NewProducts return new Product
func NewProducts(app service.ProductApp) *ProductHandler {
	return &ProductHandler{app}
}

//Create product
func (p *ProductHandler) Create(c *gin.Context) {
	var createProductCmd param.CreateProductCmd
	if error := c.ShouldBindJSON(&createProductCmd); error != nil {
		ErrorResponse(c, "invalid json")
		return
	}
	//validate the request param
	if validateErr := createProductCmd.Validate(); validateErr != nil {
		ErrorResponse(c, validateErr.Error())
		return
	}
	result := p.app.CreateProduct(&createProductCmd)
	if result {
		SuccessResponse(c, nil, "success")
	} else {
		ErrorResponse(c, "error")
	}
}
func (p *ProductHandler) Update(c *gin.Context) {
	var updateProductCmd param.UpdateProductCmd
	if error := c.ShouldBind(&updateProductCmd); error != nil {
		ErrorResponse(c, "invalid json")
		return
	}
	if validateErr := updateProductCmd.Validate(); validateErr != nil {
		ErrorResponse(c, validateErr.Error())
		return
	}
	result := p.app.UpdateProduct(&updateProductCmd)
	if result {
		SuccessResponse(c, nil, "success")
	} else {
		ErrorResponse(c, "error")
	}
}
func (p *ProductHandler) Query(c *gin.Context) {
	id, error := strconv.ParseUint(c.Query("id"),10,64)
	if error != nil {

	}
	product := p.app.Query(id)
	SuccessResponse(c,product,"success")

}
func SuccessResponse(c *gin.Context, obj interface{}, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": message,
		"data":    obj,
	})
}
func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": message,
	})
}
