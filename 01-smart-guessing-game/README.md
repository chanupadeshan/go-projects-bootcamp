# 🎯 Lecture: Go Fundamentals via a Guessing Game

This is your first Go project. Let's turn it into a **mini lecture** that covers Go fundamentals clearly and practically.

## Running the Game
```bash
go run main.go
```

**How to Play:** Guess a number between 1 and 100. The program will tell you if it's too high, too low, or correct!

---

## 1. Program Structure (Entry Point)

```go
func main() {
```

* Every Go program starts from `main()`
* It is the **entry point** (like `main()` in C/Java)

---

## 2. Importing Packages

```go
import (
	"fmt"
	"math/rand"
	"time"
)
```

### What each package does:

* `fmt` → input/output (print, scan)
* `math/rand` → random number generation
* `time` → used to create a dynamic seed

👉 **Concept:** Go uses packages instead of libraries

---

## 3. Random Number Generation

```go
rand.Seed(time.Now().UnixNano())
secretNumber := rand.Intn(100) + 1
```

### Key idea:

* `Seed()` initializes randomness
* Without it → same number every run ❌
* With it → different number every run ✅

### Logic:

* `rand.Intn(100)` → gives 0–99
* `+1` → converts to **1–100**

---

## 4. Variables and Types

```go
attempts := 0
maxAttempts := 10
```

* `:=` → short variable declaration (Go feature)
* Type is inferred automatically

👉 Example:

* `attempts` → int (inferred)
* No need to write `int` explicitly

---

## 5. Output to User

```go
fmt.Println("Welcome to the Smart Guessing Game!")
```

* `Println()` → prints with new line
* `Printf()` → formatted output

### Example from your game:

```go
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
```

The `%d` is replaced by the variable value.

---

## 6. Loop (Core Game Logic)

```go
for attempts < maxAttempts {
```

* Go has **only one loop → `for`**
* This acts like a `while` loop

👉 Logic:

* Repeat until attempts reach limit

---

## 7. Taking User Input

```go
var guess int
fmt.Scan(&guess)
```

### Important concept:

* `&guess` → **memory address (pointer)**

👉 Why?

* Go needs the address to store input value
* `guess` = the value
* `&guess` = the address where to put it

---

## 8. Conditional Logic

```go
if guess < secretNumber {
	fmt.Println("Too low!")
} else if guess > secretNumber {
	fmt.Println("Too high!")
} else {
	fmt.Printf("Congratulations! ...")
}
```

* Standard `if-else`
* Controls game feedback

---

## 9. Early Exit (Return)

```go
return
```

* Ends the program immediately when guessed correctly

---

## 10. Game Over Condition

```go
fmt.Println("Game Over! ...")
```

* Executes only if loop finishes without correct guess

---

## 🧠 Core Concepts You Learned

From this ONE project, you covered:

| Concept | Example |
|---------|---------|
| Program structure | `func main()` |
| Packages | `import ("fmt", "math/rand", "time")` |
| Variables & type inference | `attempts := 0` |
| Loops | `for attempts < maxAttempts` |
| Conditions | `if guess < secretNumber` |
| User input | `fmt.Scan(&guess)` |
| Pointers | `&guess` |
| Random numbers | `rand.Intn(100)` |
| Basic game logic | Guessing mechanics |

---

## ⚠️ Important Improvements (Real Developer Thinking)

### 1. Input Validation (VERY IMPORTANT)

**Problem:**

```go
fmt.Scan(&guess)
```

* If user enters text → program crashes ❌

**Better approach:**

```go
_, err := fmt.Scan(&guess)
if err != nil {
    fmt.Println("Invalid input! Enter a number.")
    continue
}
```

### 2. Smarter Hint System

Instead of just "Too high/low":

```go
if guess-secretNumber > 10 {
	fmt.Println("Way too high!")
}
```

👉 Improves UX (User Experience)

### 3. Difficulty Levels (Next Step Idea)

* Easy → 20 attempts
* Medium → 10 attempts
* Hard → 5 attempts

---

# 📚 Detailed Lectures

## Lecture: Variables and Types in Go

### What is a Variable?

A **variable** is a named storage location in memory. It holds a value that your program can use.

Example:

```go
var age int = 23
```

Here:
- `age` → variable name
- `int` → type
- `23` → value

### What is a Type?

A **type** tells Go what kind of data a variable can store.

Examples:
- `int` → whole numbers
- `float64` → decimal numbers
- `string` → text
- `bool` → true or false

### Why are Types Important?

Go is a **statically typed language**, meaning the type is known before the program runs. Go checks type correctness at compile time.

```go
var age int = 20
age = "hello"  // ERROR! age is int, not string
```

### Ways to Declare Variables in Go

#### Method 1: Full Form

```go
var age int = 23
```

This is the clearest form. Structure: `var variableName type = value`

#### Method 2: Let Go Infer the Type

```go
var age = 23  // Go understands 23 is int
```

#### Method 3: Short Declaration (Most Common)

