# -*- coding: utf-8 -*-

import time
import uiautomation

# action = {
#     'url': 'http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg',
#     'votes': 1,
#     'clicks': [(12, 363, 336), (17, 364, 321), (23, 366, 511), (28, 665, 249)]
# }

# jonty add for dict
nextHandleKey = 0
handleDict = {}
# state of vote, default is False
voteIngFlag = False
# for vote order
nextVoteIndex = 0

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

# create dict
def createDict():
    console = uiautomation.GetConsoleWindow()
    global nextHandleKey
    global handleDict
    index = 0
    handleDict.clear()
    while True:
        window = uiautomation.WindowControl(searchDepth=1,  ClassName='WeChatMainWndForPC', SubName=u'微信')
        if not window.Exists(0):
            log('please open windown first')
            return
        window.ShowWindow(uiautomation.ShowWindow.Maximize)
        #window.SetActive()
        curHandle = window.Handle
        if curHandle in handleDict.values():
            break
        else :
            handleDict[index] = curHandle
            log(str(index) + ' ' + str(curHandle))
            index += 1
        window.SendKeys('{ALT}{ESC}')
    console.SetActive()
    nextHandleKey = index
    log('dice is created: len=' + str(len(handleDict)) + ' nextHandleKey='+str(nextHandleKey))

# refresh dict
def refreshDict():
    console = uiautomation.GetConsoleWindow()
    # new handle
    global nextHandleKey
    log('nextHandleKey = ' + str(nextHandleKey))
    begHandle = 0
    while True:
        window = uiautomation.WindowControl(searchDepth=1,  ClassName='WeChatMainWndForPC', SubName=u'微信')
        window.ShowWindow(uiautomation.ShowWindow.Maximize)
        #window.SetActive()
        curHandle = window.Handle
        if begHandle == 0 :
            begHandle = curHandle
        elif begHandle == curHandle :
            break
        if curHandle not in handleDict.values():
            handleDict[nextHandleKey] = curHandle
            log('add: ' + str(nextHandleKey) + ' ' + str(curHandle))
            nextHandleKey += 1
        window.SendKeys('{ALT}{ESC}')
    console.SetActive()
    # invalid handle
    keyList = handleDict.copy().keys()
    for index in keyList:
        curHandle = handleDict[index]
        window = uiautomation.ControlFromHandle(curHandle)
        if window == None:
            handleDict.pop(index)
            log('pop: ' + str(index) + ' ' + str(curHandle))
    log('dice is refreshed: len=' + str(len(handleDict)) + ' nextHandleKey='+str(nextHandleKey))

def vote(url, count):
    '''投票'''
    global nextVoteIndex
    global voteIngFlag
    global handleDict
    log('vote() begin...')
    voteIngFlag = True
    console = uiautomation.GetConsoleWindow()
    while count > 0:
        if nextVoteIndex >= nextHandleKey:
            nextVoteIndex = 0
        if nextVoteIndex not in handleDict.keys():
            nextVoteIndex += 1
            continue
        window = uiautomation.ControlFromHandle(handleDict[nextVoteIndex])
        if window == None:
            log('Error: window not exist ' + str(nextVoteIndex) + ' ' + str(handleDict[nextVoteIndex]) + ', so pop it')
            log('pop: ' + str(nextVoteIndex) + ' ' + str(handleDict[nextVoteIndex]))
            handleDict.pop(nextVoteIndex)
            nextVoteIndex += 1
        else:
            window.ShowWindow(uiautomation.ShowWindow.Maximize)
            window.SetActive()
            # # 点击搜索
            # uiautomation.Win32API.MouseClick(126, 24)
            # # 输入“文件传输助手”
            # window.SendKeys(u'文件传输助手')
            # # 点击联系人
            # uiautomation.Win32API.MouseClick(147, 88)

            # 直接点击第一个联系人
            uiautomation.Win32API.MouseClick(136, 73)
            log('click')
            # 输入url
            window.SendKeys(40 * '{Ctrl}{Enter}' + 4 * (url + ' ') + '{Enter}', 0)
            # 点击输入框的上面一行文字（要求刚输入的文字就贴在输入框上方），弹出webview或浏览器
            uiautomation.Win32API.MouseClick(1661, 817)
            # 做投票动作
            # dovote3(window)
            # TODO 等待并截图，或者判断是否成功
            count -= 1
            # 窗口放到最后
            window.SendKeys('{ALT}{ESC}')
            log('vote end window: ' + str(handleDict[nextVoteIndex]))
            nextVoteIndex += 1
    console.SetActive()
    voteIngFlag = False
    log('vote() end...')

