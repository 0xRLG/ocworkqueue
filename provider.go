// Copyright 2018 Ross Guarino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ocworkqueue

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/apimachinery/pkg/util/runtime"
)

type gague struct {
	stat *stats.Int64Measure
	name string
}

func (g *gague) Inc() {
	ctx, err := tag.New(context.TODO(),
		tag.Upsert(Name, g.name))
	if err != nil{
		runtime.HandleError(err)
	}
	stats.Record(ctx, g.stat.M(1))
}

func (g *gague) Dec() {
	ctx, err := tag.New(context.TODO(),
		tag.Upsert(Name, g.name))
	if err != nil{
		runtime.HandleError(err)
	}
	stats.Record(ctx, g.stat.M(-1))
}

type counter struct {
	stat *stats.Int64Measure
	name string
}

func (c *counter) Inc() {
	ctx, err := tag.New(context.TODO(),
		tag.Upsert(Name, c.name))
	if err != nil{
		runtime.HandleError(err)
	}
	stats.Record(ctx, c.stat.M(1))
}

type summary struct {
	stat *stats.Float64Measure
	name string
}

func (s *summary) Observe(val float64) {
	ctx, err := tag.New(context.TODO(),
		tag.Upsert(Name, s.name))
	if err != nil{
		runtime.HandleError(err)
	}
	stats.Record(ctx, s.stat.M(val))
}

type metricsProvider struct{}

func (p *metricsProvider) NewDepthMetric(name string) workqueue.GaugeMetric {
	return &gague{stat: QueueDepth, name: name}
}

func (p *metricsProvider) NewAddsMetric(name string) workqueue.CounterMetric {
	return &counter{stat: QueueAdds, name: name}
}

func (p *metricsProvider) NewLatencyMetric(name string) workqueue.SummaryMetric {
	return &summary{stat: QueueLatency, name: name}
}

func (p *metricsProvider) NewWorkDurationMetric(name string) workqueue.SummaryMetric {
	return &summary{stat: QueueWorkDuration, name: name}
}

func (p *metricsProvider) NewRetriesMetric(name string) workqueue.CounterMetric {
	return &counter{stat: QueueRetries, name: name}
}

// MetricsProvider returns an implementation of the workqueue metrics provider backed by
// opencensus measurements. This provider works with all of the metrics that the
// workqueue can export (Depth,Adds,Latency,WorkDuration,Retries)
//
// You must register the views before any data is collected.
func MetricsProvider() workqueue.MetricsProvider {
	return &metricsProvider{}
}
