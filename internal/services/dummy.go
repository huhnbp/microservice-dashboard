package services

import (
	"math/rand"
	"time"
)

type ServiceMetrics struct {
	Name      string
	Status    string
	LatencyMs float64
	ErrorRate float64
	CPU       float64
	Memory    float64
	RiskScore float64
}

// List of dummy services
var Services = []ServiceMetrics{
	{Name: "auth"},
	{Name: "payments"},
	{Name: "checkout"},
}

// UpdateMetrics randomly simulates metrics for each service
func UpdateMetrics() {
	rand.Seed(time.Now().UnixNano())
	for i := range Services {
		// Random latency between 50ms - 500ms
		Services[i].LatencyMs = 50 + rand.Float64()*450

		// Random error rate 0-5%
		Services[i].ErrorRate = rand.Float64() * 0.05

		// Random CPU and Memory
		Services[i].CPU = 10 + rand.Float64()*70      // 10%-80%
		Services[i].Memory = 256 + rand.Float64()*512 // 256MB-768MB

		// RiskScore = weighted sum of latency and error rate
		Services[i].RiskScore = Services[i].LatencyMs*0.002 + Services[i].ErrorRate*100
		// Set Status based on RiskScore
		if Services[i].RiskScore < 1 {
			Services[i].Status = "OK"
		} else if Services[i].RiskScore < 3 {
			Services[i].Status = "Warning"
		} else {
			Services[i].Status = "Critical"
		}
	}
}
