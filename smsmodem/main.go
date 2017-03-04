package main

import (
	// "encoding/json"
	"flag"
	"log"
	"net"
)

var server = flag.String("server", "127.0.0.1:6080", "短信猫的server地址，默认'127.0.0.1:6080'")

func main() {
	flag.Parse()

	client, err := NewClient(*server)
	if err != nil {
		log.Printf("NewClient error: %v", err)
		return
	}

	ports, err := client.Ports()
	log.Printf("ports: %v, err: %v", ports, err)

	portInfo, err := client.PortInfo(1)
	log.Printf("portInfo: %+v, err: %v", portInfo, err)

	err = client.Task("发短信", "短信", "13770641012", "你好！", 1)
	log.Printf("task error: %v", err)

	smses, err := client.SMS(1, 5)
	log.Printf("smses: %+v, err: %v", smses, err)
}

type Client struct {
	Conn *net.TCPConn
}

func NewClient(server string) (*Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		log.Printf("net.ResolveTCPAddr error: %v", err)
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("DialTCP error: %v", err)
		return nil, err
	}

	return &Client{
		Conn: conn,
	}, nil
}

// 获取所有通道端口号指令：AP$PORTS?
// 软件返回数据格式：AP$PORTS=4,"160,161,162,163"，4表示有4个通道，160，161，162，163
func (c *Client) Ports() ([]int, error) {
	// TODO
	return nil, nil
}

// 获取单个通道的信息：AP$PINFO=端口号，如：AP$PINFO=160
// 返回数据格式：
// 84表示： 的数据长度
// 160表示端口号,  表示通道信息，以json数据返回可直接解析,phonum:表示通道手机号，imsi表示通道imsi，imei表示通道串号信息，state表示通道状态：0是表示通道就绪; 101表示通道端口打开失败;102表示未检测到手机卡;100表示正在初始化;如下图所示：
func (c *Client) PortInfo(port int) (*PortInfo, error) {
	// TODO
	return nil, nil
}

type PortInfo struct {
	State  int    `json:"state"`
	PhoNum string `json:"phonum"`
	IMSI   string `json:"imsi"`
	IMEI   string `json:'imei'`
}

// 发送任务指令：AP$TASK=数据长度,端口号,数据
// 说明: 数据格式为json格式: 如：
// {"taskname":"短信","tasktype":"短信","number":"10001","content":"102","count":1,"waittime":2｝
// 然后计算数据长度，如：{"taskname":"短信","tasktype":"短信","number":"10001","content":"102","count":1,"waittime":2｝数据长度为93，如果端口号小于或等于0表示该任务发送到所有通道，否则指定通道发送任务，指令为：
// AP$TASK=93,0,{"taskname":"短信","tasktype":"短信","number":"10001","content":"102","count":1,"waittime":2}  表示所有通道都执行短信任务，发送内容为102，接收号码为10001; 如果仅发送任务到端口160则：AP$TASK=93,160,{"taskname":"短信","tasktype":"短信","number":"10001","content":"102","count":1,"waittime":2}
func (c *Client) Task(taskname, tasktype, number, content string, waittime int) error {
	// TODO
	return nil
}

// 读取短信指令：AP$SMS=端口号,上报条数(0所上报所有);
// 返回数据格式：AP$SMS=数据长度,端口号,数据
// 如：AP$SMS=91,160,{"total":5,"time":"2016-3-10 18:0:15","number":"10001","content":"b.189.cn/HBDX参与活动。"}
// 91表示数据长度，160表示端口号，total表示总共未读短信;如果要获取全部通道短信发送：AP$SMS=0,0，如果只获取端口号160的前5条短信，则发送：AP$SMS=160,5
func (c *Client) SMS(port int, count int) ([]*SMS, error) {
	// TODO 返回结构有疑问。如果多条，怎么回？
	return nil, nil
}

type SMS struct {
	Port    int    // 如果输入port为0，则可能返回多条port不同SMS
	Total   int    `json:"total"`
	Time    string `json:"time"`
	Number  string `json:"number"`
	Content string `json:"content"`
}
