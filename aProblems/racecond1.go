/*
 * racecond1 - demonstrate a basic race condition.
 * 
 *  - Most straightforward examples: multiple code paths, with shared, read &
 *     write access to a piece of memory, and "comedy ensues".
 *  
 * "go run racecond1.go"
 */

package main

import (
    "fmt"
)

func main() {
    /*
     * First example is non-deterministic and has 3 possible outcomes:
     * 1.  go routine completes and exits resulting in nothing getting 
     *      printed
     * 2.  main completes before the go routine completes, resulting in
     *      "0" getting printed (most likely on most platforms)
     * 3.  Processing for the two paths "interleaves" such that 
     *      after the test for "0", but before the "print" statement runs,
     *      the go routine updates the variable, resulting in "1" getting
     *      printed
     */

     var data int

     go func() {
        data++
    }()

    if data == 0 {
        fmt.Printf("the value of data is %v.\n", data)
    }
}

