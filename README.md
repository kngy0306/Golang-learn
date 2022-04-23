# 参考サイト

# はじめてのGo―シンプルな言語仕様，型システム，並行処理

https://gihyo.jp/dev/feature/01/go_4beginners

## インポート時のオプション指定

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

### json to struct & struct to json

```go
import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID    int `json:"id"`
	Name  string
	Email string
}

func main() {
	person := &Person{
		ID:    1,
		Name:  "kona",
		Email: "aaa@example.com",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	// []byte→string 
	fmt.Println(string(b)) // {"id":1,"Name":"kona","Email":"aaa@example.com"}

	var p Person
	byte := []byte(`{"id":1,"name":"kona","age":5}`)
	e := json.Unmarshal(byte, &p)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(p) // {1 kona }
}
```

### io処理
os.Create()関数にファイル名を渡すと、*os.File構造体へのポインタが取得できる。  
*os.Fileは、io.ReadWriteCloserというインタフェース型であり、Read()，Write()，Close()の3つのメソッドを実装している。

```go
import (
	"log"
	"os"
)

func main() {
	// ファイルの作成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ファイルを閉じる処理
	defer file.Close()

	// Write()は[]byteを引数に取る
	message := []byte("Hello world\n")

	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	// もしくはFprintで出力先を指定する
	// _, err = fmt.Fprint(file, "hello world\n")
}
```

### io/ioutilパッケージで簡単にファイル操作

```go
import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	hello := []byte("Hello world\n")
	err := ioutil.WriteFile("./file.txt", hello, 0666) // 引数(ファイル名, データ, パーミッション)
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll()は引数にio.Readerを渡し、全てを[]byteで返す
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))
}
```

### POST

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // 処理の最後にBodyを閉じる

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil { // エラー処理
			log.Fatal(err)
		}

		// ファイル名を {id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) // ファイルを生成
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}

```

### html/templateパッケージ
Goは，text/templateパッケージとhtml/templateパッケージの2つのテンプレートエンジンが付属しており，どちらも同じインタフェースで使用できる。
html/templateパッケージは、エスケープ処理が自動に行われる。（XSSなどの対策）が容易。

```go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

var t = template.Must(template.ParseFiles("index.html"))

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // 処理の最後にBodyを閉じる

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil { // エラー処理
			log.Fatal(err)
		}

		// ファイル名を {id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) // ファイルを生成
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		// パラメータ取得
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// person生成
		person := Person{
			ID:   id,
			Name: string(b),
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}

```

```html
<body>
  <h1>{{ .ID }} : {{ .Name }}</h1>
</body>
```
