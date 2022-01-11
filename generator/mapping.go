package generator

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

/*
files:
  -
  	filename: test_temperature_data.csv
    title_row: true
    timestamp_col_index: 1
    fields:
      1: date;timestamp
      2: prov;label
      3: city;label
      4: high;value
      5: low;value
*/

type CsvMappingConfg struct {
	Files []CsvFile `yaml:"files"`
}

type CsvFile struct {
	Filename     string         `yaml:"filename"`
	TitleRow     bool           `yaml:"title_row"`
	SeriesPrefix string         `yaml:"series_prefix"`
	Fields       map[int]string `yaml:"fields"`
	OutputDir    string         `yaml:"output_dir"`
}

func newMappingConfig(mappingConf string) *CsvMappingConfg {
	m := &CsvMappingConfg{}
	yamlFile, err := ioutil.ReadFile(mappingConf)
	if err != nil {
		log.Printf("Error: %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return m
}
