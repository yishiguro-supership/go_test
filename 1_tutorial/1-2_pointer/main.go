package main

import (
    "fmt"
    "reflect"
)

func main() {
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
}
