package profile

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/pkg/helpers"
)

type ProfileManager struct {
	Environment string
}

func NewProfileManager() *ProfileManager {
	return &ProfileManager{}
}

func (m *ProfileManager) Create() error {
	helpers.PrintHeadline("Let's set up your new profile")

	profile, err := helpers.ReadText("type a name for it: ")
	if err != nil {
		return err
	}

	helpers.PrintHeadline("Let's configure your credentials for the SANDBOX environment")
	clientId, err := helpers.ReadText("client_id: ")
	if err != nil {
		return err
	}

	clientSecret, err := helpers.ReadPassword("client_secret: ")
	if err != nil {
		return err
	}

	successHeadline := fmt.Sprintf("you can select the profile by running: dock profile select %s sandbox", profile)
	helpers.PrintHeadline(successHeadline)

	fmt.Println(clientId, clientSecret)
	m.saveEnvironment()
	return nil
}

type Credential struct {
	ClientId string `json:"client_id"`
	SecretId string `json:"secret_id"`
}
type Env struct {
	Sandbox    Credential
	Production Credential
}

func (m *ProfileManager) saveEnvironment() {

	//users := map[string]Env{"profile": Env{Sandbox: Credential{ClientId: "x1", SecretId: "x2"}, Production: Credential{ClientId: "x3", SecretId: "x4"}}}

	users := map[string]Env{"profile": {
		Sandbox:    Credential{ClientId: "a", SecretId: "b"},
		Production: Credential{ClientId: "a", SecretId: "b"},
	}}

	helpers.WriteFile(users)
}
