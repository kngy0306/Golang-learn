# A-Tour-of-Go

### インポート時のオプション指定

```go
import (
	f "fmt"     　　　　　　　　　　　　　// f.~で使用可能
	. "strings" 　　　　　　　　　　　　　// stringsを省略可能

	_ "github.com/kngy0306/gosample" // 使用しない場合に_を指定する（コンパイルエラー回避）
)

func main() {
	f.Println(ToUpper("kona"))　　　　 // KONA 
}
```
### reverse関数実装

```go
func reverse(ary []int) []int {
	for i := len(ary)/2 - 1; i >= 0; i-- {
		j := len(ary) - i - 1
		ary[i], ary[j] = ary[j], ary[i]
	}
	return ary
}

func main() {
	ary1 := []int{1, 2, 3, 4, 5, 6}
	ary2 := []int{1, 2, 3, 4, 5}

	reverse(ary1)
	reverse(ary2)

	fmt.Println(ary1, ary2)
	// [6 5 4 3 2 1] [5 4 3 2 1]
}
```
