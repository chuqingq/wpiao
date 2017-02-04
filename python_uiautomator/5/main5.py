# -*- coding: utf-8 -*-

"""
操作步骤：
* 需要先打开三星手机上的insecureADB
* 修改设备串号和adb端口：bbb5fc231f5c3，5037
* 指定安卓用户：u0_a121
* 指定微信的控件ID：你的手机号码，填写密码
* 修改要操作的账号：17092560668
* 修改需要的动作：login
"""
import time
import os
import subprocess
import importlib

from uiautomator import Device

from data import accounts

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

curDeviceId = '071efe2c00e37e37';
d = Device(curDeviceId, adb_server_port=5037)

# 投票地址
url = 'http://mp.weixin.qq.com/s/Nw_Jiahy6tTswuOtPv0-Zg'

# 投票动作
actions = ['pagedown', 'pagedown', 'pagedown', 'pagedown',(300, 1230), 'pagedown', 'pagedown', 'pagedown', (300, 465)]

## ---- 内部函数

# 启动uiautomatorviewer
def startViewer():
    subprocess.Popen('/Users/chuqq/Library/Android/sdk/tools/uiautomatorviewer')

# 导出当前界面层级
def dump():
    d.dump('1.uix')
    d.screenshot('1.png')


# 启动微信
def startApp():
    log('start app...')
    os.system('adb -s '+curDeviceId+' shell am start -n com.tencent.mm/com.tencent.mm.ui.LauncherUI')

def clearData():
    os.system('adb -s '+curDeviceId+' shell am force-stop com.tencent.mm HERE')
    os.system('adb -s '+curDeviceId+' shell pm clear com.tencent.mm')

def restore(account):
    log('restore '+account)
    os.system('adb -s '+curDeviceId+' shell am force-stop com.tencent.mm HERE')
    os.system('adb -s '+curDeviceId+' shell pm clear com.tencent.mm')
    os.system('adb -s '+curDeviceId+' shell rm -rf /mnt/sdcard/tencent')
    os.system('adb -s '+curDeviceId+' shell rm -rf /mnt/sdcard/Android/data/com.tencent.mm')
    # 恢复-覆盖
    os.system('adb -s '+curDeviceId+' shell "su -c \\\"cd /data/data/; tar -xf com.tencent.mm.'+account+'.tar\\\""')
    # os.system('adb -s '+curDeviceId+' shell "su -c \\\"cd /data/data/; pwd\\\""')
    os.system('adb -s '+curDeviceId+' shell "cd /mnt/sdcard/; tar -xf tencent.'+account+'.tar"')

def backup(account):
    log('backup '+account)
    os.system('adb -s '+curDeviceId+' shell am force-stop com.tencent.mm HERE')
    os.system('adb -s '+curDeviceId+' shell "su -c \\\"cd /data/data/; tar -cf com.tencent.mm.'+account+'.tar com.tencent.mm\\\""')
    os.system('adb -s '+curDeviceId+' shell "cd /mnt/sdcard/; tar -cf tencent.'+account+'.tar tencent"')
    os.system('adb -s '+curDeviceId+' shell pm clear com.tencent.mm')
    os.system('adb -s '+curDeviceId+' shell rm -rf /mnt/sdcard/tencent')

## 养号相关的函数
def chat_with(pname):
    if d(text=pname).exists:
        d(text=pname).click.wait()
        d(className='android.widget.EditText').wait.exists()
        d(className='android.widget.EditText').click.wait()
        # d(resourceId='com.tencent.mm:id/chatting_content_et').click()
        os.system('adb -s '+curDeviceId+' shell input text "ItIsSunnyAndBeautifulDay'+str(time.time())+'"')
        d(text=u'发送').wait.exists()
        d(text=u'发送').click.wait()
        #d(resourceId='com.tencent.mm:id/chatting_send_btn').click()
        time.sleep(1)
        #d.press.back()
        d.press.back()
    else:
        log('chat_with not find text ' + pname)

def intest():
    # 等待出现同学滔滔，10秒超时
    # d(text=u'同学滔滔').wait.exists(timeout=10000)
    # Tencent news
    # d(text=u'腾讯新闻').wait.exists()
    # d(text=u'腾讯新闻').click.wait()
    # time.sleep(3)
    # d.press.back()
    # if d(text=u'腾讯新闻').exists:
    #     d(text=u'腾讯新闻').click.wait()
    #     time.sleep(3)
    #     d.press.back()
    # else:
    #     log('not find text txxw')

    # Begin Chat
    # d(text=u'同学滔滔').wait.exists()
    # chat_with('同学滔滔')
    # time.sleep(3)
    # if d(text=u'同学滔滔').exists:
    #     chat_with('同学滔滔')
    #     time.sleep(3)
    # else:
    #     log('not find text taotao')

    d(text=u'同事晴晴').wait.exists()
    chat_with('同事晴晴')
    time.sleep(1)
    # if d(text=u'同事晴晴').exists:
    #     chat_with('同事晴晴')
    #     time.sleep(3)
    # else:
    #     log('not find text qingqing')

    #sleeping for backup
    log('Begin sleeping for backup...')
    time.sleep(3600/2)

# 对一个账号进行养号：
def yanghao(account):
    try:
        restore(account)
        startApp()
        time.sleep(3)
        intest()
    except Exception as e:
        print('exception...')
        raise e
    finally:
        backup(account)
        time.sleep(1)
        log('***** done with ' + account + '\n')

def autoyh():
    index = 0
    while True:
        index = index%len(accounts)
        account = accounts[index]
        log('yanghao '+str(index)+' '+account)
        yanghao(account)
        index = index + 1

def autoyanghao(bIndex):
    index = bIndex
    while True:
        index = index%len(accounts)
        account = accounts[index]
        log('yanghao '+str(index)+' '+account)
        yanghao(account)
        index = index + 1

def imanual(index):
    if index < len(accounts):
        account = accounts[index]
        restore(account)
        startApp()
        input("\nPress Enter for backup ...")
        backup(account)
        log('***** done *****\n')
    else:
        log(str(index) + ' beyond ' + str(len(accounts)))

