package templates

import "github.com/sandronister/kubsecret_generate/configs"

type model struct {
	templateName string
}

func New(conf *configs.Enviroment) *model {
	return &model{
		templateName: conf.TemplateName,
	}
}
