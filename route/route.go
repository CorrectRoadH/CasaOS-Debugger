package route

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/CorrectRoadH/CasaOS-Debugger/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
)

type DebuggerRoute struct{}

var (
	_swagger *openapi3.T

	APIPath string
	DocPath string
)

func init() {
	swagger, err := codegen.GetSwagger()
	if err != nil {
		panic(err)
	}

	_swagger = swagger

	u, err := url.Parse(_swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	APIPath = strings.TrimRight(u.Path, "/")
	DocPath = "/doc" + APIPath
}

func GetRouter() http.Handler {
	hello := NewDebuggerService()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderXCSRFToken, echo.HeaderContentType, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderConnection, echo.HeaderOrigin, echo.HeaderXRequestedWith},
		ExposeHeaders:    []string{echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders},
		MaxAge:           172800,
		AllowCredentials: true,
	}))

	e.Use(middleware.Gzip())

	e.Use(middleware.Logger())

	e.Use(echomiddleware.OapiRequestValidatorWithOptions(_swagger, &echomiddleware.Options{
		Options: openapi3filter.Options{AuthenticationFunc: openapi3filter.NoopAuthenticationFunc},
	}))

	codegen.RegisterHandlersWithBaseURL(e, hello, APIPath)

	return e
}

func GetDocRouter(docHTML string, docYAML string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == DocPath {
			if _, err := w.Write([]byte(docHTML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		if r.URL.Path == DocPath+"/openapi.yaml" {
			if _, err := w.Write([]byte(docYAML)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}

func NewDebuggerService() codegen.ServerInterface {
	return &DebuggerRoute{}
}

func PropertiesFromQueryParams(httpCtx echo.Context) map[string]string {
	properties := make(map[string]string)

	for k, values := range httpCtx.QueryParams() {
		if len(values) > 0 {
			properties[k] = values[0]
		}
	}

	return properties
}
