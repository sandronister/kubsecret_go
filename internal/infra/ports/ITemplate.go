package ports

import (
	"bytes"

	"github.com/sandronister/kubsecret_generate/internal/dto"
)

type ITemplate interface {
	GetTemplate(kubMap dto.MapKub) (bytes.Buffer, error)
}
