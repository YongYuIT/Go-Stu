package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path"
)

//获取秘钥(private key)
func publicKey(keypath string) ssh.AuthMethod {

	key, err1 := ioutil.ReadFile(keypath)
	if err1 != nil {
		fmt.Println("读取秘钥失败", err1)
	}
	signer, err2 := ssh.ParsePrivateKey(key)
	if err2 != nil {
		fmt.Println("ssh 秘钥签名失败", err2)
	}
	return ssh.PublicKeys(signer)
}

//获取ssh连接
func GetSSHConect(ip, user string, port int, keypath string) *ssh.Client {
	con := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{publicKey(keypath)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr := fmt.Sprintf("%s:%d", ip, port)
	client, err := ssh.Dial("tcp", addr, con)
	if err != nil {
		fmt.Println("Dail failed: ", err)
		panic(err)
	}
	return client
}

//远程执行命令
func Exec_Cmd(ip, user, cmd string, port int, keypath string) int {
	client := GetSSHConect(ip, user, port, keypath)
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("创建会话失败", err)
		panic(err)
	}
	defer session.Close()
	err1 := session.Run(cmd)
	if err1 != nil {
		fmt.Println("远程执行命令失败", err1)
		return 2
	} else {
		fmt.Println("远程执行命令成功")
		return 1
	}
}

// 远程执行脚本
func Exec_Task(ip, user, localpath, remotepath string, keypath string) int {
	port := 22
	client := GetSSHConect(ip, user, port, keypath)
	UploadFile(ip, user, localpath, remotepath, port, keypath)
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("创建会话失败", err)
		panic(err)
	}
	defer session.Close()
	remoteFileName := path.Base(localpath)
	dstFile := path.Join(remotepath, remoteFileName)
	err1 := session.Run(fmt.Sprintf("/usr/bin/sh %s", dstFile))
	if err1 != nil {
		fmt.Println("远程执行脚本失败", err1)
		return 2
	} else {
		fmt.Println("远程执行脚本成功")
		return 1
	}
}

//获取ftp连接
func getftpclient(client *ssh.Client) *sftp.Client {
	ftpclient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}
	return ftpclient
}

//上传文件
func UploadFile(ip, user, localpath, remotepath string, port int, keypath string) {
	client := GetSSHConect(ip, user, port, keypath)
	ftpclient := getftpclient(client)
	defer ftpclient.Close()

	remoteFileName := path.Base(localpath)
	fmt.Println(localpath, remoteFileName)
	srcFile, err := os.Open(localpath)
	if err != nil {
		fmt.Println("打开文件失败", err)
		panic(err)
	}
	defer srcFile.Close()

	dstFile, e := ftpclient.Create(path.Join(remotepath, remoteFileName))
	if e != nil {
		fmt.Println("创建文件失败", e)
		panic(e)
	}
	defer dstFile.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("已读取到文件末尾")
				break
			} else {
				fmt.Println("读取文件出错", err)
				panic(err)
			}
		}
		dstFile.Write(buffer[:n])
		//注意，由于文件大小不定，不可直接使用buffer，否则会在文件末尾重复写入，以填充1024的整数倍
	}
	fmt.Println("文件上传成功")
}

//文件下载
func DownLoad(ip, user, localpath, remotepath string, port int, keypath string) {
	client := GetSSHConect(ip, user, port, keypath)
	ftpClient := getftpclient(client)
	defer ftpClient.Close()

	srcFile, err := ftpClient.Open(remotepath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		panic(err)
	}
	defer srcFile.Close()
	localFilename := path.Base(remotepath)
	dstFile, e := os.Create(path.Join(localpath, localFilename))
	if e != nil {
		fmt.Println("文件创建失败", e)
		panic(e)
	}
	defer dstFile.Close()
	if _, err1 := srcFile.WriteTo(dstFile); err1 != nil {
		fmt.Println("文件写入失败", err1)
		panic(err1)
	}
	fmt.Println("文件下载成功")
}
