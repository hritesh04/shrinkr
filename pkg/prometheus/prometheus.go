package prometheus

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type MonitorService struct {
	Client v1.API
	Topics map[string]*prometheus.CounterVec
}

var topics = map[string]*prometheus.CounterVec{
	"UrlVisitCount": prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "url_visit_count",
			Help: "Total number of time URL is visited",
		},
		[]string{"url"},
	),
}

func (m *MonitorService) Metrics(c *fiber.Ctx) error {
	adapter := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	adapter(c.Context())
	return nil
}

func (m *MonitorService) Increment(name, topic string) {
	m.Topics[name].WithLabelValues(topic).Inc()
}

func NewMonitorService(url string) *MonitorService {
	connection, err := api.NewClient(api.Config{
		Address: url,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	client := v1.NewAPI(connection)
	for _, topic := range topics {
		prometheus.MustRegister(topic)
	}
	return &MonitorService{
		Topics: topics,
		Client: client,
	}
}

func (m *MonitorService) GetStats(topic, step, limit string) ([]byte, error) {
	offset, err := time.ParseDuration(limit)
	if err != nil {
		return nil, fmt.Errorf("limit err")
	}
	stepDuration, err := time.ParseDuration(step)
	if err != nil {
		return nil, fmt.Errorf("step err")
	}

	start := time.Now().Add(-offset)
	end := time.Now()

	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(stepDuration)

	query := fmt.Sprintf(`url_visit_count{url=%s}`, topic)
	queryRange, _, err := m.Client.QueryRange(context.Background(), query, v1.Range{
		Start: start,
		End:   end,
		Step:  stepDuration,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(queryRange)
	value, err := queryRange.Type().MarshalJSON()
	if err != nil {
		return nil, err
	}
	return value, nil
}
