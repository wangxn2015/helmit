// Copyright 2019-present Open Networking Foundation.
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

package main

import (
	"github.com/wangxn2015/helmit/pkg/benchmark"
	"github.com/wangxn2015/helmit/pkg/registry"
	"github.com/wangxn2015/helmit/pkg/simulation"
	"github.com/wangxn2015/helmit/pkg/test"
	tests "github.com/wangxn2015/helmit/test"
	"os"
)

func main() {
	jobType := os.Getenv("JOB_TYPE")
	switch jobType {
	case "test":
		registry.RegisterTestSuite("chart", &tests.ChartTestSuite{})
		test.Main()
	case "benchmark":
		registry.RegisterBenchmarkSuite("chart", &tests.ChartBenchmarkSuite{})
		benchmark.Main()
	case "simulation":
		registry.RegisterSimulationSuite("chart", &tests.ChartSimulationSuite{})
		simulation.Main()
	}
}
