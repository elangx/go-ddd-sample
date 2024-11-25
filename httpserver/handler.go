package httpserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Output struct {
	Errno  int64       `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func HandlerFunc[RequestT any, ResponseT any](fn func(context.Context, *RequestT) (*ResponseT, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		out := &Output{
			Errno:  0,
			Errmsg: "",
		}
		ctx := c.Request.Context()
		requestPtr := reflect.New(reflect.TypeOf((*RequestT)(nil)).Elem()).Interface()
		if err := c.ShouldBind(requestPtr); err != nil {
			c.PureJSON(http.StatusBadRequest, out)
			return
		}
		responsePtr, err := fn(ctx, requestPtr.(*RequestT))
		if err != nil {
			c.PureJSON(http.StatusInternalServerError, out)
			return
		}
		c.PureJSON(http.StatusOK, responsePtr)
	}
}
