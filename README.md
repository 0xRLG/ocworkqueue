# OpenCensus Workqueue MetricsProvider

[![GoDoc][godoc-image]][godoc-url]

This package implements a `k8s.io/client-go/util/workqueue.MetricsProvider`
backed by OpenCensus measurements.

## Installation

```
$ go get -u github.com/0xRLG/ocworkqueue
```

The API of opencensus is still evolving, see their: [Deprecation Policy][deprecation-policy].
The use of vendoring or a dependency management tool is recommended.

## Prerequisites

OpenCenusus requires Go 1.8 or later, therefor this project also requires Go 1.8 or later.

## Usage

See the [GoDoc][godoc-url] for usage information.

[godoc-image]: https://godoc.org/github.com/0xRLG/ocworkqueue?status.svg
[godoc-url]: https://godoc.org/github.com/0xRLG/ocworkqueue

[deprecation-policy]:https://github.com/census-instrumentation/opencensus-go/blob/master/README.md#deprecation-policy
