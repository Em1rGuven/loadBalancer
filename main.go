package main

// Bu bir Reverse Proxy (Load Balancer).
// Bu sunucuya gelen istekler o an sırası gelen ve uygun olan node'da işlenir ve cevabı da buradan verilir.
func main() {
	server := newServer()
	server.start(8080)
}
