package main

import "github.com/hartfordfive/csv-to-openmetrics/cmd"

func init() {
	//log.SetFormatter(&logging.LogFormatPlain)
	//log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)
}

func main() {

	cmd.Execute()
}
