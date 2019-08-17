package jumper

import (
	"encoding/base64"
)

type Config struct {
	Current Current `json:"current"`
	Repos   []Repo  `json:"repos"`
}

type Repo struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	// UsernameStr is plaintext
	UsernameStr string `json:"-"`
	Password    string `json:"password"`
	// PasswordStr is plaintext
	PasswordStr string `json:"-"`
}

func (repo *Repo) SetUsernamePassword(username, password string) {
	repo.UsernameStr = username
	repo.PasswordStr = password
	repo.Username = base64.StdEncoding.EncodeToString([]byte(repo.UsernameStr))
	repo.Password = base64.StdEncoding.EncodeToString([]byte(repo.PasswordStr))
}

func (repo Repo) GetUsernamePassword() (string, string, error) {

	if repo.Username == "" {
		return "", "", nil
	}

	username, err := base64.StdEncoding.DecodeString(repo.Username)
	if err != nil {
		return "", "", err
	}
	password, err := base64.StdEncoding.DecodeString(repo.Password)
	if err != nil {
		return "", "", err
	}
	return string(username), string(password), nil
}

type Current struct {
	// Repo is current repo name
	Repo string
}
