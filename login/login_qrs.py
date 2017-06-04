# 一次性开N个二维码，一起提供给对方进行扫码登陆

import time

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

import uiautomation

# 数二维码有多少个
def count_qr():
    count = 1
    log('start: ')
    while True:
        wxlogin = uiautomation.PaneControl(searchDepth=1, foundIndex=count, ClassName='WeChatLoginWndForPC', Name=u'登录')
        if not wxlogin.Exists():
            break
        log('foundIndex: ' + str(count))
        count += 1
    log('stop: ')
    log('count: ' + str(count-1))

# 截二维码
def capture_qr():
    # 先开count个微信PC TODO 目前手工做
    # 然后对所有的登陆二维码进行截图
    log('start: ')
    res = []
    index = 1
    firstHandle = 0
    while True:
        wxlogin = uiautomation.PaneControl(searchDepth=1, ClassName='WeChatLoginWndForPC', Name=u'登录')
        if not wxlogin.Exists():
            log('FATAL wxlogin not exists')
            break
        qrfile = str(index) + '_' + str(wxlogin.Handle) + '.png'
        # 如果handle和第一个handle重复，则退出
        if firstHandle == 0:
            firstHandle = wxlogin.Handle
        elif firstHandle == wxlogin.Handle:
            break
        uiautomation.Win32API.SetForegroundWindow(wxlogin.Handle)
        time.sleep(0.2)
        wxlogin.CaptureToImage('qr/'+qrfile, x=45, y=80, width=190, height=190)
        url = 'http://mp.xxying.com:8090/' + qrfile
        res.append(url)
        log(str(index) + ': ' + url)
        index += 1
        wxlogin.SendKeys('{ALT}{ESC}')
    log('total: ' + str(index-1))
    # 目前对方要求格式是每行一个url的文本，不是jsonarray
    # print(res)
    for i in range(len(res)):
        print(res[i])
    log('stop: ')

if __name__ == '__main__':
    capture_qr()

# 启动本地的二维码图片服务
# python -m http.server 8090
# 执行reverseproxy_qr.sh建立和阿里云的ssh映射