```go
age := 23
```

* `:=` both declares and assigns
* Mostly used inside `main()` or other functions
* Cannot be used outside functions for package-level variables

### Difference between `var` and `:=`

| Feature | `var` | `:=` |
|---------|-------|------|
| Form | `var x int = 10` | `x := 10` |
| Declare & assign later | ✅ Yes | ❌ No |
| Type inference | ✅ Optional | ✅ Automatic |
| Where to use | Anywhere | Inside functions |

### Common Basic Types in Go

| Type | Values | Example |
|------|--------|---------|
| `int` | Whole numbers | `-5, 20, 1000` |
| `float64` | Decimal numbers | `3.14, 10.5, -2.7` |
| `string` | Text | `"hello"`, `"Go"` |
| `bool` | True or False | `true`, `false` |

### Default Zero Values

If you declare a variable but don't assign a value, Go gives a **zero value**:

```go
var age int        // 0
var price float64  // 0
var name string    // "" (empty)
var isReady bool   // false
```

### Type Inference

Go can often detect the type from the value:

```go
x := 10          // int
y := 3.14        // float64
name := "Chanupa"  // string
ok := true       // bool
```

### Type Conversion

If you want to convert one type to another, do it explicitly:

```go
var x int = 10
var y float64 = float64(x)
fmt.Println(y)  // 10

var a float64 = 9.8
var b int = int(a)
fmt.Println(b)  // 9 (decimal part removed)
```

### Variable Naming Best Practices

Good variable names matter:

❌ Bad:
```go
x := 10
y := 5
z := x + y
```

✅ Better:
```go
width := 10
height := 5
area := width * height
```

In your code, names like `secretNumber`, `maxAttempts`, and `guess` are excellent.

### Go Naming Convention

Go usually uses:
- `camelCase` for variable names (not `snake_case`)

Examples:
```go
secretNumber
maxAttempts
userName
totalMarks
```

### Constants vs Variables

A **variable** can change:
```go
age := 20
age = 21  // ✅ OK
```

A **constant** cannot change:
```go
const pi = 3.14
pi = 4.0  // ❌ ERROR
```

### Improvement Idea for Your Game

Your code has:
```go
maxAttempts := 10
```

You could write:
```go
const maxAttempts = 10
```

Why? Because the maximum number of attempts is fixed and won't change.

---

## Lecture: Output to the User in Go

Output is how your program **communicates** with the user.

Without output, the user cannot see what's happening.

### The `fmt` Package

For output, we use the **`fmt`** package:

```go
import "fmt"
```

The three main functions are:
* `fmt.Print()`
* `fmt.Println()`
* `fmt.Printf()`

### `fmt.Print()` - No New Line

Prints output **without automatically adding a new line** at the end:

```go
fmt.Print("Hello")
fmt.Print("World")
```

Output:
```
HelloWorld
```

### `fmt.Println()` - With New Line

Prints output and **adds a new line** at the end:

```go
fmt.Println("Hello")
fmt.Println("World")
```

Output:
```
Hello
World
```

### `fmt.Printf()` - Formatted Output

Used for **formatted output**. You can control how values appear:

```go
name := "chanupa"
age := 23
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
```

Output:
```
My name is chanupa and I am 23 years old.
```

### Why `Printf()` is Powerful

You can insert variable values inside a sentence:

```go
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
```

If `attempts = 3` and `maxAttempts = 10`, output becomes:
```
Attempt 3/10 — Enter your guess:
```

### Common Format Verbs in `Printf()`

| Verb | Use | Example |
|------|-----|---------|
| `%d` | Integer | `fmt.Printf("Age: %d", 23)` |
| `%s` | String | `fmt.Printf("Name: %s", "chanupa")` |
| `%f` | Float | `fmt.Printf("Price: %f", 99.5)` |
| `%.2f` | Float with decimals | `fmt.Printf("Price: %.2f", 99.5)` → `99.50` |
| `%t` | Boolean | `fmt.Printf("Ready: %t", true)` |
| `%T` | Type | `fmt.Printf("%T", 10)` → `int` |

### Special Characters in Output

* `\n` → new line
* `\t` → tab space
* `\"` → double quote inside text
* `%%` → percent sign

Example:
```go
fmt.Printf("Score: 95%%\n")
```

Output:
```
Score: 95%
```

### In Your Guessing Game

You used output for:

```go
fmt.Println("Welcome to the Smart Guessing Game!")
```

**Purpose:** Greet the user

```go
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
```

**Purpose:** Show progress and ask for input

```go
fmt.Println("Too low! Try again.")
fmt.Println("Too high! Try again.")
```

**Purpose:** Give feedback

```go
fmt.Printf("Congratulations! You've guessed the number %d in %d attempts!\n", secretNumber, attempts)
```

**Purpose:** Show final success message

### Good Output Principles

Output should be:
* ✅ Clear
* ✅ Short
* ✅ Helpful
* ✅ Friendly

