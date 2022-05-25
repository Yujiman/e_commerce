package middleware

import (
	"errors"
	"net/http"

	pbAuthentication "github.com/Yujiman/e_commerce/basket/gatway/internal/proto/authentication"
	helperError "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/basket/gatway/internal/server/http/helper/http"
	"github.com/Yujiman/e_commerce/basket/gatway/internal/service/authentication"
)

var (
	OauthDomainId = "oauth_domain_id"
	OauthUserId   = "oauth_user_id"
	OauthScopes   = "oauth_scopes"
)

func UserAuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, r *http.Request) {
		// Token from Header
		accessToken := r.Header.Get("AccessToken")
		if accessToken == "" {
			helperHttp.ErrorResponse(errors.New("AccessToken is not set"), response, http.StatusUnprocessableEntity)
			return
		}

		//Handler of gRPC
		data, err := authentication.CheckAccess(&pbAuthentication.CheckRequest{AccessToken: accessToken})
		if err != nil {
			helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
			return
		}

		r.Header.Del("AccessToken")

		r.Header.Add(OauthDomainId, data.DomainId)
		r.Header.Add(OauthUserId, data.UserId)
		r.Header.Add(OauthScopes, data.Scopes)

		next.ServeHTTP(response, r)
	}
}
