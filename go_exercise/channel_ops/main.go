package main

import "time"

func fn() string {
	time.Sleep(5*time.Second)
	return "xxxx"
}

func finshReq(time1 time.Duration) (ob string){
	ch := make(chan string,1)
	go func() {
		result := fn()
		ch <- result
	}()
	for  {
		select {
		case result := <- ch:
			println(result)
			//return result
		case <- time.After(time1*time.Second):
			println("---")
			//return ""
		}
	}

}

func main()  {
	res := finshReq(2)
	println("***************",res)
}


