/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	smesapcomv1alpha1 "github.com/sap/cap-operator/pkg/apis/sme.sap.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CAPApplicationVersionStatusApplyConfiguration represents an declarative configuration of the CAPApplicationVersionStatus type for use
// with apply.
type CAPApplicationVersionStatusApplyConfiguration struct {
	GenericStatusApplyConfiguration `json:",inline"`
	State                           *smesapcomv1alpha1.CAPApplicationVersionState `json:"state,omitempty"`
	FinishedJobs                    []string                                      `json:"finishedJobs,omitempty"`
}

// CAPApplicationVersionStatusApplyConfiguration constructs an declarative configuration of the CAPApplicationVersionStatus type for use with
// apply.
func CAPApplicationVersionStatus() *CAPApplicationVersionStatusApplyConfiguration {
	return &CAPApplicationVersionStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *CAPApplicationVersionStatusApplyConfiguration) WithObservedGeneration(value int64) *CAPApplicationVersionStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *CAPApplicationVersionStatusApplyConfiguration) WithConditions(values ...v1.Condition) *CAPApplicationVersionStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *CAPApplicationVersionStatusApplyConfiguration) WithState(value smesapcomv1alpha1.CAPApplicationVersionState) *CAPApplicationVersionStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithFinishedJobs adds the given value to the FinishedJobs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FinishedJobs field.
func (b *CAPApplicationVersionStatusApplyConfiguration) WithFinishedJobs(values ...string) *CAPApplicationVersionStatusApplyConfiguration {
	for i := range values {
		b.FinishedJobs = append(b.FinishedJobs, values[i])
	}
	return b
}