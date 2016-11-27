package postgres

import (
    "fmt"
    "github.com/jacinda/etl/csv_processing"
)

// read in CSV
// create schema based on filename
// 1 schema == 1 client
func Postgres() {
    fmt.Print("Inside postgres function")

    csv_lines := csv_processing.ReadCSVLineByLine()
    fmt.Print("\nNumber of lines in csv: ")
    fmt.Print(csv_lines)
}
