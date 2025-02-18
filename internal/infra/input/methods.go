package input

import (
	"fmt"
	"os"
	"strings"
)

func (m *model) GetPath() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("informe o caminho do certificado")
	}

	path := os.Args[1]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("arquivo não encontrado")
	}

	return path, nil
}

func (m *model) GeKeyboardInput(message string) string {
	fmt.Println(message)
	info, _ := m.reader.ReadString('\n')
	return strings.TrimSpace(info)
}
