package main

import (
	"fmt"
)

// Go does not have classes like java or python
// insetead we can define a struct to store data and method to add behavior 
type Item struct {
	Name     string
	Quantity int
	Price	 float64
}

// this is method to calculate total value of one item 
// TotalValue is the method name
// float64 is the return type of the method
func (i Item) TotalValue() float64 {
	return float64(i.Quantity) * i.Price
}

func main() {
	// create empty map called inventory
	// map is the data type that store key value pair
	inventory := map[string]Item{}

	for {
		fmt.Println("Inventory Manager")
		fmt.Println("1. Add Item")
		fmt.Println("2. Update Item")
		fmt.Println("3. Remove Item")
		fmt.Println("4. View Inventory")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")


		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		}

		switch choice{
		case 1:
			addItem(inventory)
		case 2:
			updateItem(inventory)
		case 3:
			removeItem(inventory)
		case 4:
			viewInventory(inventory)
		case 5:
			fmt.Println("Exiting Inventory Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// addItem function to add new item to the inventory
func addItem(inventory map[string]Item){
	var name string
	var quantity int
	var price float64

	fmt.Print("Item name: ")
	fmt.Scan(&name)
	
	fmt.Print("Quantity: ")
	fmt.Scan(&quantity)

	fmt.Print("Price: ")
	fmt.Scan(&price)

	inventory[name] = Item{Name: name, Quantity: quantity, Price: price}
	fmt.Println("Item added successfully!")
}

// updateItem function to update existing item in the inventory
func updateItem(inventory map[string]Item){
	var name string
	var quantity int
	var price float64

	fmt.Print("Item name to update: ")
	fmt.Scan(&name)

	_, exists := inventory[name]
	if !exists{
		fmt.Println("Item not found.")
		return
	}

	fmt.Print("New Quantity:")
	fmt.Scan(&quantity)
	fmt.Print("New Price:")
	fmt.Scan(&price)

	inventory[name] = Item{Name: name, Quantity: quantity, Price: price}
	fmt.Println("Item updated successfully!")
}

// removeItem function to remove item from the inventory
func removeItem(inventory map[string]Item){
	var name string

	fmt.Print("Item name to remove: ")
	fmt.Scan(&name)

	_,exists := inventory[name]
	if !exists{
		fmt.Println("Item not found.")
		return
	}

	// delete is the built in function to remove key value pair from map
	delete(inventory, name)
	fmt.Println("Item removed successfully!")
}

// viewInventory function to display all items in the inventory
func viewInventory(inventory map[string]Item){
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("Current Inventory:")
	// _ for the ignore the key
	for _, item := range inventory {
		fmt.Printf("- %s: Quantity: %d, Price: $%.2f, Total Value: $%.2f\n", item.Name, item.Quantity, item.Price, item.TotalValue())
	}
}