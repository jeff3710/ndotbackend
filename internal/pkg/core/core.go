package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeff3710/ndot/internal/pkg/errno"
)

type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`
}
type JsonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 err 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, &JsonResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
