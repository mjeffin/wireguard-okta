package okta

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
	"mjeffn/wireguard-okta/pkg/conf"
)


type OktaHandler struct {
	Conf conf.OktaServerConfig
}

type UserProfile struct {
	Email string `json:"email"`
}

func (oh OktaHandler) GetUsers()  {
	ctx, client, err := okta.NewClient(context.Background(),
		okta.WithOrgUrl(oh.Conf.OrgUrl), okta.WithToken(oh.Conf.ApiToken))
	if err != nil {
		log.Println("error initiating client",err)
	}
	users, _, err := client.Group.ListGroupUsers(ctx,oh.Conf.WireguardGroupId,nil)
	if err != nil {
		log.Println("Error fetching users of group")
	}
	fmt.Println(users[0].Profile)
	var u UserProfile
	b,_ := json.Marshal(users[0].Profile)
	json.Unmarshal(b,&u)
	fmt.Println(u)

}