import utils
import json

class Container:
	"""
		self.Command	(str)
		self.Created	(int)
		self.Id			(str)
		self.Image		(str)
		self.Names		(list)
		self.Ports		(list)
		self.Status		(str)
	"""

	def __init__(self, **entries):
		self.__dict__.update(entries)

	def __str__(self):
		s = self.Names[0].strip("/") + "\n"
		s += self.Id[:10] + "\t" + self.Image + "\n"
		s += (self.Command + "\t" + self.Status)
		return s

def getcontainers():
	handler = utils.RequestHandler()
	response = handler.request("GET", "/containers/json?all=1", None)
	if handler.haserror():
		print(response)
		return None

	containers_list = json.loads(response)
	return [Container(**container_dict) for container_dict in containers_list]

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

def containerexists(name = None, containerid = None):
	containers = getcontainers()
	if name != None:
		return (__searchbyname("/" + name, containers) != None)
	elif containerid != None:
		return (__searchbyid(containerid, containers) != None)

	return None

if __name__ == "__main__":
	#jsonresponse = getcontainers()
	#utils.printjson(jsonresponse)
	#for item in jsonresponse:
	#	print item, "\n"
	cid = getcontainerid("berserk_colden")
	if not cid:
		print "No such container found."

	print containerexists("berserk_colden")