def voteOld(url, count):
    '''投票'''
    log('vote() begin...')
    console = uiautomation.GetConsoleWindow()

    while count > 0:
        window = uiautomation.WindowControl(searchDepth=1, ClassName='WeChatMainWndForPC', SubName=u'微信')
        log('vote begin window: {0}'.format(window.Handle))
        # window.ShowWindow(uiautomation.ShowWindow.Maximize)
        # window.ShowWindow(uiautomation.ShowWindow.Restore)
        # window.MoveWindow(0,0,850,590)
        window.SetActive(waitTime=0)
        log('setactive')

        # # 点击搜索
        # uiautomation.Win32API.MouseClick(126, 24)
        # # 输入“文件传输助手”
        # window.SendKeys(u'文件传输助手')
        # # 点击联系人
        # uiautomation.Win32API.MouseClick(147, 88)

        # 直接点击第一个联系人
        uiautomation.Win32API.MouseClick(136, 73)
        log('click')
        # 输入url
        window.SendKeys(4 * (url + ' ') + '{Enter}', 0, 0)
        # 点击输入框的上面一行文字（要求刚输入的文字就贴在输入框上方），弹出webview或浏览器
        uiautomation.Win32API.MouseClick(591, 346)

        # 做投票动作
        # dovote3(window)
        # TODO 等待并截图，或者判断是否成功

        count -= 1

        # 窗口放到最后
        window.SendKeys('{ALT}{ESC}')
        log('vote end window: {0}'.format(window.Handle))
    console.SetActive()
    log('vote() end...')


# def dovote1():
#     '''方式1：录制一系列动作，然后回放'''
#     # 打开页面并最大化 TODO 第一版是直接打开浏览器，待优化
#     page = uiautomation.WindowControl(searchDepth=1, ClassName='IEWebViewWnd', SubName=u'微信')
#     page.ShowWindow(uiautomation.ShowWindow.Maximize)
#     page.SetActive()

#     # 滚动到选项
#     v = page.PaneControl(ClassName='Internet Explorer_Server')
#     for (percent, x, y) in action['clicks']:
#         v.SetScrollPercent(0, percent)
#         uiautomation.Win32API.MouseClick(x, y)

# 性能测试，结果在每秒20~30次之间
def bench(count=30):
    log('bench begin...')
    window = uiautomation.WindowControl(searchDepth = 1, ClassName = 'Notepad', SubName = '无标题 - 记事本')
    window.SetActive(waitTime=0)
    window.Maximize(waitTime=0)
    while count > 0:
        window = uiautomation.WindowControl(searchDepth = 1, ClassName = 'Notepad', SubName = '无标题 - 记事本')
        window.SetActive(waitTime=0)
        window.SendKeys('123123123123123{ENTER}', 0,0)
        window.SendKeys('{ALT}{ESC}', 0,0)
        count -= 1
    log('bench end...')



# def dovote2():
#     '''方式2：打开浏览器，识别单选控件直接select：该方案不可行，控件不能click'''
#     # 1、先获取到浏览器窗口，然后最大化
#     c = uiautomation.GetForegroundControl()
#     w = c.Convert() # c.ControlTypeName == 'WindowControl'
#     w.Maximize()
#     # 2、在窗口中获取单选控件，select
#     r = w.RadioButtonControl(Name=u'选项一在此')

