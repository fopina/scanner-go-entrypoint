package scanner

import (
	"os"
	"testing"
)

func TestDefaultBinaryPath(t *testing.T) {
	var s = Scanner{
		Name: "randomNameNoBinary",
	}
	got := s.GetDefaultBinaryPath()
	var expected = "randomNameNoBinary"
	if got != expected {
		t.Errorf("GetDefaultBinaryPath = %s; want %s", got, expected)
	}
	expected = "otherBinary"
	s.DefaultBinary = expected
	got = s.GetDefaultBinaryPath()
	if got != expected {
		t.Errorf("GetDefaultBinaryPath = %s; want %s", got, expected)
	}
}

func TestBuildOptions(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"myscanner", "-H"}
	got := s.BuildOptions()
	ParseOptions(got)
}
