package main

import "fmt"

func primaDalamRentang(a int, b int) {
    for a <= b {
        isPrime := true
        if a < 2 {
            isPrime = false
        } else {
            for i := 2; i*i <= a; i++ {
                if a%i == 0 {
                    isPrime = false
                    break
                }
            }
        }
        if isPrime {
            fmt.Print(a, " ")
        }
        a++
    }
}

func main() {
    var x, y int
    fmt.Print("Masukkan rentang (x y): ")
    fmt.Scan(&x, &y)
    primaDalamRentang(x, y)
}
