import utils
import json


class ContainerHandler(utils.RequestHandler):

    def __init__(self):
        super(ContainerHandler, self).__init__()

    def get_containers(self, **params):
        """
            Accepted parameters:
            all, limit, since, before, size, filters, exited, status
        """
        url_path = "/containers/json?" + utils.paramstr_from_dict(params)

        response = self.request("GET", url_path, None)
        if self.has_error():
            print(response)
            return None

        return json.loads(response)

    def __searchbyname(self, name, containers):
        for container in containers:
            if name in container.Names:
                return container.Id

        return None

    def __searchbyid(self, containerid, containers):
        for container in containers:
            if fixed_name in container.Names:
                return container.Names

        return None

    def get_containerid(self, name):
        fixed_name = "/" + name
        containers = self.getcontainers()
        return self.__searchbyname(fixed_name, containers)

    def container_exists(self, name=None, containerid=None):
        containers = self.getcontainers()
        if name is not None:
            return (self.__searchbyname("/" + name, containers) is not None)
        elif containerid is not None:
            return (self.__searchbyid(containerid, containers) is not None)

        return None

    def get_containerinfo(self, containerid):
        url = "/containers/" + containerid + "/json"
        response = self.request("GET", url, None)
        if self.has_error():
            print response
            return None

        return json.loads(response)


if __name__ == "__main__":
    handler = ContainerHandler()
    containers = handler.get_containers(all=True)
    if containers:
        for item in containers:
            utils.printjson(json.dumps(item))

    cinfo = handler.get_containerinfo("00731")
    utils.printjson(obj=cinfo)

    """
    cid = getcontainerid("berserk_colden")
    if not cid:
        print "No such container found."

    print containerexists("berserk_colden")
    """
