package main

import (
	"fmt"
	"github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
	"os"
)

type ContainerInfo struct {
	Names []string
	Id    string
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		name := arg
		fmt.Printf("Container %v Id : %v\n", name, getContainerId(name))
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func getContainerId(name string) (id string) {
	return sampleutils.GetContainerId(name)
}
