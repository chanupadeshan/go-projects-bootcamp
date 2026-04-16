# Lecture: Go Fundamentals via a Guessing Game

This is your first Go project. Let's learn Go fundamentals clearly and practically through this guessing game.

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
* No arguments needed for this simple program

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

| Package | Purpose |
|---------|---------|
| `fmt` | Input/output (printing, scanning user input) |
| `math/rand` | Random number generation |
| `time` | Time-related functions and current time |

**Concept:** Go uses packages to organize reusable code, similar to libraries in other languages.

---

## 3. Seeding the Random Number Generator

```go
rand.Seed(time.Now().UnixNano())
secretNumber := rand.Intn(100) + 1
```

### Understanding Randomness

**What is `rand.Seed()`?**

* Initializes the random number generator with a starting value
* `time.Now().UnixNano()` gets the current time in nanoseconds
* This ensures a different value every time the program runs

**What is `rand.Intn(100)`?**

* Generates a random integer from 0 to 99
* The `+ 1` converts it to the range 1 to 100

**Example flow:**
- `rand.Intn(100)` might give: 0, 45, 99
- `rand.Intn(100) + 1` gives: 1, 46, 100

---

## 4. Variables and Type Inference

```go
attempts := 0
maxAttempts := 10
```

### Understanding Variable Declaration

**What does `:=` mean?**

* Short variable declaration
* Both **declares** and **assigns** a value
* Type is **automatically inferred** from the value

**In this example:**
- `attempts` is inferred as `int` (because 0 is an integer)
- `maxAttempts` is inferred as `int` (because 10 is an integer)
- Go figures out the type without you writing it explicitly

**Why this matters:**
- Less typing
- Still type-safe (Go checks types at compile time)
- More readable code

---

## 5. Output to the User

```go
fmt.Println("Welcome to the Smart Guessing Game!")
fmt.Println("I have picked a number between 1 and 100. Can you guess it?")
```

### Understanding Output Functions

**What is `fmt.Println()`?**

* Prints text to the console
* Automatically adds a new line at the end
* Perfect for welcome messages and simple output

**Example in the game:**

```go
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
```

This is `fmt.Printf()` which is **formatted output**.

---

## 6. The Main Loop Structure

```go
for attempts < maxAttempts {
```

### Understanding the `for` Loop

**Key fact:** Go has **only ONE loop construct** - the `for` loop.

It works in multiple ways:
* `for condition { }` - acts like a `while` loop
* `for i := 0; i < 10; i++ { }` - acts like a traditional `for` loop
* `for { }` - infinite loop

**In this game:**
- `for attempts < maxAttempts` keeps running while attempts is less than 10
- The loop repeats until the player runs out of attempts

**What happens inside:**
1. `attempts++` - increment the attempt counter
2. Get guess from user
3. Compare guess with secret number
4. Give feedback

---

## 7. Taking User Input

```go
var guess int
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
_, err := fmt.Scan(&guess)
```

### Understanding Variable Declaration with `var`

```go
var guess int
```

**What does this do?**
- Declares a variable named `guess`
- Type is `int` (whole numbers)
- Does NOT assign a value yet
- Goes to default zero value: 0

**Why use `var` here instead of `:=`?**
- We're declaring first, assigning later
- `:=` requires assignment at declaration

### Understanding User Input

```go
fmt.Scan(&guess)
```

**Why the `&` symbol?**

* `Scan()` needs the **memory address** of the variable
* `guess` = the variable itself
* `&guess` = the address of where to store the input

**Think of it like this:**
- Variable = a mailbox
- `guess` = the mailbox itself
- `&guess` = the address of the mailbox
- `Scan()` needs the address to put mail (input) in it

### Understanding Error Handling

```go
_, err := fmt.Scan(&guess)

if err != nil {
    fmt.Println("Invalid input. Please enter a number.")
    attempts--
}
```

**What does this do?**

* `fmt.Scan()` returns two values:
  - The number of items successfully scanned
  - An error value (if something went wrong)

