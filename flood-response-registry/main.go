// Package main provides the command-line interface for the Shelter Management System.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"flood-response-registry/internal"
)

// reader is shared by the CLI input helpers so buffered input is not recreated
// for each prompt.
var reader = bufio.NewReader(os.Stdin)

// main runs the interactive menu until the user chooses to exit.
func main() {
	for {
		clearScreen()
		showMainMenu()

		switch input("Select option: ") {
		case "1":
			showMobileScreen()

		case "2":
			showCSVScreen()

		case "3":
			showNGOScreen()

		case "4":
			showRandomIntakeScreen()

		case "5":
			showRecordsScreen()

		case "6":
			showUpdateStatusScreen()

		case "7":
			fmt.Println("\nExiting Shelter Management System...")
			return

		default:
			pause("Invalid option.")
		}
	}
}

// ==================================================
// SCREEN MANAGEMENT
// ==================================================

// clearScreen clears the terminal using ANSI escape sequences.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// pause displays a message and waits for the user to press Enter.
func pause(message string) {
	fmt.Println("\n" + message)
	input("Press Enter to continue...")
}

// input displays a prompt, reads one line from standard input, and trims
// surrounding whitespace from the returned value.
func input(label string) string {
	fmt.Print(label)

	value, err := reader.ReadString('\n')
	if err != nil && len(value) == 0 {
		return ""
	}

	return strings.TrimSpace(value)
}

// showHeader prints a consistent title block for CLI screens.
func showHeader(title string) {
	fmt.Println("========================================================")
	fmt.Printf("                 %s\n", title)
	fmt.Println("========================================================")
}

// ==================================================
// MAIN MENU
// ==================================================

// showMainMenu displays the available intake, record-management, and exit actions.
func showMainMenu() {
	showHeader("SHELTER MANAGEMENT SYSTEM")

	fmt.Println(" System : Shelter Intake Simulator")
	fmt.Println(" Date   :", time.Now().Format("02 Jan 2006"))
	fmt.Println("--------------------------------------------------------")
	fmt.Println(" INTAKE SOURCES")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(" [1] Mobile Request")
	fmt.Println(" [2] Legacy CSV")
	fmt.Println(" [3] NGO Partner")
	fmt.Println(" [4] Generate Random Intake")
	fmt.Println(" [5] View All Shelter Records")
	fmt.Println(" [6] Update Shelter Status")
	fmt.Println(" [7] Exit")
	fmt.Println("--------------------------------------------------------")
}

// ==================================================
// SOURCE GENERATORS
// ==================================================

// geneID returns a closure that generates a shelter ID using the supplied
// three-letter source prefix and a random four-digit sequence.
func geneID(prefix string) func() string {
	return func() string {
		min := 1000
		max := 9999

		gen := rand.Intn(max-min+1) + min

		return fmt.Sprintf("%s-%d", prefix, gen)
	}
}

// genCap generates a capacity candidate in the inclusive range -10,000 to
// 10,000. Values outside the accepted business range are intentionally possible
// and are handled by the intake validation process.
func genCap() int {
	return rand.Intn(20_001) - 10_000
}

// generateMobile creates a MobileRequest with a generated ID and int64 capacity.
func generateMobile() internal.MobileRequest {
	mobile := geneID("MBL")

	return internal.MobileRequest{
		ShelterID: mobile(),
		Capacity:  int64(genCap()),
	}
}

// generateCSV creates a LegacyCSVRow with its capacity represented as a string.
func generateCSV() internal.LegacyCSVRow {
	csv := geneID("CSV")

	return internal.LegacyCSVRow{
		ShelterID: csv(),
		Capacity:  strconv.Itoa(genCap()),
	}
}

// generateNGO creates a PartnerPayload with its capacity represented as an int.
func generateNGO() internal.PartnerPayload {
	ngo := geneID("NGO")

	return internal.PartnerPayload{
		ShelterID: ngo(),
		Capacity:  genCap(),
	}
}

// ==================================================
// INTAKE PROCESSING
// ==================================================

// IntakeResult stores the number of generated intake records and the number
// accepted or rejected during a random intake operation.
type IntakeResult struct {
	Received int
	Accepted int
	Rejected int
}

// processIntake sends one source through the internal record-creation flow.
// It prints the outcome and returns true only when the record is accepted.
func processIntake(source internal.CapacitySource, name string) bool {
	record, err := internal.CreatedRecord(source)

	if err != nil {
		fmt.Printf("[REJECTED] %s: %v\n", name, err)
		return false
	}

	fmt.Printf(
		"[OK] %s accepted | ID: %s | Source: %s | Capacity: %d | Status: %s\n",
		name,
		record.ID,
		sourceName(record.Tag),
		record.Cap,
		statusName(record.Status),
	)

	return true
}

// runMobile generates and processes one MobileRequest.
func runMobile() bool {
	source := generateMobile()
	return processIntake(source, "Mobile Request")
}

// runCSV generates and processes one LegacyCSVRow.
func runCSV() bool {
	source := generateCSV()
	return processIntake(source, "Legacy CSV")
}

