//
// deadlock.go - deliberately force/create a deadlock situation.
//  Just so we know what one looks like and why.
//
//  Logic:
//      1. Spin off a couple of go routines that lock a common mutex and put
//      themselves to sleep while holding the resource.
//      2. main() goes to sleep waiting on "done" notifications from the go 
//          routines. 
//      
//  Net:  we have 3 go routines, none of which are dispatchable.
//
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// value encapsulates the data and its lock
type value struct {
    mu      sync.Mutex
    value   int
} 

func main() {
    var wg sync.WaitGroup 

    fmt.Println("Version", runtime.Version())
    fmt.Println("NumCPU", runtime.NumCPU())
    fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
    fmt.Println("main: NumGoroutine: ", runtime.NumGoroutine())

    printSum := func(v1, v2 *value, tag string) {

        defer fmt.Printf("%s: terminating, NumGoroutines: %d\n", 
            tag, runtime.NumGoroutine())
        defer wg.Done()
        defer fmt.Printf("%s - posting done to wait group\n", tag)

        /**********************************
        // Specifically yield to See if this helps force parallelism 
        runtime.Gosched()
        **********************************/

        fmt.Printf("%s entered. NumGoroutine: %d\n", tag, runtime.NumGoroutine())

        v1.mu.Lock()

        /*************************
        fmt.Printf("%s - now holds v1 mutex\n", tag)
        fmt.Printf("%v\n",v1)
        **************************/

        /******************************
        defer fmt.Printf("%s - unlocked v1 mutex\n", tag)
        defer fmt.Printf("%v\n",v1)
        *******************************/
        defer v1.mu.Unlock()

        time.Sleep(2*time.Second)
        v2.mu.Lock()
        
        /*******************************
        fmt.Printf("%s - now holds v2 mutex\n", tag)
        fmt.Printf("%v\n",v2)
        *******************************/

        /******************************
        defer fmt.Printf("%s - unlocked v2 mutex\n", tag)
        defer fmt.Printf("%v\n",v2)
        ******************************/

        defer v2.mu.Unlock()
        fmt.Printf("sum = %v\n", v1.value + v2.value)
    }

    var a, b value
    a.value = 1
    b.value = 2

    wg.Add(4)
    go printSum(&a, &b, "go1")
    go printSum(&a, &b, "go2")
    go printSum(&a, &b, "go3")
    go printSum(&a, &b, "go4")
    fmt.Println("main: waiting on wait group")
    wg.Wait()

}

