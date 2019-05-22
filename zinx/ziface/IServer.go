/**
 Server模块的抽象层
*/
package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
}


