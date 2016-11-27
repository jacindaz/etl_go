package csv_processing

import (
    "log"
    "encoding/csv" //Package csv reads and writes comma-separated values (CSV) files.
    "fmt"          //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
    "io"           //Package io provides basic interfaces to I/O primitives.
    "os"           //Package os provides a platform-independent interface to operating system functionality.
)

func ReadEntireCSV() [][]string {
    // START HERE
    // CAN REFACTOR/MAKE DRY AND PULL LINES 15-21 INTO SEPARATE FUNCTION
    // HAVE READCSV() FUNCTION ALSO USE THIS
    file_name := "/Users/jacindazhong/Documents/jacinda/go_workspace/src/github.com/jacinda/etl/data/oldnavy_55e7ebb5c8b09a0024005661.csv"
    f, err_opening_file := os.Open(file_name)
    defer f.Close()

    reader := csv.NewReader(f)
    reader.TrimLeadingSpace = true
    if err_opening_file != nil { log.Fatal(err_opening_file) }

    records, err := reader.ReadAll()
    if err != nil { log.Fatal(err) }

    return records
}

func ReadCSVLineByLine() int {
    file_name := "/Users/jacindazhong/Documents/jacinda/go_workspace/src/github.com/jacinda/etl/data/oldnavy_55e7ebb5c8b09a0024005661.csv"
    f, err_opening_file := os.Open(file_name)
    defer f.Close()

    reader := csv.NewReader(f)
    reader.TrimLeadingSpace = true

    if err_opening_file != nil {
        log.Fatal(err_opening_file)
    }

    lineCount := 1
    total_lines := 0
    for {
        // read just one record, but we could ReadAll() as well
        _, error_reading_file := reader.Read()
        total_lines += 1

        // end-of-file is fitted into error_reading_file
        if error_reading_file == io.EOF {
            break
        } else if error_reading_file != nil {
            log.Fatal(error_reading_file)
            fmt.Println("error_reading_file:", error_reading_file)

            lineCount += 1
            reader.Read()
            continue
        }
        lineCount += 1
    }

    return total_lines
}
