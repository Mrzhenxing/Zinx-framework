package ziface

type IRequest interface {
	//得到当前请求链接
	GetConnection () IConnection
	//	得到链接的的数据
	GetData() []byte
	//得到数据的长度
	GetDateLen() int
}
