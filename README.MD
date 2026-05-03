# Go (Golang) Learning Guide

## 📚 Introduction
Go is a statically typed, compiled programming language designed for simplicity, speed, and concurrency. It's ideal for building scalable backend applications, microservices, and CLI tools.

## 🎯 Prerequisites
- Basic programming knowledge (any language)
- A code editor (VS Code recommended)
- Command line/terminal familiarity
- Go 1.19+ installed

## 🚀 Installation & Setup

### 1. Install Go
- Visit [golang.org](https://golang.org/dl)
- Download for your OS (Windows/Mac/Linux)
- Run installer and verify: `go version`

### 2. Setup Workspace
```bash
# Create workspace
mkdir $HOME/go/src/projects

# Set GOPATH (usually automatic in Go 1.11+)
export PATH=$PATH:/usr/local/go/bin
```

### 3. Verify Installation
```bash
go version
go env
```

---

## 📖 Learning Path (Structured)

### Phase 1: Fundamentals (Week 1)
- **Variables & Data Types**: `int`, `string`, `float64`, `bool`
- **Basic Operations**: arithmetic, logical, comparison
- **Output**: `fmt.Println()`, string formatting

### Phase 2: Control Flow (Week 2)
- **If-Else Statements**: conditional logic
- **Switch-Case**: multi-way branching
- **For Loops**: iteration, ranges, while loops
- **Break & Continue**: loop control

### Phase 3: Functions & Scope (Week 3)
- **Function Declaration**: parameters, return values
- **Variable Scope**: local vs package-level
- **Defer Statement**: cleanup operations
- **Error Handling**: basics with `error` type

### Phase 4: Data Structures (Week 4)
- **Arrays**: fixed-length collections
- **Slices**: dynamic arrays
- **Maps**: key-value pairs
- **Structs**: custom data types

### Phase 5: Advanced Topics
- **Goroutines**: concurrent execution
- **Channels**: goroutine communication
- **Interfaces**: polymorphism
- **Error Handling**: `defer`, `panic`, `recover`
- **Packages & Modules**: code organization

---

## 📁 Project Structure
```
golang/
├── README.md              # This file
├── main.go               # Entry point & simple examples
├── variable-scope/       # Variable scope demonstrations
├── if-else/              # Conditional logic examples
├── switch/               # Switch statement examples
├── for-loop/             # Loop implementations
└── (Add more folders for new topics)
```

---

## 🔑 Key Concepts to Master

| Concept | Description |
|---------|------------|
| **Packages** | Code organization, `main` package, `func main()` |
| **Variables** | Declaration: `var`, `:=`, type inference |
| **Data Types** | int, float64, string, bool, arrays, slices, maps |
| **Functions** | First-class citizens, multiple returns, defer |
| **Errors** | Return error as value, not exceptions |
| **Goroutines** | Lightweight concurrency model |
| **Interfaces** | Duck typing, satisfies interface implicitly |

---

## 💡 Practice Tips

### Daily Practice
1. Write 1-2 small programs daily
2. Experiment with edge cases
3. Read Go standard library code
4. Try refactoring code to be more idiomatic

### Progression
- Start with `main.go` for basic concepts
- Create separate files/folders for each topic
- Build mini-projects (calculator, todo app, web server)
- Read others' Go code on GitHub

### Common Mistakes to Avoid
- ❌ Using `panic()` for regular error handling
- ❌ Ignoring errors with `_ = someFunction()`
- ❌ Inefficient string concatenation (use `strings.Builder`)
- ❌ Creating goroutines without proper synchronization
- ❌ Modifying maps/slices during iteration

---

## 📚 Essential Resources

### Official
- [Go Official Tour](https://tour.golang.org)
- [Go Documentation](https://golang.org/doc)
- [Standard Library Packages](https://golang.org/pkg)

### Learning
- *The Go Programming Language* by Donovan & Kernighan
- [Go by Example](https://gobyexample.com)
- YouTube: Traversy Media, TechWorld with Nana

### Practice
- [LeetCode Go Problems](https://leetcode.com) (filter by Go)
- [HackerRank Go Challenges](https://www.hackerrank.com)
- Build projects: CLI tools, REST APIs, web scrapers

---

## 🎓 Sample First Program

```go
package main

import "fmt"

func main() {
    name := "Go Learner"
    fmt.Printf("Hello, %s! Welcome to Go.\n", name)
}
```

**Run it:**
```bash
go run main.go
```

---

## 📊 Quick Commands Reference

| Command | Purpose |
|---------|---------|
| `go run main.go` | Execute Go file directly |
| `go build` | Compile to binary |
| `go build -o myapp` | Build with custom name |
| `go test` | Run tests |
| `go fmt ./...` | Format all code |
| `go vet ./...` | Check for errors |
| `go get` | Download dependencies |

---

## ✅ Milestone Checklist

- [ ] Installed Go and verified setup
- [ ] Wrote first "Hello World" program
- [ ] Mastered variables and data types
- [ ] Completed control flow (if/switch/for)
- [ ] Understood variable scope
- [ ] Wrote functions with multiple returns
- [ ] Used slices and maps
- [ ] Created custom structs
- [ ] Handled errors properly
- [ ] Built a complete mini-project

---

## 🔗 Next Steps

1. Complete this learning path sequentially
2. Build 3-5 small projects to reinforce concepts
3. Study concurrency (goroutines & channels)
4. Explore web frameworks (Gin, Echo)
5. Contribute to open-source Go projects

---

**Happy Learning! 🚀**
