package csv_processing
import (
    "fmt"
    "testing"
    "math/rand"
    "regexp"
)

func TestReadEntireCSV(t *testing.T) {
    records := ReadEntireCSV()
    header := records[0]

    if (header[0] != "Candidate Email") {
        t.Error(
            "For", "csv file",
            "expected", "Candidate Email",
            "got", header[0],
        )
    }

    if (header[1] != "Job Title") {
        t.Error(
            "For", "csv file",
            "expected", "Job Title",
            "got", header[1],
        )
    }

    random_record_index := rand.Intn(6714)
    random_record_email := records[random_record_index][0]
    matched, err := regexp.MatchString("@", random_record_email)

    if !matched {
        t.Error(
            "For", "csv file",
            "expected", fmt.Sprintf("%s to contain an email address", random_record_email),
            "got", err,
        )
    }
}

func TestReadCSVLineByLine(t *testing.T) {
    total_lines := ReadCSVLineByLine()
    if total_lines != 6714 {
        t.Error(
            "For", "csv",
            "expected", 6714,
            "got", total_lines,
        )
    }
}
