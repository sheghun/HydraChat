package main

import (
	"Hydra/hydraconfigurator"
	"fmt"
	"os"
)

type Config struct {
	TS      string  `name:"testString"`
	TB      bool    `name:"testBool"`
	TF      float64 `name:"testFloat"`
	TestInt int     `name:"testInt"`
}

func main() {
	//Get base directory
	dir, err := os.Getwd()

	configStruct := new(Config)
	err = hydraconfigurator.GetCongiguration(hydraconfigurator.CUSTOM, configStruct, dir+"/hydraconfigurator/configFile.conf")
	fmt.Println(*configStruct)

	if configStruct.TB {
		fmt.Println("Bool is true")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(4.8 * configStruct.TF)

	fmt.Println(5 * configStruct.TestInt)

	fmt.Println(configStruct.TS)
}
