package main
import "fmt"

func compare(a,b interface{}) bool{
    switch a:=a.(type){
    case int:
        switch b:=b.(type){
        case int:
            return a>b
        case float64:
            return float64(a)>b
        }
    case float64:
        switch b:=b.(type){
        case int:
            return a>float64(b)
        case float64:
            return a>b
        }
    }
    return false
}


func max(s []interface{}) interface{}{
   m := s[0]
   for _,v:=range s{
       if compare(v,m){
           m = v
       }
   }
   return m
}

func main(){
    fmt.Println("hello")
    s := []interface{}{2,9.3,2,5,8.2,88,9.02}
    m := max(s)
    fmt.Println(m)
}
