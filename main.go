package main

import (
	"flag"
	"fmt"
	"github.com/carefreex-io/installx/common"
	"log"
	"path"
)

const help = `Carefree Install Tool

	You can use this program to install carefree api framework or rpc framework.

The parameter description:
-t string
	the framework type, option values: api、rpc.
-n string
	the application name, this name will be used as the directory name.
-d string
	the install directory, default to current directory.
`

func main() {
	frameworkType := flag.String("t", "api", "the framework type, option values: api、rpc.")
	name := flag.String("n", "", "the application name, this name will be used as the directory name.")
	dir := flag.String("d", "./", "the install directory, default to current directory.")

	flag.Usage = func() {
		fmt.Print(help)
	}
	flag.Parse()

	if *name == "" {
		log.Fatalf("%v -n is required", common.ErrorStr)
	}
	options := common.Options{
		Name: *name,
		Path: path.Join(*dir, *name),
	}

	switch *frameworkType {
	case "api":
		Install(common.CarefreeInfo, options)
	case "rpc":
		Install(common.CarefreeXInfo, options)
	default:
		log.Fatalf("%v unsupported framework type's value: %v", common.ErrorStr, *frameworkType)
	}

	log.Fatalf("%v %v installed", common.SuccessStr, options.Path)
}