❌ Bad:
```go
fmt.Println("Error")
```

✅ Better:
```go
fmt.Println("Invalid input! Please enter a number between 1 and 100.")
```

---

## Lecture: Taking User Input in Go

User input means getting data from the user while the program is running.

### What is User Input?

User input is data entered by the user through the keyboard.

Examples:
* name
* age
* number guess
* password
* menu choice

### Basic Input with `fmt.Scan()`

```go
var age int
fmt.Print("Enter your age: ")
fmt.Scan(&age)
fmt.Println("Your age is:", age)
```

### Why `&age`? (The Most Important Part)

```go
fmt.Scan(&guess)
```

We do **not** write `fmt.Scan(guess)`.

We write `fmt.Scan(&guess)` because:

* `Scan()` needs the **memory address** of the variable
* `guess` → the value
* `&guess` → the address of `guess`

**Simple meaning:** `Scan()` stores the user's input into that memory location.

Think like this:
* Variable = a box
* Value = what is inside the box
* `&variable` = the location of the box

`fmt.Scan()` needs to know **where to put the input**.

### Input Variable Must Have Suitable Type

```go
var age int
fmt.Scan(&age)
```

This expects an integer input like `23`.

If user enters `hello`, that's invalid for `int`.

The variable type matters!

### Input with Different Types

**String input:**
```go
var name string
fmt.Scan(&name)
```

**Float input:**
```go
var price float64
fmt.Scan(&price)
```

**Boolean input:**
```go
var isReady bool
fmt.Scan(&isReady)
```

Valid input: `true` or `false`

### Taking Multiple Inputs

`fmt.Scan()` can read more than one value:

```go
var name string
var age int

fmt.Print("Enter name and age: ")
fmt.Scan(&name, &age)

fmt.Println("Name:", name, "Age:", age)
```

Input:
```
chanupa 23
```

Output:
```
Name: chanupa Age: 23
```

### Important Behavior

`fmt.Scan()` separates input by **spaces**.

If you do:
```go
var fullName string
fmt.Scan(&fullName)
```

And user types: `chanupa Perera`

Only `chanupa` is stored!

### When to Use `fmt.Scan()`

✅ Good for:
* One word
* Numbers
* Multiple small values

❌ Not ideal for:
* Full sentences
* Full names with spaces
* Paragraph input

### Handling Input Errors

This is **very important** for real programs:

```go
_, err := fmt.Scan(&guess)

if err != nil {
    fmt.Println("Invalid input! Please enter a number.")
    return
}
```

### Better Version for Games

In games, you usually don't want to end the whole program on bad input:

```go
_, err := fmt.Scan(&guess)

if err != nil {
    fmt.Println("Invalid input! Please enter a number.")
    continue  // Try again
}
```

### In Your Guessing Game

Your code:
```go
var guess int
fmt.Print("Attempt %d/%d - Enter your guess: ", attempts, maxAttempts)
fmt.Scan(&guess)
```

**Better version:**
```go
var guess int
fmt.Printf("Attempt %d/%d - Enter your guess: ", attempts, maxAttempts)

_, err := fmt.Scan(&guess)
if err != nil {
    fmt.Println("Invalid input! Please enter a number.")
    continue
}
```

This prevents crashes from bad input.

### Input Flow in Your Game

1. Program asks for input with `Printf()`
2. User types a number
3. `Scan()` reads it and stores in `guess`
4. Program compares `guess` with `secretNumber`
5. Program gives feedback

**Input** is the bridge between user action and program logic.

### Common Beginner Mistakes

❌ **Mistake 1:** Forgetting `&`
```go
fmt.Scan(guess)  // WRONG
```

✅ Correct:
```go
fmt.Scan(&guess)  // RIGHT
```

❌ **Mistake 2:** Wrong type expectation
```go
var age int
fmt.Scan(&age)
```
If user types `"hello"` → error

❌ **Mistake 3:** Expecting full sentence
```go
var text string
fmt.Scan(&text)
```
Input: `hello world`
Stored: `hello` (only first word!)

---

## Summary

### Key Functions

| Function | Behavior | When to Use |
|----------|----------|------------|
| `fmt.Print()` | No new line | Rarely needed |
| `fmt.Println()` | With new line | Simple output |
| `fmt.Printf()` | Formatted | Professional output |
| `fmt.Scan()` | Read input | Get user data |

### Key Concepts

| Concept | Example | Importance |
|---------|---------|-----------|
| Variable | `age := 23` | Store data |
| Type | `int`, `string`, etc. | Ensure correctness |
| `&address` | `&guess` | Pass to `Scan()` |
| Format verbs | `%d`, `%s`, `%f` | Format output |
| Error handling | `if err != nil` | Robust programs |

---

## 🚀 Next Steps

1. ✅ Understand this game thoroughly
2. ⬜ Add **input validation** so your game doesn't crash
3. ⬜ Add **difficulty levels** (easy/medium/hard)
4. ⬜ Move to next project: Inventory Manager
