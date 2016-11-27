package main

import (
    "github.com/jacinda/etl/csv_processing"
    "github.com/jacinda/etl/postgres"
)

func main() {
    csv_processing.ProcessCSV()
    postgres.Postgres()
}
