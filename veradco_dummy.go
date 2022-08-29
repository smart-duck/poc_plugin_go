package main

import (
	"fmt"
	"plugin"
	"regexp"
	// "gopkg.in/yaml.v3"
	// "io/ioutil"
	"io"
	"log"
	"archive/tar"
  "compress/gzip"
	"os"
	"encoding/base64"
	"path/filepath"
)

type VeradcoPlugin interface {
	Init(params string)
	Info() string
}

// type BinaryPackage struct {
// 	File string
// }

func ExtractTarGz(gzipStream io.Reader) string {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
			log.Fatalf("ExtractTarGz: NewReader failed: %s", err.Error())
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
			header, err := tarReader.Next()

			if err == io.EOF {
					break
			}

			if err != nil {
					log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
			}

			switch header.Typeflag {
			case tar.TypeDir:
					if err := os.Mkdir(header.Name, 0755); err != nil {
							log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
					}
			case tar.TypeReg:
					log.Printf("header.Name=%v", header.Name)
					outFile, err := os.Create(header.Name)
					if err != nil {
							log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
					}
					if _, err := io.Copy(outFile, tarReader); err != nil {
							log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
					}
					outFile.Close()

					return header.Name

			default:
					log.Fatalf(
							"ExtractTarGz: uknown type: %s in %s",
							header.Typeflag,
							header.Name)
			}

	}
	return ""
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

		// Packaged in a yaml?
		matched, err := regexp.MatchString(`\.tgz\.base64`, v.Path)

		path := v.Path

		if err == nil && matched {
			fmt.Printf("Plugin %s is packaged in tgz/Base64 format: %s\n", v.Name, v.Path)

			base64File := v.Path

			base64File = base64File[:len(base64File)-len(filepath.Ext(base64File))]

			fmt.Println(base64File)

			err :=  ExtractBase64ToFile(v.Path, base64File)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			} else {
				// fmt.Printf("B64 file: %v\n", b64File)
				file, err := os.Open(base64File)

				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}	

				path = ExtractTarGz(file)
			}

			// binaryPackage, err := ExtractBase64PluginBinary(v.Path)

			// if err == nil {
			// 	fmt.Println(binaryPackage.File)
			// 	r, err := os.Open("./file.tar.gz")
			// 	if err != nil {
			// 			fmt.Println("error")
			// 	}
			// 	ExtractTarGz(r)
			// }
		}

		fmt.Printf("Plugin path: %s\n", path)

		plug, err := plugin.Open(path)
		if err != nil {
			fmt.Printf("Unable to load plugin %s: %v\n", v.Name, err)
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

func ExtractBase64ToFile(b64File string, targetFile string) (error) {

	b, err := os.ReadFile(b64File) // just pass the file name
	if err != nil {
		return err
	}

	b64 := string(b) // convert content to a 'string'

	dec, err := base64.StdEncoding.DecodeString(b64)
    if err != nil {
      return err
    }

    f, err := os.Create(targetFile)
    if err != nil {
			return err
    }
    defer f.Close()

    if _, err := f.Write(dec); err != nil {
			return err
    }
    if err := f.Sync(); err != nil {
			return err
    }

		return nil
}

// func ExtractBase64PluginBinary(file string) (BinaryPackage, error) {

// 	result := BinaryPackage{}

// 	yfile, err := ioutil.ReadFile(file)

// 	if err != nil {

// 		log.Fatal(err)

// 		return result, err
// 	}

// 	err = yaml.Unmarshal(yfile, &result)

// 	if err != nil {
// 		log.Fatal(err)

// 		return result, err
// 	}

// 	return result, nil
// }
