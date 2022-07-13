package utils

import (
	"log-prober-go/common"
)

//
// Recover
//  @Description: 异常捕获 记录错误
//
func Recover() {
	if recoverResult := recover(); recoverResult != nil {
		common.Logger.Error(recoverResult)
	}
}
