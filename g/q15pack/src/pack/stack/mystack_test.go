package stack
import "testing"

func TestStack(t *testing.T){
    s := Newstack(7)
    if (!s.Push(2)){
        t.Log("Push fail!")
        t.Fail()
    }
    s.Push(3)
    s.Push(13)
    s.Push(9)
    s.Push(2)
    s.Push(4)
    s.Push(0)
    if (s.top != 6){
        t.Log("top val should be 6! but be %d",s.top)
    }
    if (s.Push(3)){
        t.Log("push element to a full stack, but ret is true!")
        t.Fail()
    }
    
    v,e := s.Pop()
    if (v!=0 || e!=nil){
        t.Log("pop fail!")
        t.Fail()
    }
    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()
    v,e = s.Pop()
    if (v!=2 || e!=nil){
        t.Log("pop fail!")
        t.Fail()
    }
    v,e = s.Pop()
    if (v!=0 || e==nil){
        t.Log("pop fail!")
        t.Fail()
    }
    
}
