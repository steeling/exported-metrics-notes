package config

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promscrape"
)

type PromConfig struct {
	RawConfig     promscrape.ScrapeConfig
	ParsedMetrics *promrelabel.ParsedConfigs
}
