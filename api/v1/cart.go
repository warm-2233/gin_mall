package v1

import (
	"gin_mall/response"
	"gin_mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListCart(c *gin.Context) {
	uid := c.GetUint("ID")
	response.SuccessResponse(c, service.CartService.GetCart(c.Request.Context(), uid))
}
func AddCart(c *gin.Context) {
	uid := c.GetUint("ID")
	pid, _ := strconv.Atoi(c.Query("product_id"))
	if pid == 0 {
		response.FailResponse(c, response.InvalidParamsCode, "product_id is invalid")
		return
	}
	r := service.CartService.AddCart(c.Request.Context(), uid, uint(pid))
	response.SuccessResponse(c, r)
}
func DeleteCart(c *gin.Context) {
	uid := c.GetUint("ID")
	pid, _ := strconv.Atoi(c.Query("product_id"))
	if pid == 0 {
		response.FailResponse(c, response.InvalidParamsCode, "product_id is invalid")
		return
	}

	resp := service.CartService.DeleteCart(c.Request.Context(), uid, uint(pid)) //could be more than one product id in the request.  (if it is
	// c.JSON(http.StatusOK, resp)                                     //a request for multiple products)  if it is a request for a single product id.
	response.SuccessResponse(c, resp)

}
func UpdateCart(c *gin.Context) {
	// service.CartService.UpdateCart(c.Request.Context(), c.Query("product_id"), c.Query("quantity"))
}
