package main

func main() {
	storage := &GlobalURLs
	httpHandler := GetHttpHandler(storage)
	runServer(httpHandler)
}
