package jumper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

// var ConfigDir = "~/.jumper"
// var ConfigPath = fmt.Sprintf("%s/config", ConfigDir)

// var config *Config

// func GetConfig() (result Config, err error) {
// 	if config == nil {
// 		result, err = getConfig()
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return *config, nil
// }

func rootDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/.jumper", nil
}

func path(p string) (string, error) {
	h, err := rootDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", h, p), nil
}

func configFile() (string, error) {
	return path("config.yaml")
}

func repoDir() (string, error) {
	return path("repos")
}

func InitConfig() error {

	configFile, err := configFile()
	if err != nil {
		return err
	}

	_, err = os.Stat(configFile)
	if err == nil {
		return nil
	}

	if os.IsExist(err) {
		return nil
	}

	// file is not exists , create it
	root, err := rootDir()
	if err != nil {
		return err
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.Mkdir(root, 0755)
		if err != nil {
			fmt.Printf("Error to create dir %s: %s\n", root, err.Error())
			return err
		}
	}

	_, err = os.OpenFile(configFile, os.O_CREATE, 0666)
	if err != nil && !os.IsExist(err) {
		log.Printf("Create %s error:%s", configFile, err.Error())
		return err
	}

	return nil
}

func GetConfig() (Config, error) {

	configF, err := configFile()
	if err != nil {
		return Config{}, err
	}

	bts, err := ioutil.ReadFile(configF)
	if err != nil && os.IsNotExist(err) {
		return Config{}, nil
	}

	if err != nil {
		return Config{}, err
	}

	config := &Config{}
	err = yaml.Unmarshal(bts, config)
	if err != nil {
		return Config{}, err
	}

	return *config, nil
}

func saveConfig(config Config) error {
	bts, err := yaml.Marshal(config)
	if err != nil {
		log.Printf("Save Config error:%s", err.Error())
		return err
	}

	err = InitConfig()
	if err != nil {
		return nil
	}

	configF, err := configFile()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configF, bts, 0666)
	if err != nil {
		log.Printf("Write Config error:%s", err.Error())
		return err
	}
	return nil
}

func AppendRepo(repo Repo) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	if config.Repos == nil {
		config.Repos = []Repo{}
	}

	for _, item := range config.Repos {
		if item.Name == repo.Name {
			return fmt.Errorf("repo named %s is already exists", item.Name)
		}
	}

	config.Repos = append(config.Repos, repo)
	err = saveConfig(config)
	return err
}

func RemoveRepo(name string) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	if config.Repos == nil {
		config.Repos = []Repo{}
	}

	results := []Repo{}
	for _, item := range config.Repos {
		if item.Name == name {
			continue
		}
		results = append(results, item)
	}

	config.Repos = results
	err = saveConfig(config)
	return err
}

func ListRepo() ([]Repo, error) {
	config, err := GetConfig()
	if err != nil {
		return []Repo{}, err
	}

	return config.Repos, nil
}
