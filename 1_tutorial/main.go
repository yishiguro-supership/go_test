package main

import (
    "fmt"
    "reflect"
    "myproject/pkg_a"
    "myproject/pkg_b"
)

func main() {
    // import
    fmt.Println("=== import test ===")
    Mod1()
    pkg_a.Mod2()
    pkg_b.Mod3()
    pkg_b.Mod4()
    fmt.Println("")

    // ポインタ
    fmt.Println("=== pointer test ===")
    i := 1
    fmt.Println("type of i: ", reflect.TypeOf(i))
    pi := &i // pi = iのアドレス
    fmt.Println("type of  pi: ", reflect.TypeOf(pi))
    fmt.Println("type of *pi: ", reflect.TypeOf(*pi))
    ppi := &pi // ppi = piのアドレス
    fmt.Println("type of   ppi: ", reflect.TypeOf(ppi))
    fmt.Println("type of  *ppi: ", reflect.TypeOf(*ppi))
    fmt.Println("type of **ppi: ", reflect.TypeOf(**ppi))
    fmt.Println("")

    fmt.Println("&i:", &i) // iのアドレス
    fmt.Println(" i:",  i) // i
    fmt.Println("")

    fmt.Println("&pi:", &pi) // piのアドレス
    fmt.Println(" pi:",  pi) // pi (= iのアドレス)
    fmt.Println("*pi:", *pi) // piが持つアドレスに格納されている値 (= i)
    fmt.Println("")

    fmt.Println(" &ppi:",  &ppi) // ppiのアドレス
    fmt.Println("  ppi:",   ppi) // ppi (= piのアドレス)
    fmt.Println(" *ppi:",  *ppi) // ppiが持つアドレスに格納されている値 (= pi)
    fmt.Println("**ppi:", **ppi) // (= i)
    fmt.Println("")

    // 構造体、関数/メソッド
    fmt.Println("=== struct and functoin/method test ===")
    user01 := User{"Alice", 21}
    user02 := User{age:22, name:"Bob"}
    user03 := new(User)
    user03.name = "Cathy"
    user03.age = 23

    fmt.Println("type of user01: ", reflect.TypeOf(user01))
    fmt.Println("name: ", user01.name, "  age: ", user01.age)
    fmt.Println("type of user02: ", reflect.TypeOf(user02))
    fmt.Println("name: ", user02.name, "  age: ", user02.age)
    fmt.Println("type of user03: ", reflect.TypeOf(user03))
    fmt.Println("name: ", (*user03).name, "  age: ", (*user03).age)
    fmt.Println("name: ", user03.name, "  age: ", user03.age) // ポインタのままでもフィールドにアクセス可能(syntax sugar)
    fmt.Println("")

    fmt.Println("function")
    fmt.Println(GetUserInfoFunction(user01))
    fmt.Println(GetUserInfoFunction(user02))
    fmt.Println(GetUserInfoFunction(*user03))
    //fmt.Println(GetUserInfoFunction(user03)) // 引数の型が違うのでポインタはNG

    fmt.Println("method")
    fmt.Println(user01.GetUserInfoMethod())
    fmt.Println(user02.GetUserInfoMethod())
    fmt.Println((*user03).GetUserInfoMethod())
    fmt.Println(user03.GetUserInfoMethod()) // ポインタのままでもOK
    fmt.Println("")

    // インターフェース
    fmt.Println("=== interface test ===")

    h := Human{"Tom"}
    TalkTo(h)

    d := Dog{"Hachi"}
    TalkTo(d)

    c := Cat{"Jiji"}
    TalkTo(c)

    //p := Plant{"Apple"}
    //TalkTo(p) // pはAnimalインターフェースに沿っていないためTalkTo()の引数にはなりえない
}



type User struct {
    name string
    age int
}

// function
// 第一引数：User型
// 返り値: string型
func GetUserInfoFunction(u User) string {
    return fmt.Sprint("name: ", u.name, "  age: ", u.age)
}

// method
// レシーバ：User型
// 引数：なし
// 返り値: string型
func (u User) GetUserInfoMethod() string {
    return fmt.Sprint("name: ", u.name, "  age: ", u.age)
}


// Animalインターフェース
type Animal interface {
    Speak()
}

func TalkTo(a Animal) {
    fmt.Println("Hello, what's your name?")
    a.Speak()
}

// Human型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Human struct {
    name string
}
func (h Human) Speak() {
    fmt.Println("Hello, my name is", h.name, ".")
}

// Dog型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Dog struct {
    name string
}
func (d Dog) Speak() {
    fmt.Println("Bow wow")
}

// Cat型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Cat struct {
    name string
}
func (c Cat) Speak() {
    fmt.Println("Meow")
}

// Plant型はSpeakメソッドを持たない(=Animalインターフェースに沿っていない)
type Plant struct {
    name string
}
