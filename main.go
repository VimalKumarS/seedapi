package main

import (
	"flag"
	"fmt"
	"seedapi/util"
	"seedapi/route"
	"seedapi/model"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	name = "SeedAPI"
)

func main() {

	
	//get the command line env argument --
	envPtr := flag.String("env", "", "define the env string")
	flag.Parse()
	//fmt.Println(len(os.Args), os.Args)
	if *envPtr != "" {
		fmt.Println("Parsing appsetting " + *envPtr)
		util.Logger.Info("Parsing appsetting " + *envPtr)
	}
	appsetting := util.GetConfig(*envPtr)

	// Use a single logger in the whole application.
	util.SetupLogging(appsetting)
	// Need to close it at the end.
	defer util.Logger.Sync()

	//fmt.Println(appsetting.Port)
	router := route.LoadRouter(appsetting, util.Logger)

	server := model.CreateServer(appsetting, util.Logger)
	server.Run(router)

	// graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	
	util.Logger.Info("Stopping the http server")

	if err := server.Shutdown(ctx); err != nil {
		util.Logger.Error(err.Error())
	}
}
