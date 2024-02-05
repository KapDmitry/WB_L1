package main

import (
	"fmt"
)

// Target
type MetricsProvider interface {
	GetMetrics() map[string]int
}

// Adaptee
type ThirdPartyMetricsLibrary interface {
	Fetch() map[string]interface{}
}

// Concrete Adaptee
type LegacyMetricsLibrary struct{}

func (l *LegacyMetricsLibrary) Fetch() map[string]interface{} {
	return map[string]interface{}{
		"metric1": 10,
		"metric2": 20,
		"metric3": 30,
	}
}

// Adapter
type MetricsAdapter struct {
	Adaptee ThirdPartyMetricsLibrary
}

func (ma *MetricsAdapter) GetMetrics() map[string]int {
	// Адаптация метрик к ожидаемому формату
	rawMetrics := ma.Adaptee.Fetch()
	resultMetrics := make(map[string]int)

	for key, value := range rawMetrics {
		if v, ok := value.(int); ok {
			resultMetrics[key] = v
		}
	}

	return resultMetrics
}

// Concrete Target
type SystemMetrics struct {
	Provider MetricsProvider
}

func PrintMetrics(provider MetricsProvider) {
	metrics := provider.GetMetrics()
	fmt.Println("Metrics:")
	for key, value := range metrics {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func main() {
	// Используем адаптер для работы с LegacyMetricsLibrary через MetricsProvider
	adapter := &MetricsAdapter{Adaptee: &LegacyMetricsLibrary{}}
	systemMetrics := &SystemMetrics{Provider: adapter}

	// Вывод метрик с использованием MetricsProvider
	PrintMetrics(systemMetrics.Provider)
}
