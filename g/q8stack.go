package main
import "fmt"
import "errors"

type stack struct{
   a []int;
   n int;
   top int;
}

func newstack(n int)* stack{
    var s stack
    s.n = n
    s.top = -1
    s.a = make([]int,n,n) 
    return &s
}

func (s *stack) push(e int) bool {
    if s.top < s.n-1 {
        s.top++
        s.a[s.top] = e
        return true
    }
    return false
}

func (s *stack) pop() (int, error) {
    if s.top == -1 {
        return 0,errors.New("empty stack")
    }
    e := s.a[s.top]
    s.top--
    return e, nil
}

func (s *stack) String() string {
    var str string
    var a string
    for i:=0; i<=s.top; i++{
        a = fmt.Sprintf("%d",s.a[i])
        if i==0 {
            str += a
        }else{
            str += " " + a
        }
    }
    str += "\n"
    return str
}
/*
func (s *stack) String() string {
    var str string
    str = fmt.Sprintf("%d", s.a.([]interface{}))
}
*/
func main(){
    fmt.Println("ok");
    s := newstack(10);
    var i int
    for i= 0; i<10; i++ {
        r := s.push(i)
        if r {
            fmt.Println("push ",i," ok!")
        }else{
            fmt.Println("push ",i," fail!")
        }
    }
    r := s.push(i)
    if r {
        fmt.Println("push ",i," ok!")
    }else{
        fmt.Println("push ",i," fail!")
    }
    fmt.Printf("stack = %v", s)

    for i:=0; i<10; i++ {
        a,err := s.pop()
        if err==nil {
            fmt.Println("pop ", a, "  ok!")
        }else{
            fmt.Println("pop fail, err=",err.Error())
        }
    }
    a,err := s.pop()
    if err==nil {
        fmt.Println("pop ", a, "  ok!")
    }else{
        fmt.Println("pop fail, err=",err.Error())
    }

    fmt.Printf("stack = %v", s)

}
