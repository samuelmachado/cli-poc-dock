package config

import (
	"github.com/merci-app/dock-sdk-backend"
	"github.com/merci-app/dock-sdk-backend/client"
)

func SDKClient(apiClient *client.Master, token string) {

	clientConfig := &dock.RequestHandlerConfig{
		Hooks: loggerHooks(),
	}
	apiClient.Init(token, nil, clientConfig)
}
