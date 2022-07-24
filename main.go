package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
	fn "github.com/thoas/go-funk"
	"github.com/xanzy/go-gitlab"
)

const (
	MMDOMAIN     = ""
	MMTOKEN      = ""
	MMCHANNELID  = ""
	GITLAB_URL   = ""
	GITLAB_TOKEN = ""
)

type User struct {
	Username string
	Email    string
}

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
	// for _, user := range users {
	// 	fmt.Println(user.Username)
	// }

	git, err := gitlab.NewClient(GITLAB_TOKEN, gitlab.WithBaseURL(GITLAB_URL))
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fu := fn.Filter(fn.Map(users, func(user *model.User) User {
		return User{user.Username, user.Email}
	}), func(user User) bool {
		u, _, _ := git.Search.Users(user.Username, &gitlab.SearchOptions{})
		return !fn.Contains(fn.Map(u, func(u *gitlab.User) string {
			return u.Username
		}), user.Username)
	})
	fmt.Println(fu)
}
