package templates

import (
	"github.com/sandronister/kubsecret_generate/configs"
	"github.com/sandronister/kubsecret_generate/internal/infra/ports"
)

type model struct {
	template string
}

func New(conf *configs.Enviroment) ports.ITemplate {
	return &model{
		template: conf.Template,
	}
}
