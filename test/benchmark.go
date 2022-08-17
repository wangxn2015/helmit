// Copyright 2020-present Open Networking Foundation.
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

package test

import (
	"time"

	"github.com/wangxn2015/helmit/pkg/benchmark"
	"github.com/wangxn2015/helmit/pkg/helm"
	"github.com/wangxn2015/helmit/pkg/input"
)

// ChartBenchmarkSuite benchmarks a Helm chart
type ChartBenchmarkSuite struct {
	benchmark.Suite
	value input.Source
}

// SetupSuite :: benchmark
func (s *ChartBenchmarkSuite) SetupSuite(b *input.Context) error {
	atomix := helm.Chart("kubernetes-controller").
		Release("atomix-controller").
		Set("scope", "Namespace")

	err := atomix.Install(true)
	if err != nil {
		return err
	}

	err = atomix.Uninstall()
	if err != nil {
		return err
	}
	return nil

}

// SetupWorker :: benchmark
func (s *ChartBenchmarkSuite) SetupWorker(b *input.Context) error {
	s.value = input.RandomString(8)
	return nil
}

// BenchmarkTest :: benchmark
func (s *ChartBenchmarkSuite) BenchmarkTest(b *benchmark.Benchmark) error {
	println(s.value.Next().String())
	time.Sleep(time.Second)
	return nil
}
