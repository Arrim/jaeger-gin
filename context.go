package jaegergin

import (
	"context"
	"github.com/gin-gonic/gin"
)

const ContextTracerKey = "Tracer-context"

func InjectSpanInGinContext(ctx context.Context, gCtx *gin.Context) {
	gCtx.Set(ContextTracerKey, ctx)
}

func GetSpanFromContext(ctx context.Context) context.Context {
	val := ctx.Value(ContextTracerKey)
	if sp, ok := val.(context.Context); ok {
		return sp
	}

	return context.Background()
}
