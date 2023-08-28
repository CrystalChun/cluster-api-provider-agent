/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

type AgentReference struct {
	// Name is unique within a namespace to reference an agent resource.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace defines the space within which the agent name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// AgentMachineTemplateResource describes the data needed to create an AgentMachine from a template
type AgentMachineTemplateResource struct {
	// Spec is the specification of the desired behavior of the machine.
	Spec AgentMachineSpec `json:"spec"`
}
