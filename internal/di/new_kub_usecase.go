package di

import (

	"github.com/sandronister/kubsecret_generate/configs"
	"github.com/sandronister/kubsecret_generate/internal/infra/input"
	"github.com/sandronister/kubsecret_generate/internal/infra/output"
	"github.com/sandronister/kubsecret_generate/internal/infra/templates"
	"github.com/sandronister/kubsecret_generate/internal/usecase"
)

func NewKubUsecase(enviroment *configs.Enviroment) *usecase.KubSecretUsecase {
	inputObj := input.New()
	outputObj := output.New(enviroment)
	templateObj := templates.New(enviroment)
	return usecase.NewKubSecretUsecase(inputObj, outputObj, templateObj)
}
