package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

const (
	MMDOMAIN     = ""
	MMTOKEN      = ""
	MMCHANNELID  = ""
	GITLAB_URL   = ""
	GITLAB_TOKEN = ""
)

func main() {
	client := model.NewAPIv4Client(MMDOMAIN)
	client.SetToken(MMTOKEN)

	users, resp, err := client.GetUsersInChannel(MMCHANNELID, 0, 200, "")
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	if resp.StatusCode != 200 {
		fmt.Errorf("Error: %v", resp)
	}
	for _, user := range users {
		fmt.Println(user.Username)
	}

	// git, err := gitlab.NewClient(GITLAB_URL)
}
