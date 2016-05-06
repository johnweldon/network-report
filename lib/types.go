package report

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"
)

// Hub describes a step in an MTR route, and associated metrics.
type Hub struct {
	Index        int           `xml:"COUNT,attr"`
	Host         string        `xml:"HOST,attr"`
	Sent         stringInteger `xml:"Snt"`
	Dropped      stringInteger `xml:"Drop"`
	Average      stringFloat   `xml:"Avg"`
	Best         stringFloat   `xml:"Best"`
	Worst        stringFloat   `xml:"Wrst"`
	StdDeviation stringFloat   `xml:"StDev"`
}

// MTR encapsulates an mtr network report.
type MTR struct {
	Source      string `xml:"SRC,attr"`
	Destination string `xml:"DST,attr"`
	Tests       int    `xml:"TESTS,attr"`
	Hops        []Hub  `xml:"HUB"`
}

type stringInteger int64
type stringFloat float64

func (s *stringInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return err
		}
		if cd, ok := t.(xml.CharData); ok {
			str := strings.TrimSpace(string(cd))
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			*s = stringInteger(i)
		}
	}
	return nil
}

func (s *stringFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return err
		}
		if cd, ok := t.(xml.CharData); ok {
			str := strings.TrimSpace(string(cd))
			i, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return err
			}
			*s = stringFloat(i)
		}
	}
	return nil
}
