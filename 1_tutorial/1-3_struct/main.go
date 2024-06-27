package main

import (
    "fmt"
    "reflect"
)

func main() {
    fmt.Println("=== struct and functoin/method test ===")
    user01 := User{"Alice", 21}
    user02 := User{age:22, name:"Bob"}
    user03 := new(User)
    user03.name = "Cathy"
    user03.age = 23

    fmt.Println("type of user01: ", reflect.TypeOf(user01))
    fmt.Println("name: ", user01.name, "  age: ", user01.age)
    fmt.Println("")

    fmt.Println("type of user02: ", reflect.TypeOf(user02))
    fmt.Println("name: ", user02.name, "  age: ", user02.age)
    fmt.Println("")

    fmt.Println("type of user03: ", reflect.TypeOf(user03))
    fmt.Println("name: ", (*user03).name, "  age: ", (*user03).age)
    fmt.Println("name: ", user03.name, "  age: ", user03.age) // ポインタのままでもフィールドにアクセス可能(syntax sugar)
    fmt.Println("")

    fmt.Println("function")
    fmt.Println(GetUserInfoFunction(user01))
    fmt.Println(GetUserInfoFunction(user02))
    //fmt.Println(GetUserInfoFunction(user03)) // 引数の型が違うのでポインタはNG
    fmt.Println(GetUserInfoFunction(*user03))
    fmt.Println("")

    fmt.Println("method")
    fmt.Println(user01.GetUserInfoMethod())
    fmt.Println(user02.GetUserInfoMethod())
    fmt.Println(user03.GetUserInfoMethod()) // ポインタのままでもOK
    fmt.Println((*user03).GetUserInfoMethod())
    fmt.Println("")

    user01.GrowOlder1()
    user01.GrowOlder2() // ポインタレシーバでも値型から呼べる(syntax sugar)
    user03.GrowOlder1() // 値レシーバでもポインタ型から呼べる(syntax sugar)
    user03.GrowOlder2()
    fmt.Println("")

    // 値レシーバの場合、レシーバの値は更新不可
    user02.GrowOlder1()
    user02.GrowOlder1()
    user02.GrowOlder1()
    user02.GrowOlder1()
    user02.GrowOlder1()
    fmt.Println("")

    // ポインタレシーバの場合、レシーバの値は更新可能
    user02.GrowOlder2()
    user02.GrowOlder2()
    user02.GrowOlder2()
    user02.GrowOlder2()
    user02.GrowOlder2()
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

// 値レシーバ
// uをコピーして関数内で利用する(関数内のu ≠ レシーバのu)
// uを更新をしてもレシーバの値は変わらない
func (u User) GrowOlder1() {
    beforeAge := u.age
    u.age++
    afterAge := u.age
    fmt.Println("name:", u.name, "age:", beforeAge, "→", afterAge)
}

// ポインタレシーバ
// uを参照して関数内で利用する(関数内のu = レシーバのu)
// uを更新した場合レシーバの値も変わる
func (u *User) GrowOlder2() {
    beforeAge := u.age
    u.age++
    afterAge := u.age
    fmt.Println("name:", u.name, "age:", beforeAge, "→", afterAge)
}
