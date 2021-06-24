package jwt

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"strings"
)

func init() {
	context.SetHandlerName("github.com/iris-contrib/middleware/jwt.*", "iris-contrib.jwt")
}

var (
	// ErrTokenMissing is the error value that it's returned when
	// a token is not found based on the token extractor.
	ErrTokenMissing = errors.New("required authorization token not found")

	// ErrTokenInvalid is the error value that it's returned when
	// a token is not valid.
	ErrTokenInvalid = errors.New("token is invalid")

	// ErrTokenExpired is the error value that it's returned when
	// a token value is found and it's valid but it's expired.
	ErrTokenExpired = errors.New("token is expired")
)

// A function called whenever an error is encountered
type ErrorHandler func(iris.Context, error)

// TokenExtractor is a function that takes a context as input and returns
// either a token or an error.  An error should only be returned if an attempt
// to specify a token was found, but the information was somehow incorrectly
// formed.  In the case where a token is simply not present, this should not
// be treated as an error.  An empty string should be returned in that case.
type TokenExtractor func(iris.Context) (string, error)

// FromAuthHeader is a "TokenExtractor" that takes a give context and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(ctx iris.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

// FromParameter returns a function that extracts the token from the specified query string parameter
func FromParameter(param string) TokenExtractor {
	return func(ctx iris.Context) (string, error) {
		return ctx.URLParam(param), nil
	}
}

// FromFirst returns a function that runs multiple token extractors and takes the first token it finds
func FromFirst(extractors ...TokenExtractor) TokenExtractor {
	return func(ctx iris.Context) (string, error) {
		for _, ex := range extractors {
			token, err := ex(ctx)
			if err != nil {
				return "", err
			}
			if token != "" {
				return token, nil
			}
		}
		return "", nil
	}
}

// OnError is the default error handler.
// Use it to change the behavior for each error.
// See `Config.ErrorHandler`.
func OnError(ctx iris.Context, err error) {
	if err == nil {
		return
	}

	ctx.StopExecution()
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.WriteString(err.Error())
}

func Logf(ctx iris.Context, format string, args ...interface{}) {
	ctx.Application().Logger().Debugf(format, args...)
}
