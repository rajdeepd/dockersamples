import sys
import containers as c
import utils

__author__ = 'sriram'


def print_help():
    print("Usage:")
    print("      python main.py [help | container] <commands> <options>")
    print("      <commands> - [list | info | start | stop]")


def main():
    if len(sys.argv) < 3:
        print_help()
        return

    if sys.argv[1] == "help":
        print_help()
    elif sys.argv[1] == "container":
        handler = c.ContainerHandler()
        if sys.argv[2] == "list":
            containers = handler.get_containers(all=True)
            utils.printjson(obj=containers)
        elif sys.argv[2] == "info":
            container_info = handler.get_containerinfo(sys.argv[3])
            utils.printjson(obj=container_info)
        elif sys.argv[2] == "start":
            response = handler.start_container(sys.argv[3])
            if response:
                print("Container %s started successfully!" % sys.argv[3])
        elif sys.argv[2] == "stop":
            response = handler.stop_container(sys.argv[3])
            if response:
                print("Container %s stopped successfully!" % sys.argv[3])


if __name__ == "__main__":
    main()
