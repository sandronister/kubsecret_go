package output

import (
	"bytes"
	"os"
	"os/exec"
)

func (m *model) GenerateCert(path string) error {
	cmd := exec.Command("openssl", "pkcs12", "-in", path, "-clcerts", "-nokeys", "-out", "cert_gen.crt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (m *model) GenerateKey(path string) error {
	cmd := exec.Command("openssl", "pkcs12", "-in", path, "-nocerts", "-out", "cert_gen.key")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (m *model) GenerateRSAKey() error {
	cmd := exec.Command("openssl", "rsa", "-in", "cert_gen.key", "-out", "cert_gen_rsa.key")
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

func (m *model) GenerateTLsCert() (string, error) {
	return m.generateBase64("cert_gen.crt")
}

func (m *model) GenerateTLsKey() (string, error) {
	return m.generateBase64("cert_gen_rsa.key")
}

func (m *model) SaveFile(content []byte) error {
	return os.WriteFile(m.fileOutput, content, 0644)
}
