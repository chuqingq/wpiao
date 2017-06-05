# 每次http请求返回不同的二维码，循环使用

import time

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

handles = []
index = 0

import uiautomation

def prepare_qr():
    global handles
    global index
    handles = []
    index = 0
    log('prepare_qr(): ')
    count = 1
    while True:
        wxlogin = uiautomation.PaneControl(searchDepth=1, foundIndex=count, ClassName='WeChatLoginWndForPC', Name=u'登录')
        if not wxlogin.Exists():
            break
        log('foundIndex: ' + str(count))
        handles.append(wxlogin.Handle)
        count += 1
    log('count: ' + str(count-1))
    log('prepare_qr() handles: ' + str(handles))

def get_qrcode(handle):
    # # 点击多开
    # duokai = uiautomation.WindowControl(searchDepth=1, SubName=u'电脑端微信多开')
    # duokai.SetActive()
    # time.sleep(1)
    # # edit = duokai.EditControl(LocalizedControlType=u'编辑')
    # # edit.Click(simulateMove=False)
    # # duokai.SendKeys('{Home}'+'{Shift}{End}'+'C:\\WeChat')
    # qidong = duokai.ButtonControl(Name=u'启动微信')
    # qidong.Click(simulateMove=False)
    # 获取二维码
    # wxlogin = uiautomation.PaneControl(searchDepth=2, ClassName='WeChatLoginWndForPC', Name=u'登录')
    wxlogin = uiautomation.ControlFromHandle(handle)
    time.sleep(0.1)
    ret = uiautomation.Win32API.SetForegroundWindow(handle)
    log('ret: ' + str(ret))
    time.sleep(0.1)
    # 可能是已登录，所以点击“切换账号”
    # time.sleep(1)
    # wxlogin.Click(ratioY=0.9, simulateMove=False)
    # time.sleep(3)
    qrfile = 'qr/qr.png'
    wxlogin.CaptureToImage(qrfile, x=45, y=80, width=190, height=190)
    wxlogin.SendKeys('{ALT}{ESC}')
    f = open(qrfile, 'rb')
    b = f.read()
    f.close()
    return b

# from http.server import HTTPServer,BaseHTTPRequestHandler 

# class MyHttpHandler(BaseHTTPRequestHandler):
#     def do_GET(self):
#         global handles
#         global index
#         if self.path != '/qr':
#             self.wfile.write(b'Not Found')
#             return
#         print('path: ' + self.path + ', index: ' + str(index))
#         content = get_qrcode(handles[index])
#         self.send_response(200)
#         self.send_header("Content-Type", "image/png")
#         self.send_header("Content-Length", str(len(content)))
#         self.end_headers()
#         self.wfile.write(content)
#         # 处理index
#         index += 1
#         if index >= len(handles):
#             prepare_qr()

# if __name__ == '__main__':
#     # 加载目前所有的handle
#     prepare_qr()
#     # 通过http使用时，每次不同的handle
#     httpd = HTTPServer(('',8090), MyHttpHandler)
#     httpd.serve_forever()

import tornado.ioloop
import tornado.web

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        global handles
        global index
        print('index: ' + str(index))
        content = get_qrcode(handles[index])
        self.set_header("Content-Type", "image/png")
        self.set_header("Content-Length", str(len(content)))
        self.write(content)
        # 处理index
        index += 1
        if index >= len(handles):
            prepare_qr()

if __name__ == "__main__":
    prepare_qr()
    app = tornado.web.Application([
        (r"/qr", MainHandler),
    ])
    app.listen(8090)
    tornado.ioloop.IOLoop.current().start()
