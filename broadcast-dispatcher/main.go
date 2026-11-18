package main

import (
	"fmt"
	"log"

	dispatcher "broadcast-dispatcher/internal"
)

func main() {
	printHeader("REGIONAL EMERGENCY BROADCAST DISPATCHER")

	// 1. CHANNEL CONFIGURATION
	printSection("CHANNEL CONFIGURATION")

	channels := dispatcher.Channel{
		true,  // SMS
		true,  // Cell Broadcast
		false, // Radio
		true,  // Satellite
	}

	fmt.Println("Initial Configuration:")
	printChannels(channels)

	updatedChannels, err := dispatcher.Config(channels, 2, true)
	if err != nil {
		log.Fatal("Channel configuration failed:", err)
	}

	channels = updatedChannels

	fmt.Println("\nUpdated Configuration:")
	printChannels(channels)

	// 2. RECIPIENT QUEUE
	printSection("RECIPIENT QUEUE")

	queue := dispatcher.Queue{
		"REG-001",
		"REG-002",
		"REG-003",
	}

	incoming := dispatcher.Incoming{
		"REG-004",
		"REG-005",
	}

	fmt.Printf("Initial Queue : %v\n", queue)
	fmt.Printf("Incoming      : %v\n", incoming)

	queue = dispatcher.Recipients(queue, incoming)

	fmt.Printf("Updated Queue : %v\n", queue)
	fmt.Printf("Queue Length  : %d\n", len(queue))

	// 3. LIVE PAGE VIEW
	printSection("LIVE PAGE VIEW")

	start, end := 0, 3

	view, err := dispatcher.View(queue, start, end)
	if err != nil {
		log.Fatal("Page view failed:", err)
	}

	fmt.Printf("Page Range : [%d:%d]\n", start, end)
	fmt.Printf("View       : %v\n", view)
	fmt.Printf("Length     : %d\n", len(view))
	fmt.Printf("Capacity   : %d\n", cap(view))

	// 4. SNAPSHOT
	printSection("SNAPSHOT")

	snapshot, err := dispatcher.Audit(queue, start, end)
	if err != nil {
		log.Fatal("Snapshot creation failed:", err)
	}

	fmt.Printf("Snapshot : %v\n", snapshot)
	fmt.Printf("Length   : %d\n", len(snapshot))
	fmt.Printf("Capacity : %d\n", cap(snapshot))

	// 5. BACKING ARRAY DEMONSTRATION
	printSection("BACKING ARRAY DEMONSTRATION")

	fmt.Println("Before modification:")
	fmt.Printf("Queue    : %v\n", queue)
	fmt.Printf("View     : %v\n", view)
	fmt.Printf("Snapshot : %v\n", snapshot)

	fmt.Println("\nModifying the first element through View...")

	view[0] = "REG-001-UPDATED"

	fmt.Printf("Queue    : %v\n", queue)
	fmt.Printf("View     : %v\n", view)
	fmt.Printf("Snapshot : %v\n", snapshot)

	// 6. APPEND CAPACITY DEMONSTRATION
	printSection("APPEND CAPACITY DEMONSTRATION")

	fmt.Println("Appending a new recipient to View...")

	view = append(view, "REG-006")

	fmt.Printf("Queue      : %v\n", queue)
	fmt.Printf("View       : %v\n", view)
	fmt.Printf("View Length: %d\n", len(view))
	fmt.Printf("View Cap   : %d\n", cap(view))

	printFooter()
}

func printHeader(title string) {
	fmt.Println("==================================================")
	fmt.Printf("       %s\n", title)
	fmt.Println("==================================================")
}

func printSection(title string) {
	fmt.Println("\n--------------------------------------------------")
	fmt.Printf(" %s\n", title)
	fmt.Println("--------------------------------------------------")
}

func printChannels(channels dispatcher.Channel) {
	names := [...]string{
		"SMS",
		"Cell Broadcast",
		"Radio",
		"Satellite",
	}

	for i, name := range names {
		status := "DISABLED"

		if channels[i] {
			status = "ENABLED"
		}

		fmt.Printf("  %-18s : %s\n", name, status)
	}
}

func printFooter() {
	fmt.Println("\n==================================================")
	fmt.Println("             SIMULATION COMPLETED")
	fmt.Println("==================================================")
}
