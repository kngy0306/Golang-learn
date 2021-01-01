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

### map操作
- 二つ目の戻り値を受け取ると、値が存在するかbool値が返る。
- for文で取り出すとき、順序は保証されない

```go
func main() {
	var month map[int]string = map[int]string{}

	month[1] = "January"
	month[2] = "February"

	v, e := month[1]
	fmt.Println(v, e) // January true

	delete(month, 1)

	v, e = month[1]
	fmt.Println(v, e) // " " false
	
	for key, value := range month {
    		fmt.Printf("%d %s\n", key, value)
	}
}
```

### ポインタ
- **&iでiのポインタを取得**
- ***iでiのポインタが指す先の値を取得**

```go
func callByValue(i int) {
	i++
}

func callByRef(i *int) {
	*i++
}

func main() {
	i := 0
	callByValue(i)
	fmt.Println(i) // 0

	callByRef(&i)
	fmt.Println(i) // 1
}
```

### defer構文
関数スコープのreturnの直後に実行される

```go
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "sample.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	fmt.Println(string(b))
}
```

### コンストラクタ
Newで始まる関数を定義し，その内部で構造体を生成するのが通例

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func main() {
	task := NewTask(1, "write blog")
	fmt.Println(task.Detail) // write blog
}
```

### メソッド
レシーバに紐付けられた関数はメソッドに

```go
func (task *Task) Finish() {
	task.done = true
}

func main() {
	blogTask := NewTask(1, "write blog")
	blogTask.Finish()
	fmt.Println(blogTask.done) // true
}
```

### 型の埋め込み

```go
type User struct {
	FirstName string
	LastName  string
}

func (u *User) PrintFullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

type Task struct {
	ID     int
	Detail string
	done   bool
	*User  // User型の埋め込み
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
	return task
}

func main() {
	task := NewTask(1, "write blog", "mirei", "sasaki")

	fmt.Println(task.FirstName, task.LastName)
	fmt.Println(task.PrintFullName())
	// mirei sasaki
}
```
