package main

import "sync"
import "fmt"
import "time"

type Object struct {
    //data
}

func (obj *Object) Update(wg *sync.WaitGroup) {
    //update data
    time.Sleep(time.Second)
    fmt.Println("Update done")
    wg.Done()
    return
}

func main() {
    var wg sync.WaitGroup
    list := make([]Object, 5)
    for {
        for _, object := range list {
            wg.Add(1)
            go object.Update(&wg)
        }
        //now everything has been updated. start again
        wg.Wait()
        fmt.Println("Group done")
    }
}
