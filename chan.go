package main

import "os"
import "strconv"

func process(c chan int, out chan []byte) {
    for i := range c {
        out <- []byte(strconv.Itoa(i))
    }
}

func write(out chan []byte, f *os.File) {
    for b := range out {
        f.Write(b)
    }
}

func main() {
    c := make(chan int)
    out := make(chan []byte)
    f, err := os.Create("test.txt")
    if err != nil {
        panic(err)
    }
    go process(c, out)
    go write(out, f)

    for i := range make([]int, 10) {
        c <- i
    }
    c <- 1001
    close(c)
}
