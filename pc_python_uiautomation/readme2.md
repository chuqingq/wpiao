pc版本，使用python的uiautomation驱动。

https://github.com/chuqingq/wpiao/tree/master/pc_python_uiautomation

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
python
>>> import main
>>> main.voteall()
```

# TODO

* ~~批量手机投票（已完成）~~
* 如何支持新pc版？不用控件的Scroll
* pc版登陆时间较长不点击可能会自动退出，如何解决？
	* 知道哪个窗口退出了：ControlType:	UIA_PaneControlTypeId (0xC371) Name:	"登录" ClassName:	"WeChatLoginWndForPC"
	* （不退出时是：Name:	"微信" ControlType:	UIA_WindowControlTypeId (0xC370) ClassName:	"WeChatMainWndForPC"）
	* 知道是哪个账号：？？

* 养号：和30天以上的互相
* 养号：PC版本，自动发送消息
* 测试：每天50个
* 投票：支持PC新版本
	* 方案：已确认：PC新版本也使用了系统的证书机制和代理机制。因此可以用类似fiddler的方案拿到投票的页面Url，后面用代码实现即可
		* fiddler安装FiddlerCertMaker.exe插件，用于https注入
		* 通过Tool->fiddler options->HTTPS->actions->TrustRootCertificate导出fiddler的根证书
		* 把上述证书导入到系统中
	* 1、根据根证书制作一个模拟mp.weixin.qq.com:443的服务端（openssl 根证书 制作服务端证书）（）
		* openssl genrsa -out privkey.pem 2048
		* openssl req -new -x509 -key privkey.pem -out cacert.pem -days 1095
		* 修改cacert.pem为cacert.cer，在server.js（nodejs）中使用
		* 双击cer文件导入到“受信任的根证书颁发机构”，IE下正常，chrome下还是没辙。
		* TODO 使用phantomjs多实例来执行投票动作
	* 2、把根证书导入系统
	* TODO 3、制作一个proxy，把到mp.weixin.qq.com:443的HTTPS CONNECT重定向到步骤1的服务端（goproxy）
	* 4、步骤1的服务端每收到如下请求，就模拟浏览器进行交互。。。
