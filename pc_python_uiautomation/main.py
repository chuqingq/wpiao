# -*- coding: utf-8 -*-

import websocket
# import thread
import time
import socket
import json
import threading

import vote

def on_message(ws, message):
    print('on_message: ' + message)
    msg = json.loads(message)
    if msg['cmd'] == 'vote':
        # 开始投票 {"cmd":"vote", "url": "http://wxxxx", "votes": 100}
        vote.vote(msg['url'], msg['votes'])
        # 投票结束 通知voter cmd: vote_finish
        vote_finish = json.dumps({'cmd': 'vote_finish', 'url': msg['url'], 'votes': msg['votes']})
        ws.send(vote_finish)
        print('vote_finish: ' + vote_finish)
    elif msg['cmd'] == 'train':
        vote.train()
    else:
        print('cmd('+msg['cmd']+') is invalid')
    # 其他动作：TODO

def on_error(ws, error):
    print('on_error: ', error)

def on_close(ws):
    print("on_close, will reconnect in 5 seconds... ")
    time.sleep(5)
    start_websocket()

def on_open(ws):
    print('on_open: ')
    # 发送当前微信账号数量
    # {"pc": "wp-001", "account_count": 10}
    res = {"cmd": "login", "name": socket.getfqdn(), "accounts": len(vote.handleDict)}
    print("login: ", res)
    ws.send(json.dumps(res))
    # def run(*args):
    #     for i in range(3):
    #         time.sleep(1)
    #         ws.send("Hello %d" % i)
    #     time.sleep(1)
    #     ws.close()
    #     print "thread terminating..."
    # thread.start_new_thread(run, ())

def start_websocket():
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://127.0.0.1:8080/api/ws/runner",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
    ws.on_open = on_open
    ws.run_forever()

# print('start_websocket...')
# t = threading.Thread(target=start_websocket, args=())
# t.start()

if __name__ == "__main__":
    # websocket.enableTrace(True)
    # ws = websocket.WebSocketApp("ws://127.0.0.1:8080/api/ws/pc",
    #                           on_message = on_message,
    #                           on_error = on_error,
    #                           on_close = on_close)
    # ws.on_open = on_open
    # ws.run_forever()
    start_websocket()
