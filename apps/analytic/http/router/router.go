package router

import (
	"analytic/docs"
	"analytic/http/rest"
	fastHttpRouter "github.com/fasthttp/router"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/http/httppath"
	"github.com/khrulsergey/chain-processor/pkg/http/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"os"
	"os/exec"
)

// AppRouterHandler is responsible for routing components
type AppRouterHandler struct {
	versionController *rest.SystemController
	logger            logger.Logger
	appConfig         config.AppConfig
}

//	@title			Swagger Entitlements-engine API
//	@version		1.0
//	@description	This is a microservice serving permission mission
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// NewRouterHandler constructs a router
func NewRouterHandler(logger logger.Logger, appConfig config.AppConfig,
	versionController *rest.SystemController) *AppRouterHandler {
	return &AppRouterHandler{
		versionController: versionController,
		logger:            logger,
		appConfig:         appConfig,
	}
}

// InitRouter registers routing patterns in the global app router
func (h *AppRouterHandler) InitRouter() *fastHttpRouter.Router {
	// Auto-generate Swagger documentation
	err := generateSwaggerDocs()
	if err != nil {
		h.logger.Fatalf("Failed to generate Swagger docs: %v", err)
	}

	r := fastHttpRouter.New()
	r.GlobalOPTIONS = middleware.CORSDisablingMiddleware(func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
	})
	r.GET(httppath.AnalyticBasePath, h.Index)

	// Define router with Swagger
	docs.SwaggerInfo.Version = h.appConfig.AppVersion
	docs.SwaggerInfo.Host = h.appConfig.ServiceExternalUrl
	docs.SwaggerInfo.Schemes = []string{h.appConfig.ServiceProtocol}
	docs.SwaggerInfo.BasePath = httppath.AnalyticBasePath

	//Swagger
	r.GET("/docs/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))

	// Define base group
	baseGroup := r.Group(httppath.AnalyticBasePath)

	// Define api group
	baseGroup.GET(httppath.VersionPath, h.versionController.GetVersion)

	return r
}

func (h *AppRouterHandler) Index(ctx *fasthttp.RequestCtx) {
	_, err := ctx.WriteString("Welcome to API!")
	if err != nil {
		h.logger.Panic("Can't write response", err)
	}
}

// generateSwaggerDocs runs the swag command to generate documentation
func generateSwaggerDocs() error {
	cmd := exec.Command("swag", "init", "--dir", "./cmd,./http/rest")
	cmd.Dir = "." // Set the working directory if needed
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
