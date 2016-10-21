# -*- coding: utf-8 -*-  
import time
from uiautomator import Device

def log(str):
    print(time.strftime('%Y-%m-%d %H:%M:%S') + ': ' + str)

log('connect to device...')
d = Device('0710ad7b00f456bb', adb_server_host='127.0.0.1', adb_server_port=55037)
# d = Device('071efe2c00e37e37', adb_server_host='127.0.0.1', adb_server_port=5037)

# 打开微信应用
log('open weixin app...')
d.press.home() # 确保只有一个桌面，点击1下home键，回到主界面
d(text=u'微信').click.wait() # 打开微信

# 退出之前的会话（如果左上角有“返回”，则点击）
log('return back to main activity...')
while d(description=u'返回').exists:
    d(description=u'返回').click.wait()

# 退出前面的微信账号
log('logout last account...')
if d(text=u'我').exists:
    d(text=u'我').click.wait() # 点击右下角的“我”
    d(text=u'设置').click.wait() # 点击“设置”
    d(text=u'退出').click.wait() # 点击“退出”
    d(text=u'退出当前帐号').click.wait() # 退出当前账号
    d(text=u'退出').click.wait() # 确认退出

# # 如果是第一次，需要点击左下角的“登录”。如果已登录别的账号，也没有影响
# if d(text=u'登录').exists:
#     d(text=u'登录').click.wait() # 点击登录

# 如果已有账号，则点击“更多”到输入账号页面；否则，点击登录，才能输入账号
while not d(text=u'更多').exists and not d(text=u'登录').exists:
    time.sleep(0.1)
    log('"more" and "login" not exists...')

if d(text=u'更多').exists:
    log('click "more"...')
    d(text=u'更多').click.wait()
    d(resourceId='com.tencent.mm:id/e9').click.wait() # 点击“切换账号”

if d(text=u'登录').exists:
    log('click "login"...')
    d(text=u'登录').click.wait()

# 在输入账号页面登录
account = '17092560668'
password = '580608.Chu4'
log('login "' + account + '"...')

d(text=u'你的手机号码').set_text(account) # 输入账号 TODO 建议不要用text区分，可能已填入内容
d(resourceId='com.tencent.mm:id/fo').set_text(password) # 收入密码
d(text=u'登录').click.wait() # 登录
while not d(description=u'搜索').exists: # TODO 有可能判断的时候还不存在
    log('"search" not exists...')
    if d(text=u'否').exists:
        d(text=u'否').click.wait() # 不使用通讯录（不是每次都有）
    time.sleep(0.1)

# 关注并进入公众号
weixinid = u'la365dichanjiajuwang'
log('enter weixinid "' + weixinid + '"...')

d(description=u'搜索').click.wait() # 点击右上角的“搜索”
d(resourceId='com.tencent.mm:id/fo').set_text(weixinid) # 不能输入中文u'六安楼市'，只能输入微信公众号id
time.sleep(0.5) # 可能已关注，但是不能立刻搜索到
if d(textContains=(u'微信号: '+weixinid)).exists:
    log('weixinid already exists...')
    d(textContains=(u'微信号: '+weixinid)).click.wait()
else:
    log('weixinid not exists...')
    d(textContains=u'搜一搜').click.wait()
    # 点击第一个搜索结果可能无效，需要等到进入公众号后才能关注
    while not d(textContains=u'功能介绍').exists:
        log('"introduce" not exists...')
        time.sleep(0.1)
        d.click(550, 550) # 点击第一个搜索结果
    # TODO 避免死循环
    while not d(text=u'关注').exists:
        log('"focus" not appears...')
        d.swipe(300,1000, 300, 300, 2)
    d(text=u'关注').click.wait() # 点击“关注”

# 输入内容
while not d(resourceId='com.tencent.mm:id/a1o').exists:
    log('"editText" not exists...')
    if d(description=u'消息').exists:
        log('"message" exists...')
        d(description=u'消息').click.wait()
    if d(description=u'切换到键盘').exists:
        log('"switch to keyboard" exists...')
        d(description=u'切换到键盘').click.wait()
d(resourceId='com.tencent.mm:id/a1o').set_text(u'1018') # 输入投票内容

while not d(text=u'发送').exists:
    log('"send" not exists...')
    time.sleep(0.1)
d(text=u'发送').click.wait() # 点击“发送”

# 截图记录结果
log('screenshot...')
d.screenshot(account + '_' + weixinid + '.png')

# 取消关注（暂时不做，可能不是每个都能取消）
# 退出登录（放在最前面做）
