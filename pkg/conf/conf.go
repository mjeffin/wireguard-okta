package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type OktaServerConfig struct {
	ApiToken         string `json:"api_token" split_words:"true"`
	WireguardGroupId string `json:"wireguard_group_id" split_words:"true"`
	OrgUrl string `json:"org_url" split_words:"true"`
}

func GetOktaServerConfig() (OktaServerConfig,error) {
	var o OktaServerConfig
	err := envconfig.Process("oktaserver",&o)
	if err != nil {
		log.Println("Cannot read okta server config")
		return OktaServerConfig{}, err
	}
	return o,nil
}
func LoadEnvFile()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Local env file not found")
	}
}