package v1

import (
	"gin_mall/dto"
	"gin_mall/response"
	"gin_mall/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Tags		商品相关路由
//	@Produce	json
//	@Summary	创建商品
//	@Param		Authorization	header		string			true	"登录凭证"
//	@Param		file			formData	file			true	"文件"
//	@Param		ProductDto		formData	dto.ProductDto	true	"参数"
//	@Success	200				{object}	response.Response{data=vo.ProductVo}
//	@Router		/api/v1/product [post]
func ProductCreate(c *gin.Context) {
	form, _ := c.MultipartForm()
	fs, ok := form.File["file"]
	if !ok || len(fs) == 0 {
		response.FailResponse(c, response.InvalidParamsCode, "file form 不能为空")
		return
	}

	uid := c.GetUint("ID")

	request := dto.ProductDto{}
	if err := c.ShouldBind(&request); err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}

	resp := service.ProductService.Create(c.Request.Context(), request, uid, fs)
	c.JSON(http.StatusOK, resp)
}

// 获取产品列表
// type: 1==new  2==

//	@Tags		商品相关路由
//	@Produce	json
//	@Summary	获取产品列表
//	@Param		ProductListDto	query		dto.ProductListDto	flase	"参数"
//	@Success	200				{object}	response.Response{data=vo.ProductListVo}
//	@Router		/api/v1/product [get]
func ListProduct(c *gin.Context) {
	var request = dto.ProductListDto{}

	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}
	if request.Size == 0 {
		request.Size = 10
	}

	resp := service.ProductService.List(c.Request.Context(), request)
	c.JSON(http.StatusOK, resp)
}

// SearchProduct
func SearchProduct(c *gin.Context) {
	GetProduct(c)
}

//	@Tags		商品相关路由
//	@Produce	json
//	@Summary	获取产品信息
//	@Param		id	path		int	true	"产品ID"
//	@Success	200	{object}	response.Response{data=vo.ProductVo}
//	@Router		/api/v1/product/{id} [get]
func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}

	resp := service.GetProductById(c.Request.Context(), uint(id))
	c.JSON(http.StatusOK, resp)
}