#     # Name:   "选项一在此"
#     # ControlType:    UIA_RadioButtonControlTypeId (0xC35D)
#     # LocalizedControlType:   "单选按钮"
#     # 3、获取投票控件，invoke
#     # Name:   "投票"
#     # ControlType:    UIA_ButtonControlTypeId (0xC350)
#     # LocalizedControlType:   "按钮"

# def dovote3(window):
#     '''TODO 注册默认浏览器，不弹框，且后台进行投票'''
#     # while True:
#     #     c = uiautomation.GetForegroundControl()
#     #     w = c.Convert()
#     #     if w.Handle != window.Handle:
#     #         break
#     # w.SendKeys('{Ctrl}w', 0,0)
#     pass

def color():
    console = uiautomation.GetConsoleWindow()
    index = 0
    while True:
        if voteIngFlag:
            log('voteIngFlag ...')
            time.sleep(6)
            continue
        if index >= nextHandleKey:
            break
        if index not in handleDict.keys():
            index += 1
            continue
        window = uiautomation.ControlFromHandle(handleDict[index])
        if window == None:
            log('Error: window not exist ' + str(index) + ' ' + str(handleDict[index]) + ', so pop it')
            handleDict.pop(index)
            log('pop: ' + str(index) + ' ' + str(curHandle))
            index += 1
        else:
            window.ShowWindow(uiautomation.ShowWindow.Maximize)
            window.SetActive()
            log('train begin: ' + str(handleDict[index]) + ' ' + str(index))
            # 直接点击第1个联系人
            time.sleep(3)
            log('...: ')
            uiautomation.Win32API.MouseClick(136, 73)
            print(uiautomation.Win32API.GetPixelColor(136,73,handleDict[index]))
            # 直接点击第2个联系人
            time.sleep(3)
            log('...: ')
            uiautomation.Win32API.MouseClick(136, 129)
            print(uiautomation.Win32API.GetPixelColor(136,129,handleDict[index]))
            # 收尾窗口放到最后
            index += 1
            window.SendKeys('{ALT}{ESC}')
            log('train end window ' + str(handleDict[index]) + ' \n')
            log('end\n')
            time.sleep(10)
    console.SetActive()
    log('>>>> train() end...')

# train with handle order by dict-key
def train():
    log('>>>> train() begin...')
    console = uiautomation.GetConsoleWindow()
    index = 0
    while True:
        if voteIngFlag:
            time.sleep(600)
            continue
        if index >= nextHandleKey:
            index = 0
        if index not in handleDict.keys():
            index += 1
            continue
        window = uiautomation.ControlFromHandle(handleDict[index])
        if window == None:
            log('Error: window not exist ' + str(index) + ' ' + str(handleDict[index]) + ', so pop it')
            handleDict.pop(index)
            log('pop: ' + str(index) + ' ' + str(curHandle))
            index += 1
        else:
            window.ShowWindow(uiautomation.ShowWindow.Maximize)
            window.SetActive()
            log('train begin: ' + str(handleDict[index]) + ' ' + str(index))
            # 直接点击第1个联系人
            time.sleep(3)
            uiautomation.Win32API.MouseClick(136, 73)
            # window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好想是个A风筝的线条固阳A.{Enter}')
            # 直接点击第2个联系人
            time.sleep(3)
            uiautomation.Win32API.MouseClick(136, 129)
            # window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好想是个B风筝的线条固阳B.{Enter}')
            # 直接点击第3个联系人
            time.sleep(3)
            uiautomation.Win32API.MouseClick(136, 181)
            # window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好想是个C风筝的线条固阳C.{Enter}')
            # 直接点击第4个联系人
            time.sleep(3)
            uiautomation.Win32API.MouseClick(136, 245)
            # window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好想是个D风筝的线条固阳D.{Enter}')
            # 直接点击第5个联系人
            time.sleep(3)
            uiautomation.Win32API.MouseClick(136, 317)
            # window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好想是个E风筝的线条固阳E.{Enter}')
            # 收尾窗口放到最后
            window.SendKeys('{ALT}{ESC}')
            log('train end ' + str(handleDict[index]) + ' \n')
            index += 1
            time.sleep(300) #1776
    console.SetActive()
    log('>>>> train() end...')
 
