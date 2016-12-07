# -*- coding: utf-8 -*-
import os

from uiautomator import Device

d = Device('200ac4ae', adb_server_port=5037)

actions = []

def record():
    while True:
        # input
        action = raw_input('input action: ')
        print 'action: ' + action
        # page down
        if action == 'd' or action == '':
            d.drag(300, 1260, 300, 160,2)
            actions.append('pagedown')
        # page up
        elif action == 'u': # page up
            d.drag(300, 160, 300, 1260,2)
            actions.append('pageup')
        # click
        elif action == 'c': # click
            d.screenshot(u'1.png')
            d.dump(u'1.uix')
            print 'start uiautomatorviewer...'
            # os.system('D:\\Android\\sdk\\tools\\uiautomatorviewer.bat 1.xml 1.png')
            os.system('/Users/chuqq/Library/Android/sdk/tools/uiautomatorviewer')
            try:
                click = raw_input('click x,y=\n')
                actions.append(eval(click))
            except SyntaxError:
                continue
        # end
        elif action == 'e': # end
            break
        else:
            print 'invalid action: ' + action
    print actions

def replay():
    for action in actions:
        if action == 'pagedown':
            d.drag(300, 1260, 300, 160,2)
        elif action == 'pageup':
            d.drag(300, 160, 300, 1260,2)
        else:
            x, y = action
            d.click(x, y)
    print 'replay end'


record()
replay()
