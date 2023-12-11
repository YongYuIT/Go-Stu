原项目地址：https://github.com/topfreegames/pitaya/tree/main/examples/demo/chat

# 注意

原项目中，这里默认是 pitaya.Cluster 即集群模式，会由于连不上NAT服务而导致直接启动失败。需改成单机模式 pitaya.Standalone

# 使用Web端交互

http://localhost:3251/web/

# 使用 pitaya-cli 与服务交互

~~~
pitaya-cli
connect localhost:3250
request chat.room.join
notify chat.room.message {"name":"2","content":"11111"}
~~~
