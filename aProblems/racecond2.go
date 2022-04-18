/*
 * racecond2 - Attempt to "fix" the basic race condition via a delay.
 *
 *  This "appears" to work (the result is "generally" NO output), but
 *   it is a placebo, inefficient, and leaves the door open for 
 *   trouble down the road if/when someone adds functionality to the go routine.
 *  
 * "go run racecond2.go"
 */

package main

import (
    "fmt"
    "time"
)

func main() {
    /*
     * Introduce a one second "sleep" in the main routine to give
     *  the go routine a chance to run before we exit.
     *  Looks good on paper, and "seems" to work, but you are still
     *  depending on non-deterministic behaviour:  you have no idea
     *  what else is going on in the system.
     * (Also note this introduces the static inefficiency of ALWAYS pausing
     *  one of the threads of execution)
     *
     * Some old-timers might tend towards this trick, as it is reminicent
     *  of raising the dispatching priority of an I/O bound execution path
     *  to prevent compute bound paths from "starving" the I/O bound paths.
     *  Great logic on an IBM 370 in 1985, but even if it were still valid
     *  we are dealing with two compute paths here!
     */

     var data int

     go func() {
        data++
    }()

    time.Sleep(1*time.Second)
    if data == 0 {
        fmt.Printf("the value of data is %v\n", data)
    }

    if data == 0 {
        fmt.Printf("the value of data is %v.\n", data)
    }
}

