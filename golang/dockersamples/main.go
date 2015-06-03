package main

import (
	"fmt"
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
				fmt.Println("Please specify container name after 7")
			} else {
				name := os.Args[2]
				containerState := getContainerState(name)
				fmt.Println("Returned State: %#v \n", containerState)
			}
		} else if i == 8 {
			fmt.Printf("Unpause a Container\n")
			if len(os.Args) != 3 {
				fmt.Println("Please specify container name after 8")
			} else {
				name := os.Args[2]
				unpauseContainer(name)
				//fmt.Println("Returned State: %#v \n", containerState)
			}
		} else if i == 9 {
			fmt.Printf("Start a Container with volume\n")
			if len(os.Args) != 5 {
				fmt.Println("Please specify container name, path of the volume in container and host path for the volume after 9")
			} else {
				name := os.Args[2]
				path := os.Args[3]
				hostPath := os.Args[4]
				startContainerWithVolume(name, path, hostPath)
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
				"8 [name] : Unpause Container\n" +
				"9 [name] [path] [host path] : Start Container with volume\n"
		fmt.Printf("%v", help)
	}
}
