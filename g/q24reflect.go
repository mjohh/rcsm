package main
import "fmt"
import "reflect"

type Person struct{
    Name string //public member
    age int
}

func Set(i interface{}){
    if i,ok := i.(*Person); ok{
        r := reflect.ValueOf(i)
        r.Elem().Field(0).SetString("Albert Einstein")
    }
}

func main(){
    p := new(Person)
    Set(p)
    fmt.Println(p)
}
