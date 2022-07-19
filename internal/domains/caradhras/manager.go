package caradhras

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/merci-app/dock-sdk-backend/client"
	"github.com/samuelmachado/cli-poc-dock/internal/domains/caradhras/config"
)

type infoDTO struct {
	Version string
}

type ManagerI interface {
}

type CaradhasManager struct {
	Client      *client.Master
	token       string
	Environment string
}

//New Caradhras
func NewCaradhasManager() *CaradhasManager {

	apiClient := &client.Master{}

	return &CaradhasManager{Client: apiClient}
}

func (m *CaradhasManager) Auth(token string) {
	m.token = token
	err := m.AuthFromFile() // will always fail to force authenticate on sdk
	if err != nil {
		config.SDKClient(m.Client, m.token)
		m.WriteKey()
	}

}
func (m *CaradhasManager) WriteKey() {
	//TODO: Must write the credential to a file
}

func (m *CaradhasManager) AuthFromFile() error {
	//TODO: checks that the credential exists and is valid. If not, return an error
	return ErrAuthRequired
}

//Ping retrieves intormation about the Caradhras connection
func (m *CaradhasManager) Ping() string {
	return fmt.Sprintf("You are logged with key %s", m.Client.GetKey())
}

func (m *CaradhasManager) Version() string {
	// for academic purposes we will simulate an error in 50% of cases
	rand.Seed(time.Now().UnixNano())
	randomNumber := (rand.Intn(10-1) + 1)
	if randomNumber%2 == 0 {
		return ErrAuthRequired.Error()
	}
	//
	return fmt.Sprintf("the version of Caradhras is vX.X.X")
}
