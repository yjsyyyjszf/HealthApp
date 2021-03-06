// Units
package Units

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	PRE_STUDYUID = "1.2.826.0.1.3680043.9.7604."
)

type OFHashValue struct {
	first  int32
	second int32
}

func GetStudyHashDir(studyuid string) string {
	vl := CreateHashValue([]byte(studyuid), strings.Count(studyuid, "")-1)
	dir := "/" + strconv.Itoa((int)(vl.first)) + "/" + strconv.Itoa((int)(vl.second))
	return dir
}

func GetCurrentTime() string {
	st := time.Now().String()
	st = strings.ReplaceAll(st, "-", "")
	st = strings.ReplaceAll(st, ":", "")
	st = strings.ReplaceAll(st, ".", "")
	st = strings.ReplaceAll(st, " ", "")
	st = st[0:21]
	return st
	// fmt.Println(st)
	// fmt.Println(time.Now())
	// return tm.Format("20060102130405000000")
}

func GetCurrentFormatTime() string {
	st := time.Now().String()
	st = st[0:19]
	return st
}

func GetRandUID() string {
	rand.Seed(int64(time.Now().UnixNano()))
	return strconv.Itoa(rand.Int())
}

func GetNewStudyUID() string {
	return PRE_STUDYUID + GetCurrentTime()
}

func CreateHashValue(buffer []byte, length int) OFHashValue {
	//2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
	//101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179
	//181 191 193 197 199
	// var a string = "Runoob"
	const M1 int32 = 71
	const M2 int32 = 79

	const D1 int32 = 197
	const D2 int32 = 199
	var value int32 = 0
	var hashvalue OFHashValue
	var i int = 0
	for i = 0; i < length; i++ {
		// println("--buffer[i]--")
		// println(buffer[i])
		// println(buffer[i] & 0xFF)
		// println(value * (int32)(M1))
		value = value*M1 + (int32)(buffer[i])
		// println(value)
	}
	value = value % D1
	if value < 0 {
		value = value + D1
	}
	hashvalue.first = value

	value = 0
	for i = 0; i < length; i++ {
		value = (int32)(value*M2) + (int32)(buffer[i])
	}
	value = value % D2
	if value < 0 {
		value = value + D2
	}
	hashvalue.second = value
	return hashvalue
}

// type OFHashValue struct {
// 	first  uint
// 	second uint
// }

// func GetStudyHashDir(studyuid string) string {
// 	vl := CreateHashValue([]byte(studyuid), (uint)(strings.Count(studyuid, "")-1))
// 	dir := "/" + strconv.Itoa((int)(vl.first)) + "/" + strconv.Itoa((int)(vl.second))
// 	return dir
// }

// func CreateHashValue(buffer []byte, length int) OFHashValue {
// 	//2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
// 	//101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179
// 	//181 191 193 197 199
// 	// var a string = "Runoob"
// 	println("--length--")
// 	println(length)
// 	const M1 uint = 71
// 	const M2 uint = 79

// 	const D1 uint = 197
// 	const D2 uint = 199
// 	var value uint = 0
// 	var hashvalue OFHashValue
// 	var i uint = 0
// 	for i = 0; i < length; i++ {
// 		println("--buffer[i]--")
// 		println(buffer[i])
// 		println((uint)(buffer[i] & 0xFF))
// 		println(value * M1)
// 		value = value*M1 + (uint)(buffer[i]&0xFF)
// 		println(value)
// 	}
// 	value = value % D1
// 	if value < 0 {
// 		value = value + D1
// 	}
// 	hashvalue.first = value

// 	value = 0
// 	for i = 0; i < length; i++ {
// 		value = value*M2 + (uint)(buffer[i]&0xFF)
// 	}
// 	value = value % D2
// 	if value < 0 {
// 		value = value + D2
// 	}
// 	hashvalue.second = value
// 	return hashvalue
// }
