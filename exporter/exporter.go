package exporter

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type IExporter interface {
	CSV(rows [][]string, fileName string)
}

type Exporter struct{}

func (e *Exporter) CSV(rows [][]string, fileName string) {
	file, err := os.Create(fmt.Sprintf("data/%s", fileName))
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(bufio.NewWriter(file))
	w.WriteAll(rows) // calls Flush internally
	if err := w.Error(); err != nil {
		log.Fatalln("error export csv:", err)
	}

	w.Flush()
}

func New() IExporter {
	return &Exporter{}
}
