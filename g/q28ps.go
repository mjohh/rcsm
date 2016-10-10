package main
import(
    "bytes"
    "log"
    "os/exec"
    "strings"
    "fmt"
)



func main(){
    cmd := exec.Command("ps","-e", "-opid ppid")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil{
       log.Fatal(err)
    }

    // to store ppid and it's children pid
    processes:=make(map[string][]string)
  
    // discard title
    out.ReadString('\n')
    lines := make([][]string,0)
    for {
        line,err := out.ReadString('\n')
        if err != nil {
            break;
        }
        tokens := strings.Split(line, " ")
        l := make([]string,0)
        for _,v := range tokens{
            if(v!="" && v!="\t"){
                if v[len(v)-1]=='\t'||v[len(v)-1]=='\n'{
                    l = append(l,v[:len(v)-1])
                }else{
                    l = append(l,v)
                }
            }
        }
        lines = append(lines, l)
    }

    for _,v := range lines{
        ppid := v[1]
        children,ok := processes[ppid]
        // this ppid has been handled in the past
        if ok {
            continue
        }
        // create children slice for new ppid
        children = make([]string,0)
        for _,v2 := range lines{
            pid := v2[0]
            if ppid== v2[1]{
                children = append(children,pid)
            }
        }
        processes[ppid] = children
    }
    
    for key,val := range processes{
        fmt.Printf("Pid %s has %d children:%v\n",key,len(val),val)
    }
    
}
