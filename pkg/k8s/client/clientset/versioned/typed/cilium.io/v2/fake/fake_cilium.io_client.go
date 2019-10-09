// Copyright 2017-2019 Authors of Cilium
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCiliumV2 struct {
	*testing.Fake
}

func (c *FakeCiliumV2) CiliumClusterwideNetworkPolicies() v2.CiliumClusterwideNetworkPolicyInterface {
	return &FakeCiliumClusterwideNetworkPolicies{c}
}

func (c *FakeCiliumV2) CiliumEndpoints(namespace string) v2.CiliumEndpointInterface {
	return &FakeCiliumEndpoints{c, namespace}
}

func (c *FakeCiliumV2) CiliumIdentities() v2.CiliumIdentityInterface {
	return &FakeCiliumIdentities{c}
}

func (c *FakeCiliumV2) CiliumNetworkPolicies(namespace string) v2.CiliumNetworkPolicyInterface {
	return &FakeCiliumNetworkPolicies{c, namespace}
}

func (c *FakeCiliumV2) CiliumNodes() v2.CiliumNodeInterface {
	return &FakeCiliumNodes{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCiliumV2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
