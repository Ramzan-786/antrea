// Copyright 2019 Antrea Authors
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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AntreaAgentInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Version         string                 `json:"version,omitempty"`         // Antrea binary version
	PodRef          corev1.ObjectReference `json:"podRef,omitempty"`          // The Pod that Antrea Agent is running in
	NodeRef         corev1.ObjectReference `json:"nodeRef,omitempty"`         // The Node that Antrea Agent is running in
	NodeSubnet      []string               `json:"nodeSubnet,omitempty"`      // Node subnet
	OVSInfo         OVSInfo                `json:"ovsInfo,omitempty"`         // OVS Information
	LocalPodNum     int32                  `json:"localPodNum,omitempty"`     // The number of Pods which the agent is in charge of
	AgentConditions []AgentCondition       `json:"agentConditions,omitempty"` // Agent condition contains types like AgentHealthy
}

type OVSInfo struct {
	Version    string           `json:"version,omitempty"`
	BridgeName string           `json:"bridgeName,omitempty"`
	FlowTable  map[string]int32 `json:"flowTable,omitempty"` // Key: flow table name, Value: flow number
}

type AgentConditionType string

const (
	AgentHealthy           AgentConditionType = "AgentHealthy"           // Status is always set to be True and LastHeartbeatTime is used to check Agent health status.
	ControllerConnectionUp AgentConditionType = "ControllerConnectionUp" // Status False is caused by the reasons: AgentControllerConnectionDown
)

type AgentCondition struct {
	Type              AgentConditionType     `json:"type"`              // One of the AgentConditionType listed above, AgentHealthy or ControllerConnectionUp
	Status            corev1.ConditionStatus `json:"status"`            // Mark certain type status, one of True, False, Unknown
	LastHeartbeatTime metav1.Time            `json:"lastHeartbeatTime"` // The timestamp when AntreaAgentInfo is created/updated, ideally heartbeat interval is 60s
	Reason            string                 `json:"reason,omitempty"`  // Brief reason
	Message           string                 `json:"message,omitempty"` // Human readable message indicating details
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AntreaAgentInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AntreaAgentInfo `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AntreaControllerInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Version                     string                      `json:"version,omitempty"`                     // Antrea binary version
	PodRef                      corev1.ObjectReference      `json:"podRef,omitempty"`                      // The Pod that Antrea Controller is running in
	NodeRef                     corev1.ObjectReference      `json:"nodeRef,omitempty"`                     // The Node that Antrea Controller is running in
	ServiceRef                  corev1.ObjectReference      `json:"serviceRef, omitempty"`                 // Antrea Controller Service
	NetworkPolicyControllerInfo NetworkPolicyControllerInfo `json:"networkPolicyControllerInfo,omitempty"` // NetworkPolicy information
	ConnectedAgentNum           int32                       `json:"connectedAgentNum,omitempty"`           // Number of agents which are connected to this controller
	ControllerConditions        []ControllerCondition       `json:"controllerConditions,omitempty"`        // Controller condition contains types like ControllerHealthy
}

type NetworkPolicyControllerInfo struct {
	PolicyNum        int32 `json:"policyNum,omitempty"`
	AddressGroupNum  int32 `json:"addressGroupNum,omitempty"`
	ApplyingGroupNum int32 `json:"applyingGroupNum,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AntreaControllerInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AntreaControllerInfo `json:"items"`
}

type ControllerConditionType string

const (
	ControllerHealthy ControllerConditionType = "ControllerHealthy" // Status is always set to be True and LastHeartbeatTime is used to check Controller health status.
)

type ControllerCondition struct {
	Type              ControllerConditionType `json:"type"`              // One of the ControllerConditionType listed above, controllerHealthy
	Status            corev1.ConditionStatus  `json:"status"`            // Mark certain type status, one of True, False, Unknown
	LastHeartbeatTime metav1.Time             `json:"lastHeartbeatTime"` // The timestamp when AntreaControllerInfo is created/updated, ideally heartbeat interval is 60s
	Reason            string                  `json:"reason,omitempty"`  // Brief reason
	Message           string                  `json:"message,omitempty"` // Human readable message indicating details
}
