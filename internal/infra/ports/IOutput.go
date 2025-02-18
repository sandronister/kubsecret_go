package ports

type IOutput interface {
	GenerateCert(path string) error
	GenerateKey(path string) error
	GenerateRSAKey() error
	GenerateTLsCert() (string, error)
	GenerateTLsKey() (string, error)
	SaveFile(content []byte) error
}
