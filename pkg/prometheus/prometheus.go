package prometheus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/dto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type MonitorService struct{
	Topics map[string]*prometheus.CounterVec
}

// var UrlVisitCount = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "url_visit_count",
// 		Help: "Total number of time URL is visited",
// 	},
// 	[]string{"url"},
// )

// .\prometheus.exe --storage.tsdb.retention.time=30d  for rentiontime

func (m *MonitorService)Metrics(c *fiber.Ctx)error {
	adapter := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	adapter(c.Context())
	return nil
}

func (m *MonitorService)GetCounter(name string)*prometheus.CounterVec{
    return m.Topics[name]
}

func InitMetrics(topics map[string]*prometheus.CounterVec)*MonitorService {
	for _,topic := range topics {
		prometheus.MustRegister(topic)
	}
	return &MonitorService{
		Topics: topics,
	}
}


func (m *MonitorService) GetStats(topic, step, limit string) (*dto.QueryResponse, error) {
    // Set the time range based on the limit
    offset, err := time.ParseDuration(limit)
    fmt.Println(offset)
    if err != nil {
        return nil, fmt.Errorf("limit err")
    }
    now := time.Now().UTC()
    end := now
    start := now.Add(-offset)

    fmt.Println(now,end,start)
    // Construct the Prometheus query
    query := fmt.Sprintf("sum(url_visit_count{job=\"prometheus\",url=\"%s\"})&start=%s&end=%s&step=%s",
        topic, start.Format(time.RFC3339), end.Format(time.RFC3339), step)

    fmt.Println(query)

    // Construct the URL for the Prometheus query endpoint
    url := fmt.Sprintf("http://localhost:9090/api/v1/query?query=%s", query)

    fmt.Println(url)

    // Send the HTTP request to the Prometheus query endpoint
    client := http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    // Read the response body
    var body any
    er := json.NewDecoder(res.Body).Decode(&body)
    // body, err := ioutil.ReadAll(res.Body)
    if er != nil {
        return nil, er
    }

    bodyBytes,err := json.Marshal(body)
    if err != nil {
        return nil,err
    }
    
    // Parse the response body as JSON
    var response dto.QueryResponse
    if err := json.Unmarshal(bodyBytes, &response); err != nil {
        return nil, fmt.Errorf("unmarshal error")
    }

    fmt.Println(response)

    return &response, nil
}