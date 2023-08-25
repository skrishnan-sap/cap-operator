/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sap/cap-operator/pkg/apis/sme.sap.com/v1alpha1"
)

// CAPTenantOperationStepApplyConfiguration represents an declarative configuration of the CAPTenantOperationStep type for use
// with apply.
type CAPTenantOperationStepApplyConfiguration struct {
	Name              *string           `json:"name,omitempty"`
	Type              *v1alpha1.JobType `json:"type,omitempty"`
	ContinueOnFailure *bool             `json:"continueOnFailure,omitempty"`
}

// CAPTenantOperationStepApplyConfiguration constructs an declarative configuration of the CAPTenantOperationStep type for use with
// apply.
func CAPTenantOperationStep() *CAPTenantOperationStepApplyConfiguration {
	return &CAPTenantOperationStepApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CAPTenantOperationStepApplyConfiguration) WithName(value string) *CAPTenantOperationStepApplyConfiguration {
	b.Name = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *CAPTenantOperationStepApplyConfiguration) WithType(value v1alpha1.JobType) *CAPTenantOperationStepApplyConfiguration {
	b.Type = &value
	return b
}

// WithContinueOnFailure sets the ContinueOnFailure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContinueOnFailure field is set to the value of the last call.
func (b *CAPTenantOperationStepApplyConfiguration) WithContinueOnFailure(value bool) *CAPTenantOperationStepApplyConfiguration {
	b.ContinueOnFailure = &value
	return b
}