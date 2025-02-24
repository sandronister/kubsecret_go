package usecase

import (
	"github.com/sandronister/kubsecret_generate/internal/dto"
	"github.com/sandronister/kubsecret_generate/internal/infra/ports"
)

type KubSecretUsecase struct {
	input    ports.IInput
	output   ports.IOutput
	template ports.ITemplate
}

func NewKubSecretUsecase(input ports.IInput, output ports.IOutput, template ports.ITemplate) *KubSecretUsecase {
	return &KubSecretUsecase{
		input:    input,
		output:   output,
		template: template,
	}
}

func (k *KubSecretUsecase) generateFiles(path, password string) error {
	err := k.output.GenerateCert(path, password)
	if err != nil {
		return err
	}

	err = k.output.GenerateKey(path, password)
	if err != nil {
		return err
	}

	err = k.output.GenerateRSAKey()
	if err != nil {
		return err
	}

	return nil
}

func (k *KubSecretUsecase) Generate() error {
	if !k.output.IsFolder() {
		err := k.output.CreateFolder()
		if err != nil {
			return err
		}
	}

	path, err := k.input.GetPath()

	if err != nil {
		return err
	}

	name := k.input.GeKeyboardInput("Informe o nome")
	namespace := k.input.GeKeyboardInput("Informe o namespace")
	password := k.input.GeKeyboardInput("Informe a senha")

	err = k.generateFiles(path, password)

	if err != nil {
		return err
	}

	tlsCert, err := k.output.GenerateTLsCert()
	if err != nil {
		return err
	}

	tlsKey, err := k.output.GenerateTLsKey()
	if err != nil {
		return err
	}

	mapKub := dto.MapKub{
		Name:      name,
		Namespace: namespace,
		Cert:      tlsCert,
		Key:       tlsKey,
	}

	result, err := k.template.GetTemplate(mapKub)

	if err != nil {
		return err
	}

	err = k.output.DeleteFolder()

	if err != nil {
		return err
	}

	return k.output.SaveFile(result.Bytes())

}
