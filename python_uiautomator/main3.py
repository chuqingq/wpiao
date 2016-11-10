# -*- coding: utf-8 -*-  
import time
import os
from uiautomator import Device

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

log('connect to device...')
# d = Device('0710ad7b00f456bb', adb_server_host='127.0.0.1', adb_server_port=55037)
d = Device('071efe2c00e37e37', adb_server_host='127.0.0.1', adb_server_port=5037)

accounts = {
    '17092560668': '580608.Chu4',
}

# 确保所有的账号都登录了。只提前做一次
for (n, p) in accounts.items():
    login(n, p)

def login(account, password):
    os.system('adb shell rm -rf /data/data/com.tencent.mm')
    os.system('adb shell mkdir -p /data/data/com.tencent.mm')
    os.system('adb shell ln -s /data/app-lib/com.tencent.mm-1 /data/data/com.tencent.mm/lib')
    os.system('adb shell chown u0_a126:u0_a126 /data/data/com.tencent.mm')
    os.system('adb shell am start -n com.tencent.mm/com.tencent.mm.ui.LauncherUI')
    # 如果已有账号，则点击“更多”到输入账号页面；否则，点击登录，才能输入账号
    while not d(text=u'登录').exists:
        time.sleep(0.1)
        # log('wait for "登录" exists...')
    d(text=u'登录').click.wait()
    # 在输入账号页面登录
    log('login "' + account + '"...')
    d(text=u'你的手机号码').set_text(account) # 输入账号 TODO 建议不要用text区分，可能已填入内容
    d(resourceId='com.tencent.mm:id/fo').set_text(password) # 收入密码
    d(text=u'登录').click.wait() # 登录
    os.system('adb shell am force-stop com.tencent.mm HERE')
    os.system('adb shell mv /data/data/com.tencent.mm /data/data/com.tencent.mm.'+account)


for (n, p) in accounts.items():
    vote(n, p)


def vote(account, password):
    # 移动号码的目录到/data/data/com.tencent.mm
    log('移动' + account + '的目录到/data/data/com.tencent.mm')
    os.system('mv /data/data/com.tencent.mm.'+account+' /data/data/com.tencent.mm')
    # 用adb启动微信
    log('am start app...')
    os.system('adb shell am start -n com.tencent.mm/com.tencent.mm.ui.LauncherUI')
    # 等待群聊出现
    while not d(text=u'群聊').exists:
        log('wait for "群聊" exists...')
        if d(text=u'否').exists:
            d(text=u'否').click.wait() # 不使用通讯录（不是每次都有）
        time.sleep(0.1)
    # 进入群组
    d(text=u'群聊').click()
    log('enter "群聊"...')
    # 关注并进入公众号
    weixinid = u'六安楼市'
    log('enter weixinid "' + weixinid + '"...')
    d(text=weixinid).click()
    # 滑到最下面，才会出现“关注”或者“进入公众号”
    d.swipe(300,1000, 300, 300, 2)
    d.swipe(300,1000, 300, 300, 2)
    d.swipe(300,1000, 300, 300, 2)
    if d(text=u'关注').exists:
        d(text=u'关注').click.wait() # 点击“关注”
    elif d(text=u'进入公众号').exists:
        d(text=u'进入公众号').click.wait() # 点击“关注”
    else:
        log('ERROR "关注" or "进入公众号" not exists...')
    # 输入内容
    while not d(resourceId='com.tencent.mm:id/a1o').exists:
        log('wait for "editText" exists...')
        if d(description=u'消息').exists:
            log('"message" exists...')
            d(description=u'消息').click.wait()
        if d(description=u'切换到键盘').exists:
            log('"switch to keyboard" exists...')
            d(description=u'切换到键盘').click.wait()
    d(resourceId='com.tencent.mm:id/a1o').set_text(u'1018') # 输入投票内容
    # 等待发送出现，点击发送
    while not d(text=u'发送').exists:
        log('wait for "发送" exists...')
        time.sleep(0.1)
    d(text=u'发送').click.wait() # 点击“发送”
    # 截图记录结果
    log('screenshot...')
    d.screenshot(account + '_' + weixinid + '.png')
    # 停止微信并保存账号状态
    os.system('adb shell am force-stop com.tencent.mm HERE')
    os.system('mv /data/data/com.tencent.mm /data/data/com.tencent.mm.'+account)

# 取消关注（暂时不做，可能不是每个都能取消）
# 退出登录（放在最前面做）
