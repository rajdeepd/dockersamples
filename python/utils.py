import httplib
import socket
import json

class UnixHTTPConnection(httplib.HTTPConnection):
	"""
	HTTPConnection object which connects to a unix socket.
	"""

	def __init__(self, sock_path):
		httplib.HTTPConnection.__init__(self, "localhost")
		self.sock_path = sock_path

	def connect(self):
		self.sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
		self.sock.connect(self.sock_path)

class RequestHandler:
	def __init__(self):
		self.errno = 0

	def request(self, http_method, path, data):
		self.errno = 0

		try:
			conn = UnixHTTPConnection("/var/run/docker.sock")
			conn.connect()
			conn.request(http_method, path, body = json.dumps(data))
			response = conn.getresponse()

			if response.status != 200:
				self.errno = response.status

			return response.read()
		except Exception as e:
			self.errno = -1
			self.msg = str(e)
			return str(e)
		finally:
			conn.close()

	def haserror(self):
		return (self.errno != 0 and self.errno != 200)

def printjson(jsonstr):
	obj = json.loads(jsonstr)
	print(json.dumps(obj, indent = 2, sort_keys = True))