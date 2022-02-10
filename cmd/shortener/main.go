package main

func main() {
	httpHandler := GetHttpHandler()
	runServer(httpHandler)
}
