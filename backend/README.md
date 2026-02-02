# ascii-art

## Summary
`ascii-art` is a command-line application written in Go that transforms input text into artistic ASCII renditions. It allows users to convert strings into stylized ASCII art using predefined font files, providing a fun and creative way to display text in the terminal.

## Key Features
*   **ASCII Art Generation**: Converts input text into multi-line ASCII art.
*   **Custom Font Support**: Utilizes `.txt` files (like `standard.txt`) as ASCII art font definitions.
*   **Command-Line Argument Parsing**: Handles input strings and optional font file specifications directly from the command line.
*   **Robust Input Handling**: Includes checks for file existence and format.
*   **Comprehensive Testing**: Features unit and integration tests to ensure reliable ASCII art generation.

## Tech stack
*   Go
*   testing
*   os
*   strings
*   fmt

## Installation
To get a local copy up and running, follow these simple steps.

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/elvisotieno0/ascii-art.git
    cd ascii-art
    ```
2.  **Download dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Build the application**:
    ```bash
    go build -o ascii-art .
    ```

## Folder Structure
```
ascii-art/
├── Lib/                    # Library containing core logic and utility functions
│   ├── check/              # Functions for command-line argument and file validation
│   │   └── checkArgs.go    # Validates and parses command-line arguments
│   ├── print/              # Functions responsible for generating and printing ASCII art
│   │   └── printAscii.go   # Main function for converting data to ASCII art and printing
│   ├── process/            # Utilities for processing ASCII art font files
│      └── processFile.go  # Reads and parses the ASCII art font file
├── Tests/                  # Contains all test cases and related assets
│   ├── createfiles.sh      # Script to generate test input/output files
│   ├── input1.txt          # Example input files for testing
│   ├── output1.txt         # Expected output files for testing
│   └── ...                 # Additional test input/output files
├── Ascii_test.go           # Top-level integration tests for the application
├── go.mod                  # Go module definition and dependency management file
├── main.go                 # Main application entry point
├── max-docs                # Supplementary documentation or notes
├── standard.txt            # Default ASCII art font definition file
└── testdocs.md             # Documentation specifically for testing
```

## API Documentation (Usage & Core Functions)
This project is a command-line interface (CLI) application. The primary way to interact with it is through the `main.go` executable.

### Basic Usage
To run the application with a given string and the default font (`standard.txt`):
```bash
./ascii-art "Hello World"
```

To specify a different font file (e.g., `shadow.txt` if available):
```bash
./ascii-art "Hello World" shadow.txt
```

### Core Functions

*   **`main.go`**:
    *   The entry point of the application. It orchestrates argument parsing, file processing, and ASCII art generation.

*   **`check.FileEdgeCase(file string) (string, bool)`**:
    *   Performs validation checks on the specified font file, ensuring it has a `.txt` extension, exists, and is not empty.
    *   **Returns**: A status message and a boolean indicating if the file passed all checks.

*   **`print.AsciiArt(data string, dataList [][][]rune)`**:
    *   The central function for generating ASCII art. It takes the input string and a pre-processed list of ASCII character representations (from a font file) and prints the artistic output to standard output. It correctly handles newline characters in the input data.

*   **`printChar.go` (Implicit Function)**:
    *   Likely contains logic to map individual characters to their multi-line ASCII art representation, which is then used by `print.AsciiArt`.

*   **`process.ProcessFile(file string) ([][][]rune, error)`**:
    *   Reads and processes the specified font `.txt` file, converting its content into a structured format (`[][][]rune`) that can be easily used by the `print.AsciiArt` function. Handles file reading errors.