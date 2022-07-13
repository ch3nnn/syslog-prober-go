package handler

//
// ParseHandler
//  @Description: 正则表达式 Json Csv 等实现解析处理程序接口
//
type ParseHandler interface {

	//
	// Handle
	//  @Description: 解析处理方法
	//
	Handle() (message []byte)

	//
	// Output 输出数据方法
	//  @Description:
	//
	Output(message []byte)
}
