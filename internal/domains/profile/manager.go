package profile

import (
	"fmt"
	"os"

	"github.com/samuelmachado/cli-poc-dock/pkg/helpers"
)

type ProfileManager struct {
	Environment string
}

func NewProfileManager() *ProfileManager {
	return &ProfileManager{}
}

func (m *ProfileManager) Create() error {
	stdin := os.Stdin
	headline := helpers.CreateHeadline("Let's set up your new profile")
	fmt.Println(headline)

	profile, err := helpers.ReadText("type a name for it: ", stdin)
	if err != nil {
		return err
	}

	headline = helpers.CreateHeadline("Let's configure your credentials for the SANDBOX environment")
	fmt.Println(headline)
	clientId, err := helpers.ReadText("client_id: ", stdin)
	if err != nil {
		return err
	}

	clientSecret, err := helpers.ReadPassword("client_secret: ", os.Stdin)
	if err != nil {
		return err
	}

	successHeadline := fmt.Sprintf("you can select the profile by running: dock profile select %s sandbox", profile)
	headline = helpers.CreateHeadline(successHeadline)
	fmt.Println(headline)

	fmt.Println(clientId, clientSecret)
	return nil
}
func (m *ProfileManager) Select(name string) {

}

func (m *ProfileManager) Delete(name string) {

}
