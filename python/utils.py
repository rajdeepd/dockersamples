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


class RequestHandler(object):
    def __init__(self):
        self.err_no = 0

    def request(self, http_method, path, data):
        self.err_no = 0

        try:
            conn = UnixHTTPConnection("/var/run/docker.sock")
            conn.connect()
            conn.request(http_method, path, body=json.dumps(data), headers={"Content-type": "application/json"})
            response = conn.getresponse()

            if response.status != 200:
                self.err_no = response.status

            return response.read()
        except Exception as e:
            self.err_no = -1
            self.msg = str(e)
            return str(e)
        finally:
            conn.close()

    def has_error(self):
        return (self.err_no != 0 and self.err_no != 200 and self.err_no != 204)


def printjson(jsonstr=None, obj=None):
    if obj is None:
        obj = json.loads(jsonstr)
    print(json.dumps(obj, indent=2, sort_keys=True))


def paramstr_from_dict(params):
    params_str = ""
    for key in params.keys():
        params_str += ("&" + key + "=" + str(params[key]))

    return params_str
