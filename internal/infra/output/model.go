package output

import "github.com/sandronister/kubsecret_generate/configs"

type model struct {
	fileOutput string
}

func New(env *configs.Enviroment) *model {
	return &model{
		fileOutput: env.FileOutput,
	}
}
