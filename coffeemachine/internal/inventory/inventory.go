package inventory

import (
	"coffeemachine/pkg/models"
	"sync"
)

// Inventory manages the items in the vending machine.
type Inventory struct {
	mu    sync.Mutex              // Mutex to ensure thread-safe access to inventory
	items map[string]*models.Item // Map to store items by their ID
}

// NewInventory creates a new Inventory instance with an initialized item map.
func NewInventory() *Inventory {
	return &Inventory{
		items: make(map[string]*models.Item), // Initialize the item map
	}
}

// AddItem adds a new item to the inventory.
func (inv *Inventory) AddItem(item *models.Item) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.items[item.ID] = item // Add or update item in the map
}

// GetItem retrieves an item from the inventory by its ID.
// It returns the item and a boolean indicating if the item was found.
func (inv *Inventory) GetItem(id string) (*models.Item, bool) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	item, exists := inv.items[id]
	return item, exists
}

// DeductItem decreases the quantity of an item by 1.
// It returns true if the deduction was successful, false otherwise (e.g., if the item is out of stock).
func (inv *Inventory) DeductItem(id string) bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	// Check if the item exists and if there is at least 1 in stock
	if item, exists := inv.items[id]; exists && item.Quantity > 0 {
		item.Quantity--
		return true
	}
	return false
}

// GetAllAvailableItems returns a slice of items that have a quantity greater than 0.
func (inv *Inventory) GetAllAvailableItems() []*models.Item {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	availableItems := []*models.Item{}
	for _, item := range inv.items {
		if item.Quantity > 0 {
			availableItems = append(availableItems, item)
		}
	}
	return availableItems
}
