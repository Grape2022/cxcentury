启动 graphviz

`dot -version`



启动压测

go -wrk -c=400 -t=8 -n=100000 http://localhost:6060/ApiPprof

执行测试

 go tool pprof --text http://127.0.0.1:6060/debug/pprof/profile

执行后，会自动监听30秒

开启监听后，登录

登录

http://127.0.0.1:6060/debug/pprof/allocs?debug=1





生成火焰图

go tool pprof -http=":8081" [binary] [profile]

go tool pprof -http localhost:9295 C:\Users\Administrator\pprof\pprof.___go_build_main_go.exe.samples.cpu.015.pb.gz

运行

go tool pprof http://localhost:6060/debug/pprof/profile

命令

```
top 10 列出前10
web   调用 graphviz 生成svg图片，然后打开
list  查看具体的函数分析
pdf   命令可以生成可视化的pdf文件
help  命令可以提供所有pprof支持的命令说明
```

web路径：
file:///C:/Users/Administrator/AppData/Local/Temp/pprof008.svg

pdf路径：在项目目录下

例如：D:\ApiPprof\cxcenturyPprof
