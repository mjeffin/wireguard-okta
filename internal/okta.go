package internal

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"log"
	"mjeffn/wireguard-okta/pkg/conf"
)

type userProfile struct {
	Email string `json:"email"`
}

// GetAllowedUsers calls Okta API and returns the list of email addresses for the specified group
// TODO - check about okta pagination and the default number of users in one response.
// uses json marshaling and unmarshalling to extract the emails from user profiles which is map[string]interface{}.
// not sure if it's the best way, but it gets the job done.
func GetAllowedUsers() ([]string, error) {
	config, err := conf.GetOktaServerConfig()
	if err != nil {
		log.Println("Error fetching okta config")
	}
	ctx, client, err := okta.NewClient(context.Background(),
		okta.WithOrgUrl(config.OrgUrl), okta.WithToken(config.ApiToken))
	if err != nil {
		log.Println(err)
		return nil, errors.New("error initiating okta handler")
	}
	users, _, err := client.Group.ListGroupUsers(ctx, config.WireguardGroupId, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error fetching users of given group from okta")
	}
	var oupl []*okta.UserProfile
	for _, u := range users {
		oupl = append(oupl, u.Profile)
	}
	b, err := json.Marshal(oupl)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error marshaling user profiles")
	}
	var ups []userProfile
	err = json.Unmarshal(b, &ups)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error extracting emails from okta users")
	}
	var ues []string
	for _, up := range ups {
		ues = append(ues, up.Email)
	}
	return ues, nil
}
