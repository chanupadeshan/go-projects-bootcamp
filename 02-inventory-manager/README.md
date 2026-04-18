# 📦 Inventory Manager (Go CLI Project)

This is a simple **Inventory Management System** built using Go. It runs in the terminal and allows users to manage items with basic operations like adding, updating, deleting, and viewing inventory.

---

## 🚀 Running the Program

```bash
go run main.go
```

---

## 🧩 Features

* ➕ Add Item
* ✏️ Update Item
* ❌ Remove Item
* 📋 View Inventory
* 🚪 Exit Program

---

## 🧠 Core Concepts Covered

| Concept      | Description                                |
| ------------ | ------------------------------------------ |
| Struct       | Custom data type to store item details     |
| Map          | Key-value storage for inventory            |
| Functions    | Separate logic for each operation          |
| Methods      | Attached behavior to struct (`TotalValue`) |
| Loop         | Menu runs continuously                     |
| Switch       | Handles user choices                       |
| Input/Output | Interaction with user                      |

---

## 📦 Data Structure

### Item Struct

```go
type Item struct {
Name     string
Quantity int
Price    float64
}
```

### Inventory Storage

```go
inventory := map[string]Item{}
```

* **Key** → Item Name (`string`)
* **Value** → Item struct

---

## 🔢 Method: Total Value

```go
func (i Item) TotalValue() float64 {
return float64(i.Quantity) * i.Price
}
```

### Purpose

Calculates total value of an item:

```
Total Value = Quantity × Price
```

---

## 🔄 Program Flow

1. Start program (`main()`)
2. Initialize empty inventory
3. Display menu
4. Take user input
5. Execute selected operation
6. Repeat until user exits

---

## 🧪 Function Breakdown

### 1. Add Item

* Takes name, quantity, price
* Stores item in map

### 2. Update Item

* Checks if item exists
* Updates quantity and price

### 3. Remove Item

* Checks if item exists
* Deletes item from map

### 4. View Inventory

* Displays all items
* Shows total value using method

---

## 🧾 Example Output

```
Current Inventory:
- Pen: Quantity: 10, Price: $25.50, Total Value: $255.00
- Book: Quantity: 3, Price: $100.00, Total Value: $300.00
```

---

## ⚠️ Notes

* Map order is not guaranteed (items may display in different order)
* Input validation can be improved
* Item name is used as unique key

---

## 📚 What You Learned

* Structs and methods in Go
* Maps for storing data
* Function-based program design
* CLI interaction (input/output)
* Basic program flow control

---

## 💡 Summary

This project demonstrates how to build a simple real-world CLI application using Go fundamentals. It is a strong foundation for learning backend logic and data handling.
