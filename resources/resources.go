package resources

import (
	_ "embed"
	"strings"
)

//go:embed igns.txt
var igns []byte

//go:embed kmsigns.txt
var kmsIGNs []byte

func LoadIGNs() []string {
	return strings.Split(string(igns), "\r\n")
}

func LoadKMSIGNs() []string {
	return strings.Split(string(kmsIGNs), "\r\n")
}
