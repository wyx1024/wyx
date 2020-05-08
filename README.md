# wyx


运行分布式爬虫:

启动ItemSaver服务

go run distributed_crawler/rpc/server/Item/rpc.go -itemsaver_host=PORT


启动Worker服务(可起多个worker服务)

go run distributed_crawler/rpc/server/worker/rpc.go -worker_host=PORT_WORKER1

go run distributed_crawler/rpc/server/worker/rpc.go -worker_host=PORT_WORKER2


启动程序

go run distributed_crawler/main.go -item_host=":PORT" -worker_hosts=":PORT_WORKER,:PORT_WORKER2"