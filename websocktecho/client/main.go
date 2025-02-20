package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/coder/websocket"
)

func main() {
	ctx := context.Background()
	conn, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "Client closing connection")

	fmt.Println("Connected to the echo server. Type messages to send (Ctrl+C to exit):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		// Send message to the server
		err = conn.Write(ctx, websocket.MessageText, []byte(message))
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}

		// Read echoed message from the server
		_, data, err := conn.Read(ctx)
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		fmt.Printf("Echoed: %s\n", data)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from input: %v", err)
	}
}
