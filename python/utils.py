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


def sock_request(http_method, path, data):
	try:
		conn = UnixHTTPConnection("/var/run/docker.sock")
		conn.connect()
		conn.request(http_method, path)
		response = conn.getresponse()
		print json.loads(response.read())
	except Exception as e:
		print e
	finally:
		conn.close()

if __name__ == "__main__":
	sock_request("GET", "/containers/json?all=1", None)