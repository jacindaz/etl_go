package csv_processing
import (
    "testing"
)

func TestReadCSV(t *testing.T) {
    total_lines := ReadCSV()
    if total_lines != 6714 {
        t.Error(
            "For", "csv",
            "expected", 6714,
            "got", total_lines,
        )
    }
}

func TestProcessCSV(t *testing.T) {
    records := ProcessCSV()
    header := records[0]

    if (header[0] != "Candidate Email") {
        t.Error(
            "For", "csv file",
            "expected", "Candidate Email",
            "got", header[0],
        )
    }
}
