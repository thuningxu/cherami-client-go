// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metrics

// MetricName is the name of the metric
type MetricName string

// MetricType is the type of the metric, which can be one of the 3 below
type MetricType int

// MetricTypes which are supported
const (
	Counter MetricType = iota
	Timer
	Gauge
)

const (
	// ServiceNameTagName is the tag name to identify partner service which uses Chermai client
	ServiceNameTagName = "serviceName"
	// DeploymentTagName is the tag name to identify current deployment name
	DeploymentTagName = "deployment"

	// PublishMessageRate is the rate of message wrote to input
	PublishMessageRate = "cherami.publish.message.rate"
	// PublishMessageFailedRate is the rate of message try writting to input but failed
	PublishMessageFailedRate = "cherami.publish.message.failed"
	// PublishMessageLatency is the latency of message wrote to input
	PublishMessageLatency = "cherami.publish.message.latency"
	// PublishAckRate is the rate of ack got from input
	PublishAckRate = "cherami.publish.ack.rate"
	// PublishReconfigureRate is the rate of reconfiguration happening
	PublishReconfigureRate = "cherami.publish.reconfigure.rate"
	// PublishNumConnections is the number of connections with input
	PublishNumConnections = "cherami.publish.connections"
	// PublishNumInflightMessagess is the number of inflight messages hold locally by publisher
	PublishNumInflightMessagess = "cherami.publish.message.inflights"

	// ConsumeMessageRate is the rate of message got from output
	ConsumeMessageRate = "cherami.consume.message.rate"
	// ConsumeCreditRate is the rate of credit sent to output
	ConsumeCreditRate = "cherami.consume.credit.rate"
	// ConsumeCreditFailedRate is the rate of credit try sending to output but failed
	ConsumeCreditFailedRate = "cherami.consume.credit.failed"
	// ConsumeCreditLatency is the latency of credit sent to output
	ConsumeCreditLatency = "cherami.consume.credit.latency"
	// ConsumeAckRate is the rate of ack sent to output
	ConsumeAckRate = "cherami.consume.ack.rate"
	// ConsumeAckFailedRate is the rate of ack try sending to output but failed
	ConsumeAckFailedRate = "cherami.consume.ack.failed"
	// ConsumeNackRate is the rate of nack sent to output
	ConsumeNackRate = "cherami.consume.nack.rate"
	// ConsumeReconfigureRate is the rate of reconfiguration happening
	ConsumeReconfigureRate = "cherami.consume.reconfigure.rate"
	// ConsumeNumConnections is the number of connections with output
	ConsumeNumConnections = "cherami.consume.connections"
	// ConsumeLocalCredits is the number of credit hold locally by consumer
	ConsumeLocalCredits = "cherami.consume.credit.local"
)

// MetricDefs contains definition of metrics to its type mapping
var MetricDefs = map[MetricName]MetricType{
	PublishMessageRate:          Counter,
	PublishMessageFailedRate:    Counter,
	PublishMessageLatency:       Timer,
	PublishAckRate:              Counter,
	PublishReconfigureRate:      Counter,
	PublishNumConnections:       Gauge,
	PublishNumInflightMessagess: Gauge,

	ConsumeMessageRate:      Counter,
	ConsumeCreditRate:       Counter,
	ConsumeCreditFailedRate: Counter,
	ConsumeCreditLatency:    Timer,
	ConsumeAckRate:          Counter,
	ConsumeAckFailedRate:    Counter,
	ConsumeNackRate:         Counter,
	ConsumeReconfigureRate:  Counter,
	ConsumeNumConnections:   Gauge,
	ConsumeLocalCredits:     Gauge,
}
