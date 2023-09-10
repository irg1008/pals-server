package auth

import (
	"log"

	"github.com/go-pkgz/auth/provider"
)

type EmailProviderOpts struct {
	Name  string
	Check func(user, pass string) (bool, error)
}

func (s *AuthService) AddEmailProvider(opts *EmailProviderOpts) {
	s.AddDirectProvider(opts.Name, provider.CredCheckerFunc(opts.Check))
	s.refreshProviders()
}

type GoogleProviderOpts struct {
	ClientID     string
	ClientSecret string
}

func (s *AuthService) AddGoogleProvider(opts *GoogleProviderOpts) {
	s.AddProvider("google", opts.ClientID, opts.ClientSecret)
	s.refreshProviders()
}

type AppleProviderOpts struct {
	*provider.AppleConfig
	PrivateKey string
}

func (opts *AppleProviderOpts) LoadPrivateKey() (b []byte, e error) {
	return []byte(opts.PrivateKey), nil
}

func (s *AuthService) AddCustomAppleProvider(opts *AppleProviderOpts) {
	err := s.AddAppleProvider(*opts.AppleConfig, opts)
	if err != nil {
		log.Fatal(err)
	}
	s.refreshProviders()
}

func (s *AuthService) refreshProviders() {
	s.Roles = NewRoles(s.Service.Middleware())
}
