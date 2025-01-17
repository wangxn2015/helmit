// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"

	"github.com/onosproject/helmit/pkg/job"
	"github.com/onosproject/helmit/pkg/registry"
	"google.golang.org/grpc"
)

// newCoordinator returns a new test coordinator
func newCoordinator(config *Config) (*Coordinator, error) {
	return &Coordinator{
		config: config,
		runner: job.NewNamespace(config.Namespace),
	}, nil
}

// Coordinator coordinates workers for suites of tests
type Coordinator struct {
	config *Config
	runner *job.Runner
}

// Run runs the tests
func (c *Coordinator) Run() (int, error) {
	var returnCode int
	for iteration := 1; iteration <= c.config.Iterations || c.config.Iterations < 0; iteration++ {
		suites := c.config.Suites
		if len(suites) == 0 || suites[0] == "" {
			suites = registry.GetTestSuites()
		}
		returnCode = 0
		for _, suite := range suites {
			jobID := newJobID(c.config.ID+"-"+strconv.Itoa(iteration), suite)
			env := c.config.Env
			if env == nil {
				env = make(map[string]string)
			}
			env[testTypeEnv] = string(testTypeWorker)
			config := &Config{
				Config: &job.Config{
					ID:              jobID,
					Namespace:       c.config.Config.Namespace,
					ServiceAccount:  c.config.Config.ServiceAccount,
					Image:           c.config.Config.Image,
					ImagePullPolicy: c.config.Config.ImagePullPolicy,
					Executable:      c.config.Config.Executable,
					Context:         c.config.Config.Context,
					Values:          c.config.Config.Values,
					ValueFiles:      c.config.Config.ValueFiles,
					Env:             env,
					Timeout:         c.config.Config.Timeout,
					NoTeardown:      c.config.Config.NoTeardown,
					Secrets:         c.config.Config.Secrets,
					Args:            c.config.Config.Args,
				},
				Suites:     []string{suite},
				Tests:      c.config.Tests,
				Iterations: c.config.Iterations,
				Args:       c.config.Args,
			}
			task := &WorkerTask{
				runner: c.runner,
				config: config,
			}
			status, err := task.Run()
			if err != nil {
				return status, err
			} else if returnCode == 0 {
				returnCode = status
			}
		}
		if returnCode == 0 {
			return 0, nil
		}
	}
	return returnCode, nil
}

// newJobID returns a new unique test job ID
func newJobID(testID, suite string) string {
	return fmt.Sprintf("%s-%s", testID, suite)
}

// WorkerTask manages a single test job for a test worker
type WorkerTask struct {
	runner *job.Runner
	config *Config
}

// Run runs the worker job
func (t *WorkerTask) Run() (int, error) {
	job := &job.Job{
		Config:    t.config.Config,
		JobConfig: t.config,
		Type:      testJobType,
	}

	err := t.runner.StartJob(job)
	if err != nil {
		return 0, err
	}

	address := fmt.Sprintf("%s:5000", job.ID)
	conn, err := grpc.Dial(address,
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()),
		grpc.WithInsecure())

	if err != nil {
		return 0, err
	}

	client := NewWorkerServiceClient(conn)
	_, err = client.RunTests(context.Background(), &TestRequest{
		Suite: t.config.Suites[0],
		Tests: t.config.Tests,
		Args:  t.config.Args,
	})

	if err != nil {
		return 0, err
	}

	status, err := t.runner.WaitForExit(job)
	if err != nil {
		return 0, err
	}
	return status, err
}
