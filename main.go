package main

import (
	"github.com/echenim/gost/logz"
)


func main(){
//var info logz.Logha = logz.New(logz.Information)
var debug logz.Logha = logz.New(logz.Debug)
 //writeMessage(info)
  writeMessage(debug)
}

func writeMessage(logger logz.Logha) {
 logger.Info("Hello, Platform")
 logger.Debug("Hello, Platform")
}