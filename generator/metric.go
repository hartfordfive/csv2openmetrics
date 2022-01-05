package generator

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
)

type Metric struct {
	Name      string
	Timestamp string
	Labels    map[string]string
	Value     float64
	Type      string
	Help      string
}

// ValueCanBeInt is used to determine if the float value can really be represented as an int
func ValueCanBeInt(val float64) bool {
	if val-math.Trunc(val) == 0 {
		return true
	}
	return false
}

func convertToIntString(val float64) string {
	return strconv.FormatInt(int64(val), 10)
}

func (m *Metric) ToString() string {

	fmt.Println("Date passed in:", m.Timestamp)
	ts, err := dateparse.ParseLocal(m.Timestamp)
	if err != nil {
		panic(err.Error())
	}

	if len(m.Labels) >= 1 {
		flattendLabels := []string{}
		for k, v := range m.Labels {
			flattendLabels = append(flattendLabels, fmt.Sprintf("%s=\"%s\"", k, v))
		}
		if ValueCanBeInt(m.Value) {
			return fmt.Sprintf("%s{%s} %s %s", m.Name, strings.Join(flattendLabels, ", "), convertToIntString(m.Value), ts)
		} else {
			return fmt.Sprintf("%s{%s} %f %s", m.Name, strings.Join(flattendLabels, ", "), m.Value, ts)
		}

	} else {

		if ValueCanBeInt(m.Value) {
			return fmt.Sprintf("%s %s", m.Name, convertToIntString(m.Value))
		} else {
			return fmt.Sprintf("%s %f", m.Name, m.Value)
		}
	}
}
