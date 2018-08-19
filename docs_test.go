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

package ocworkqueue_test

import (
	"github.com/0xRLG/ocworkqueue"
	"go.opencensus.io/stats/view"
	"k8s.io/client-go/util/workqueue"
)

func Example() {

	// You must register the views to collect data.
	view.Register(ocworkqueue.DefaultViews...)

	// Register the metrics provider before you create the queue.
	workqueue.SetProvider(ocworkqueue.MetricsProvider())
	q := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "test_queue")

	// use the queue
	_ = q
}
