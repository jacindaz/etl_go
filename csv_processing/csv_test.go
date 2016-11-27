package csv_processing
import (
    "testing"
)

func TestReadCSV(t *testing.T) {
    if ReadCSV() != 6714 {
        t.Error(
            "For", "csv",
            "expected", 6714,
            "got", ReadCSV(),
        )
    }
}
