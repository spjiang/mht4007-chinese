package analysis_group

import (
	"errors"
	"fmt"
	"strings"
	"telegraphTranslator/config"
)

// 解析电报主函数
func ProcessMessageMain(message string) (data string, err error) {
	var (
		messageRule config.MessageRule
		role []int
	)

	fmt.Println("电报内容: ", message)

	// 按照"-"或者换行符进行切分
	newMessage := strings.ReplaceAll(message, "\n", "")
	messageArr := strings.Split(newMessage, "-")
	if 0 == len(messageArr) {
		err = errors.New(fmt.Sprintf("电报输入错误，请检查"))
		fmt.Println(err)
		return
	}
	fmt.Println()
	// 编组3 解析电报类型
	group3Info, err := AnalysisGroup3(messageArr[0])
	if err != nil {
		return
	}
	fmt.Println(group3Info)
	// 对逻辑确认报特殊处理
	if strings.HasPrefix(messageArr[0], "LAM") {
		data = group3Info
		return
	}

	for _, v := range config.Config.MessageRuleList {
		if v.MessageType == messageArr[0] {
			messageRule = v
		}
	}

	if messageRule.MessageType == "" {
		err = errors.New(fmt.Sprintf("未找到正确的电报类型[Type: %s]", messageArr[0]))
		return
	}

	for _, value := range messageRule.RuleList {
		if len(value) == len(messageArr)-1 {
			role = value
		}
	}

	if len(role) == 0 {
		err = errors.New(fmt.Sprintf("数据格式错误，请检查[Message: %s]\n", message))
		return
	}

	data += group3Info
	index := 1
	blank := "    -"
	for _, value := range role {
		processRsp := ""
		// 按照编组类型进行解析
		switch value {
		case 3:
			// 编组3已处理
		case 5:
			processRsp, err = AnalysisGroup5(messageArr[index])
		case 7:
			processRsp, err = AnalysisGroup7(messageArr[index])
		case 8:
			processRsp, err = AnalysisGroup8(messageArr[index])
		case 9:
			processRsp, err = AnalysisGroup9(messageArr[index])
		case 10:
			processRsp, err = AnalysisGroup10(messageArr[index])
		case 13:
			processRsp, err = AnalysisGroup13(messageArr[index])
		case 14:
			processRsp, err = AnalysisGroup14(messageArr[index])
		case 15:
			processRsp, err = AnalysisGroup15(messageArr[index])
		case 16:
			processRsp, err = AnalysisGroup16(messageArr[index])
		case 17:
			processRsp, err = AnalysisGroup17(messageArr[index])
		case 18:
			processRsp, err = AnalysisGroup18(messageArr[index])
		case 19:
			processRsp, err = AnalysisGroup19(messageArr[index])
		case 20:
			processRsp, err = AnalysisGroup20(messageArr[index])
		case 21:
			processRsp, err = AnalysisGroup21(messageArr[index])
		case 22:
			processRsp, err = AnalysisGroup22(messageArr[index])
		default:
			err = errors.New(fmt.Sprintf("当前未对编组%s的内容进行解析\n", value))
			return
		}
		if err != nil {
			return
		}
		data += blank + processRsp
		index++
	}

	return
}
