package main

import (
	"fmt"
	"github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
)

func startContainerWithVolume(name string, path string, hostPath string) () {
	containerId := getContainerId(name)
	fmt.Printf("startContainer :Container Id : %v, Volume Path : %s, Host Path : %s", containerId, path, hostPath)
	if containerId == "" {
		fmt.Printf("Invalid Container name %v\n", name)
	}else {

		config := map[string]interface{}{
			"Binds": []string{hostPath + ":" + path},
		}

		body, err := sampleutils.SockRequest("POST", "/containers/" + containerId + "/start", config)
		if err != nil  {
			fmt.Printf("\nerror %v\n", err)
		}else {
			fmt.Printf("\nbody: %#v\n", body)
		}
	}
}
