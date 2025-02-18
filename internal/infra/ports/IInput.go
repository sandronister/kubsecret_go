package ports

type IInput interface {
	GetPath() (string, error)
	GeKeyboardInput(message string) string
}
