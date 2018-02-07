package main

import ()
import (
	"fmt"
	"time"
)

//func Producer (queue chan<- int){
//	for i:= 0; i < 10; i++ {
//		queue <- i
//		fmt.Println("produce:",i)
//		time.Sleep(time.Second)
//	}
//}
//
//func Consumer( queue <-chan int){
//	for i :=0; i < 10; i++{
//		v := <- queue
//		fmt.Println("receive:", v)
//		time.Sleep(time.Second)
//	}
//}
//
//func main(){
//	queue := make(chan int, 1)
//	go Producer(queue)
//	go Consumer(queue)
//	time.Sleep(time.Second*100) //让Producer与Consumer完成
//}

//func main() {
//	ch := make(chan int, 2)
//	go test(ch)
//	for i := range ch {
//		fmt.Println(i)
//	}
//}
//
//func test(ch chan int) {
//	ch <- 1
//	ch <- 2
//	ch <- 3
//	close(ch)
//}

//func main() {
//	var N = 100
//	type Empty interface{}
//	var empty Empty
//	data := make([]float64, N)
//	for i := range data {
//		data[i] = float64(i)
//	}
//	res := make([]string, N)
//	sem := make(chan Empty, N)
//	for i, xi := range data {
//		//为匿名函数传参，避免并行协程i和xi并发冲突
//		go func(i int, xi float64) {
//			res[i] = fmt.Sprintf("%d-%f", i, xi)
//			sem <- empty
//		}(i, xi)
//		//闭包引用，会有并发冲突
//		//go func() {
//		//res[i] = fmt.Sprintf("%d-%f", i, xi)
//		//sem <- empty
//		//}()
//	}
//	// wait for goroutines to finish
//	for i := 0; i < N; i++ {
//		<-sem
//	}
//	for _, v := range res {
//		fmt.Println(v)
//	}
//}

//append and slice
//func main(){
//
//	l1 := []int{1,2,3}
//	fmt.Println(len(l1),cap(l1))
//
//	l2 := []int{4,5,6}
//	l1 =append(l1,l2...)
//	fmt.Println(l1)
//	fmt.Println(len(l1),cap(l1))
//	fmt.Println(len(l2),cap(l2))
//	l3 := l1[2:5]
//	fmt.Println(len(l3),cap(l3))
//	l3 = append(l3, 7)
//	fmt.Println(len(l3),cap(l3))
//	fmt.Println("-----")
//
//	fmt.Println(l1,l2,l3)
//
//	l4 := make([]int,3)
//	copy(l4, l1[1:4])
//
//	l4[0] = 0
//	fmt.Println(l1,l4)
//}

func main() {
	go ticker()
	time.Sleep(time.Hour * 1)
}

func ticker() {
	tick := time.Tick(time.Minute * 1)
	for {
		select {
		case <-tick:
			fmt.Println(time.Now())
		}
	}
}
