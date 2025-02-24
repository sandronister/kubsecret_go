package output

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func (m *model) IsFolder() bool {
	stat, err := os.Stat(m.folderOutput)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

func (m *model) CreateFolder() error {
	err := os.MkdirAll(m.folderOutput, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (m *model) DeleteFolder() error {
	err := os.RemoveAll(m.folderOutput)
	if err != nil {
		return err
	}
	return nil
}

func (m *model) GenerateCert(path, password string) error {
	cmd := exec.Command("openssl", "pkcs12", "-in", path, "-clcerts", "-nokeys", "-out", fmt.Sprintf("%s/cert_gen.crt", m.folderOutput), "-passin", fmt.Sprintf("pass:%s", password), "-passout", fmt.Sprintf("pass:%s", password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (m *model) GenerateKey(path, password string) error {
	cmd := exec.Command("openssl", "pkcs12", "-in", path, "-nocerts", "-out", fmt.Sprintf("%s/cert_gen.key", m.folderOutput), "-passin", fmt.Sprintf("pass:%s", password), "-passout", fmt.Sprintf("pass:%s", password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (m *model) GenerateRSAKey() error {
	cmd := exec.Command("openssl", "rsa", "-in", fmt.Sprintf("%s/cert_gen.key", m.folderOutput), "-out", fmt.Sprintf("%s/cert_gen_rsa.key", m.folderOutput))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (m *model) generateBase64(file string) (string, error) {
	var hash bytes.Buffer
	cmd := exec.Command("base64", file, "-w0")
	cmd.Stdout = &hash
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return hash.String(), nil
}

func (m *model) GetTLsCert() (string, error) {
	return m.generateBase64(fmt.Sprintf("%s/cert_gen.key", m.folderOutput))
}

func (m *model) GetTLsKey() (string, error) {
	return m.generateBase64(fmt.Sprintf("%s/cert_gen_rsa.key", m.folderOutput))
}

func (m *model) SaveFile(content []byte) error {
	return os.WriteFile(m.fileOutput, content, 0644)
}
