package logz

import (
	"fmt"
	"log"
)

type Default struct{
	minLevel LogLevel
	loggers map[LogLevel] *log.Logger
	triggerPanic bool
}

func (l *Default) MinLogLevel() LogLevel {
    return l.minLevel
}

func (l *Default) write(level LogLevel, message string) {
 if (l.minLevel <= level) {
 l.loggers[level].Output(2, message)
 }
}
func (l *Default) Trace(msg string) {
 l.write(Trace, msg)
}
func (l *Default) Tracef(template string, vals ...interface{}) {
 l.write(Trace, fmt.Sprintf(template, vals...))
}
func (l *Default) Debug(msg string) {
 l.write(Debug, msg)
}

func (l *Default) Debugf(template string, vals ...interface{}) {
 l.write(Debug, fmt.Sprintf(template, vals...))
}
func (l *Default) Info(msg string) {
 l.write(Information, msg)
}
func (l *Default) Infof(template string, vals ...interface{}) {
 l.write(Information, fmt.Sprintf(template, vals...))
}
func (l *Default) Warn(msg string) {
 l.write(Warning, msg)
}
func (l *Default) Warnf(template string, vals ...interface{}) {
 l.write(Warning, fmt.Sprintf(template, vals...))
}

func (l *Default) Panic(msg string) {
 l.write(Fatal, msg)
 if (l.triggerPanic) {
	 panic(msg)
 }
}
func (l *Default) Panicf(template string, vals ...interface{}) {
 formattedMsg := fmt.Sprintf(template, vals...)
 l.write(Fatal, formattedMsg)
 if (l.triggerPanic) {
 panic(formattedMsg)
 }
}