package util

import (
	"fmt"
	"net"

	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/valyala/fasthttp"
)

const (
	XForwardedForHeader = "X-Forwarded-For"
)

func GetSenderInfo(ctx *fasthttp.RequestCtx) model.SenderInfo {
	senderIP := net.ParseIP(string(ctx.Request.Header.Peek(XForwardedForHeader)))
	if senderIP == nil {
		senderIP = ctx.RemoteIP()
	}
	senderDevice := string(ctx.UserAgent())
	return model.SenderInfo{
		IP:     senderIP,
		Device: senderDevice,
	}
}

func GetRequiredQueryArgs(ctx *fasthttp.RequestCtx, keys ...string) (map[string]string, error) {
	requiredArgs := make(map[string]string, len(keys))
	missingKeys := make([]string, 0)
	for _, key := range keys {
		value := ctx.QueryArgs().Peek(key)
		if len(value) == 0 {
			missingKeys = append(missingKeys, key)
			continue
		}
		requiredArgs[key] = string(value)
	}
	if len(missingKeys) != 0 {
		return nil, fmt.Errorf("missing required keys: %v", missingKeys)
	}
	return requiredArgs, nil
}

func GetRequiredFormFields(ctx *fasthttp.RequestCtx, fields ...string) (map[string]string, error) {
	requiredFields := make(map[string]string, len(fields))
	missingFields := make([]string, 0)
	for _, field := range fields {
		value := ctx.FormValue(field)
		if len(value) == 0 {
			missingFields = append(missingFields, field)
			continue
		}
		requiredFields[field] = string(value)
	}
	if len(missingFields) != 0 {
		return nil, fmt.Errorf("missing required form fields: %v", missingFields)
	}
	return requiredFields, nil
}
