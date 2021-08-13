package jaegergin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func StartServerSpanMiddleware(c *gin.Context) {
	spanName := fmt.Sprintf("%s %s", c.Request.Method, c.FullPath())

	sCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span, ctx := opentracing.StartSpanFromContext(c, spanName, ext.RPCServerOption(sCtx))
	defer span.Finish()

	InjectSpanInGinContext(ctx, c)
	c.Next()

	ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
	ext.HTTPMethod.Set(span, c.Request.Method)
	ext.HTTPUrl.Set(span, c.Request.URL.EscapedPath())
}
