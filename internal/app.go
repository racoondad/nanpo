package internal

type Application interface {
	Run()
}

func NewApplication() Application {
	globalAppOnce.Do(func() {})
	return globalApp
}
