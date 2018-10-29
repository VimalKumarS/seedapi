package main

import (
	"flag"
	"fmt"
	"seedapi/util"
	"seedapi/model"
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
	}
	appsetting := util.GetConfig(*envPtr)
	//fmt.Println(appsetting.Port)
	router := util.LoadRouter(appsetting)

	server := model.CreateServer(appsetting)
	server.Run(router)
}