def trainOld():
    '''养号'''
    log('>>>> train() begin...')
    console = uiautomation.GetConsoleWindow()

    while True:
        window = uiautomation.WindowControl(searchDepth=1,  ClassName='WeChatMainWndForPC', SubName=u'微信')
        window.ShowWindow(uiautomation.ShowWindow.Maximize)
        window.SetActive()
        log('train begin window: {0}'.format(window.Handle))

        # # 点击搜索
        # uiautomation.Win32API.MouseClick(126, 24)
        # # 输入“文件传输助手”
        # window.SendKeys(u'文件传输助手')
        # # 点击联系人
        # uiautomation.Win32API.MouseClick(147, 88)

        # 直接点击第一个联系人
        uiautomation.Win32API.MouseClick(136, 73)
        # 输入url
        window.SendKeys(time.strftime('%Y-%m-%d %H:%M:%S')+' 你好你好！！！{Enter}')

        # 窗口放到最后
        window.SendKeys('{ALT}{ESC}')
        log('train end window: {0}'.format(window.Handle))
        time.sleep(1800)
    console.SetActive()
    log('>>>> train() end...')


def chat(name=u'同事晴晴'):
    '''养号聊天'''
    window = uiautomation.WindowControl(
        searchDepth=1, ClassName='WeChatMainWndForPC', SubName=u'微信')
    window.ShowWindow(uiautomation.ShowWindow.Maximize)
    window.SetActive()
    # 点击搜索
    uiautomation.Win32API.MouseClick(126, 24)
    # 输入对方名字
    window.SendKeys(name)
    # 点击对象名称
    uiautomation.Win32API.MouseClick(147, 88)
    window.SendKeys(u'你好你好！！！' + '{Enter}')
    # 这个窗口排到最后
    window.ShowWindow(uiautomation.ShowWindow.ShowMinNoActive)
    pass


def record(url):
    '''录制'''
    console = uiautomation.GetConsoleWindow()
    window = uiautomation.WindowControl(
        searchDepth=1, ClassName='WeChatMainWndForPC', SubName=u'微信')
    window.ShowWindow(uiautomation.ShowWindow.Maximize)
    window.SetActive()
    # 点击“文件传输助手”（要求置顶）
    uiautomation.Win32API.MouseClick(170, 75)

    window.SendKeys(4 * (url + ' ') + '{Enter}')
    uiautomation.Win32API.MouseClick(1130, 480)

    page = uiautomation.WindowControl(
        searchDepth=1, ClassName='IEWebViewWnd', SubName=u'微信')
    page.ShowWindow(uiautomation.ShowWindow.Maximize)
    page.SetActive()

    # 滚动到选项
    v = page.PaneControl(ClassName='Internet Explorer_Server')
    clicks = []
    while True:
        console.SetActive()
        # input(u'1.请把webview滚动到合适的未知，然后按回车: ')
        raw_input(u'1. scroll webview to proper percent, and press enter: ')
        percent = v.CurrentVerticalScrollPercent()
        print(percent)
        v.SetScrollPercent(0, percent)
        console.SetActive()
        # input(u'2.请把鼠标放在webview中需要点击的位置，然后按回车: ')
        c = raw_input(u'2. put cursor to proper position, and press enter: ')
        (x, y) = uiautomation.Win32API.GetCursorPos()
        print((x, y))
        clicks.append((percent, x, y))
        # c = input(u'3.还需要继续操作，请输入c并按回车：')
        b = raw_input(
            u'3. if want to break, input b and press enter, or press enter for exit: ')
        if b == 'b':
            break
    page.Close()
    print(clicks)
    # TODO wait
    votes = input('4. how many votes you want to do? ')
    action = {  # 被认为是局部变量不会保存到全局变量中
        'url': url,
        'clicks': clicks,
        # TODO wait
        'votes': votes
    }
    print('action: ')
    print(action)
    
createDict()
