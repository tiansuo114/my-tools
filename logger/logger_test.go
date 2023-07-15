package logger

import "testing"

func TestLogger(t *testing.T) {
	l := NewLogger()

	//l.WriteLog("[Emergency] emergency level log")
	l.WriteLog("[Alert] alert level log")
	l.WriteLog("[Critical] critical level log")
	l.WriteLog("[Error] error level log")
	l.WriteLog("[Warning] warning level log")
	l.WriteLog("[Notice] notice level log")
	l.WriteLog("debug level log")
}
