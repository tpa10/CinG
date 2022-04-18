/*
 * access-sync.go - demonstrate "a" method around race conditions using a "sync" mutex
 *
 *  Not necessarily the "best" way to solve this problem, but it does make processing
 *  more dereministic.
 *
 *  Note that this ONLY ensures that our code paths don't step on each other's toes.
 *  It does NOTHING to ensure any kind of deterministic behavior.
 *  
 *  Also note that now everyone has to "remember" to surround references to the "data"
 *  variable with calls to "lock" and "unlock".
 *
 *  "go run access-sync.go"
 */

package main

import (
    "fmt"
    "sync"
)

func main() {
    
    var myMutex sync.Mutex
    var data int

    go func() {
        myMutex.Lock() // gain the mutex or block until we can get it.
        data++
        myMutex.Unlock()   // release the mutex for use by others.
    }()

    myMutex.Lock()   // gain the mutex or block until we can get it.
    if data == 0 {
        fmt.Printf("the value of data is %v\n", data)
    } else {
        fmt.Printf("the value of data is %v\n", data)
    }
    myMutex.Unlock()   // release to mutex for use by others.
}
