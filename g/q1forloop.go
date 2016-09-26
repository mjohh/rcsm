package main
import "fmt"

func forloop(){
    
    for i:=0; i<10; i++ {
        fmt.Println(i)
    }
}

func gotoloop(){

    i := 0
loop:
   fmt.Println(i)
   if i++; i< 10{
      goto loop
   }

}

func sliceloop(){
    var s [10]int
    for i:=0; i<10; i++{
        fmt.Println(s[i])
    }
    fmt.Println("use len()")
    for i:=0; i<len(s); i++{
        fmt.Println(s[i])
    }
    fmt.Println("use range")
    for i,v:=  range s{
        fmt.Println(i,v);
    }
    fmt.Println("initialized array")
    s1 := []int{1,2,3,4,5,6,7,8,9,10}
    for _,v := range s1{
        fmt.Println(v);
    }
    fmt.Println("make slice")
    s2 := make([]int, 10)
    for i,v := range s2{
        fmt.Println(i,v);
    }
}

func main() {
    forloop()
    fmt.Println("---------")
    gotoloop()
    fmt.Println("---------")
    sliceloop()
}
