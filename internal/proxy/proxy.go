package proxy

import (
	"bufio"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/steeling/prom-filter/internal/config"
	"github.com/steeling/prom-filter/internal/metrics"
)

type ConfigGetter interface {
	GetConfig() map[string]config.PromConfig
}

type FilterProxy struct {
	ConfigGetter

	nodeIP string
}

// ServeHTTP is the handler used to proxy HTTP requests.
func (p *FilterProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Debug("request received", "uri", r.RequestURI)
	metrics.RequestCount.Inc()
	job := r.URL.Query().Get("job")
	conf := p.GetConfig()
	c, _ := conf[job]
	if len(c.RawConfig.StaticConfigs) == 0 || len(c.RawConfig.StaticConfigs[0].Targets) == 0 {
		http.Error(w, "no targets found", http.StatusNotFound)
		metrics.ProxyErrorCount.WithLabelValues("no_targets_found").Inc()
		return
	}

	client := http.Client{
		Timeout: c.RawConfig.ScrapeTimeout.Duration(),
	}

	uri := url.URL{
		Scheme: c.RawConfig.Scheme,
		Host:   c.RawConfig.StaticConfigs[0].Targets[0],
		Path:   c.RawConfig.MetricsPath,
	}

	resp, err := client.Get(uri.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		metrics.ProxyErrorCount.WithLabelValues("http_error").Inc()
		return
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		// parse line and see if we keep it, then return it.
		labels := toLabels(line)
		c.ParsedMetrics.Apply(labels, 0)
	}

}
