# Magical Arena

## Overview

This project simulates battles between players in a magical arena. Players have health, strength, and attack attributes, and they use dice to determine attack and defense outcomes.

## Structure

- `main.go`: Entry point of the program.
- `pkg/model`: Contains the `Die` and `Player` structs.
- `pkg/service`: Contains the `Arena` struct and its methods.
- `tests`: Contains unit tests for the code.

## Running the Code

1. Use `go run main.go` to execute the program.
2. Use `go test -v ./tests` to execute the unit test in the program

## Additional Information

- The code is written in Go and uses only standard libraries.
- Unit tests are provided in the `tests` directory.

