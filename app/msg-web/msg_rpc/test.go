package main

import "fmt"

//func main() {
//	//配置连连接参数(无加密)
//	dial, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	defer dial.Close()
//	//创建客户端连接
//	client := msg_rpc.NewMsgClient(dial)
//	//通过客户端调用方法
//	res, err := client.Ping(context.Background(), &msg_rpc.PingRequest{Ping: "xxx"})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(res)
//
//}

func main() {
	slice := make([]int, 0)

	for i := 0; i < 512; i++ {
		slice = append(slice, i)
	}

	newSlice := append(slice, 5000)
	fmt.Printf("Before Pointer = %p, len = %d, cap = %d\n", &slice, len(slice), cap(slice))
	fmt.Printf("Before Pointer = %p, len = %d, cap = %d\n", &newSlice, len(newSlice), cap(newSlice))
}
