package gomamayo

type templateData struct {
	Repos    map[string]RepoList
	Websites []*Website
}

func (c *Config) ToTemplateData() templateData {
	data := templateData{
		Repos:    make(map[string]RepoList),
		Websites: c.Websites,
	}

	for _, r := range c.Repos {
		if _, ok := data.Repos[r.Lang]; !ok {
			data.Repos[r.Lang] = make(RepoList, 0)
		}
		data.Repos[r.Lang] = append(data.Repos[r.Lang], r)
	}

	return data
}
