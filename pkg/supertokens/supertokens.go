package supertokens

import (
	"fmt"
	"irg1008/pals/pkg/config"

	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Init(basePath string, appName string, config *config.Config) error {

	email := fmt.Sprintf("%s@%s", "noreply", config.Domain)
	smtpSettings := emaildelivery.SMTPSettings{
		Host: config.EmailHost,
		From: emaildelivery.SMTPFrom{
			Name:  appName,
			Email: email,
		},
		Port:     config.EmailPort,
		Username: &config.EmailUser,
		Password: config.EmailPass,
		Secure:   !config.IsDev,
	}

	return supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: config.AuthCoreUrl,
			APIKey:        config.AuthCoreKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         appName,
			APIDomain:       config.APIUrl,
			WebsiteDomain:   config.ClientUrl,
			APIBasePath:     &basePath,
			WebsiteBasePath: &basePath,
		},
		RecipeList: []supertokens.Recipe{
			dashboard.Init(nil),
			session.Init(nil),
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired,
				EmailDelivery: &emaildelivery.TypeInput{
					Service: emailverification.MakeSMTPService(emaildelivery.SMTPServiceConfig{
						Settings: smtpSettings,
					}),
				},
			}),
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				EmailDelivery: &emaildelivery.TypeInput{
					Service: thirdpartyemailpassword.MakeSMTPService(emaildelivery.SMTPServiceConfig{
						Settings: smtpSettings,
					}),
				},
				Providers: []tpmodels.ProviderInput{
					// We have provided you with development keys which you can use for testing.
					// IMPORTANT: Please replace them with your own OAuth keys for production use.
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
									ClientSecret: "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",
								},
							},
						},
					},
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "apple",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID: "4398792-io.supertokens.example.service",
									AdditionalConfig: map[string]interface{}{
										"keyId":      "7M48Y4RYDL",
										"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
										"teamId":     "YWQCXGJRJL",
									},
								},
							},
						},
					},
				},
			}),
		},
	})
}
