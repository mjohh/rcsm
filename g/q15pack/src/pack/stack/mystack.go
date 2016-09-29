/*
    This is a pakage notes sample
*/
package stack
import "fmt"
import "errors"

type stack struct{
   a []int;
   n int;
   top int;
}

// create a stack with max n elements
func Newstack(n int)* stack{
    var s stack
    s.n = n
    s.top = -1
    s.a = make([]int,n,n) 
    return &s
}

// push a int element to stack. It returns 
// true is successful and false otherwise
func (s *stack) Push(e int) bool {
    if s.top < s.n-1 {
        s.top++
        s.a[s.top] = e
        return true
    }
    return false
}

// pop a int element from stack. It returns
// element and nil if successful, otherwise it returns
// 0 and error
func (s *stack) Pop() (int, error) {
    if s.top == -1 {
        return 0,errors.New("empty stack")
    }
    e := s.a[s.top]
    s.top--
    return e, nil
}

// formating printf
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