// runNGO generates and processes one PartnerPayload.
func runNGO() bool {
	source := generateNGO()
	return processIntake(source, "NGO Partner")
}

// ==================================================
// INDIVIDUAL SOURCE SCREENS
// ==================================================

// showMobileScreen previews and processes one generated mobile intake request.
func showMobileScreen() {
	clearScreen()
	showHeader("MOBILE REQUEST")

	source := generateMobile()

	fmt.Printf(" Shelter ID : %s\n", source.ShelterID)
	fmt.Printf(" Capacity   : %d\n", source.Capacity)

	fmt.Println("\n--------------------------------------------------------")
	fmt.Println("Processing intake...")
	fmt.Println("--------------------------------------------------------")

	processIntake(source, "Mobile Request")

	pause("Mobile intake process completed.")
}

// showCSVScreen previews and processes one generated legacy CSV intake row.
func showCSVScreen() {
	clearScreen()
	showHeader("LEGACY CSV")

	source := generateCSV()

	fmt.Printf(" Shelter ID : %s\n", source.ShelterID)
	fmt.Printf(" Capacity   : %q\n", source.Capacity)

	fmt.Println("\n--------------------------------------------------------")
	fmt.Println("Processing intake...")
	fmt.Println("--------------------------------------------------------")

	processIntake(source, "Legacy CSV")

	pause("CSV intake process completed.")
}

// showNGOScreen previews and processes one generated NGO partner payload.
func showNGOScreen() {
	clearScreen()
	showHeader("NGO PARTNER")

	source := generateNGO()

	fmt.Printf(" Shelter ID : %v\n", source.ShelterID)
	fmt.Printf(" Capacity   : %v\n", source.Capacity)

	fmt.Println("\n--------------------------------------------------------")
	fmt.Println("Processing intake...")
	fmt.Println("--------------------------------------------------------")

	processIntake(source, "NGO Partner")

	pause("NGO intake process completed.")
}

// ==================================================
// RANDOM INTAKE SCREEN
// ==================================================

// showRandomIntakeScreen generates one payload per source, processes all three,
// and displays the accepted and rejected totals.
func showRandomIntakeScreen() {
	clearScreen()
	showHeader("RANDOM INTAKE")

	mobile := generateMobile()
	csv := generateCSV()
	ngo := generateNGO()

	fmt.Println("\n[ MOBILE REQUEST ]")
	fmt.Printf(" Shelter ID : %s\n", mobile.ShelterID)
	fmt.Printf(" Capacity   : %d\n", mobile.Capacity)

	fmt.Println("\n[ LEGACY CSV ]")
	fmt.Printf(" Shelter ID : %s\n", csv.ShelterID)
	fmt.Printf(" Capacity   : %q\n", csv.Capacity)

	fmt.Println("\n[ NGO PARTNER ]")
	fmt.Printf(" Shelter ID : %v\n", ngo.ShelterID)
	fmt.Printf(" Capacity   : %v\n", ngo.Capacity)

	fmt.Println("\n--------------------------------------------------------")
	fmt.Println("Processing intake...")
	fmt.Println("--------------------------------------------------------")

	result := IntakeResult{
		Received: 3,
	}

	if processIntake(mobile, "Mobile Request") {
		result.Accepted++
	} else {
		result.Rejected++
	}

	if processIntake(csv, "Legacy CSV") {
		result.Accepted++
	} else {
		result.Rejected++
	}

	if processIntake(ngo, "NGO Partner") {
		result.Accepted++
	} else {
		result.Rejected++
	}

	fmt.Println("\n--------------------------------------------------------")
	fmt.Printf("Total Received : %d\n", result.Received)
	fmt.Printf("Accepted       : %d\n", result.Accepted)
	fmt.Printf("Rejected       : %d\n", result.Rejected)

	pause("Random intake process completed.")
}

// ==================================================
// RECORD MANAGEMENT SCREEN
// ==================================================

// showRecordsScreen displays the shelter registry and opens the selected
// record-management action.
func showRecordsScreen() {
	for {
		clearScreen()
		showHeader("SHELTER RECORDS")

		showAvailableShelters()

		fmt.Println("\n [1] Change Status")
		fmt.Println(" [2] View Details")
		fmt.Println(" [3] Back")

		switch input("\nSelect option: ") {
		case "1":
			showUpdateStatusScreen()

		case "2":
			showRecordDetailsScreen()

		case "3":
			return

		default:
			pause("Invalid option. Choose 1, 2, or 3.")
		}
	}
}

