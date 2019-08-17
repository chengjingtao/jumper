package jumper

import (
	"fmt"
	"io"
	"log"
	"os"

	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"gopkg.in/src-d/go-git.v4"
	httptransport "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Area struct {
	Name    string   `json:"-"`
	Servers []Server `json:"servers`
}

type Server struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	IP   string `json:"ip"`
	User string `json:"user"`
}

func InspectRepo(name string) ([]Area, error) {
	if _, err := ensureRepoExist(name); err != nil {
		return []Area{}, err
	}

	repoDir, err := repoDir()
	if err != nil {
		return []Area{}, err
	}

	dir := fmt.Sprintf("%s/%s", repoDir, name)
	pathes := []string{}
	err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		if !strings.HasSuffix(f.Name(), ".yaml") {
			return nil
		}

		pathes = append(pathes, path)
		return nil
	})

	if err != nil {
		log.Printf("error to range files in repo %s,error:%s\n", name, err.Error())
		return []Area{}, err
	}

	results := []Area{}
	for _, path := range pathes {

		bts, err := ioutil.ReadFile(path)
		if err != nil && os.IsNotExist(err) {
			return []Area{}, nil
		}
		area := Area{}

		err = yaml.Unmarshal(bts, &area)
		if err != nil {
			fmt.Printf("Unmarshal file %s error:%s", path, err.Error())
			continue
		}
		segs := strings.Split(path, "/")
		area.Name = strings.TrimRight(segs[len(segs)-1], ".yaml")

		results = append(results, area)
	}

	return results, nil
}

func ensureRepoExist(name string) (*Repo, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	var targetRepo *Repo
	for _, repo := range config.Repos {
		if repo.Name == name {
			targetRepo = &repo
		}
	}

	if targetRepo == nil {
		return nil, fmt.Errorf("repo %s is not exist", name)
	}

	return targetRepo, nil
}

func UpdateRepo(name string, out io.Writer) error {

	targetRepo, err := ensureRepoExist(name)
	if err != nil {
		return err
	}
	repoDir, err := repoDir()
	if err != nil {
		return err
	}

	if _, err = os.Stat(repoDir); os.IsNotExist(err) {
		err = os.Mkdir(repoDir, 0755)
		if err != nil {
			log.Printf("Create dir %s error:%s", repoDir, err.Error())
			return err
		}
	}

	path := repoDir + "/" + name
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		opt, err := cloneOptions(*targetRepo, out)
		if err != nil {
			return err
		}
		_, err = git.PlainClone(path, false, opt)
		return err
	}

	repository, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	wr, err := repository.Worktree()
	if err != nil {
		log.Printf("get git repository %s worktree error:%s", name, err.Error())
		return err
	}
	opt, err := pullOptions(*targetRepo, out)
	if err != nil {
		return err
	}
	err = wr.Pull(opt)
	if err != nil {
		log.Printf("Pull git repository %s error:%s", name, err.Error())
		return err
	}

	return nil
}

func cloneOptions(repo Repo, out io.Writer) (*git.CloneOptions, error) {
	opt := &git.CloneOptions{
		URL:      repo.Url,
		Progress: out,
	}
	username, password, err := repo.GetUsernamePassword()
	if err != nil {
		return nil, err
	}
	if username != "" {
		opt.Auth = &httptransport.BasicAuth{
			Username: username,
			Password: password,
		}
	}
	return opt, nil
}

func pullOptions(repo Repo, out io.Writer) (*git.PullOptions, error) {
	opt := &git.PullOptions{
		RemoteName: "origin",
		Progress:   out,
	}
	username, password, err := repo.GetUsernamePassword()
	if err != nil {
		return nil, err
	}
	if username != "" {
		opt.Auth = &httptransport.BasicAuth{
			Username: username,
			Password: password,
		}
	}
	return opt, nil
}
