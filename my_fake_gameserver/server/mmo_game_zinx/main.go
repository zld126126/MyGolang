package main

import (
	"fmt"
	"my_fake_gameserver/mmo_game_zinx/apis"
	"my_fake_gameserver/mmo_game_zinx/core"
	"my_fake_gameserver/zinx/ziface"
	"my_fake_gameserver/zinx/znet"
)

// 当前客户端建立连接之后的hook函数
func OnConnecionAdd(conn ziface.IConneciton) {
	//创建一个Player对象
	player := core.NewPlayer(conn)

	//给客户端发送MsgID:1的消息: 同步当前Player的ID给客户端
	player.SyncPid()

	//给客户端发送MsgID:200的消息: 同步当前Player的初始位置给客户端
	player.BroadCastStartPosition()

	//将当前新上线的玩家添加到WorldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定一个Pid 玩家ID的属性
	conn.SetProperty("pid", player.Pid)

	//同步周边玩家，告知他们当前玩家已经上线，广播当前玩家的位置信息
	player.SyncSurrounding()

	fmt.Println("=====> Player pid = ", player.Pid, " is arrived <=====")
}

func main() {
	//创建zinx server句柄
	s := znet.NewServer("MMO Game Zinx")

	//连接创建和销毁的HOOK钩子函数
	s.SetOnConnStart(OnConnecionAdd)

	//注册一些路由业务
	s.AddRouter(2, &apis.WorldChatApi{})

	//启动服务
	s.Serve()
}
