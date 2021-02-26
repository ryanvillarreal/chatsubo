package server

// this will be the main package that will start up the server

import (
	// builtins
	"os"
	"path"

	// 3rd party
	"github.com/ryanvillarreal/chatsubo/lib/logging"
	"github.com/ryanvillarreal/chatsubo/server/listeners"

)




func ServerStart() {
	logging.InitLogger()
	//logging.Logger.Println("C2 Server Starting...")
	logging.ServerLogger.Println("C2 Server Starting...")


	// setup global vars
	cwd, err := os.Getwd()
	if err != nil {
		logging.ErrorLogger.Println(err.Error())
		os.Exit(10)
	}


	goRoot := os.Getenv("GOROOT")
	logging.Logger.Println("Main.go goroot: ", goRoot)
	if goRoot == "" {
		newGoRoot := path.Join(cwd, "goroot", "go")
		os.Setenv("GOROOT", newGoRoot)
	}
	goPath := os.Getenv("GOPATH")
	logging.Logger.Println("Main.go gopath: ", goPath)
	if goPath == "" {
		newGoPath := path.Join(cwd, "gopath")
		os.Setenv("GOPATH", newGoPath)
	}

	// if all good start the server
	// specify which listener to use here. 
	listeners.StartTCPServer()

}
