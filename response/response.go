package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 返回值结构体
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r *Response) SetCode(code int) {
	r.Status = code
}

func (r *Response) SetMessage(m string) {
	r.Message = m
}
func (r *Response) SetData(data interface{}) {
	r.Data = data
}

// 请求失败的返回信息
func FailResponse(c *gin.Context, code int, err string) {
	c.JSON(http.StatusOK, Response{
		Status:  code,
		Message: err,
		Data:    nil,
	})
}

// 参数绑定失败返回
func InvalidParamsResponse(c *gin.Context, err string) {
	if err == "" {
		err = "参数错误"
	}
	c.JSON(http.StatusOK, Response{
		Status:  InvalidParamsCode,
		Data:    nil,
		Message: err,
	})
}

func SuccessResponse(c *gin.Context, resp interface{}){
	c.JSON(http.StatusOK, resp)
}

// new response
func NewResponse() Response {
	return Response{
		Status:  SuccessCode,
		Message: Success,
	}
}
