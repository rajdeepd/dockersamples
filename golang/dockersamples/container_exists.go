package main

import (
	"encoding/json"
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
		fmt.Printf("Container %v exists : %v\n", name, ContainerExists(name))
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func ContainerExists(name string) (b bool) {
	_, body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
	var inspectJSON []sampleutils.ContainerInfo
	if err = json.Unmarshal(body, &inspectJSON); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}
	_, err = sampleutils.GetContainerInfoFor(name, inspectJSON)

	return err == nil
}
