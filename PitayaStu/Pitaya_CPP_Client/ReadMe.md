# 编译得到 libpitaya 动态链接库

1. 安装CMake

下载：https://cmake.org/download/

到libpitaya目录下新建build目录，进入build目录执行

~~~
cmake .. -DBUILD_SHARED_LIBS=ON -DCMAKE_BUILD_TYPE=Release
cmake --build .
~~~

在build\Debug下即可找到动态链接库

# 如何添加 libpitaya 的引用

1. 添加libpitaya头文件引用

找到解决方案资源管理器里面的C++项目

点击项目顶部扳手按钮（属性）

配置属性-->C/C++-->常规-->附加包含目录

将libpitaya项目源码中libpitaya\include目录添加进去

2. 添加libpitaya动态链接库的引用

找到解决方案资源管理器里面的C++项目

点击项目顶部扳手按钮（属性）

属性配置-->链接器-->常规-->附加库目录

将libpitaya\build\Debug目录添加进去

属性配置-->链接器-->输入-->附加依赖项

将pitaya-windows.lib添加进去

将libpitaya\build\Debug\pitaya-windows.dll文件拷贝一份到项目x64\Debug目录下，使其在运行时能加载pitaya-windows.dll

