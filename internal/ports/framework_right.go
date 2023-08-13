package ports

type DB interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error
}