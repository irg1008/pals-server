package supertokens

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				if err := next(c); err != nil {
					c.Error(err)
				}
			})
			supertokens.Middleware(handler).ServeHTTP(c.Response(), c.Request())
			return nil
		}
	}
}

func GetHeaders() []string {
	return supertokens.GetAllCORSHeaders()
}

// func CORSMiddleware() echo.MiddlewareFunc {
// 	baseHeaders := []string{echo.HeaderContentType}
// 	headers := append(baseHeaders, supertokens.GetAllCORSHeaders()...)
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			if c.Request().Method == http.MethodOptions {
// 				c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, strings.Join(headers, ","))
// 				return nil
// 			}
// 			return next(c)
// 		}
// 	}
// }

func VerifySessionWithOptions(opts *sessmodels.VerifySessionOptions) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				c.Set("session", session.GetSessionFromRequestContext((r.Context())))
				next(c)
			})

			session.VerifySession(opts, handler).ServeHTTP(c.Response(), c.Request())
			return nil
		}
	}
}

func VerifySession() echo.MiddlewareFunc {
	return VerifySessionWithOptions(nil)
}
