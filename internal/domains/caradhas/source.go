package caradhas

import (
	"fmt"

	"github.com/merci-app/dock-sdk-backend/client"
	"github.com/samuelmachado/cli-poc-dock/internal/domains/caradhas/config"
)

type infoDTO struct {
	Version string
}

type SourceI interface {
}

type CaradhasManager struct {
	Client *client.Master
	Token  string
}

//New Caradhas
func NewCaradhasManager(token string) *CaradhasManager {

	apiClient := &client.Master{}
	config.SDKClient(apiClient, token)

	return &CaradhasManager{Client: apiClient, Token: token}
}

//Info retrieves intormation about the caradhas version
func (m *CaradhasManager) Ping() string {
	return fmt.Sprintf("You are logged with key %s", m.Token)
}
