package dummy

import (
	"math/rand"

	"github.com/huhnbp/microservice-dashboard/graph/model"
)

var Services = []model.Service{
	{Name: "auth", Status: "OK"},
	{Name: "payments", Status: "OK"},
	{Name: "checkout", Status: "OK"},
	{Name: "inventory", Status: "OK"},
	{Name: "notifications", Status: "OK"},
}

func GetDummyMetrics(s *model.Service) {
	if rand.Float64() < 0.05 {
		s.Status = "DOWN"
	} else {
		s.Status = "OK"
	}

	// Random metrics
	s.LatencyMs = rand.Float64()*400 + 50
	s.ErrorRate = rand.Float64() * 0.1
	s.CPUUsage = rand.Float64() * 100
	s.MemoryUsage = rand.Float64()*500 + 100
	s.Uptime = int32(rand.Intn(86400))
	s.RequestRate = int32(rand.Intn(2000) + 50)
}
