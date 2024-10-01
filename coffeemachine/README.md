# Coffee Vending Machine

## Overview

The **Coffee Vending Machine** is a CLI-based vending machine simulation built in Go. It allows users to select from a variety of coffee options, insert cash, and receive their selected item. The machine tracks inventory, processes payments, and provides feedback to the user on successful purchases or errors (such as insufficient funds or out-of-stock items).

The application is built to be modular and uses Go's concurrency features to manage inventory in a thread-safe manner.

## Features

- **Inventory Management**: 
  - Tracks available items (e.g., different types of coffee) and their quantities.
  - Provides real-time feedback on available stock.
  
- **Payment Processing**:
  - Accepts cash input for items.
  - Verifies the amount inserted against the item price.
  - Provides change when necessary.

- **User Interaction via CLI**:
  - Lists available items.
  - Allows the user to select an item by ID and insert cash.
  - Handles error cases like insufficient funds or out-of-stock items.
  
- **Thread-Safe Inventory**: 
  - Uses a mutex to ensure thread-safe access to the inventory, making it safe for concurrent operations.

## How to Run the Application

### Prerequisites

- **Go (Golang)**: Make sure you have Go installed on your system (version 1.16 or above is recommended). You can download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/coffeemachine.git
   cd coffeemachine
   ```
2. Initialize Go modules:
   ```bash
   go mod tidy
   ```
3. To run the Coffee Vending Machine app, use the following command:
   ```bash
   go run cmd/coffeemachine/main.go
   ```