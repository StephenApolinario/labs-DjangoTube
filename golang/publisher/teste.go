func main() {`     **`go`** `fmt.Println("Hello from another goroutine")     fmt.Println("Hello from the main goroutine")      // At this point the program execution stops and all     // active goroutines are killed. }