package analysis_group

import (
	"fmt"
	"telegraphTranslator/global"
)

// 编组1001，获取日期
func AnalysisGroup1001(message string) (data string, err error) {
	fmt.Println(fmt.Sprintf("AnalysisGroup1001-message:%s", message))
	if true == global.GlobalVar.PrintDebugInfo {
		fmt.Println("GetGroup1001Info：", message)
		defer func() {
			fmt.Println("GetGroup1001Info resp: ", data)
		}()
	}

	// 起飞时间(可选)
	depTime := ""
	if len(message) == 4 {
		// todo 待处理AFIL
		// 处理起飞时间
		depTime = "航班日期: " + message[0:2] + "-" + message[2:]
	}

	if "" != depTime {
		data = depTime + "\n"
	}

	return
}
