package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hashicorp/vault/api"
)

const (
	serviceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	tokenAuth               = "auth/kubernetes/login"
)

func getServiceAccountToken() (string, error) {
	jwt, err := ioutil.ReadFile(serviceAccountTokenPath)
	if err != nil {
		return "", err
	}
	return string(jwt), nil
}

type Client struct {
	cli *api.Client
}

func getClient(addr string) (Whisper, error) {
	config := api.DefaultConfig()
	if strings.HasPrefix(addr, "http://") {
		config.Address = addr
	} else {
		config.Address = "http://" + addr
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// saRole  is a name that the service account binding,
// it binds some policy which limits to access data.
func (c *Client) getLoginToken(saRole string) (string, error) {
	saToken, err := getServiceAccountToken()
	if err != nil {
		return "", err
	}
	options := map[string]interface{}{
		"jwt":  saToken,
		"role": saRole,
	}

	loginSecret, err := c.cli.Logical().Write(tokenAuth, options)
	if err != nil {
		return "", err
	}

	if loginSecret != nil && loginSecret.Auth != nil {
		return loginSecret.Auth.ClientToken, nil
	}
	return "", fmt.Errorf("no token")
}

type Secret struct {
	data *api.Secret
}

func (c *Client) GetSecret(saRole, keyPath string) (*Secret, error) {
	loginToken, err := c.getLoginToken(saRole)
	if err != nil {
		return nil, err
	}
	c.cli.SetToken(loginToken)
	sec, err := c.cli.Logical().Read(keyPath)
	return &Secret{sec}, err
}
