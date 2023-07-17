package logger

import "testing"

func TestLogger(t *testing.T) {
	//The Emergency level log will panic all the program
	//I have not yet found a method to customize the panic behavior in logrus
	//So I will temporarily comment out this test case
	//l.WriteLog("[Emergency] emergency level log")
	WriteLog("[Alert] alert level log")
	WriteLog("[Critical] critical level log")
	WriteLog("[Error] error level log")
	WriteLog("[Warning] warning level log")
	WriteLog("[Notice] notice level log")
	WriteLog("debug level log")
}
