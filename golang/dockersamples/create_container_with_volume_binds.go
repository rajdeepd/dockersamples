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
		CreateContainerWithVolumeBinds(name)
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func CreateContainerWithVolumeBinds(name string) {
	config := map[string]interface{}{
		"Image":     "busybox",
		"Volumes":   map[string]struct{}{"/tmp": {}},
		"OpenStdin": true,
	}

	status, _, err := sampleutils.SockRequest("POST", "/containers/create?name="+name, config)
	fmt.Printf("status: %v\n", status)
	if err != nil {
		fmt.Printf("Error while creating the Container: %v\n", err)
		return
	}

	bindPath := sampleutils.RandomUnixTmpDirPath("test")
	fmt.Printf("BindPath: %v\n", bindPath)

	config = map[string]interface{}{
		"Binds": []string{bindPath + ":/tmp"},
	}
	containerId := sampleutils.GetContainerId(name)
	status, _, err = sampleutils.SockRequest("POST", "/containers/"+containerId+"/start", config)
	fmt.Printf("Status of the call: %v\n", status)
}
