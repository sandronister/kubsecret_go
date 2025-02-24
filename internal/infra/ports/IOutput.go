package ports

type IOutput interface {
	GenerateCert(path, password string) error
	GenerateKey(path, password string) error
	GenerateRSAKey() error
	GenerateTLsCert() (string, error)
	GenerateTLsKey() (string, error)
	SaveFile(content []byte) error
	IsFolder() bool
	CreateFolder() error
	DeleteFolder() error
}
