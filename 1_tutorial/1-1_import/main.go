package main

import (
    "fmt"
    "myproject/pkg_a"
    "myproject/pkg_b"
)

func main() {
    fmt.Println("=== import test ===")
    Mod1()
    pkg_a.Mod2()
    pkg_b.Mod3()
    pkg_b.Mod4()
}
