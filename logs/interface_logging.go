package logs

type LogLevel int 

const(
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)