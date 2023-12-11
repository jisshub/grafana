package authn

import "github.com/grafana/grafana/pkg/util/errutil"

var (
	ErrTokenNeedsRotation  = errutil.Unauthorized("session.token.rotate", errutil.WithLogLevel(errutil.LevelDebug))
	ErrUnsupportedClient   = errutil.BadRequest("auth.client.unsupported")
	ErrClientNotConfigured = errutil.BadRequest("auth.client.notConfigured")
	ErrUnsupportedIdentity = errutil.NotImplemented("auth.identity.unsupported")
	ErrInvalidIdentityID   = errutil.BadRequest("auth.identity.invalid")
	ErrExpiredAccessToken  = errutil.Unauthorized("oauth.expired-token", errutil.WithPublicMessage("OAuth access token expired"))
)
