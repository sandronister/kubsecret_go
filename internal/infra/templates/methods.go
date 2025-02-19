package templates

import (
	"bytes"
	"text/template"

	"github.com/sandronister/kubsecret_generate/internal/dto"
)

func (m *model) templateParse(name string) (*template.Template, error) {

	tmpl, err := template.New(name).Parse(m.template)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (m *model) GetTemplate(kubMap dto.MapKub) (bytes.Buffer, error) {
	var result bytes.Buffer
	tmpl, err := m.templateParse("secret")

	if err != nil {
		return bytes.Buffer{}, err
	}

	err = tmpl.Execute(&result, kubMap)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return result, nil

}
