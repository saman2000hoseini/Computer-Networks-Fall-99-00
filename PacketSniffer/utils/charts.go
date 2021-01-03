package utils

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/wcharczuk/go-chart"
	"os"
)

const (
	pieChartPath = "_PieChart.png"
)

func DrawPieChart(data map[string]uint64, path string) {
	var values []chart.Value
	for key, value := range data {
		values = append(values, chart.Value{
			Label: key,
			Value: float64(value),
		})
	}

	pie := chart.PieChart{
		Canvas: chart.Style{
			FillColor: chart.ColorAlternateBlue,
		},
		Values: values,
	}

	ch := bytes.NewBuffer([]byte{})

	if err := pie.Render(chart.PNG, ch); err != nil {
		return
	}

	writeChart(ch.Bytes(), path+pieChartPath)
}

func writeChart(chart []byte, path string) {
	file, err := os.Create(path)
	if err != nil {
		logrus.Errorf("error creating chart file: %s", err.Error())
		return
	}
	defer file.Close()

	file.Write(chart)
}
