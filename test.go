//https://www.hackerrank.com/contests/visa-codeurway-2017/challenges/visa-and-string-encryption

package main
import (
     "fmt"
     "bufio"
    "strconv"
    "strings"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var x int
    for scanner.Scan() {
        if x==0{
            x,_= strconv.Atoi(scanner.Text())  
            if x<=1 || x>=100{
                return
            }
            break
        }
        for x>0{
            var i int
            var v int16
            var c int16
            text:=strings.Split(scanner.Text()," ") // split string and window size
            y,_ := strconv.Atoi(text[1]) //read window size
            runes:=[]rune(text[0]) // string to char            
            if len(text[0])<=1 && len(text[0])>=10000{
                return
            }
            for y<len(text[0]){
                for i < y {
                    switch runes[i] {
                        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
                           v++     
                        default:
                            c++
                    }
                    fmt.Printf("%d\n",v*c)
                    i++                
                }
                i--
                y=y+i
            }
            fmt.Println("")
            x--
            
        }
    }    
}
