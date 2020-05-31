package main

import "sync"

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
    go func() {
        res1 := make(chan int, n)
        res2 := make(chan int, n)
        wg := new(sync.WaitGroup)
        wg.Add(2 * n)
        go func() {
            for i := 0; i < n; i++ {
                res1 <- f(<-in1)
                wg.Done()
            }
        }()
        go func() {
	    for i := 0; i < n; i++ {
                res2 <- f(<-in2)
                wg.Done()
            }
        }()
        wg.Wait()
        for a := range res1 {
            out <- a + <-res2
        }
    }()
}
