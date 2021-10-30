package internal

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
	"mjeffn/wireguard-okta/pkg/conf"
)


type OktaHandler struct {
	Conf conf.OktaServerConfig
}

type userProfile struct {
	Email string `json:"email"`
}

// GetAllowedEmails calls Okta API and returns the list of email addresses for the specified group
// TODO - check about okta pagination and the default number of users in one response.
// uses json marshaling and unmarshalling to extract the emails from user profiles which is map[string]interface{}.
// not sure if it's the best way, but it gets the job done.
func (oh OktaHandler) GetAllowedEmails() ([]string,error) {
	ctx, client, err := okta.NewClient(context.Background(),
		okta.WithOrgUrl(oh.Conf.OrgUrl), okta.WithToken(oh.Conf.ApiToken))
	if err != nil {
		log.Println(err)
		return nil, errors.New("error initiating okta handler")
	}
	users, _, err := client.Group.ListGroupUsers(ctx,oh.Conf.WireguardGroupId,nil)
	if err != nil {
		log.Println( err)
		return nil,errors.New("error fetching users of given group from okta")
	}
	var oupl []*okta.UserProfile
	for _,u := range users {
		oupl = append(oupl,u.Profile)
	}
	b, err := json.Marshal(oupl)
	if err != nil {
		log.Println(err)
		return nil,errors.New("error marshaling user profiles")
	}
	var ups []userProfile
	err = json.Unmarshal(b,&ups)
	if err != nil {
		log.Println(err)
		return nil,errors.New("error extracting emails from okta users")
	}
	var ues []string
	for _, up := range ups {
		ues = append(ues,up.Email)
	}
	return ues,nil
}
