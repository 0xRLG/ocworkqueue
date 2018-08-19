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
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	NameAttribute = "k8s.client-go.workqueue.name"
)

// Name tag is determined by the name the queue is given when created.
var (
	Name, _ = tag.NewKey(NameAttribute)
)

// This package provides the following Measures for use in custom views.
var (
	QueueDepth        = stats.Int64("opencensus.io/k8s/client-go/workqueue/depth", "Current depth of the workqueue", stats.UnitDimensionless)
	QueueAdds         = stats.Int64("opencensus.io/k8s/client-go/workqueue/adds", "Total number of items added to the queue", stats.UnitDimensionless)
	QueueLatency      = stats.Float64("opencensus.io/k8s/client-go/workqueue/latency", "How long an item stays in a workqueue", stats.UnitMilliseconds)
	QueueWorkDuration = stats.Float64("opencensus.io/k8s/client-go/workqueue/work_duration", "How long processing an item from a workqueue takes", stats.UnitMilliseconds)
	QueueRetries      = stats.Int64("opencensus.io/k8s/client-go/workqueue/retries", "Total number of items re-added to the workqueue", stats.UnitDimensionless)
)

// The default distributions used by the views in this package.
var DefaultMillisecondsDistribution = view.Distribution(0, 0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)

// This package provides the following convenience views. You must register views before
// any data is collected.
var (
	QueueDepthView = &view.View{
		Name:        "opencensus.io/k8s/client-go/workqueue/depth",
		Description: "Sum of items in the queue",
		Measure:     QueueDepth,
		TagKeys:     []tag.Key{Name},
		Aggregation: view.Sum(),
	}
	QueueAddsView = &view.View{
		Name:        "opencensus.io/k8s/client-go/workqueue/adds",
		Description: "Sum of items added to the queue",
		Measure:     QueueAdds,
		TagKeys:     []tag.Key{Name},
		Aggregation: view.Sum(),
	}
	QueueLatencyView = &view.View{
		Name:        "opencensus.io/k8s/client-go/workqueue/latency",
		Description: "Distribution of how long items stay in the workqueue",
		Measure:     QueueLatency,
		TagKeys:     []tag.Key{Name},
		Aggregation: DefaultMillisecondsDistribution,
	}
	QueueWorkDurationView = &view.View{
		Name:        "opencensus.io/k8s/client-go/workqueue/workduration",
		Description: "Distribution of how long items take to be processed",
		Measure:     QueueWorkDuration,
		TagKeys:     []tag.Key{Name},
		Aggregation: DefaultMillisecondsDistribution,
	}
	QueueRetriesView = &view.View{
		Name:        "opencensus.io/k8s/client-go/workqueue/retries",
		Description: "Sum of items re-added to the workqueue",
		Measure:     QueueRetries,
		TagKeys:     []tag.Key{Name},
		Aggregation: view.Sum(),
	}
)

// DefaultViews are the default metrics exported by this package.
var DefaultViews = []*view.View{
	QueueDepthView,
	QueueAddsView,
	QueueLatencyView,
	QueueWorkDurationView,
	QueueRetriesView,
}
