/**
* @Author: Aceld(刘丹冰)
* @Date: 2019/5/22 16:43
* @Mail: danbing.at@gmail.com
*/
package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
 */
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	//直接connect 服务器得到一个 已经建立好的conn句柄
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start errr", err)
		return
	}

	for {
		//写
		_, err := conn.Write([]byte("Hello Zinx..."))
		if err != nil {
			fmt.Println("write conn err,", err)
			return
		}

		//读
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf(" servar call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 *time.Second)
	}


}
