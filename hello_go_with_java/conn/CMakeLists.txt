cmake_minimum_required(VERSION 3.14)
project(conn)

set(CMAKE_CXX_STANDARD 14)

include_directories(/home/yong/Desktop/java_env/jdk-8u211-linux-x64/jdk1.8.0_211/include)
include_directories(/home/yong/Desktop/java_env/jdk-8u211-linux-x64/jdk1.8.0_211/include/linux)

add_library(conn SHARED library.cpp library.h com_thinking_utils_Native.cpp com_thinking_utils_Native.h)