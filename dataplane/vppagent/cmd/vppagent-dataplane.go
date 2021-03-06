// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
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
	"os"
	"os/signal"
	"syscall"

	"github.com/networkservicemesh/networkservicemesh/dataplane/pkg/common"
	"github.com/networkservicemesh/networkservicemesh/dataplane/vppagent/pkg/vppagent"
	"github.com/sirupsen/logrus"
)

func main() {
	// Capture signals to cleanup before exiting
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go common.BeginHealthCheck()

	vppagent := vppagent.CreateVPPAgent()

	registration := common.CreateDataplane(vppagent)

	select {
	case <-c:
		logrus.Info("Closing Dataplane Registration")
		registration.Close()
	}
}
