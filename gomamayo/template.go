package gomamayo

import "strings"

type templateData struct {
	Codes     map[string]CodeList
	Websites  []*Website
	RepoCount int
}

func (c *Config) ToTemplateData() templateData {
	data := templateData{
		Codes:     make(map[string]CodeList),
		Websites:  c.Websites,
		RepoCount: len(c.Codes),
	}

	for _, r := range c.Codes {
		if _, ok := data.Codes[r.Lang]; !ok {
			data.Codes[r.Lang] = make(CodeList, 0)
		}

		r.Desc = strings.ReplaceAll(r.Desc, "\n", "  \n  ")

		data.Codes[r.Lang] = append(data.Codes[r.Lang], r)
	}

	return data
}
