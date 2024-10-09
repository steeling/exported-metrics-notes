package main

import (
	"fmt"
	"os"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promscrape"
	"gopkg.in/yaml.v2"
)

func main() {
	// promrelabel.ParseRelabelConfigs
	data, err := os.ReadFile("./testdata/scrape.yaml")
	check(err)
	var cfgs []promscrape.ScrapeConfig

	yaml.Unmarshal(data, &cfgs)
	// fmt.Println(string(data))
	for _, cfg := range cfgs {
		labeler, err := promrelabel.ParseRelabelConfigs(cfg.MetricRelabelConfigs)
		check(err)
		fmt.Println(labeler)
		labeler.Apply()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
