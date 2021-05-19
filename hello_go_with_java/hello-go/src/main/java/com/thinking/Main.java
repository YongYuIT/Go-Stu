package com.thinking;

import com.thinking.utils.Native;

public class Main {
    public static void main(String[] args) {
        System.load("/yong/codes/yong_research/Go-Stu/hello_go_with_java/hello-go/libs/libconn.so");
        Native.SayHello();
    }
}
