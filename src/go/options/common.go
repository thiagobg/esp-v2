// Copyright 2019 Google Cloud Platform Proxy Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"time"
)

// CommonOptions describes the possible overrides used by both the ADS bootstrapper and the config generator.
// By defining all the common options in one struct, we prevent duplicate flag initialization and reduce repeated code.
type CommonOptions struct {
	// Flags for envoy
	AdminPort int

	// Flags for tracing
	EnableTracing             bool
	TracingProjectId          string
	TracingStackdriverAddress string
	TracingSamplingRate       float64
	TracingIncomingContext    string
	TracingOutgoingContext    string

	// Flags for metadata
	NonGCP                 bool
	MetadataFetcherTimeout time.Duration
}

// DefaultCommonOptions returns CommonOptions with default values.
//
// The default values are expected to match the default values from the flags.
func DefaultCommonOptions() CommonOptions {
	return CommonOptions{
		AdminPort:                 8001,
		EnableTracing:             false,
		MetadataFetcherTimeout:    5 * time.Second,
		NonGCP:                    false,
		TracingProjectId:          "",
		TracingStackdriverAddress: "",
		TracingSamplingRate:       0.001,
		TracingIncomingContext:    "",
		TracingOutgoingContext:    "",
	}
}
