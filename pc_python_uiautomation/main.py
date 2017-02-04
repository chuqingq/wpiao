# -*- coding: utf-8 -*-

import time
import uiautomation

url = 'http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg'

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

log('begin...')

foundIndex = 0

while True:
	foundIndex += 1
	# 每次取到下一个微信窗口:foundIndex=1
	window = uiautomation.WindowControl(searchDepth = 1, foundIndex=foundIndex, ClassName = 'WeChatMainWndForPC', SubName = u'微信')

	if not window.Exists():
		break
	print('begin window...', foundIndex, window.Handle)
	window.ShowWindow(uiautomation.ShowWindow.Maximize)
	window.SetActive()

	# 点击“文件传输助手”（要求置顶）
	uiautomation.Win32API.MouseClick(170, 75)
	window.SendKeys(4*(url+' ')+'{Enter}')

	# 点击输入框的上面一行文字（要求刚输入的文字就贴在输入框上方）
	uiautomation.Win32API.MouseClick(1130, 480)

	# 打开页面并最大化
	page = uiautomation.WindowControl(searchDepth = 1, ClassName = 'IEWebViewWnd', SubName = u'微信')
	page.ShowWindow(uiautomation.ShowWindow.Maximize)
	page.SetActive()

	# 滚动到选项
	v = page.PaneControl(ClassName='Internet Explorer_Server')
	v.SetScrollPercent(0,13)
	# 点击选项
	uiautomation.Win32API.MouseClick(364,211)
	# 滚动到“投票”
	v.SetScrollPercent(0,28)
	# 点击投票
	uiautomation.Win32API.MouseClick(672,245)

	# 关闭web页面
	page.Close()
	print('end window...', foundIndex, window.Handle)

log('end...')
