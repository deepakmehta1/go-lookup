package machine

import (
	"coffeemachine/internal/inventory"
	"coffeemachine/internal/payment"
	"coffeemachine/pkg/models"
	"fmt"
)

// VendingMachine represents the vending machine
type VendingMachine struct {
	Inventory        *inventory.Inventory
	PaymentProcessor *payment.PaymentProcessor
}

// NewVendingMachine creates a new vending machine instance.
func NewVendingMachine(inv *inventory.Inventory, pp *payment.PaymentProcessor) *VendingMachine {
	return &VendingMachine{
		Inventory:        inv,
		PaymentProcessor: pp,
	}
}

// PurchaseItem handles the purchase flow.
func (vm *VendingMachine) PurchaseItem(itemID string, cash float64) error {
	// Accept cash
	vm.PaymentProcessor.AcceptCash(cash)

	// Retrieve item
	item, exists := vm.Inventory.GetItem(itemID)
	if !exists {
		return fmt.Errorf("item not available")
	}

	// Verify payment
	if err := vm.PaymentProcessor.VerifyAmount(item.Price); err != nil {
		return err
	}

	// Deduct item from inventory
	if !vm.Inventory.DeductItem(itemID) {
		return fmt.Errorf("item out of stock")
	}

	fmt.Printf("Dispensing item: %s\n", item.Name)
	return nil
}

// GetAvailableItems returns a list of all items currently available in the vending machine.
func (vm *VendingMachine) GetAvailableItems() []*models.Item {
	return vm.Inventory.GetAllAvailableItems() // Call the method from the Inventory struct
}
