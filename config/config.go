package config

type MessageRule struct {
	MessageType string
	RuleList    [][]int // 电报编组规则
}

var Config = struct {
	MessageRuleList []MessageRule
}{
	MessageRuleList: []MessageRule{
		{
			MessageType: "ARR",
			RuleList: [][]int{
				{7, 13, 17},
				{7, 13, 16, 17},
			},
		},
		{
			MessageType: "DEP",
			RuleList: [][]int{
				{7, 13, 16, 18},
			},
		},
		{
			MessageType: "FPL",
			RuleList: [][]int{
				{7, 8, 9, 10, 13, 15, 16, 18},
			},
		},
		{
			MessageType: "CHG",
			RuleList: [][]int{
				{7, 13, 16, 18, 22},
			},
		},
		{
			MessageType: "CNL",
			RuleList: [][]int{
				{7, 13, 16, 18},
			},
		},
		{
			MessageType: "DLA",
			RuleList: [][]int{
				{7, 13, 16, 18},
			},
		},
		{
			MessageType: "EST",
			RuleList: [][]int{
				{7, 13, 14, 16},
			},
		},
		{
			MessageType: "CPL",
			RuleList: [][]int{
				{7, 8, 9, 10, 13, 14, 15, 16, 18},
			},
		},
		{
			MessageType: "CDN",
			RuleList: [][]int{
				{7, 13, 16, 22},
			},
		},
		{
			MessageType: "ACP",
			RuleList: [][]int{
				{7, 13, 16},
			},
		},
		{
			MessageType: "RQP",
			RuleList: [][]int{
				{7, 13, 16, 18},
			},
		},
		{
			MessageType: "RQS",
			RuleList: [][]int{
				{7, 13, 16, 18},
			},
		},
		{
			MessageType: "SPL",
			RuleList: [][]int{
				{7, 13, 16, 18, 19},
			},
		},
		{
			MessageType: "ALR",
			RuleList: [][]int{
				{5, 7, 8, 9, 10, 13, 15, 16, 18, 19, 20},
			},
		},
		{
			MessageType: "RCF",
			RuleList: [][]int{
				{7, 21},
			},
		},
		{
			MessageType: "PLN",
			RuleList: [][]int{
				{1001, 7, 8, 9, 13, 16, 15, 18},
			},
		},
	},
}

type Group3_unit struct {
	MessageType string
	Desc        string
}

var Group3_info = struct {
	Info []Group3_unit
}{
	Info: []Group3_unit{
		{
			"ARR",
			"落地报",
		},
		{
			"DEP",
			"起飞报",
		},
		{
			"FPL",
			"領航计划报",
		},
		{
			"CHG",
			"修改領航报",
		},
		{
			"CNL",
			"取消领航报",
		},
		{
			"DLA",
			"延误报",
		},
		{
			"EST",
			"预计飞越报",
		},
		{
			"CPL",
			"飞行计划变更报",
		},
		{
			"CDN",
			"管制协调报",
		},
		{
			"ACP",
			"协调接受报",
		},
		{
			"LAM",
			"逻辑确认报",
		},
		{
			"RQP",
			"请求飞行计划报",
		},
		{
			"RQS",
			"请求领航计划补充信息报",
		},
		{
			"SPL",
			"领航计划补充信息报",
		},
		{
			"ALR",
			"告警报",
		},
		{
			"RCF",
			"无线电通信失效报",
		},
		{
			"PLN",
			"飞行预报",
		},
	},
}
