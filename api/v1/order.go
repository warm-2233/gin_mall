package v1

import (
	"gin_mall/dto"
	"gin_mall/response"
	"gin_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Tags		订单操作
//	@Produce	json
//	@Summary	获取订单列表
//	@Param		Authorization	header	string	true	"登录凭证"
//	@Router		/api/v1/orders  [get]
func ListOrders(c *gin.Context) {
	uid := c.GetUint("ID")
	c.JSON(http.StatusOK, service.OrderService.List(c.Request.Context(), uid))
}

//	@Tags		订单操作
//	@Produce	json
//	@Summary	创建订单
//	@Param		Authorization	header	string				true	"登录凭证"
//	@Param		OrderCreateDto	body	dto.OrderCreateDto	true	"创建订单参数"
//	@Router		/api/v1/orders  [post]
func OrderCreate(c *gin.Context) {
	request := dto.OrderCreateDto{}
	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}

	uid := c.GetUint("ID")
	resp := service.OrderService.Create(c.Request.Context(), uid, request)

	c.JSON(http.StatusOK, resp)
}

//	@Tags		订单操作
//	@Produce	json
//	@Summary	支付订单
//	@Param		Authorization	header	string			true	"登录凭证"
//	@Param		OrderPayDto		body	dto.OrderPayDto	true	"支付订单参数"
//	@Router		/api/v1/pay   [post]
func OrderPay(c *gin.Context) {
	request := dto.OrderPayDto{}
	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}
	uid := c.GetUint("ID")

	resp := service.OrderService.Pay(c.Request.Context(), uid, request.OrderId)
	c.JSON(http.StatusOK, resp)
}

func OrderUpdate(c *gin.Context) {

}
func OrderDelete(c *gin.Context) {

}
func GetOrder(c *gin.Context) {

}
