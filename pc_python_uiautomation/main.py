# -*- coding: utf-8 -*-

import time
# import logging
import uiautomation

# logging.basicConfig(filename='main.log',level=logging.DEBUG)

action = {
    'url': 'http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg',
    'votes': 25,
    'clicks': [(12, 363, 336), (17, 364, 321), (23, 366, 511), (28, 665, 249)]
}


def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)


def voteall():
    log('voteall() begin...')
    # 确认要投票的所有窗口
    windows = []
    foundIndex = 1
    while True:
        log('window: {0}'.format(foundIndex))
        # 每次取到下一个微信窗口:foundIndex=1
        window = uiautomation.WindowControl(
            searchDepth=1, foundIndex=foundIndex, ClassName='WeChatMainWndForPC', SubName=u'微信')
        if action['votes'] < 0 or not window.Exists():
            break
        window.ShowWindow(uiautomation.ShowWindow.Maximize)
        foundIndex += 1
        action['votes'] -= 1
        windows.append(window)
    log('windows[{0}]: {1}'.format(len(windows), windows))

    for window in windows:
        if not window.Exists():
            break
        log('begin window: {0}'.format(window.Handle))
        window.SetActive()

        # 点击“文件传输助手”（要求置顶）
        uiautomation.Win32API.MouseClick(170, 75)
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

        # 关闭web页面
        page.Close()
        log('end window: {0}'.format(window.Handle))

    log('voteall() end...')

# 录制


def record(url):
    console = uiautomation.GetConsoleWindow()
    window = uiautomation.WindowControl(
        searchDepth=1, foundIndex=1, ClassName='WeChatMainWndForPC', SubName=u'微信')
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
