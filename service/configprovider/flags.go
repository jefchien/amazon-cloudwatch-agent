// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package configprovider

import (
	"errors"
	"flag"
	"strings"

	"go.opentelemetry.io/collector/featuregate"
)

const (
	configFlag       = "otelconfig"
	featureGatesFlag = "feature-gates"
)

type configFlagValue struct {
	values []string
	sets   []string
}

func (s *configFlagValue) Set(val string) error {
	s.values = append(s.values, val)
	return nil
}

func (s *configFlagValue) String() string {
	return "[" + strings.Join(s.values, ", ") + "]"
}

func RegisterFlags(flagSet *flag.FlagSet) {
	registerFlags(flagSet, featuregate.GlobalRegistry())
}

func registerFlags(flagSet *flag.FlagSet, reg *featuregate.Registry) {
	cfgs := new(configFlagValue)
	flagSet.Var(cfgs, configFlag, "Locations to the OTEL config file(s), note that only a"+
		" single location can be set per flag entry e.g. `--otelconfig=file:/path/to/first --otelconfig=file:path/to/second`.")

	flagSet.Func("set",
		"Set arbitrary component config property. The component has to be defined in the config file and the flag"+
			" has a higher precedence. Array config properties are overridden and maps are joined. Example --set=processors.batch.timeout=2s",
		func(s string) error {
			idx := strings.Index(s, "=")
			if idx == -1 {
				// No need for more context, see TestSetFlag/invalid_set.
				return errors.New("missing equal sign")
			}
			cfgs.sets = append(cfgs.sets, "yaml:"+strings.TrimSpace(strings.ReplaceAll(s[:idx], ".", "::"))+": "+strings.TrimSpace(s[idx+1:]))
			return nil
		})

	reg.RegisterFlags(flagSet)
}

func GetConfigFlag(flagSet *flag.FlagSet) []string {
	f := flagSet.Lookup(configFlag)
	if f == nil {
		return nil
	}
	cfv, ok := f.Value.(*configFlagValue)
	if !ok {
		return nil
	}
	return append(cfv.values, cfv.sets...)
}

func GetFeatureGatesFlag(flagSet *flag.FlagSet) string {
	f := flagSet.Lookup(featureGatesFlag)
	if f == nil || f.Value == nil {
		return ""
	}
	return f.Value.String()
}
