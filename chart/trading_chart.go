package chart

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateTradingChart(path string) (*charts.Line, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	var dates []string
	var prices []opts.LineData

	for i, row := range rows {
		if i == 0 {
			continue
		}
		price, _ := strconv.ParseFloat(row[1], 64)
		dates = append(dates, row[0])
		prices = append(prices, opts.LineData{Value: price})
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Trading Price Chart",
		}),
	)

	line.SetXAxis(dates).
		AddSeries("Price", prices)

	return line, nil
}