// showAvailableShelters displays the current records in a stable ID order.
// It only reads the registry and does not modify any record.
func showAvailableShelters() {
	fmt.Println("\nAVAILABLE SHELTERS")
	fmt.Println("--------------------------------------------------------")

	if len(internal.Dict) == 0 {
		fmt.Println("No shelter records available.")
		return
	}

	ids := make([]internal.ShelterID, 0, len(internal.Dict))
	for id := range internal.Dict {
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	fmt.Printf("%-15s %-10s %-12s %-12s\n", "SHELTER ID", "SOURCE", "CAPACITY", "STATUS")
	fmt.Println("--------------------------------------------------------")

	for _, id := range ids {
		record := internal.Dict[id]
		fmt.Printf(
			"%-15s %-10s %-12d %-12s\n",
			record.ID,
			sourceName(record.Tag),
			record.Cap,
			statusName(record.Status),
		)
	}

	fmt.Println("--------------------------------------------------------")
	fmt.Printf("Total Records: %d\n", len(internal.Dict))
}

// ==================================================
// RECORD DETAILS
// ==================================================

// showRecordDetailsScreen looks up a shelter by ID and displays its stored fields.
func showRecordDetailsScreen() {
	clearScreen()
	showHeader("SHELTER RECORD DETAILS")

	showAvailableShelters()

	if len(internal.Dict) == 0 {
		pause("There are no records to view.")
		return
	}

	id := internal.ShelterID(input("\nEnter Shelter ID (or 0 to cancel): "))
	if id == "0" {
		return
	}

	record, exists := internal.Dict[id]
	if !exists {
		pause("Record not found. Check the ID in the list and try again.")
		return
	}

	fmt.Println("\nRECORD INFORMATION")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(" Shelter ID :", record.ID)
	fmt.Println(" Source     :", sourceName(record.Tag))
	fmt.Println(" Capacity   :", record.Cap)
	fmt.Println(" Status     :", statusName(record.Status))
	fmt.Println(" Updated    :", record.Update.Format(time.DateTime))
	fmt.Println("--------------------------------------------------------")

	pause("Record details displayed.")
}

// ==================================================
// STATUS MANAGEMENT
// ==================================================

// showUpdateStatusScreen changes a selected record's status. A successful
// change also refreshes the record's Update timestamp; invalid or unchanged
// selections leave the record untouched.
func showUpdateStatusScreen() {
	clearScreen()
	showHeader("UPDATE SHELTER STATUS")

	showAvailableShelters()

	if len(internal.Dict) == 0 {
		pause("There are no records to update.")
		return
	}

	id := internal.ShelterID(input("\nEnter Shelter ID to update (or 0 to cancel): "))
	if id == "0" {
		return
	}

	record, exists := internal.Dict[id]
	if !exists {
		pause("Record not found. No changes were made.")
		return
	}

	fmt.Println("\nSELECTED SHELTER")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(" Shelter ID    :", record.ID)
	fmt.Println(" Source        :", sourceName(record.Tag))
	fmt.Println(" Capacity      :", record.Cap)
	fmt.Println(" Current Status:", statusName(record.Status))
	fmt.Println(" Last Updated  :", record.Update.Format(time.DateTime))
	fmt.Println("--------------------------------------------------------")

	fmt.Println("\nAVAILABLE STATUS")
	fmt.Println(" [1] PENDING")
	fmt.Println(" [2] VERIFIED")
	fmt.Println(" [3] REJECTED")
	fmt.Println("\nUpdate rule:")
	fmt.Println("- Select one of the listed statuses.")
	fmt.Println("- Invalid choices will not modify the record.")
	fmt.Println("- Selecting the current status will leave the record unchanged.")
	fmt.Println("- A successful change updates the record timestamp.")

	choice := input("\nSelect new status (or 0 to cancel): ")
	if choice == "0" {
		pause("Status update cancelled. No changes were made.")
		return
	}

	var newStatus internal.ReportStatus
	switch choice {
	case "1":
		newStatus = internal.PENDING
	case "2":
		newStatus = internal.VERIFIED
	case "3":
		newStatus = internal.REJECTED
	default:
		pause("Invalid status choice. No changes were made.")
		return
	}

	oldStatus := record.Status
	if oldStatus == newStatus {
		fmt.Println("\nNo update required.")
		fmt.Printf("Shelter %s is already %s.\n", record.ID, statusName(record.Status))
		pause("Record remains unchanged.")
		return
	}

	record.Status = newStatus
	record.Update = time.Now()

	fmt.Println("\nSTATUS UPDATE SUCCESSFUL")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(" Shelter ID     :", record.ID)
	fmt.Println(" Source         :", sourceName(record.Tag))
	fmt.Println(" Capacity       :", record.Cap)
	fmt.Println(" Previous Status:", statusName(oldStatus))
	fmt.Println(" New Status     :", statusName(record.Status))
	fmt.Println(" Updated At     :", record.Update.Format(time.DateTime))
	fmt.Println("--------------------------------------------------------")

	pause("Status update completed.")
}

// ==================================================
// HELPERS
// ==================================================

// sourceName converts a source tag into its display label.
func sourceName(tag internal.Tags) string {
	switch tag {
	case internal.MOBILE:
		return "MOBILE"
	case internal.CSV:
		return "CSV"
	case internal.NGO:
		return "NGO"
	default:
		return "UNKNOWN"
	}
}

// statusName converts a report status into its display label.
func statusName(status internal.ReportStatus) string {
	switch status {
	case internal.UNKNOWN:
		return "UNKNOWN"

	case internal.PENDING:
		return "PENDING"

	case internal.VERIFIED:
		return "VERIFIED"

	case internal.REJECTED:
		return "REJECTED"

	default:
		return "UNDEFINED"
	}
}
