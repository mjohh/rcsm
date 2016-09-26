package main
import "fmt"


func printtriangle(){
    j:=1
    i:=0
    for {
        for k:=0; k<j; k++{
            fmt.Print("A")
            if i++; i==100{
                return    
            }
        }
        fmt.Print("\n")
        j++
    }
}


func printtriangle2(){
    str := "A"
    for i:=0; i<100; i++{
        fmt.Println(str)
        str += "A"
    }
}

func main() {
    printtriangle()
    fmt.Println("-----------------------")
    printtriangle2()
}
