package csv_processing
import (
    "testing"
)

func TestReadCSV(t *testing.T) {
    if readCSV() != 6714 {
        t.Error(
            "For", "csv",
            "expected", 6714,
            "got", readCSV(),
        )
    }
}
