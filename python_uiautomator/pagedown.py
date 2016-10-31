from uiautomator import Device

d = Device('071efe2c00e37e37', adb_server_host='127.0.0.1', adb_server_port=5037)

# 翻页
d.drag(300,1000, 300, 300,2)

# 截图
d.screenshot('1.png')
d.dump("1.xml")

# 点击
d.click(300, 1000)
