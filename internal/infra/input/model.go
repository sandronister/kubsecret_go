package input

import (
	"bufio"
	"os"

	"github.com/sandronister/kubsecret_generate/internal/infra/ports"
)

type model struct {
	reader *bufio.Reader
}

func New() ports.IInput {
	return &model{
		reader: bufio.NewReader(os.Stdin),
	}
}
