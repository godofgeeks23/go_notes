package main

func main() {
	// We can call sayHello directly because it's in the same package
	// for this to work - the module must be initialized with "go mod init <name>" and run using "go run ."
	sayHello("Alice")
}
