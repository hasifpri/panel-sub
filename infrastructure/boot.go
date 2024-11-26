package infrastructure

import "context"

func Initialize() {
	InitializeConfiguration()
	InitializeZap()
	InitializeValidator()
	InitializeTracer()
	InitializePostgresql()
	InitializeChannelEngine()
}

func CloseAll(ctx context.Context) {
	RDS.Close()
	ZAPLOGGER.Sync()
	CloseTracer(ctx)
	ClosePostgresql()
}
