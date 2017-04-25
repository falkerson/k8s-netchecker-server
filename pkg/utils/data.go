// Copyright 2017 Mirantis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// AgentInfo is payload structure for keepalive message received from agent.
type AgentInfo struct {
	ReportInterval int                 `json:"report_interval"`
	NodeName       string              `json:"nodename"`
	PodName        string              `json:"podname"`
	HostDate       time.Time           `json:"hostdate"`
	LastUpdated    time.Time           `json:"last_updated"`
	LookupHost     map[string][]string `json:"nslookup"`
	IPs            map[string][]string `json:"ips"`
}

// CheckConnectivityInfo is payload structure for server answer to connectivity
// check request.
type CheckConnectivityInfo struct {
	Message  string   `json="message"`
	Absent   []string `json="outdated,omitempty"`
	Outdated []string `json="absent,omitempty"`
}

// AgentMetrics contains Prometheus entities and agent data required for
// reporting metrics for particular agent.
type AgentMetrics struct {
	ErrorCount           prometheus.Counter
	ReportCount          prometheus.Counter
	PodName              string
	ErrorsFromLastReport int
}
