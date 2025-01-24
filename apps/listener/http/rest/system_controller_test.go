package rest

import (
	"encoding/json"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"listener/http/rest/dto"
	"testing"
	"time"
)

func TestGetVersion_Success(t *testing.T) {
	// Setup
	appVersion := "1.0.0"
	mockConfig := &config.AppConfig{AppVersion: appVersion}
	mockLogger := new(logger.MockLogger)
	controller := NewSystemController(mockConfig, mockLogger)

	// Mock request and response
	ctx := &fasthttp.RequestCtx{}
	ctx.Response.ResetBody()

	// Call the method
	controller.GetVersion(ctx)

	// Assert response status code is 200
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())

	// Assert response content type is JSON
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	// Decode response body
	var response dto.VersionResponseDTO
	err := json.Unmarshal(ctx.Response.Body(), &response)
	assert.NoError(t, err)

	// Assert app version
	assert.Equal(t, appVersion, response.AppVersion)

	// Assert current timestamp is a valid Unix timestamp (within a reasonable range)
	currentTime := time.Now().Unix()
	assert.LessOrEqual(t, response.CurrentTimestamp, currentTime)
	assert.GreaterOrEqual(t, response.CurrentTimestamp, currentTime-10) // Allow a small delay
}
