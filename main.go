package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 21; i < 20000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("110.242.68.66:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了\n", address)
		}(i)
	}
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d second", elapsed)
	wg.Wait()
}

// func main() {
// 	for i := 21; i < 200; i++ {
// 		address := fmt.Sprintf("8.7.198.46:%d", i)
// 		// address := fmt.Sprintf("110.242.68.66:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			fmt.Printf("%s关闭了\n", address)
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%s打开了\n", address)
// 	}
// }
