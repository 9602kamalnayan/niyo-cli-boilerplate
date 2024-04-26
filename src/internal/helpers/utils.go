package helpers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func getValueFromContext(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}
	value, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}

func GetDefaultValueFromContext(ctx context.Context, contextKey string) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(string(contextKey))
	} else {
		return getValueFromContext(ctx, string(contextKey))
	}
}

func GetSpanIdFromContext(ctx context.Context) string {
	var spanCtx = trace.SpanContextFromContext(ctx)
	if ginCtx, ok := ctx.(*gin.Context); ok {
		spanCtx = trace.SpanContextFromContext(ginCtx.Request.Context())
	}
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}
	return ""
}

func GetTraceIdFromContext(ctx context.Context) string {
	var spanCtx = trace.SpanContextFromContext(ctx)
	if ginCtx, ok := ctx.(*gin.Context); ok {
		spanCtx = trace.SpanContextFromContext(ginCtx.Request.Context())
	}
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}
