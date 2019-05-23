package net

import "Zinx/zinx/ziface"

type Requset struct {
	//	链接信息
	conn ziface.IConnection

	//数据内容
	data []byte
	//数据长度
	len int
}

func NewRequest(conn ziface.IConnection,data []byte,len int ) ziface.IRequest{
	req:=&Requset{
		conn:conn,
		data:data,
		len:len,
	}
	return req
}




//得到当前请求链接
func (r*Requset)GetConnection () ziface.IConnection{
	 return r.conn

}
//	得到链接的的数据
func (r*Requset)GetData() []byte{
	return r.data

}
//得到数据的长度
func (r*Requset)GetDateLen() int{
	return r.len

}