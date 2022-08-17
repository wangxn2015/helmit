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
	"context"
	"testing"
	"time"

	"github.com/wangxn2015/helmit/pkg/kubernetes"

	"github.com/stretchr/testify/assert"
	"github.com/wangxn2015/helmit/pkg/helm"
	"github.com/wangxn2015/helmit/pkg/test"
)

// ChartTestSuite is a test for chart deployment
type ChartTestSuite struct {
	test.Suite
}

// TestLocalInstall tests a local chart installation
func (s *ChartTestSuite) TestLocalInstall(t *testing.T) {
	atomix := helm.Chart("kubernetes-controller").
		Release("atomix-controller").
		Set("scope", "Namespace")
	err := atomix.Install(true)
	assert.NoError(t, err)

	raft := helm.Chart("raft-storage-controller").
		Release("raft-storage-controller").
		Set("scope", "Namespace")

	err = raft.Install(true)
	assert.NoError(t, err)

	topo := helm.Chart("onos-topo").
		Release("onos-topo").
		Set("store.controller", "atomix-controller-kubernetes-controller:5679")
	err = topo.Install(true)
	assert.NoError(t, err)

	client := kubernetes.NewForReleaseOrDie(topo)

	pods, err := client.CoreV1().Pods().List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, pods, 2)

	deployment, err := client.AppsV1().
		Deployments().
		Get(context.Background(), "onos-topo")
	assert.NoError(t, err)

	pods, err = deployment.Pods().List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, pods, 1)
	pod := pods[0]
	err = pod.Delete(context.Background())
	assert.NoError(t, err)

	err = deployment.Wait(context.Background(), 1*time.Minute)
	assert.NoError(t, err)

	pods, err = deployment.Pods().List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, pods, 1)
	assert.NotEqual(t, pod.Name, pods[0].Name)

	services, err := client.CoreV1().Services().List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, services, 2)

	err = atomix.Uninstall()
	assert.NoError(t, err)

	err = raft.Uninstall()
	assert.NoError(t, err)

	err = topo.Uninstall()
	assert.NoError(t, err)
}

// TestRemoteInstall tests a remote chart installation
func (s *ChartTestSuite) TestRemoteInstall(t *testing.T) {
	kafka := helm.Chart("kafka", "http://storage.googleapis.com/kubernetes-charts-incubator").
		Release("kafka").
		Set("replicas", 1).
		Set("zookeeper.replicaCount", 1)
	err := kafka.Install(true)
	assert.NoError(t, err)

	err = kafka.Uninstall()
	assert.NoError(t, err)
}
