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
