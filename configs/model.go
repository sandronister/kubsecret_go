package configs

type Enviroment struct {
	Template     string
	FileOutput   string
	FolderOutPut string
}

func LoadEnviroment() *Enviroment {
	return &Enviroment{
		Template: `kind: Secret
apiVersion: v1
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
data:
  tls.crt: >-
    {{.Cert}}
  tls.key: >-
    {{.Key}}`,
		FileOutput:   "kubsecret.yaml",
		FolderOutPut: "certs",
	}

}
