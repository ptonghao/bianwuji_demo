/*
 * @Author: Jimpu
 * @Description: string util
 */

package utils

import (
	"fmt"
	"unicode"
)

// input 是否是数字
func NumberValid(input string) bool {
	if input == "" {
		return false
	}

	for _, v := range []rune(input) {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

// FindBinarySearch 二分查找法,找到 最接近value值的下标
func FindBinarySearch(datas []int64, leftPtr int64, rightPtr int64, findNum int64) int64 {
	if leftPtr > rightPtr || leftPtr < 0 || rightPtr < 0 {
		return -1
	}

	dLen := len(datas)
	//先求出中间的指针位置
	mPtr := (leftPtr + rightPtr) / 2

	// 取最接近的值
	if mPtr+1 < int64(dLen) && (datas[mPtr] < findNum && datas[mPtr+1] > findNum) {
		if 2*findNum > datas[mPtr+1]+datas[mPtr] {
			return int64(mPtr + 1)
		}
		return int64(mPtr)
	}
	fmt.Println(fmt.Sprintf("FindBinarySearch len=%v, leftPtr=%v, rightPtr=%v, mPtr=%v", len(datas), leftPtr, rightPtr, mPtr))
	if datas[mPtr] > findNum {
		return FindBinarySearch(datas, leftPtr, mPtr-1, findNum)
	} else if datas[mPtr] < findNum {
		return FindBinarySearch(datas, mPtr+1, rightPtr, findNum)
	} else {
		return int64(mPtr)
	}
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
