package main

import (
	"fmt"
	"plugin"
)

type VeradcoPlugin interface {
	Init(params string)
	Info() string
}

func main() {
    fmt.Println("Starting dummy veradco!")

	// Load conf from yaml
	conf, err := ReadConf("/home/lobuntu/go/src/test_plugin/veradco.yaml")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	} else {
		fmt.Printf("Conf: %v\n", conf)
	}

	// Execute plugins
	for _, v := range conf.Plugins {
		fmt.Printf("Loading plugin %s\n", v.Name)
		plug, err := plugin.Open(v.Path)
		if err != nil {
			fmt.Printf("Unable to load plugin %s: %v", v.Name, err)
			continue
		}
		pluginHandler, err := plug.Lookup("VeradcoPlugin")
		if err != nil {
			fmt.Printf("Unable to find handler for plugin %s: %v\n", v.Name, err)
			continue
		}

		var veradcoPlugin VeradcoPlugin

		veradcoPlugin, ok := pluginHandler.(VeradcoPlugin)
		if !ok {
			fmt.Printf("Plugin %s does not implement awaited interface\n", v.Name)
		} else {
			fmt.Printf("Run plugin %s\n", v.Name)
			veradcoPlugin.Init(v.Params)
			fmt.Printf("Plugin %s info: %s\n", v.Name, veradcoPlugin.Info())
		}

	}
}
