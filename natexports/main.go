package natexports

import (
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"strings"
)

var cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "custom_cpu_usage",
	Help: "CPU usage from NATS",
})

func main() {
	// Регистрируем метрику
	prometheus.MustRegister(cpuUsage)

	// Подключаемся к NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("metrics.cpu", func(m *nats.Msg) {
		parts := strings.Fields(string(m.Data))
		if len(parts) == 2 {
			if val, err := strconv.ParseFloat(parts[1], 64); err == nil {
				cpuUsage.Set(val)
			}
		}
	})

	// Экспонируем метрики
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
