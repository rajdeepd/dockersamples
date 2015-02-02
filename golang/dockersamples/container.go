package main

import (
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
	"os"
	"strconv"
)

func main() {

    if len(os.Args) >= 2 {
        arg := os.Args[1]
        i, err := strconv.Atoi(arg)
        if err != nil {
               // handle error
            fmt.Println(err)
            os.Exit(2)
        }

        if i == 0 {
            fmt.Printf("Execute getContainers\n")
            getContainers()
        } else if i == 1 {
            fmt.Printf("Execute containerExists\n")
            if len(os.Args) != 3 {
               fmt.Println("Please specify container name after 1")
            } else {
               name := os.Args[2]
               value := containerExists(name)
               fmt.Printf("Container %v exists : %v\n", name, value)
           }

        } else if i == 2 {
            fmt.Printf("Execute createContainer\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 2")
            } else {
                name := os.Args[2]
                value := containerExists(name)

                if value  {
                    fmt.Printf("Container %v exists : %v\n", name, value)
                } else {
                    createContainer(name)
                }
            }
        }else if i == 3 {
            fmt.Printf("Print Container Id\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 3")
            } else {
                name := os.Args[2]
                printContainerId(name)
            }
        }else if i == 4 {
            fmt.Printf("Start Container Id\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 4")
            } else {
                name := os.Args[2]
                startContainer(name)
            }
        } else if i == 5 {
            fmt.Printf("Show Container processes\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 5")
            } else {
                name := os.Args[2]
                listProcesses(name)
            }
        } else if i == 6 {
            fmt.Printf("Pause the Container\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 6")
            } else {
                name := os.Args[2]
                pauseContainer(name)
            }
        } else if i == 7 {
            fmt.Printf("Get Container State\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 6")
            } else {
                name := os.Args[2]
                containerState := getContainerState(name)
                fmt.Println("Returned State: %#v \n", containerState)
            }
        } else if i == 8 {
            fmt.Printf("Unpause a Container\n")
            if len(os.Args) != 3 {
                fmt.Println("Please specify container name after 6")
            } else {
                name := os.Args[2]
                unpauseContainer(name)
                //fmt.Println("Returned State: %#v \n", containerState)
            }
        }else {
            fmt.Printf("Please specify valid test no as a command line argument\n")
        }
    } else {
         fmt.Printf("Please specify valid test no as a command line argument\n")
         help := "0        : getContainers\n" +
                 "1 [name] : containerExists\n" +
                 "2 [name] : createContainer\n" +
                 "3 [name] : Print Container Id\n" +
                 "4 [name] : StartContainer\n" +
                 "5 [name] : List Processes\n" +
                 "6 [name] : Pause Container\n" +
                 "7 [name] : Get Container State\n" +
                 "8 [name] : Unpause Container\n"
         fmt.Printf("%v", help)
    }
}

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
	}
	body, err := sampleutils.SockRequest("POST", "/containers/create?name="+name, config)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	var resp *sampleutils.ResponseCreateContainer
	if err = json.Unmarshal(body, &resp); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}
	sampleutils.PrettyPrint(resp)
}

func containerExists(name string) (b bool) {
	body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
	var inspectJSON []struct {
		Names []string
	}
	if err = json.Unmarshal(body, &inspectJSON); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}
	containerExists := false
	var noOfContainers = len(inspectJSON)
	for i := 0; i < noOfContainers; i++ {
		nameLocal := inspectJSON[i].Names[0]
		name_trimmed := nameLocal[1:len(nameLocal)]
		if name_trimmed == name {
			containerExists := true
			return containerExists
		}
	}
	return containerExists
}

func printContainerId(name string) {
	body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
	var inspectJSON []struct {
		Names []string
		Id    string
	}
	if err = json.Unmarshal(body, &inspectJSON); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}
	containerExists := false
	var noOfContainers = len(inspectJSON)
	for i := 0; i < noOfContainers; i++ {
		nameLocal := inspectJSON[i].Names[0]
		name_trimmed := nameLocal[1:len(nameLocal)]
		if name_trimmed == name {
			containerExists = true
			containerId := inspectJSON[i].Id
			fmt.Printf("Container Id : %v", containerId)
		}
	}
	if containerExists == false {
		fmt.Printf("Container %v doesn't exist\n", name)
	}
}

func getContainerId(name string) (id string) {

        containerId := ""
        body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
        var inspectJSON []struct {
           Names []string
           Id string
        }
         if err = json.Unmarshal(body, &inspectJSON); err != nil {
        	fmt.Printf("unable to unmarshal response body: %v", err)
        }
        //containerExists := false
        var noOfContainers = len(inspectJSON)
        for i := 0; i < noOfContainers; i++ {
           nameLocal := inspectJSON[i].Names[0]
           name_trimmed := nameLocal[1:len(nameLocal)]
           fmt.Printf("name %v\n", name)
           fmt.Printf("name_trimmed %v\n", name_trimmed)
           if name_trimmed == name {
               containerId := inspectJSON[i].Id
               fmt.Printf("Container Id : %v", containerId)
               return containerId
           }
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
	} else {
		body, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/top", nil)
		if err != nil {
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
	} else {
		body, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/pause", nil)
		if err != nil {
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
