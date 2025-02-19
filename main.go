package main

import (
	"fmt"

	"github.com/sandronister/kubsecret_generate/configs"
	"github.com/sandronister/kubsecret_generate/internal/di"
)

func main() {

	env := configs.LoadEnviroment()

	usecase := di.NewKubUsecase(env)

	err := usecase.Generate()

	if err != nil {
		fmt.Print(err)
		return
	}

}
