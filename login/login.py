
import os
import re

# 用zbar命令从二维码文件file中获取文本内容
def zbar_scan_qrcode(file):
	# http://blog.csdn.net/bh20077/article/details/7019060
    zbar_path = '"C:\\Program Files (x86)\\ZBar\\bin\\zbarimg.exe"'
    f = os.popen(zbar_path +' -D ' + file)
    output = f.read()
    # output = 'QR-Code:http://weixin.qq.com/x/AaLjMsCHlOuuN3HIXH61\n'
    return re.match(r'QR-Code:(.*)\n', output).group(1)

import uiautomation
import time

# 用微信多开启动一个新的微信PC，根据二维码生成微信登陆链接
def get_qrcode():
    # 点击多开
    duokai = uiautomation.WindowControl(searchDepth=1, SubName=u'电脑端微信多开')
    duokai.SetActive()
    time.sleep(1)
    edit = duokai.EditControl(LocalizedControlType=u'编辑')
    edit.Click(simulateMove=False)
    duokai.SendKeys('{Home}'+'{Shift}{End}'+'C:\\WeChat')
    qidong = duokai.ButtonControl(Name=u'启动微信')
    qidong.Click(simulateMove=False)
    # 查找新的微信PC TODO
    wxlogin = uiautomation.PaneControl(searchDepth=2, ClassName='WeChatLoginWndForPC', Name=u'登录')
    # 可能是已登录，所以点击“切换账号”
    time.sleep(1)
    wxlogin.Click(ratioY=0.9, simulateMove=False)
    # window = uiautomation.WindowControl(searchDepth=1, SubName=u'启动微信') # ClassName='WeChatMainWndForPC', 
    # window.SetActive()
    time.sleep(3)
    qrfile = './qr.png'
    wxlogin.CaptureToImage(qrfile)
    # f = open('./1.png', 'rb')
    # b = f.read()
    # f.close()
    return zbar_scan_qrcode(qrfile)

from http.server import HTTPServer,BaseHTTPRequestHandler 

class MyHttpHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        content = get_qrcode()
        # self.send_response(200)
        # self.send_header("Content-Type", "image/png")
        # self.send_header("Content-Length", str(len(content)))
        # self.end_headers()
        self.wfile.write(content.encode(encoding="utf-8"))

if __name__ == '__main__':
    httpd = HTTPServer(('',8080),MyHttpHandler)
    httpd.serve_forever()
