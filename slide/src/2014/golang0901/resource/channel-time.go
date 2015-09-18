select {
case m := <-c:
        handle(m)
case <-time.After(5 * time.Minute): // 5分後に値が返ってくる
        fmt.Println("timed out")
}