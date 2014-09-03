package main

import "fmt"

func main() {
	var m1 string = "Hello world!" // 変数は 変数名 型の順番で宣言
	m2 := "Hello Go!"              // := で型推論
	fmt.Println(m1)
	fmt.Println(m2)
}
