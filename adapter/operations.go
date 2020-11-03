// Copyright 2020 Layer5, Inc.
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

package adapter

const (
	OperationDescriptionKey  = "description"
	OperationVersionKey      = "version"
	OperationTemplateNameKey = "templateName"
	OperationServiceNameKey  = "serviceName"
)

type OperationRequest struct {
	OperationName     string
	Namespace         string
	Username          string
	CustomBody        string
	IsDeleteOperation bool
	OperationID       string
}

type Operation struct {
	Type       int32             `json:"type,string,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type Operations map[string]*Operation

func (h *BaseHandler) ListOperations() (Operations, error) {
	operations := make(Operations)
	err := h.Config.Operations(&operations)
	if err != nil {
		return nil, err
	}
	return operations, nil
}
