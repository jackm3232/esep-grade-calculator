# ESEP Grade Calculator

This library is a simple utility to calculate your grade in a course.
It lets you add assignments, exams, and essays, and then calculates a final letter grade based on the weighted average of each assignment type.

## Tech Stack:
- **Language**: Go 1.24+
- **Testing**: Go's built-in "testing" package

## Project Structure:
- **go.mod** – Declares Go module and specifies Go version.
-	**grade_calculator.go** – Implements the grade calculator that lets you add assignments, exams, and essays, and then calculates a final letter grade based on the weighted average of each assignment type.
- **grade_calculator_test.go** – Contains unit tests for the grade calculator to make sure it functions correctly.

## How It Works:
- Each grade falls into one of three categories:
  - **Assignments** (50% weight)
  - **Exams** (35% weight)
  - **Essays** (15% weight)
- The calculator computes a weighted average of the scores and converts it into a final letter grade:
  - 90-100 --> A
  - 80-89 --> B
  - 70-79 --> C
  - 60-69 --> D
  - 0-59 --> F

## Setting Up Project
- [Install Go](https://go.dev/dl/)
- Make sure Go's bin directory is added to your system's PATH
- This command will return the version if Go is set up correctly: 
```
go version
```

## Run Unit Tests:
```
cd grade_calculator/
go test ./...
```

## Run Unit Tests With Code Coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```