package gomamayo

import (
	"os"
	"slices"

	"github.com/pelletier/go-toml/v2"
)

type Repo struct {
	Lang  string `toml:"lang"`
	Label string `toml:"label"`
	Url   string `toml:"url"`
}

type Website struct {
	Title string `toml:"title"`
	Url   string `toml:"url"`
}

type RepoList []*Repo

type Config struct {
	Repos    RepoList  `toml:"repos"`
	Websites []*Website `toml:"websites"`
}

func (l *RepoList) GetLangList() []string {
	var list []string
	for _, r := range *l {
		list = append(list, r.Lang)
	}

	return slices.Compact(list)

}

func Read(path string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	toml.Unmarshal(data, &config)
	return &config, nil
}
