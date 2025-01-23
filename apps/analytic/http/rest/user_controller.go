package rest

import (
	"analytic/service"
	"errors"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/util"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const (
	authenticationCodeQueryParam = "activation_code"

	activationFormUsernameFieldName = "username"
	activationFormPasswordFieldName = "password"
)

type UserController struct {
	userService service.UserService

	log *zap.SugaredLogger
}

func NewUserController(userService service.UserService, log *zap.SugaredLogger) *UserController {
	return &UserController{
		userService: userService,
		log:         log,
	}
}

// ActivateUser godoc
//
//	@Summary		Initialize user's login and password
//	@Description	activates user by setting up new login and password
//	@Tags			user
//
//	@Accept			x-www-form-urlencoded
//
//	@Param			activation_code	query		string	true	"code which is sent to user's email during registration (embedded in link)"
//
//	@Param			username		formData	string	true	"user login"	Format(email)
//	@Param			password		formData	string	true	"user password"	Format(password)
//
//	@Success		204				"user credentials are updeted"
//	@Failure		400
//	@Failure		401	"unable to authenticate with activation  code"
//	@Failure		409	"login is taken already"
//	@Failure		500
//
//	@Router			/activate [post]
func (c *UserController) ActivateUser(ctx *fasthttp.RequestCtx) {
	params, err := util.GetRequiredQueryArgs(ctx, authenticationCodeQueryParam)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	requiredFormFields, err := util.GetRequiredFormFields(
		ctx,
		activationFormUsernameFieldName,
		activationFormPasswordFieldName,
	)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	err = c.userService.ActivateUser(
		model.UserActivationCreds{
			Login:          requiredFormFields[activationFormUsernameFieldName],
			Password:       requiredFormFields[activationFormPasswordFieldName],
			ActivationCode: params[authenticationCodeQueryParam],
		},
		util.GetSenderInfo(ctx),
	)
	if errors.Is(err, service.ErrInvalidUserActivationCode) {
		ctx.Error(err.Error(), fasthttp.StatusUnauthorized)
		return
	}
	if errors.Is(err, service.ErrUserLoginAlreadyTaken) {
		ctx.Error(err.Error(), fasthttp.StatusConflict)
		return
	}
	if err != nil {
		ctx.Error("something went wrong", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.SetStatusCode(fasthttp.StatusNoContent)
}
