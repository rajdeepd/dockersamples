import utils
import json


def getcontainers(**params):
    """
        Accepted parameters:
        all, limit, since, before, size, filters, exited, status
    """
    handler = utils.RequestHandler()
    url_path = "/containers/json?" + utils.paramstr_from_dict(params)

    response = handler.request("GET", url_path, None)
    if handler.haserror():
        print(response)
        return None

    return json.loads(response)


def __searchbyname(name, containers):
    for container in containers:
        if name in container.Names:
            return container.Id

    return None


def __searchbyid(containerid, containers):
    for container in containers:
        if fixed_name in container.Names:
            return container.Names

    return None


def getcontainerid(name):
    fixed_name = "/" + name
    containers = getcontainers()
    return __searchbyname(fixed_name, containers)


def containerexists(name=None, containerid=None):
    containers = getcontainers()
    if name is not None:
        return (__searchbyname("/" + name, containers) is not None)
    elif containerid is not None:
        return (__searchbyid(containerid, containers) is not None)

    return None


def getcontainerinfo(containerid):
    url = "/containers/" + containerid + "/json"
    handler = utils.RequestHandler()
    response = handler.request("GET", url, None)
    if handler.haserror():
        print response
        return None

    return json.loads(response)


if __name__ == "__main__":
    containers = getcontainers(all=True)
    if containers:
        for item in containers:
            utils.printjson(json.dumps(item))

    cinfo = getcontainerinfo("00731")
    utils.printjson(obj=cinfo)

    """
    cid = getcontainerid("berserk_colden")
    if not cid:
        print "No such container found."

    print containerexists("berserk_colden")
    """
