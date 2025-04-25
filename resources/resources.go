package resources

import (
	_ "embed"
	"strings"
)

//go:embed igns.txt
var igns []byte

func LoadIGNs() []string {
	return strings.Split(string(igns), "\n")
}
