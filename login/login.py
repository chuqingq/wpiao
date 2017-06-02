# import uiautomation
from http.server import HTTPServer,BaseHTTPRequestHandler 

class MyHttpHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.wfile.write(b'hello world')

if __name__ == '__main__':
    httpd = HTTPServer(('',8080),MyHttpHandler)
    httpd.serve_forever()
