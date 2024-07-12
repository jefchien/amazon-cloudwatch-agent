// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package defaultcomponents

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"golang.org/x/exp/maps"

	"github.com/aws/amazon-cloudwatch-agent/internal/util/collections"
)

func TestComponents(t *testing.T) {
	factories, err := Factories()
	assert.NoError(t, err)
	wantReceivers := []string{
		"awscontainerinsightreceiver",
		"awsxray",
		"jmx",
		"otlp",
		"tcplog",
		"udplog",
		// ADOT-only
		"awsecscontainermetrics",
		"filelog",
		"jaeger",
		"kafka",
		"prometheus",
		"statsd",
		"zipkin",
	}
	gotReceivers := collections.MapSlice(maps.Keys(factories.Receivers), component.Type.String)
	assert.Equal(t, len(wantReceivers), len(gotReceivers))
	for _, typeStr := range wantReceivers {
		assert.Contains(t, gotReceivers, typeStr)
	}

	wantProcessors := []string{
		"awsapplicationsignals",
		"batch",
		"cumulativetodelta",
		"ec2tagger",
		"filter",
		"gpuattributes",
		"metricstransform",
		"resourcedetection",
		"resource",
		"transform",
		// ADOT-only
		"attributes",
		"deltatorate",
		"groupbytrace",
		"k8sattributes",
		"memory_limiter",
		"experimental_metricsgeneration",
		"probabilistic_sampler",
		"span",
		"tail_sampling",
	}
	gotProcessors := collections.MapSlice(maps.Keys(factories.Processors), component.Type.String)
	assert.Equal(t, len(wantProcessors), len(gotProcessors))
	for _, typeStr := range wantProcessors {
		assert.Contains(t, gotProcessors, typeStr)
	}

	wantExporters := []string{
		"awscloudwatchlogs",
		"awsemf",
		"awscloudwatch",
		"awsxray",
		"debug",
	}
	gotExporters := collections.MapSlice(maps.Keys(factories.Exporters), component.Type.String)
	assert.Equal(t, len(wantExporters), len(gotExporters))
	for _, typeStr := range wantExporters {
		assert.Contains(t, gotExporters, typeStr)
	}

	wantExtensions := []string{
		"agenthealth",
		"awsproxy",
		// ADOT-only
		"ecs_observer",
		"file_storage",
		"health_check",
		"memory_ballast",
		"pprof",
		"sigv4auth",
		"zpages",
	}
	gotExtensions := collections.MapSlice(maps.Keys(factories.Extensions), component.Type.String)
	assert.Equal(t, len(wantExtensions), len(gotExtensions))
	for _, typeStr := range wantExtensions {
		assert.Contains(t, gotExtensions, typeStr)
	}
}
