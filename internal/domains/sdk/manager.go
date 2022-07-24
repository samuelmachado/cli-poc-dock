package sdk

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/merci-app/dock-sdk-backend/client"
	"github.com/samuelmachado/cli-poc-dock/internal/domains/profile"
	"github.com/samuelmachado/cli-poc-dock/internal/domains/sdk/config"
)

type SdkManagerI interface {
}

type SdkManager struct {
	Client         *client.Master
	profileManager *profile.ProfileManager
	token          string //TODO: remove this
}

//NewSdkManager
func NewSdkManager() *SdkManager {

	apiClient := &client.Master{}
	profileManager := profile.NewProfileManager()

	return &SdkManager{Client: apiClient, profileManager: profileManager}
}

func (m *SdkManager) Auth(token string) {
	m.token = token
	config.SDKClient(m.Client, m.token)
}

//Ping retrieves intormation about the connection
func (m *SdkManager) Ping() string {
	return fmt.Sprintf("You are logged with key %s", m.Client.GetKey())
}

func (m *SdkManager) Version() string {
	// for academic purposes we will simulate an error in 50% of cases
	rand.Seed(time.Now().UnixNano())
	randomNumber := (rand.Intn(10-1) + 1)
	if randomNumber%2 == 0 {
		return ErrAuthRequired.Error()
	}
	//
	return fmt.Sprintf("the version of SDK is vX.X.X")
}
