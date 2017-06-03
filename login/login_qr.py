import uiautomation

def get_qrcode():
	# 点击多开
	duokai = uiautomation.WindowControl(searchDepth=1, SubName=u'电脑端微信多开')
	duokai.SetActive()
	edit = duokai.EditControl(LocalizedControlType=u'编辑')
	edit.Click()
	duokai.SendKeys('{Home}'+'{Shift}{End}'+'C:\\WeChat')
	qidong = duokai.ButtonControl(Name=u'启动微信')
	qidong.Click()

	# 获取二维码
    window = uiautomation.WindowControl(searchDepth=1, SubName=u'启动微信') # ClassName='WeChatMainWndForPC', 
    window.SetActive()
    window.CaptureToImage("./1.png")
    f = open('./1.png', 'rb')
    b = f.read()
    f.close()
    return b

from http.server import HTTPServer,BaseHTTPRequestHandler 

class MyHttpHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        content = get_qrcode()
        self.send_response(200)
        self.send_header("Content-Type", "image/png")
        self.send_header("Content-Length", str(len(content)))
        self.end_headers()
        self.wfile.write(content)

if __name__ == '__main__':
    httpd = HTTPServer(('',8080),MyHttpHandler)
    httpd.serve_forever()
