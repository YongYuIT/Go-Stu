package check_sum

import (
	"crypto/md5"
	"fmt"
	"net"
)

func GetMd5Check() []byte {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("cannot read config from net" + err.Error())
	}
	thisMac := ""
	for index, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		if mac != nil {
			fmt.Println(inter.Name, "-->", mac)
			if index == 0 {
				thisMac = mac.String()
			}
		}
	}
	sum := md5.Sum([]byte("this is spider for thk checked --> " + thisMac))
	return sum[:]
}
