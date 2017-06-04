import socket

hostname = socket.getfqdn(socket.gethostname())
localip = socket.gethostbyname(hostname)

print(localip)
