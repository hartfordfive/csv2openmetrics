package generator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Metric struct {
	Name   string
	Labels map[string]string
	Value  float64
	Type   string
	Help   string
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
	if len(m.Labels) >= 1 {
		flattendLabels := []string{}
		for k, v := range m.Labels {
			flattendLabels = append(flattendLabels, fmt.Sprintf("%s=\"%s\"", k, v))
		}
		if ValueCanBeInt(m.Value) {
			return fmt.Sprintf("%s{%s} %s", m.Name, strings.Join(flattendLabels, ", "), convertToIntString(m.Value))
		} else {
			return fmt.Sprintf("%s{%s} %f", m.Name, strings.Join(flattendLabels, ", "), m.Value)
		}

	} else {

		if ValueCanBeInt(m.Value) {
			return fmt.Sprintf("%s %s", m.Name, convertToIntString(m.Value))
		} else {
			return fmt.Sprintf("%s %f", m.Name, m.Value)
		}
	}
}
