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

	"github.com/wangxn2015/helmit/pkg/helm"
	"github.com/wangxn2015/helmit/pkg/simulation"
)

// ChartSimulationSuite :: simulation
type ChartSimulationSuite struct {
	simulation.Suite
}

// SetupSimulation :: simulation
func (s *ChartSimulationSuite) SetupSimulation(sim *simulation.Simulator) error {
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

// ScheduleSimulator :: simulation
func (s *ChartSimulationSuite) ScheduleSimulator(sim *simulation.Simulator) {
	sim.Schedule("foo", s.SimulateFoo, 1*time.Second, 1)
	sim.Schedule("bar", s.SimulateBar, 5*time.Second, 1)
	sim.Schedule("baz", s.SimulateBaz, 30*time.Second, 1)
}

// SimulateFoo :: simulation
func (s *ChartSimulationSuite) SimulateFoo(sim *simulation.Simulator) error {
	println(sim.Arg("foo").String("<none>"))
	return nil
}

// SimulateBar :: simulation
func (s *ChartSimulationSuite) SimulateBar(sim *simulation.Simulator) error {
	println(sim.Arg("bar").String("<none>"))
	return nil
}

// SimulateBaz :: simulation
func (s *ChartSimulationSuite) SimulateBaz(sim *simulation.Simulator) error {
	println(sim.Arg("baz").String("<none>"))
	return nil
}
