package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"coffeemachine/internal/machine"
)

// StartCLI launches the command-line interface.
func StartCLI(vm *machine.VendingMachine) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Available items:")
		items := vm.GetAvailableItems()
		for _, item := range items {
			fmt.Printf("ID: %s, Name: %s, Price: %.2f, Quantity: %d\n",
				item.ID, item.Name, item.Price, item.Quantity)
		}

		fmt.Print("Enter item ID to purchase (or 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		fmt.Print("Insert cash amount: ")
		scanner.Scan()
		cashInput := scanner.Text()
		cash, err := strconv.ParseFloat(cashInput, 64)
		if err != nil {
			fmt.Println("Invalid cash amount.")
			continue
		}

		err = vm.PurchaseItem(input, cash)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			refund := vm.PaymentProcessor.Refund()
			if refund > 0 {
				fmt.Printf("Refunded: %.2f\n", refund)
			}
		} else {
			refund := vm.PaymentProcessor.Refund()
			if refund > 0 {
				fmt.Printf("Please collect your change: %.2f\n", refund)
			}
		}

		fmt.Println()
	}
}
