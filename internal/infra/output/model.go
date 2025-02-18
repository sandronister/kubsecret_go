package output

import (
	"github.com/sandronister/kubsecret_generate/configs"
	"github.com/sandronister/kubsecret_generate/internal/infra/ports"
)

type model struct {
	fileOutput string
}

func New(env *configs.Enviroment) ports.IOutput {
	return &model{
		fileOutput: env.FileOutput,
	}
}
