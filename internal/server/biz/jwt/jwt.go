package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

var (
	jwtParser = new(jwt.Parser)
)

// JwtService the middleware for JSON Web tokens authentication method
type JwtService struct {
	Config Config
}

// New constructs a new Secure instance with supplied options.
func NewJwtService() *JwtService {
	var mySecret = []byte("HS2JDFKhu7Y1av7b")

	config := Config{
		ContextKey:   DefaultContextKey,
		Extractor:    FromAuthHeader,
		ErrorHandler: OnError,

		SigningMethod: jwt.SigningMethodHS256,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
	}

	return &JwtService{Config: config}
}

// Serve the middleware's action
func (m *JwtService) Serve(ctx iris.Context) {
	err := m.CheckJWT(ctx)
	if err != nil {
		m.Config.ErrorHandler(ctx, err)
		return
	}

	// aaron: check timeout

	// If everything ok then call next.
	ctx.Next()
}

// CheckJWT the main functionality, checks for token
func (m *JwtService) CheckJWT(ctx iris.Context) error {
	if !m.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// Use the specified token extractor to extract a token from the request
	token, err := m.Config.Extractor(ctx)

	// If debugging is turned on, log the outcome
	if err != nil {
		Logf(ctx, "Error extracting JWT: %v", err)
		return err
	}

	Logf(ctx, "Token extracted: %s", token)

	// If the token is empty...
	if token == "" {
		// CheckStatus if it was required
		if m.Config.CredentialsOptional {
			Logf(ctx, "No credentials found (CredentialsOptional=true)")
			// No error, just no token (and that is ok given that CredentialsOptional is true)
			return nil
		}

		// If we get here, the required token is missing
		Logf(ctx, "Error: No credentials found (CredentialsOptional=false)")
		return ErrTokenMissing
	}

	// Now parse the token

	parsedToken, err := jwtParser.Parse(token, m.Config.ValidationKeyGetter)
	// CheckStatus if there was an error in parsing...
	if err != nil {
		Logf(ctx, "Error parsing token: %v", err)
		return err
	}

	if m.Config.SigningMethod != nil && m.Config.SigningMethod.Alg() != parsedToken.Header["alg"] {
		err := fmt.Errorf("Expected %s signing method but token specified %s",
			m.Config.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		Logf(ctx, "Error validating token algorithm: %v", err)
		return err
	}

	// CheckStatus if the parsed token is valid...
	if !parsedToken.Valid {
		Logf(ctx, "Token is invalid")
		// m.Config.ErrorHandler(ctx, ErrTokenInvalid)
		return ErrTokenInvalid
	}

	if m.Config.Expiration {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			if expired := claims.VerifyExpiresAt(time.Now().Unix(), true); !expired {
				Logf(ctx, "Token is expired")
				return ErrTokenExpired
			}
		}
	}

	Logf(ctx, "JWT: %v", parsedToken)

	// If we get here, everything worked and we can set the
	// user property in context.
	ctx.Values().Set(m.Config.ContextKey, parsedToken)

	return nil
}
