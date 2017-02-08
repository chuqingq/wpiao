20170204 验证PC版是否可行，稳定性是否有保证？
====

https://github.com/go-vgo/robotgo#examples

# 安装

go get github.com/go-vgo/robotgo

可能需要安装MinGW or other GCC

最新版本缺少png.h:

    D:\temp\weixin_vote\robotgo>go get -u github.com/go-vgo/robotgo
    # github.com/go-vgo/robotgo
    In file included from ./bitmap/../base/io_c.h:4:0,
                     from ./bitmap/goBitmap.h:25,
                     from ..\..\gopath\src\github.com\go-vgo\robotgo\robotgo.go:27:
    ./bitmap/../base/png_io_c.h:4:17: fatal error: png.h: No such file or directory
    compilation terminated.

