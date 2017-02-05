# -*- coding: utf-8 -*-

import time
import uiautomation

action = {
    'url': 'http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg',
    'votes': 3,
    'clicks': [(12, 363, 336), (17, 364, 321), (23, 366, 511), (28, 665, 249)]
}


def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)


def voteall():
    '''投票'''
    log('voteall() begin...')
    console = uiautomation.GetConsoleWindow()

    while action['votes'] > 0:
        window = uiautomation.WindowControl(
            searchDepth=1, ClassName='WeChatMainWndForPC', SubName=u'微信')
        if not window.Exists(0):
            log('ERROR: there is no weixin window')
            break
        log('begin window: {0}'.format(window.Handle))
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
        # 输入url
        window.SendKeys(4 * (action['url'] + ' ') + '{Enter}')

        # 点击输入框的上面一行文字（要求刚输入的文字就贴在输入框上方）
        uiautomation.Win32API.MouseClick(1130, 480)

        # 打开页面并最大化
        page = uiautomation.WindowControl(
            searchDepth=1, ClassName='IEWebViewWnd', SubName=u'微信')
        page.ShowWindow(uiautomation.ShowWindow.Maximize)
        page.SetActive()

        # 滚动到选项
        v = page.PaneControl(ClassName='Internet Explorer_Server')
        for (percent, x, y) in action['clicks']:
            v.SetScrollPercent(0, percent)
            uiautomation.Win32API.MouseClick(x, y)
        # TODO wait

        action['votes'] -= 1
        # 关闭web页面
        page.Close()
        window.ShowWindow(uiautomation.ShowWindow.ShowMinNoActive)  # 排到最后
        log('end window: {0}'.format(window.Handle))
    console.SetActive()
    log('voteall() end...')


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
