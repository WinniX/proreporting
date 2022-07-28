package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	"github.com/winnix/proreporting/api/apaproxy"
	"github.com/winnix/proreporting/api/domain"
	"github.com/winnix/proreporting/api/service"
	"github.com/winnix/proreporting/api/templates"
	"github.com/winnix/proreporting/api/web"
	"go.mongodb.org/mongo-driver/mongo"
)

type apaleoClaims struct {
	UserId      uuid.UUID `json:"sub"`
	AccountCode string    `json:"account_code"`
	Expiry      int64     `json:"exp"`
}

func (c *apaleoClaims) Valid() error {
	if len(c.AccountCode) == 0 {
		return errors.New("account_code claim is empty")
	}

	if c.UserId == uuid.Nil {
		return errors.New("sub claim is empty")
	}

	return nil
}

func Install(ctx *web.Context) web.HandlerResult {
	ctx.GinCtx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/auth/apaleo", ctx.Config.Host))
	return nil
}

func AuthApaleo(ctx *web.Context) web.HandlerResult {
	gothic.BeginAuthHandler(ctx.GinCtx.Writer, ctx.GinCtx.Request)
	return nil
}

func AuthApaleoCallback(ctx *web.Context) web.HandlerResult {
	user, err := gothic.CompleteUserAuth(ctx.GinCtx.Writer, ctx.GinCtx.Request)
	if err != nil {
		ctx.Logger.Warnf("apaleo authorization failed: %s\n", err.Error())
		return ServerError()
	}

	claims := apaleoClaims{}
	jwt.ParseWithClaims(user.AccessToken, &claims, nil)
	if err = claims.Valid(); err != nil {
		fmt.Println(user.AccessToken)
		ctx.Logger.Errorf("error parsing access token: %s\n", err.Error())
		return ServerError()
	}

	userId := claims.UserId

	userTokenService := service.NewUserTokenService(claims.AccountCode)
	userToken, err := userTokenService.GetCurrent(userId)
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.Logger.Errorf("error getting current user token: %s\n", err.Error())
		return ServerError()
	}

	if err != nil && err == mongo.ErrNoDocuments {
		userToken = &domain.UserToken{}
		userToken.UserUID = userId
		userToken.AccountCode = claims.AccountCode
	}

	userToken.AccessToken = user.AccessToken
	userToken.RefreshToken = user.RefreshToken
	userToken.ExpiresAt = time.Unix(claims.Expiry, 0).UTC()
	if err := userTokenService.SaveUserToken(userToken); err != nil {
		ctx.Logger.Errorf("failed saving user token: %s\n", err.Error())
		return ServerError()
	}

	ctx.SetCurrentUser(userId, userToken.AccountCode)

	proxy, err := apaproxy.NewApaProxy(ctx)
	if err != nil {
		ctx.Logger.Errorf("failed initializing apaleo proxy: %s\n", err.Error())
		return ServerError()
	}

	_, err = proxy.IntegrationProxy.CreateIntegration(fmt.Sprintf("%s/home", ctx.Config.Host))
	if err != nil {
		ctx.Logger.Errorf("failed saving ui integration: %s\n", err.Error())
		return ServerError()
	}

	redirectURL, exists := web.ExtractFromSession(ctx.GinCtx, web.RedirectURLSessionKey)
	if exists {
		if url, ok := redirectURL.(string); ok {
			ctx.GinCtx.Redirect(http.StatusTemporaryRedirect, url)
			ctx.GinCtx.AbortWithStatus(http.StatusTemporaryRedirect)
			return nil
		}

		ctx.Logger.Errorf("redirect URL is invalid: %s\n", redirectURL)
	}

	return Ok(&templates.InstallPage{
		AppUrl: "https://app.apaleo.com/",
	})
}
