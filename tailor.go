/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : tailor.go
#   Last Modified : 2021-06-25 15:08
#   Describe      :
#
# ====================================================*/

package tailor

import (
	"strings"
)

/* 根据需要返回的key列表，裁剪返回结果单个结果 */
func ClipOne(keys []string, result Object) Object {
	if nil == result { // 没数据，直接返回nil
		return nil
	}
	cliped := make(Object)
	for _, key := range keys {
		if _, exist := result[key]; exist {
			cliped[key] = result[key]
		}
	}
	return cliped
}

/* 根据需要返回的key列表，裁剪返回结果数组 */
func ClipMore(keys []string, results []Object) []Object {
	if 0 == len(results) { // 没数据，直接返回nil
		return nil
	}
	clipeds := make([]Object, 0, 20)
	for _, result := range results {
		cliped := ClipOne(keys, result)
		clipeds = append(clipeds, cliped)
	}
	return clipeds
}

/* 增加合并数据函数 */
func Merge(first, second Object) Object {
	for k, v := range second {
		first[k] = v
	}
	return first
}

/* 隐藏名字 */
func HideName(key string, result Object) {
	if nil == result {
		return
	}
	// 姓名部分字隐藏显示为“*”,姓名只有两个字的隐藏姓,超过两字的只显示头1个字和最后一个字
	if _, exist := result[key]; exist {
		name := result[key].(string)
		splitName := strings.Split(name, "")
		if len(splitName) == 2 {
			splitName[0] = "*"
		} else {
			for i := 1; i < len(splitName)-1; i++ {
				splitName[i] = "*"
			}
		}
		name = strings.Join(splitName, "")
		result[key] = name
	}
}

/* 隐藏身份证 */
func HideIdCardNo(key string, result Object) {
	if nil == result {
		return
	}
	// 身份证号码中间部分数字隐藏显示为“*”,只显示开头4位和最后4位
	if _, exist := result[key]; exist {
		idCardNo := result[key].(string)
		splitIdCardNo := strings.Split(idCardNo, "")
		for i := 4; i < len(idCardNo)-4; i++ {
			splitIdCardNo[i] = "*"
		}
		result[key] = strings.Join(splitIdCardNo, "")
	}
}

/* 对返回结果做过滤，隐藏个人身份信息 */
func HideIdentityInformation(result Object) {
	if nil == result {
		return
	}
	HideName("name", result)
	HideIdCardNo("id_card_no", result)
}
