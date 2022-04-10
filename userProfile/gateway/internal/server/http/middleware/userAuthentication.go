package middleware

//
//import (
//	"errors"
//	"net/http"
//
//	pbAuthentication "github.com/Yujiman/e_commerce/userProfile/gatway/internal/proto/authentication"
//	helperError "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/error"
//	helperHttp "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/http"
//	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/authentication"
//)
//
//var (
//	OauthDomainId = "oauth_domain_id"
//	OauthUserId   = "oauth_user_id"
//	OauthScopes   = "oauth_scopes"
//)
//
//func UserAuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
//	return func(response http.ResponseWriter, r *http.Request) {
//		// Token from Header
//		accessToken := r.Header.Get("X-Satrap-1")
//		if accessToken == "" {
//			helperHttp.ErrorResponse(errors.New("X-Satrap-1 is not set"), response, http.StatusUnprocessableEntity)
//			return
//		}
//
//		//Handler of gRPC
//		data, err := authentication.CheckAccess(&pbAuthentication.CheckRequest{AccessToken: accessToken})
//		if err != nil {
//			helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
//			return
//		}
//
//		r.Header.Del("X-Satrap-1")
//
//		r.Header.Add(OauthDomainId, data.DomainId)
//		r.Header.Add(OauthUserId, data.UserId)
//		r.Header.Add(OauthScopes, data.Scopes)
//
//		next.ServeHTTP(response, r)
//	}
//}
