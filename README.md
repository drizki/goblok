# Goblok

[![GoDoc](https://godoc.org/github.com/drizki/goblok?status.png)](https://godoc.org/github.com/drizki/goblok)
![Go CI](https://github.com/drizki/goblok/actions/workflows/go.yml/badge.svg)

Goblok is a Go library designed for filtering and analyzing profane words in Bahasa Indonesia. It provides functionalities to filter out profane words, add or remove words from the dictionary, and analyze text to detect profane content.

## Installation

To use Goblok, import the package into your Go project:

```bash
go get github.com/drizki/goblok
```

## Usage

### Importing the Package

```go
import "github.com/drizki/goblok"
```

### Creating a New Goblok Instance

You can create a new instance of `Goblok` with a custom dictionary of profane words:

```go
// Create a Goblok instance with a custom dictionary
goblok := goblok.New([]string{"exampleProfaneWord1", "exampleProfaneWord2"})
```

If you don't provide a dictionary, Goblok will use the default set of profane words.

### Adding and Removing Profane Words

To add or remove profane words from the dictionary:

```go
// Add a new profane word
wordAdded := goblok.Add("newProfaneWord")
fmt.Println("Added:", wordAdded)

// Remove a profane word
wordRemoved := goblok.Remove("existingProfaneWord")
fmt.Println("Removed:", wordRemoved)
```

### Filtering Text

Use the `Filter` method to replace profane words with a mask character:

```go
// Filter text, replacing profane words with asterisks
filteredText := goblok.Filter("This is a bad example.", '*')
fmt.Println("Filtered Text:", filteredText)
```

You can specify a different mask character if desired:

```go
// Filter text with a custom mask character
filteredText := goblok.Filter("This is another bad example.", '#')
fmt.Println("Filtered Text:", filteredText)
```

### Analyzing Text

The `Analyze` method provides detailed information about profane words in the text:

```go
// Analyze the text for profane words
analysis, err := goblok.Analyze("This is an example with bad words.")
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println("Original Input:", analysis.Input)
fmt.Println("Filtered Text:", analysis.Filtered)
fmt.Println("Profane Words Detected:", analysis.Profane)
fmt.Println("Locations of Profane Words:", analysis.Locations)
```

### Checking for Profanity

To check if the input text contains any profane words:

```go
// Check if the text contains profane words
containsProfanity := goblok.Contains("This text contains a bad word.")
fmt.Println("Contains Profanity:", containsProfanity)
```

## Structs and Methods

### Goblok

- `New(dictionary ...[]string) *Goblok` - Creates a new `Goblok` instance with an optional custom dictionary.
- `Add(word string) string` - Adds a new profane word to the dictionary.
- `Remove(word string) string` - Removes a profane word from the dictionary.
- `Filter(input string, mask ...rune) string` - Filters the input text, replacing profane words with the specified mask character.
- `Analyze(input string) (Analysis, error)` - Analyzes the input text and returns details about profane words and their locations.
- `Contains(input string) bool` - Checks if the input text contains any profane words.

### Analysis

- `Input string` - The original input text.
- `Profane []string` - List of profane words detected in the text.
- `Filtered string` - The input text with profane words replaced by the mask character.
- `Locations []Location` - Locations of profane words in the original text.

### Location

- `Word string` - The profane word.
- `Index int` - The starting index of the profane word in the original text.

## License

This library is open source and available under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on the [GitHub repository](https://github.com/drizki/goblok).
