// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package configprovider

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpsprovider"
	"go.opentelemetry.io/collector/confmap/provider/yamlprovider"
	"go.opentelemetry.io/collector/otelcol"
)

func GetConfigProviderSettings(configURIs ...string) otelcol.ConfigProviderSettings {
	providers := []confmap.Provider{
		fileprovider.NewWithSettings(confmap.ProviderSettings{}),
		envprovider.NewWithSettings(confmap.ProviderSettings{}),
		yamlprovider.NewWithSettings(confmap.ProviderSettings{}),
		httpprovider.NewWithSettings(confmap.ProviderSettings{}),
		httpsprovider.NewWithSettings(confmap.ProviderSettings{}),
		s3provider.New(),
	}
	return otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:       configURIs,
			Converters: []confmap.Converter{expandconverter.New(confmap.ConverterSettings{})},
			Providers:  toProviderMap(providers),
		},
	}
}

func toProviderMap(providers []confmap.Provider) map[string]confmap.Provider {
	providerMap := make(map[string]confmap.Provider, len(providers))
	for _, provider := range providers {
		providerMap[provider.Scheme()] = provider
	}
	return providerMap
}
