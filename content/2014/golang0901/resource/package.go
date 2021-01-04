package main

import f "fmt" // 利用するpackageの名前を変えることができる

type Dragon struct {
	Name string
}

// 先頭小文字の関数がprivate, 大文字がpublicになる
func getDragon() Dragon { // HL
	return Dragon{Name: "ヤマタノオロチ"}
}

func main() {
	dragon := getDragon()
	// "fmt" package のpublic method Printlnを呼ぶ
	f.Println(dragon, "") // HL
}
