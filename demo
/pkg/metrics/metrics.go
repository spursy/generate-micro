package metrics

import (
	gearMetrics "git.5th.im/lb-public/gear/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	taskDurationsSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "bi",
			Name:      "social_search_task_duration_microseconds",
			Help:      "Summary of the Task call duration in microseconds",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			},
		},
		[]string{"task"},
	)

	taskCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "bi",
			Name:      "social_search_data_task",
			Help:      "social search task execution status",
		},
		[]string{"task", "source", "status"},
	)
)

func Register() {
	prometheus.MustRegister(taskDurationsSummary)
	prometheus.MustRegister(taskCounter)
	gearMetrics.Register()
}

// TaskMeasureContext context of the
type TaskMeasureContext struct {
	timer *prometheus.Timer
}

// MeasureTask measure a method call duration into Prometheus, start a context
// You must call to complete.
// example:
//
// measure := MeasureTask("taskname")
// defer measure.Done()
func MeasureTask(task string) (measure TaskMeasureContext) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(seconds float64) {
		ms := seconds * 1000000 // make microseconds
		taskDurationsSummary.WithLabelValues(task).Observe(ms)
	}))

	measure = TaskMeasureContext{
		timer: timer,
	}
	return
}

// Done to complete method measure, and record duration into Prometheus
func (measure *TaskMeasureContext) Done() {
	measure.timer.ObserveDuration()
}

// defer measure.Done()
func MeasureMethod(module, method string) (measure gearMetrics.MethodMeasureContext) {
	measure = gearMetrics.MeasureMethod(module, method)
	return
}

func CountTask(taskName string, sourceName string, err error) {
	var counterStatus string
	if err != nil {
		counterStatus = "failed"
	} else {
		counterStatus = "success"
	}
	taskCounter.WithLabelValues(taskName, sourceName, counterStatus).Inc()
}
