pc版本，使用python的uiautomation驱动。

# 1. 安装python uiautomation

https://github.com/yinkaisheng/Python-UIAutomation-for-Windows

```
pip install uiautomation
```

# 2. 登陆所有账号到PC版

* 登陆所有账号到PC版
* 清理退出的账号并重新登陆
* 设置为：不自动升级微信、保留聊天记录等


# 3. 运行脚本

```
python main.py
```

# TODO

* ~~批量手机投票（已完成）~~
* 如何支持新pc版？不用控件的Scroll
* pc版登陆时间较长不点击可能会自动退出，如何解决？
	* 知道哪个窗口退出了：ControlType:	UIA_PaneControlTypeId (0xC371) Name:	"登录" ClassName:	"WeChatLoginWndForPC"
	* （不退出时是：Name:	"微信" ControlType:	UIA_WindowControlTypeId (0xC370) ClassName:	"WeChatMainWndForPC"）
	* 知道是哪个账号：？？
