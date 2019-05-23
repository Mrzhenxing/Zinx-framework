/**
* @Author: Aceld(刘丹冰)
* @Date: 2019/5/22 17:19
* @Mail: danbing.at@gmail.com
*/
package net

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//具体的TCP链接模块
type Connection struct {
	//当前链接的原生套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前的链接状态
	isClosed bool

	//当前链接所绑定的业务处理方法
	handleAPI ziface.HandleFunc
}

/*
初始化链接方法
 */
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) ziface.IConnection {
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		handleAPI:callback_api,
		isClosed:false,
	}

	return c
}

//针对链接读业务的方法
func (c *Connection) StartReader() {
	//从对端读数据
	fmt.Println("Reader go is startin....")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is = ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			break
		}

		//将数据 传递给我们 定义好的Handle Callback方法
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "Handle is error", err)
			break
		}
	}

}


//启动链接
func (c *Connection) Start() {
	fmt.Println("Conn Start（）  ... id = ", c.ConnID)
	//先进行读业务
	go c.StartReader()

	//TODO 进行写业务
}

//停止链接
func (c *Connection) Stop() {
	fmt.Println("c. Stop() ... ConnId = ", c.ConnID)

	//回收工作
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//关闭原生套接字
	_ = c.Conn.Close()
}

//获取链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取conn的原生socket套接字
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取远程客户端的ip地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据给对方客户端
func (c *Connection) Send(data []byte, cnt int) error {

	if _, err := c.Conn.Write(data[:cnt]); err != nil {
		fmt.Println("send buf error")
		return err
	}

	return nil
}