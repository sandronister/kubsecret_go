package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type MapKub struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Informe o caminho do certificado")
		return
	}

	path := os.Args[1]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Arquivo não encontrado")
		return
	}

	fmt.Println("Caminho do certificado: ", path)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe o nome")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Informe o namespace")
	namespace, _ := reader.ReadString('\n')
	namespace = strings.TrimSpace(namespace)

	cmd := exec.Command("openssl", "pkcs12", "-in", path, "-clcerts", "-nokeys", "-out", "cert_prd.crt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic("Erro ao executar o comando")
	}
	cmd = exec.Command("openssl", "pkcs12", "-in", path, "-nocerts", "-out", "cert_prd.key")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic("Erro ao executar o comando")
	}

	cmd = exec.Command("openssl", "rsa", "-in", "cert_prd.key", "-out", "cert_prd_rsa.key")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic("Erro ao executar o comando")
	}

	var firstOut, secondOut bytes.Buffer
	cmd = exec.Command("base64", "cert_prd.key", "-w0")
	cmd.Stdout = &firstOut
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		panic("Erro ao executar o comando")
	}

	cmd = exec.Command("base64", "cert_prd_rsa.key", "-w0")
	cmd.Stdout = &secondOut
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic("Erro ao executar o comando")
	}

	tmplContent, err := os.ReadFile("templatekub.yaml")
	if err != nil {
		panic("Erro ao ler o arquivo de template")
	}

	tmpl, err := template.New("secret").Parse(string(tmplContent))
	if err != nil {
		panic("Erro ao criar o template")
	}

	var result bytes.Buffer
	maps := MapKub{
		Name:      name,
		Namespace: namespace,
		Cert:      firstOut.String(),
		Key:       secondOut.String(),
	}

	err = tmpl.Execute(&result, maps)
	if err != nil {
		panic("Erro ao executar o template")
	}

	err = os.WriteFile("kubsecret.yaml", result.Bytes(), 0644)
	if err != nil {
		panic("Erro ao escrever o arquivo")
	}

	fmt.Println("Arquivo kubsecret.yaml criado com sucesso")

}
