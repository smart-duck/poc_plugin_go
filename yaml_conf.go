package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Plugin struct {
	Name string
	Path string
	Params string
}

type VeradcoCfg struct {
	Banner string
	Plugins []Plugin
}

func ReadConf(cfgFile string) (VeradcoCfg, error) {

	result := VeradcoCfg{}

	yfile, err := ioutil.ReadFile(cfgFile)

	if err != nil {

		log.Fatal(err)

		return result, err
	}

	err = yaml.Unmarshal(yfile, &result)

	if err != nil {
		log.Fatal(err)

		return result, err
	}

	return result, nil
}

// func main() {
  
// 	fmt.Println("Starting dummy veradco!")

// 	// Load conf from yaml
// 	conf, err := ReadConf("veradco.yaml")
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	} else {
// 		fmt.Printf("Conf: %v\n", conf)
// 	}
// }