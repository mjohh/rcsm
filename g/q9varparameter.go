package main

import "fmt"

func print(a ...int){
    for _,v := range a {
        fmt.Println(v)
    }
}

func main(){
    print(1,2,3,5,3,9,10,1)
}
