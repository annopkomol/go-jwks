package jwks

import dlog "log"

type Logger interface {
	Print(v ...interface{})
	Fatal(v ...interface{})
	Printf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

var log Logger = defaultLogger{}

type defaultLogger struct {
}

func (d defaultLogger) Print(v ...interface{}) {
	dlog.Print(v)
}

func (d defaultLogger) Fatal(v ...interface{}) {
	dlog.Fatal(v)
}

func (d defaultLogger) Printf(format string, v ...interface{}) {
	dlog.Printf(format, v)
}

func (d defaultLogger) Fatalf(format string, v ...interface{}) {
	dlog.Fatalf(format, v)
}

func WithLogger(logger Logger) Option {
	return func(client JWKSClient) JWKSClient {
		log = logger
		return client
	}
}