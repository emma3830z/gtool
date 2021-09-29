package gtool

import (
	"log"
	"math"
	"strconv"
	"strings"

	GOset "github.com/deckarep/golang-set"
)

// IntInSlice 判斷數值是否在 slice 中
func IntInSlice(i int, list []int) bool {
	for _, data := range list {
		if data == i {
			return true
		}
	}
	return false
}

// StringInSlice 判斷字串是否在 slice 中
func StringInSlice(s string, list []string) bool {
	for _, data := range list {
		if data == s {
			return true
		}
	}
	return false
}

// UniqueIntList 排除重複數字
func UniqueIntList(list []int) (uniqueList []int) {
	for _, i := range list {
		if !IntInSlice(i, uniqueList) {
			uniqueList = append(uniqueList, i)
		}
	}
	return
}

// UniqueStrList 排除重複字串
func UniqueStrList(list []string) (uniqueList []string) {
	for _, s := range list {
		if !StringInSlice(s, uniqueList) {
			uniqueList = append(uniqueList, s)
		}
	}
	return
}

// StrInSliceWithContain 判斷字串是否在 slice 中有部份相同的存在
func StrInSliceWithContain(s string, list []string) bool {
	for _, data := range list {
		if strings.Contains(data, s) {
			return true
		}
	}
	return false
}

// SplitIntSlice 指定每組數量，將 []int 切分成多組
func SplitIntSlice(list []int, cutNum int) (listSplit [][]int) {
	if cutNum == 0 {
		// 每組數量 0，直接回傳空
		return
	}

	chap := math.Ceil(float64(len(list)) / float64(cutNum))
	for i := 0; i < int(chap); i++ {
		head := i * cutNum
		tail := (i + 1) * cutNum

		if tail > len(list) {
			tail = len(list)
		}

		newSlice := list[head:tail]
		listSplit = append(listSplit, newSlice)
	}
	return
}

// SplitStringSlice 指定每組數量，將 []string 切分成多組
func SplitStringSlice(list []string, cutNumber int) (listSplit [][]string) {
	if cutNumber == 0 {
		// 每組數量 0，直接回傳空
		return
	}

	chap := math.Ceil(float64(len(list)) / float64(cutNumber))
	for i := 0; i < int(chap); i++ {
		head := i * cutNumber
		tail := (i + 1) * cutNumber

		if tail > len(list) {
			tail = len(list)
		}

		newSlice := list[head:tail]
		listSplit = append(listSplit, newSlice)
	}
	return
}

// ToInterfaceSlice 將資料轉為轉換成 []interface{}
// 支援型態: []int, []string, []float64
func ToInterfaceSlice(source interface{}) (r []interface{}) {
	switch source.(type) {
	case []int:
		data := source.([]int)
		r = make([]interface{}, len(data))
		for i := range r {
			r[i] = interface{}(data[i])
		}
		return
	case []string:
		data := source.([]string)
		r = make([]interface{}, len(data))
		for i := range r {
			r[i] = interface{}(data[i])
		}
		return
	case []float64:
		data := source.([]float64)
		r = make([]interface{}, len(data))
		for i := range r {
			r[i] = interface{}(data[i])
		}
		return
	}

	return
}

// ToInterfaceSlice 將資料轉為轉換成 []int
// 支援型態: []string, []interface{}
func ToIntSlice(source interface{}) (r []int) {
	switch source.(type) {
	case []string:
		var err error

		data := source.([]string)
		r = make([]int, len(data))
		for i := range r {
			r[i], err = strconv.Atoi(data[i])
			if err != nil {
				log.Println("ToIntSlice 字串轉數字失敗, source:", source)
				continue
			}
		}
		return
	case []interface{}:
		data := source.([]interface{})
		r = make([]int, len(data))
		for i, v := range data {
			switch v.(type) {
			case int:
				r[i] = v.(int)
			}
		}
		return
	}

	return
}

// ToStringSlice 將資料轉為轉換成 []string
// 支援型態: []int, []interface{}
func ToStringSlice(source interface{}) (r []string) {
	switch source.(type) {
	case []int:
		data := source.([]int)
		r = make([]string, len(data))
		for i := range r {
			r[i] = strconv.Itoa(data[i])
		}
		return
	case []interface{}:
		data := source.([]interface{})
		r = make([]string, len(data))
		for i, v := range data {
			switch v.(type) {
			case string:
				r[i] = v.(string)
			}
		}
	}

	return
}

// DifferSet 兩個[]int 取差集
// 回傳資料為亂序 且 會自動過濾重複資料
func DifferSet(source []int, target []int) []int {
	sourceSet := GOset.NewSetFromSlice(ToInterfaceSlice(source))
	targetSet := GOset.NewSetFromSlice(ToInterfaceSlice(target))
	operatedSet := sourceSet.Difference(targetSet)
	result := operatedSet.ToSlice()
	return ToIntSlice(result)
}

// IntersectSet 兩個[]int 取交集
// 回傳資料為亂序 且 會自動過濾重複資料
func IntersectSet(source []int, target []int) []int {
	sourceSet := GOset.NewSetFromSlice(ToInterfaceSlice(source))
	targetSet := GOset.NewSetFromSlice(ToInterfaceSlice(target))
	operatedSet := sourceSet.Intersect(targetSet)
	result := operatedSet.ToSlice()
	return ToIntSlice(result)
}
