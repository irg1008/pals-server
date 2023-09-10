package auth

func MockService() *AuthService {
	opts := BaseOptions{
		AppName:   "mock",
		JWTSecret: "secret",
		URL:       "http://localhost:8001",
		Local:     true,
	}

	service := NewAuthService(&opts)

	service.AddEmailProvider(&EmailProviderOpts{
		Name: "email",
		Check: func(user, pass string) (bool, error) {
			return true, nil
		},
	})

	service.AddGoogleProvider(&GoogleProviderOpts{
		ClientID:     "id",
		ClientSecret: "secret",
	})

	return service
}
