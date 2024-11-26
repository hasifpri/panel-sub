package applicationservice

type ILogger interface {
	Info(process, correlationID, step, key string, val any)
	Error(process, correlationID, step, key string, val any)
}
