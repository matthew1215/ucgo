package iutil

import (
	"strings"
)

const DEVICE_IOS string = "ios"
const DEVICE_ANDROID string = "android"

const PROID_NORMAL = "normal" //正式版

var proidList = map[string]int{
	PROID_NORMAL: 1,
}

// 是否在proid列表中
func IsInProidList(proid string) bool {
	if _, ok := proidList[proid]; ok {
		return true
	} else {
		return false
	}
}

// 获取设备类型
func GetOS(os string) string {
	var device string
	if strings.Contains(os, DEVICE_IOS) {
		device = DEVICE_IOS
	} else {
		device = DEVICE_ANDROID
	}
	return device
}
