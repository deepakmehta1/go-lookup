package main

import (
	"coffeemachine/internal/inventory"
	"coffeemachine/internal/machine"
	"coffeemachine/internal/payment"
	"coffeemachine/internal/ui"
	"coffeemachine/pkg/models"
)

func main() {
	// Initialize inventory
	inv := inventory.NewInventory()
	inv.AddItem(&models.Item{ID: "C1", Name: "Espresso", Price: 2.00, Quantity: 10})
	inv.AddItem(&models.Item{ID: "C2", Name: "Americano", Price: 2.50, Quantity: 8})
	inv.AddItem(&models.Item{ID: "C3", Name: "Latte", Price: 3.00, Quantity: 5})
	inv.AddItem(&models.Item{ID: "C4", Name: "Cappuccino", Price: 3.00, Quantity: 5})
	inv.AddItem(&models.Item{ID: "C5", Name: "Mocha", Price: 3.50, Quantity: 4})

	// Initialize payment processor
	pp := payment.NewPaymentProcessor()

	// Initialize vending machine
	vm := machine.NewVendingMachine(inv, pp)

	// Start CLI
	ui.StartCLI(vm)
}
