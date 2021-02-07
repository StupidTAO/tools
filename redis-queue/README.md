# summary
使用go语言代码链接reids服务，并用go语言实现一个redis队列，保持先进先出

# step
1.命令行启动redis服务
`/usr/local/bin/redis-server /usr/local/etc/redis.conf`
2.运行go语言本地代码
1).启动生产者 `go run redis_queue.go pro`
2).启动消费者 `go run redis_queue.go pro`
