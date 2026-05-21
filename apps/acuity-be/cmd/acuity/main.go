// cmd/acuity/main.go
package main

import (
	"Acuity/internal/infrastructure"
	"Acuity/internal/transport"
)

func main() {
	classifier := infrastructure.NewClassifier()
	server := transport.NewServer(classifier)
	server.Start()
}
