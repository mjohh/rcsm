package main
import "fmt"
import "errors"
import "reflect"

func sort(a interface{}, b interface{}) (error, interface{}, interface{}){
    if reflect.TypeOf(a) != reflect.TypeOf(b){
         return errors.New("different types could not compare"),nil,nil
    }


    switch a:=a.(type){
    case int:
        fmt.Println("case int1")
        switch b:=b.(type){
        case int:
            fmt.Println("case int2")
            if a > b{
                return nil, b, a
            }else{
                return nil, a, b
                
            }
        }
    }
    return errors.New("non-val type could not compare!"),nil,nil
}

func main(){
    r,a,b:=sort(9,2)
    fmt.Println(r,a,b)
}
