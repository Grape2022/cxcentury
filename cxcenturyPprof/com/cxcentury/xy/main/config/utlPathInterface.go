package config

type DataValue struct {
	Http    string
	Https   string
	Host    string
	Port    string
	Api     string
	Url     string
	Path    *string
	UrlPath map[string]string
}
