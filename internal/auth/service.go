package auth

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/eveisesi/skillz/internal/cache"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/oauth2"
)

type API interface {
	esi
	user
}

type esiAuth struct {
	oauthConfig *oauth2.Config
	jwks        jwk.Set
}

type userAuth struct {
	rsaKey      *rsa.PrivateKey
	tokenExpiry time.Duration
}

type Service struct {
	userAuth userAuth
	esiAuth  esiAuth
	client   *http.Client
	cache    cache.AuthAPI
}

func New(
	client *http.Client, cache cache.AuthAPI,
	esiOAuth *oauth2.Config, userPrivateKey *rsa.PrivateKey,
	esiAuthJWKSEndpoint string,
	userTokenExpiry time.Duration,
) *Service {
	s := &Service{
		client: client,
		cache:  cache,
		esiAuth: esiAuth{
			oauthConfig: esiOAuth,
		},
		userAuth: userAuth{
			rsaKey:      userPrivateKey,
			tokenExpiry: userTokenExpiry,
		},
	}

	err := s.initializeESIJWKSet(esiAuthJWKSEndpoint)
	if err != nil {
		panic(fmt.Errorf("internal.Auth: %w", err))
	}

	return s
}
