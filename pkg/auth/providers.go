package auth

import "github.com/go-pkgz/auth/provider"

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

func (s *AuthService) refreshProviders() {
	s.Roles = NewRoles(s.Service.Middleware())
}
