/*
Copyright The Karmada Authors.

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

package helper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
)

// NewPropagationPolicy will build a PropagationPolicy object.
func NewPropagationPolicy(ns, name string, rsSelectors []policyv1alpha1.ResourceSelector, placement policyv1alpha1.Placement) *policyv1alpha1.PropagationPolicy {
	return &policyv1alpha1.PropagationPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
		Spec: policyv1alpha1.PropagationSpec{
			ResourceSelectors: rsSelectors,
			Placement:         placement,
		},
	}
}

// NewClusterPropagationPolicy will build a ClusterPropagationPolicy object.
func NewClusterPropagationPolicy(policyName string, rsSelectors []policyv1alpha1.ResourceSelector, placement policyv1alpha1.Placement) *policyv1alpha1.ClusterPropagationPolicy {
	return &policyv1alpha1.ClusterPropagationPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name: policyName,
		},
		Spec: policyv1alpha1.PropagationSpec{
			ResourceSelectors: rsSelectors,
			Placement:         placement,
		},
	}
}

// NewOverridePolicy will build a OverridePolicy object.
func NewOverridePolicy(namespace, policyName string, rsSelectors []policyv1alpha1.ResourceSelector, clusterAffinity policyv1alpha1.ClusterAffinity, overriders policyv1alpha1.Overriders) *policyv1alpha1.OverridePolicy {
	return &policyv1alpha1.OverridePolicy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      policyName,
		},
		Spec: policyv1alpha1.OverrideSpec{
			ResourceSelectors: rsSelectors,
			TargetCluster:     &clusterAffinity,
			Overriders:        overriders,
		},
	}
}
