package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	MAX  = 100
	PORT = 8545
)

// scan start scan ip port
func scan(part string) {
	ips := sip(part)
	ch := make(chan string, MAX)
	wg := sync.WaitGroup{}
	result := make([]string, 0)
	for _, ip := range ips {
		wg.Add(1)
		ch <- ip
		go func() {
			defer func() {
				wg.Done()
			}()
			host := <-ch
			address := host + ":" + fmt.Sprintf("%d", PORT)
			conn, err := net.DialTimeout("tcp", address, 5*time.Second)
			if err != nil {
				return
			}
			result = append(result, address)
			defer conn.Close()
		}()
	}
	wg.Wait()
	write(result)
}

func write(address []string) {
	file, err := os.OpenFile("result.out", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer file.Close()
	for _, add := range address {
		if _, err = fmt.Fprintln(file, add); err != nil {
			fmt.Println("write file err", err)
			return
		}
	}
}
