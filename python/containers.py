import utils
import json

def getcontainers():
	handler = utils.RequestHandler()
	response = handler.request("GET", "/containers/json?all=1", None)
	if handler.haserror():
		print(response)
		return None

	return response

if __name__ == "__main__":
	jsonresponse = getcontainers()
	utils.printjson(jsonresponse)