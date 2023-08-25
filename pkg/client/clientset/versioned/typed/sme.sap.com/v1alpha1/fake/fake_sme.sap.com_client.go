/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cap-operator contributors
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/sap/cap-operator/pkg/client/clientset/versioned/typed/sme.sap.com/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSmeV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSmeV1alpha1) CAPApplications(namespace string) v1alpha1.CAPApplicationInterface {
	return &FakeCAPApplications{c, namespace}
}

func (c *FakeSmeV1alpha1) CAPApplicationVersions(namespace string) v1alpha1.CAPApplicationVersionInterface {
	return &FakeCAPApplicationVersions{c, namespace}
}

func (c *FakeSmeV1alpha1) CAPTenants(namespace string) v1alpha1.CAPTenantInterface {
	return &FakeCAPTenants{c, namespace}
}

func (c *FakeSmeV1alpha1) CAPTenantOperations(namespace string) v1alpha1.CAPTenantOperationInterface {
	return &FakeCAPTenantOperations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSmeV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}