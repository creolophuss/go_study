package main
import (
    "fmt"
    "reflect"
)
//longest common subsequence
func main(){
    const a string = "xyzabjkcdfgh"
    const b string = "lkjhacbopcdmh"
    
    const m = len(a)
    const n = len(b)


    fmt.Println(reflect.TypeOf(a[0]))

    var x = [m+1][n+1]int{}

    for i := 0; i <=m; i++ {
        x[i][0] = 0
    }

    for i := 0; i <=n; i++ {
        x[0][i] = 0
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++{
            if a[i-1] == b[j-1] {
                x[i][j] = x[i-1][j-1] + 1
            } else {
                if x[i][j-1] > x[i-1][j] {
                    x[i][j] = x[i][j-1]
                } else {
                    x[i][j] = x[i-1][j]
                }
            }
        }
    }

    l := x[m][n]
    t := l
    s := make([]byte, t)

    for i, j := m, n; i > 0 && j > 0;  {
            if a[i-1] == b[j-1] {
                s[t-1] = byte(a[i-1])
                t--
                i--
                j--
            } else {
                if x[i][j-1] > x[i-1][j] {
                    j--
                } else {
                    i--
                }
            }
        }

    fmt.Println(string(s[0:l]))
}
