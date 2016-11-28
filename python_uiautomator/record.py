# -*- coding:UTF-8 -*-
import os

from uiautomator import Device

d = Device('0709000500e28864', adb_server_port=55037)

# 参数
actions = []

def record():
    while True:
        # input
        input = raw_input('action:')
        # page down
        if input == 'd' or input == '':
            d.drag(300, 1000, 300, 300,2)
            actions.append('pagedown')
        # page up
        elif input == 'u': # page up
            d.drag(300, 300, 300, 1000,2)
            actions.append('pageup')
        # click
        elif input == 'c': # click
            d.screenshot('1.png')
            d.dump("1.uix")
            print 'start uiautomatorviewer...'
            os.system('D:\\Android\\sdk\\tools\\uiautomatorviewer.bat 1.xml 1.png')
            try:
                click = raw_input('click x,y=')
                actions.append(eval(click))
            except SyntaxError:
                continue
        # end
        elif input == 'e': # end
            break

    print actions

def replay():
    for action in actions:
        if action == 'pagedown':
            d.drag(300, 1000, 300, 300,2)
        elif action == 'pageup':
            d.drag(300, 300, 300, 1000,2)
        else:
            x, y = action
            d.click(x, y)
    print 'replay end'


record()
replay()
