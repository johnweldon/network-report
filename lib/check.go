package report

import (
	"bytes"
	"encoding/xml"
	"os/exec"
)

// Check types can test a destination (host name or IP address) and
// return an MTR report.
type Check interface {
	Test(destination string) (*MTR, error)
}

// NewChecker returns a Check that tests the real network.
func NewChecker() Check {
	return networkCheck{}
}

// NewMockChecker returns a Check that just uses sample data.
func NewMockChecker() Check {
	return mockCheck{}
}

type networkCheck struct{}

func (networkCheck) Test(destination string) (*MTR, error) {
	cmd := exec.Command("mtr", "--xml", "-o", "LSD", destination)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return unmarshal(out.Bytes())
}

type mockCheck struct{}

func (mockCheck) Test(destination string) (*MTR, error) {
	return unmarshal([]byte(sample))
}

func unmarshal(data []byte) (*MTR, error) {
	v := MTR{}
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