* `_, err :=` means:
  - `_` ignores the first value (we don't need it)
  - `err` captures the error value

* `if err != nil` means:
  - If there's an error (err is not empty)
  - Print a message and decrement attempts

**Why decrement attempts?**
- If user enters invalid input, we don't count it as a real attempt
- They get to try again

---

## 8. Conditional Logic - Comparing Numbers

```go
if guess < secretNumber {
    fmt.Println("Too low! Try again.")
} else if guess > secretNumber {
    fmt.Println("Too high! Try again.")
} else {
    fmt.Printf("Congratulations! You've guessed the number %d in %d attempts!\n", secretNumber, attempts)
    return
}
```

### Understanding `if-else` Statements

**Three paths in this game:**

1. **If guess is too low:**
   ```go
   if guess < secretNumber {
       fmt.Println("Too low! Try again.")
   }
   ```
   - Compare using `<` operator
   - Give feedback

2. **Else if guess is too high:**
   ```go
   else if guess > secretNumber {
       fmt.Println("Too high! Try again.")
   }
   ```
   - Compare using `>` operator
   - Give feedback

3. **Else (guess equals the secret number):**
   ```go
   else {
       fmt.Printf("Congratulations! You've guessed the number %d in %d attempts!\n", secretNumber, attempts)
       return
   }
   ```
   - Player won!
   - Print success message with the values
   - `return` exits the entire program

### Understanding `return`

```go
return
```

* Stops execution of the current function
* Since we're in `main()`, this ends the entire program
* The loop stops immediately

---

## 9. Game Over Message

```go
fmt.Println("Game Over! You've used all your attempts. The secret number was:", secretNumber)
```

**When does this execute?**

* Only if the `for` loop ends naturally
* That means the player used all 10 attempts without guessing correctly
* Prints the secret number so the player can see what they were trying to find

**Flow:**
- Loop runs while `attempts < maxAttempts`
- If attempts reaches 10, the loop condition is false
- Loop exits
- This line executes

---

## Core Concepts Summary

| Concept | What It Does | Example |
|---------|-------------|---------|
| `package main` | Entry point of program | `package main` |
| `import` | Load external code packages | `import ("fmt", "math/rand")` |
| `rand.Seed()` | Initialize randomness | `rand.Seed(time.Now().UnixNano())` |
| `rand.Intn()` | Generate random number | `rand.Intn(100)` |
| `:=` | Declare variable with type inference | `attempts := 0` |
| `var` | Declare variable with explicit type | `var guess int` |
| `fmt.Println()` | Print with new line | `fmt.Println("Hello")` |
| `fmt.Printf()` | Print with formatting | `fmt.Printf("%d", value)` |
| `fmt.Scan()` | Get user input | `fmt.Scan(&guess)` |
| `for` loop | Repeat code | `for attempts < max { }` |
| `if-else` | Conditional execution | `if x > 5 { } else { }` |
| `return` | Exit function/program | `return` |

---

# Detailed Lectures

## Lecture 1: Variables and Types in Go

### What is a Variable?

A **variable** is a named storage location in memory where your program stores data.

Example:
```go
age := 23
```

Here:
- `age` = name of the storage location
- `23` = the value stored
- `int` = the type (inferred)

### What is a Type?

A **type** tells Go what kind of data a variable can store.

Examples:
- `int` → whole numbers (-1000, 0, 42, 1000)
- `float64` → decimal numbers (3.14, 9.8, 0.5)
- `string` → text ("hello", "Go", "123")
- `bool` → true or false (true, false)

### Why Types Matter in Go

Go is **statically typed**, meaning:
- Every variable has a type
- The type is known before the program runs
- Go checks that you use types correctly
- You can't assign the wrong type to a variable

Example:
```go
var age int = 20
age = "hello"  // ERROR - can't assign string to int
```

### Variable Declaration Methods

**Method 1: Full form with `var`**
```go
var age int = 23
```
- Clearest and most explicit
- Shows the type clearly
- Can declare without assigning

**Method 2: Type inference with `var`**
```go
var age = 23
```
- Go figures out it's `int` from the value 23
- Still using `var` keyword

**Method 3: Short declaration with `:=` (most common inside functions)**
```go
age := 23
```
- Both declares and assigns
- Type is inferred
- Most concise
- Only works inside functions

### Default Values (Zero Values)

If you declare a variable but don't assign a value:

```go
var age int        // becomes 0
var name string    // becomes "" (empty)
var isReady bool   // becomes false
var price float64  // becomes 0.0
```

Each type has a zero value.

### Type Conversion

If you need to change a value from one type to another:

```go
var x int = 10
var y float64 = float64(x)  // Convert int to float64

var a float64 = 9.8
var b int = int(a)  // Convert float64 to int (becomes 9, decimal removed)
```

Conversion is **explicit** - you must write it.

### Variable Naming Conventions

In Go, use `camelCase` for variable names:

Good:
```go
secretNumber
maxAttempts
userName
playerScore
```

Not typical in Go:
```go
secret_number
MAX_ATTEMPTS
UserName
```

Good variable names:
- Clearly describe what they store
- Use meaningful words
- Make code easy to understand

### Example from Your Game

```go
attempts := 0         // Type: int, Value: 0
maxAttempts := 10     // Type: int, Value: 10
var guess int         // Type: int, Value: 0 (default)
```

---

## Lecture 2: Output to User in Go

### What is Output?

Output is how your program **communicates** with the user by showing information on the screen.

### The `fmt` Package

The `fmt` package provides functions for printing output:

```go
import "fmt"
```

### Three Main Output Functions

#### 1. `fmt.Print()` - No New Line

```go
fmt.Print("Hello")
fmt.Print("World")
```

Output:
```
HelloWorld
```

**When to use:** Rarely. Usually you want new lines.

#### 2. `fmt.Println()` - With New Line

```go
fmt.Println("Hello")
fmt.Println("World")
```

Output:
```
Hello
World
```

**When to use:** Most common. Simple output that needs readability.

#### 3. `fmt.Printf()` - Formatted Output

```go
name := "chanupa"
age := 23
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
```

Output:
```
My name is chanupa and I am 23 years old.
```

**When to use:** When you need to insert variable values into text.

### Format Verbs in `fmt.Printf()`

Format verbs are placeholders that get replaced by values:

| Verb | Type | Example |
|------|------|---------|
| `%d` | Integer | `fmt.Printf("Age: %d", 23)` → `Age: 23` |
| `%s` | String | `fmt.Printf("Name: %s", "Go")` → `Name: Go` |
| `%f` | Float | `fmt.Printf("Price: %f", 9.99)` → `Price: 9.990000` |
| `%.2f` | Float (2 decimals) | `fmt.Printf("Price: %.2f", 9.99)` → `Price: 9.99` |
| `%t` | Boolean | `fmt.Printf("Ready: %t", true)` → `Ready: true` |
| `%T` | Type | `fmt.Printf("%T", 10)` → `int` |

### Special Characters in Output

- `\n` → new line
- `\t` → tab space
- `\\` → backslash
- `\"` → double quote
- `%%` → percent sign

Example:
```go
fmt.Printf("Score: 95%%\n")
```

Output:
```
Score: 95%
```

### Output in Your Game

Your game uses output at different stages:

**Welcome message:**
```go
fmt.Println("Welcome to the Smart Guessing Game!")
```
- Greets the player

**Progress prompt:**
```go
fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
```
- Shows current attempt and max attempts
- Asks for input

**Feedback:**
```go
fmt.Println("Too low! Try again.")
fmt.Println("Too high! Try again.")
```
- Tells player if guess is too high or low

**Success:**
```go
fmt.Printf("Congratulations! You've guessed the number %d in %d attempts!\n", secretNumber, attempts)
```
- Celebrates when player wins

**Game over:**
```go
fmt.Println("Game Over! You've used all your attempts. The secret number was:", secretNumber)
```
- Tells player the game ended and reveals the number

---

## Lecture 3: Taking User Input in Go

### What is User Input?

User input is data that the user enters (usually from keyboard) while the program is running.

Examples:
- Number to guess
- Player name
- Menu choice
- Search term

### The `fmt.Scan()` Function

```go
var guess int
fmt.Scan(&guess)
```

**What does it do?**
- Waits for user to type something
- Reads what the user typed
- Stores it in the variable
- Continues the program

### The `&` Symbol (Address Operator)

This is very important:

```go
fmt.Scan(&guess)
```

We use `&guess`, NOT just `guess`.

**Why?**

Think of it like mailing something:
- `guess` = a mailbox
- `&guess` = the address of the mailbox
- `fmt.Scan()` needs the address to put data into that mailbox

Go's `Scan()` function needs to know **where** to put the input value.

### Input with Different Types

**Reading an integer:**
```go
var age int
fmt.Scan(&age)
```
User enters: `25`
Result: `age = 25`

**Reading a string:**
```go
var name string
fmt.Scan(&name)
```
User enters: `Alice`
Result: `name = "Alice"`

**Reading a float:**
```go
var price float64
fmt.Scan(&price)
```
User enters: `9.99`
Result: `price = 9.99`

**Reading a boolean:**
```go
var isReady bool
fmt.Scan(&isReady)
```
User enters: `true`
Result: `isReady = true`

### Reading Multiple Values

```go
var name string
var age int

fmt.Scan(&name, &age)
```

User enters: `Alice 25`
Result: `name = "Alice"`, `age = 25`

### Important Behavior - Spaces

`fmt.Scan()` stops reading at spaces:

```go
var fullName string
fmt.Scan(&fullName)
```

User enters: `Alice Smith`
Result: `fullName = "Alice"` (only first word!)

This is how `Scan()` separates multiple inputs.

### Error Handling with `fmt.Scan()`

```go
_, err := fmt.Scan(&guess)

if err != nil {
    fmt.Println("Invalid input. Please enter a number.")
}
```

**What does this mean?**

- `fmt.Scan()` returns two values
- First value: how many items were successfully read (we ignore with `_`)
- Second value: error information

If something goes wrong (user enters text when expecting number):
- `err` is not `nil`
- We can detect and handle it

### Example in Your Game

```go
_, err := fmt.Scan(&guess)

if err != nil {
    fmt.Println("Invalid input. Please enter a number.")
    attempts--
}
```

**Flow:**
1. Try to read an integer into `guess`
2. If it fails → show error message
3. Decrement attempts so invalid input doesn't count

---

## Lecture 4: Loops in Go

### What is a Loop?

A loop is code that **repeats** multiple times.

### The `for` Loop

Go has **one main loop: `for`**

It works in three ways:

#### Type 1: Condition-based (like while)

```go
for attempts < maxAttempts {
    // code repeats while attempts < maxAttempts
}
```

In your game, this repeats until the player runs out of attempts.

#### Type 2: Traditional for loop

```go
for i := 0; i < 10; i++ {
    // runs 10 times (i goes from 0 to 9)
}
```

Three parts:
- `i := 0` - initialize
- `i < 10` - condition
- `i++` - increment

#### Type 3: Infinite loop

```go
for {
    // runs forever (need break to exit)
}
```

### Your Game's Loop

```go
for attempts < maxAttempts {
    attempts++
    
    // Ask for guess
    // Compare with secret number
    // Give feedback
    // If correct, return (exit)
}

// Only reaches here if loop ends normally
```

**Loop execution:**
1. Check: is `attempts < maxAttempts`?
2. If yes: run code inside
3. Return to step 1
4. If no: exit loop

**In your game:**
- Loop runs 10 times maximum
- But if player guesses correctly, `return` exits the loop early
- If player uses all 10 attempts, loop ends naturally

---

## Lecture 5: Conditional Logic (if-else)

### What are Conditionals?

Conditionals let your code **make decisions** based on conditions.

### The `if` Statement

```go
if guess < secretNumber {
    fmt.Println("Too low! Try again.")
}
```

**How it works:**
1. Check the condition: `guess < secretNumber`
2. If true: run the code inside `{ }`
3. If false: skip it

### The `if-else` Statement

```go
if guess < secretNumber {
    fmt.Println("Too low!")
} else {
    fmt.Println("Not too low!")
}
```

**How it works:**
1. Check condition
2. If true: run first block
3. If false: run second block

### The `if-else if-else` Chain

Your game uses this:

```go
if guess < secretNumber {
    fmt.Println("Too low!")
} else if guess > secretNumber {
    fmt.Println("Too high!")
} else {
    fmt.Println("Correct!")
}
```

**How it works:**
1. Check first condition
2. If true: run first block and stop
3. If false: check second condition
4. If true: run second block and stop
5. If false: run final block

**In your game:**
- First check: if guess is too low
- Second check: if guess is too high
- Final: if neither (meaning guess equals secret)

### Comparison Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `<` | Less than | `x < 10` |
| `>` | Greater than | `x > 5` |
| `<=` | Less than or equal | `x <= 10` |
| `>=` | Greater than or equal | `x >= 0` |
| `==` | Equal | `x == 5` |
| `!=` | Not equal | `x != 0` |

### Logical Operators

Combine multiple conditions:

- `&&` means AND (both must be true)
- `||` means OR (at least one must be true)
- `!` means NOT (reverse the result)

Example:
```go
if x > 0 && x < 100 {
    // x is between 0 and 100
}
```

---

## 🚀 Summary Table

| Concept | In Your Game | Purpose |
|---------|-------------|---------|
| Random seed | `rand.Seed(time.Now().UnixNano())` | Make number different each run |
| Random number | `rand.Intn(100) + 1` | Pick number 1-100 |
| Variables | `attempts := 0` | Store game state |
| Output | `fmt.Println()`, `fmt.Printf()` | Communicate with player |
| Loop | `for attempts < maxAttempts` | Repeat game rounds |
| Input | `fmt.Scan(&guess)` | Get player's guess |
| Conditions | `if guess < secretNumber` | Decide what feedback to give |
| Early exit | `return` | End game when player wins |

---

## You Have Learned

From this one simple game project, you've covered:

- ✅ Program structure and entry point
- ✅ Packages and imports
- ✅ Variables and type inference
- ✅ Random number generation
- ✅ Output functions (print, println, printf)
- ✅ Formatted output with format verbs
- ✅ User input with Scan()
- ✅ Error handling basics
- ✅ Loops (for)
- ✅ Conditionals (if-else)
- ✅ Comparison operators
- ✅ Program flow control (return)

These are **fundamental concepts** that you'll use in every Go program!

---

