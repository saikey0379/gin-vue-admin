package utils

import (
	"sort"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/rulefmt"
)

type RuleGroups struct {
	Groups []RuleGroup `yaml:"groups"`
}

type RuleGroup struct {
	Name     string         `yaml:"name"`
	Interval model.Duration `yaml:"interval,omitempty"`
	Limit    int            `yaml:"limit,omitempty"`
	Rules    []rulefmt.Rule `yaml:"rules"`
}

type RuleMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SortRuleGroupsByName(rgs []RuleGroup) []RuleGroup {
	result := make([]RuleGroup, len(rgs))
	copy(result, rgs)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name > result[j].Name // 对字符串使用小于操作符会进行字典排序
	})

	return result
}

func SortRulesByName(rs []rulefmt.Rule) []rulefmt.Rule {
	result := make([]rulefmt.Rule, len(rs))
	copy(result, rs)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Record < result[j].Record // 对字符串使用小于操作符会进行字典排序
	})

	return result
}

type RuleGroupsSorter []RuleGroup

func (rgs RuleGroupsSorter) Len() int {
	return len(rgs)
}

func (rgs RuleGroupsSorter) Less(i, j int) bool {
	// 在这里定义你的排序规则
	return rgs[i].Name < rgs[j].Name // 假设你的排序规则是基于MyStruct的一个字段的大小
}

func (rgs RuleGroupsSorter) Swap(i, j int) {
	rgs[i], rgs[j] = rgs[j], rgs[i]
}
