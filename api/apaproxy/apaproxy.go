package apaproxy

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/winnix/proreporting/api/domain"
	"github.com/winnix/proreporting/api/service"
	"github.com/winnix/proreporting/api/web"
)

type ApaProxy struct {
	Token            *domain.UserToken
	TokenService     service.UserTokenService
	AuthProvider     *openidConnect.Provider
	IntegrationProxy IntegrationClient
	BookingProxy     BookingClient
}

func NewApaProxy(ctx *web.Context) (*ApaProxy, error) {
	authProvider, err := web.GetApaleoAuthProvider(ctx.Config)
	if err != nil {
		return nil, err
	}

	ap := &ApaProxy{
		TokenService: service.NewUserTokenService(ctx.User.AccountCode),
		AuthProvider: authProvider,
	}
	if err := ap.setToken(ctx.User.ID); err != nil {
		return nil, err
	}

	ap.initProxyClients(ctx.Ctx)

	return ap, nil
}

func (ap *ApaProxy) setToken(userID uuid.UUID) error {
	token, err := ap.TokenService.GetCurrent(userID)
	if err != nil {
		return err
	}

	ap.Token = token

	if err := ap.refreshTokenIfExpired(); err != nil {
		return err
	}

	return nil
}

func (ap *ApaProxy) initProxyClients(baseCtx context.Context) {
	ap.IntegrationProxy = newIntegrationClient(baseCtx, ap.Token.AccessToken)
	ap.BookingProxy = newBookingClient(baseCtx, ap.Token.AccessToken)
}

func (ap *ApaProxy) refreshTokenIfExpired() error {
	if delta, _ := time.ParseDuration("5m"); ap.Token.ExpiresAt.Sub(time.Now().UTC()) >= delta {
		return nil
	}

	newToken, err := ap.AuthProvider.RefreshToken(ap.Token.RefreshToken)
	if err != nil {
		return fmt.Errorf("refresh token failed: %s", err.Error())
	}

	ap.Token.AccessToken = newToken.AccessToken
	ap.Token.RefreshToken = newToken.RefreshToken
	ap.Token.ExpiresAt = newToken.Expiry.UTC()

	if err := ap.TokenService.SaveUserToken(ap.Token); err != nil {
		return err
	}

	return nil
}
