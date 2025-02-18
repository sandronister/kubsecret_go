package input

import (
	"bufio"
	"os"
)

type model struct {
	reader *bufio.Reader
}

func New() *model {
	return &model{
		reader: bufio.NewReader(os.Stdin),
	}
}
