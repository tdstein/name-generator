# Name Generator Package

The `namegen` package is a simple and flexible tool for generating random names or strings based on predefined choices. It provides choices out of the box and allows for customizable separators between generated name parts.

## Features

- **Random Name Generation**: Generate names by randomly selecting parts from provided choices.
- **Choice Datasets**: Provides datasets of predefined choices.
- **Arbitrary Name Length**: Supports selection from an arbitrary number of choice datasets.
- **Customizable Separators**: Define a custom separator to be used between name parts.

## Installation

To install the `namegen` package, use `go get`:

```bash
go get github.com/tdstein/namegen
```

## Usage

### Import the Package

```go
import (
    "github.com/tdstein/namegen"
)
```

### Create a Name Generator

To create a new name generator, initialize it with sets of choices. Each set represents a part of the name.

```go
titles := []string{"Mr.", "Ms.", "Dr."}
names := []string{"John", "Jane", "Alex"}
surnames := []string{"Doe", "Smith", "Johnson"}

ng, err := namegen.NewNameGenerator(titles, names, surnames)
if err != nil {
    log.Fatal(err)
}
```

### Generate a Name

Generate a random name using the `Generate` method. You can specify a custom separator.

```go
name := ng.Generate(" ")
fmt.Println(name) // Output could be: "Ms. Alex Smith"
fmt.Println(name) // Output could be: "Dr. Jane Doe"
fmt.Println(name) // Output could be: "Dr. Alex Doe"
```

### Choice Datasets

Load pre-installed choice datasets.

`choices, err := namegen.LoadChoices("adjectives")`

- **[`adjectives`](./data/adjectives.txt)**
- **[`amphibians`](./data/amphibians.txt)**
- **[`birds`](./data/birds.txt)**
- **[`fish`](./data/fish.txt)**
- **[`mammals`](./data/mammals.txt)**
- **[`places`](./data/places.txt)**
- **[`reptiles`](./data/reptiles.txt)**
- **[`things`](./data/things.txt)**

### Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/tdstein/namegen"
)

func main() {
    adjectives, err := namegen.LoadChoices("adjectives")
    if err != nil {
        log.Fatal(err)
    }

    things, err := namegen.LoadChoices("things")
    if err != nil {
        log.Fatal(err)
    }

    ng, err := namegen.NewNameGenerator(adjectives, things)
    if err != nil {
        log.Fatal(err)
    }

    name := ng.Generate("-")
    fmt.Println(name) // Output could be: "delightful-sunglasses"
}
```
## License

This project is licensed under the MIT License. See the [`LICENSE`](./LICENSE) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Contact

For any questions or issues, please open an issue on GitHub or contact the author directly.
