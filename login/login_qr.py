import uiautomation
import time

def get_qrcode():
    # 点击多开
    duokai = uiautomation.WindowControl(searchDepth=1, SubName=u'电脑端微信多开')
    duokai.SetActive()
    time.sleep(1)
    # edit = duokai.EditControl(LocalizedControlType=u'编辑')
    # edit.Click(simulateMove=False)
    # duokai.SendKeys('{Home}'+'{Shift}{End}'+'C:\\WeChat')
    qidong = duokai.ButtonControl(Name=u'启动微信')
    qidong.Click(simulateMove=False)
    # 获取二维码
    wxlogin = uiautomation.PaneControl(searchDepth=2, ClassName='WeChatLoginWndForPC', Name=u'登录')
    # 可能是已登录，所以点击“切换账号”
    time.sleep(1)
    wxlogin.Click(ratioY=0.9, simulateMove=False)
    time.sleep(3)
    qrfile = './qr.png'
    wxlogin.CaptureToImage(qrfile, x=45, y=80, width=190, height=190)
    f = open(qrfile, 'rb')
    b = f.read()
    f.close()
    return b

from http.server import HTTPServer,BaseHTTPRequestHandler 

class MyHttpHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        print('path: ' + self.path)
        if self.path != '/qr':
            self.wfile.write(b'Not Found')
            return
        content = get_qrcode()
        self.send_response(200)
        self.send_header("Content-Type", "image/png")
        self.send_header("Content-Length", str(len(content)))
        self.end_headers()
        self.wfile.write(content)

if __name__ == '__main__':
    httpd = HTTPServer(('',8080),MyHttpHandler)
    httpd.serve_forever()
