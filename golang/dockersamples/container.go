package main

import (
    "fmt"
    "encoding/json"
    "errors"
    "github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
)

func getContainers() {
    body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
    var respJSON *sampleutils.ResponseJSON
    if err = json.Unmarshal(body, &respJSON); err != nil {
       fmt.Printf("unable to unmarshal response body: %v", err)
    }
    sampleutils.PrettyPrint(respJSON)
}

func createContainer(name string) {
    config := map[string]interface{}{
       "Image": "busybox",
       "OpenStdin": true,
    }
    body, err := sampleutils.SockRequest("POST", "/containers/create?name="+name, config)
    if err != nil  {
       fmt.Printf("error %v", err)
    }

    var resp *sampleutils.ResponseCreateContainer
    if err = json.Unmarshal(body, &resp); err != nil {
       fmt.Printf("unable to unmarshal response body: %v", err)
    }
    sampleutils.PrettyPrint(resp)
}

func getContainerInfoFor(name string, inspectJSON []sampleutils.ContainerInfo) (sampleutils.ContainerInfo, error) {
    var noOfContainers = len(inspectJSON)
    for i := 0; i < noOfContainers; i++ {
        nameLocal := inspectJSON[i].Names[0]
        name_trimmed := nameLocal[1:len(nameLocal)]
        if name_trimmed == name {
            return inspectJSON[i], nil
        }
    }

    var value sampleutils.ContainerInfo

    return value, errors.New("No container information found for " + name)
}

func containerExists(name string) (b bool) {
    body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)

    var inspectJSON []sampleutils.ContainerInfo
    if err = json.Unmarshal(body, &inspectJSON); err != nil {
    	fmt.Printf("unable to unmarshal response body: %v", err)
    }

    _, err = getContainerInfoFor(name, inspectJSON)

    return err == nil
}

func printContainerId(name string) () {
    fmt.Printf("Container Id : %v\n", getContainerId(name))
}

func getContainerId(name string) (id string) {
    containerId := ""
    body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
    var inspectJSON []sampleutils.ContainerInfo

    if err = json.Unmarshal(body, &inspectJSON); err != nil {
        fmt.Printf("unable to unmarshal response body: %v", err)
    }

    container, err := getContainerInfoFor(name, inspectJSON)

    if err == nil {
        containerId = container.Id
    }

    return containerId
}

func startContainer(name string) () {
        containerId := getContainerId(name)
        fmt.Printf("startContainer :Container Id : %v", containerId)
        if containerId == "" {
            fmt.Printf("Invalid Container name %v\n", name)
        }else {
            body, err := sampleutils.SockRequest("POST", "/containers/" + containerId + "/start", nil)
            if err != nil  {
                fmt.Printf("\nerror %v\n", err)
            }else {
                fmt.Printf("\nbody: %#v\n", body)
            }
        }
}
func listProcesses(name string) {
        containerId := getContainerId(name)
        if containerId == "" {
            fmt.Printf("Invalid Container name %v\n", name)
        }else {
            body, err := sampleutils.SockRequest("POST", "/containers/" + containerId + "/top", nil)
            if err != nil  {
                fmt.Printf("error %v\n", err)
            } else {
              fmt.Printf("body: %#v\n", body)
            }
        }
}

func pauseContainer(name string) {
        containerId := getContainerId(name)
        if containerId == "" {
            fmt.Printf("Invalid Container name %v\n", name)
        }else {
            body, err := sampleutils.SockRequest("POST", "/containers/" + containerId + "/pause", nil)
            if err != nil  {
                fmt.Printf("\nerror %v\n", err)
            } else {
              fmt.Printf("body: %#v\n", body)
            }
        }
}

func unpauseContainer(name string) {
    containerId := getContainerId(name)
    if containerId == "" {
        fmt.Printf("Invalid Container name %v\n", name)
    } else {
        body, err := sampleutils.SockRequest("POST", "/containers/" + containerId + "/unpause", nil)
        if err != nil  {
            fmt.Printf("\nerror %v\n", err)
        } else {
            fmt.Printf("body: %#v\n", body)
        }
    }
}
func getContainerState(name string) (containerState string) {
    containerId := ""
    containerId = getContainerId(name)
    ///containers/4fa6e0f0c678/json
    if containerId != "" {
        body, err := sampleutils.SockRequest("GET", "/containers/" + containerId+ "/json", nil)
        var inspectJSON struct {
            Id string
            State string
        }
        if err = json.Unmarshal(body, &inspectJSON); err != nil {
            fmt.Printf("unable to unmarshal response body: %v", err)
        }
        containerState := string(inspectJSON.State)

        return containerState
    } else {
        fmt.Printf("Container doesn't exist")
        return ""
    }
}
