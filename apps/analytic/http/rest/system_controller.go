package rest

import (
	"analytic/http/rest/dto"
	"encoding/json"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/valyala/fasthttp"
	"time"
)

type SystemController struct {
	config *config.AppConfig
	log    logger.Logger
}

func NewSystemController(config *config.AppConfig, log logger.Logger) *SystemController {
	return &SystemController{
		config: config,
		log:    log,
	}
}

// GetVersion godoc
//
//	@Summary		Get application version
//	@Description	Returns the current version of the application and the current timestamp
//	@Tags			System
//
//	@Produce		json
//
//	@Success		200	{object}	dto.VersionResponseDTO
//	@Failure		500
//
//	@Router			/system/version [get]
func (c *SystemController) GetVersion(ctx *fasthttp.RequestCtx) {
	response := dto.VersionResponseDTO{
		AppVersion:       c.config.AppVersion,
		CurrentTimestamp: time.Now().Unix(),
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx.Response.BodyWriter()).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		c.log.Error(err.Error())
	}
}
