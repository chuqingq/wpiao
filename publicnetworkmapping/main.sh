# 先建立本机和服务器的信任关系，然后执行此脚本，确保公网映射一直工作。 
# TODO 判断自己的IP:192.168.31.72

while True
do
    echo `date` "will try reconnect in 2 seconds..."
    ssh -R 0.0.0.0:8080:192.168.31.72:8080 root@121.41.103.23
    sleep 2 
done

