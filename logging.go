package logging

type Logger interface {
	Close() error
	Debug(interface{})
	Debugf(string, ...interface{})
	NewError(string) error
	NewErrorf(string, ...interface{}) error
	Error(error) bool
	DebugJSON(interface{})
	ErrorJSON(interface{})
	Reflect(interface{})
	Panic(interface{})
	Fatal(interface{})
}
