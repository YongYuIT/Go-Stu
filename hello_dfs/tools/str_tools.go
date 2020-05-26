package tools

import "strconv"

func printHexStr(bytes []byte) string {
	reslut := ""
	for i := 0; i < len(bytes); i++ {
		reslut += strconv.FormatInt(int64(bytes[i]&0xff), 16) + " "
	}
	return reslut
}
