package interfaces

type Logger interface {
	PrintlnUser(messages ...any)
	PrintlnProduct(messages ...any)
}
