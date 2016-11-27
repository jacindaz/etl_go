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
