package log

import "log"

type Log struct {
	Log_enabled bool
}

func (l *Log) Printf(format string, a ...any) {

	if l.Log_enabled {
		log.Printf(format, a...)
	}
}

func (l *Log) Println(a ...any) {

	if l.Log_enabled {
		log.Println(a...)
	}
}

func (l *Log) Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}
