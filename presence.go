package main

import (
	"github.com/hugolgst/rich-go/client"
)

func Connect(appID string) error {
	return client.Login(appID)
}

func UpdateStatus(state string) error {
	return client.SetActivity(client.Activity{
		State: state,
	})
}

func Disconnect() {
	client.Logout()
}
