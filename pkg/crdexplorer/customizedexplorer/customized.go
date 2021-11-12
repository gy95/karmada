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

package customizedexplorer

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	webhookutil "k8s.io/apiserver/pkg/util/webhook"

	configv1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	workv1alpha2 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha2"
	"github.com/karmada-io/karmada/pkg/crdexplorer/customizedexplorer/configmanager"
	"github.com/karmada-io/karmada/pkg/util/informermanager"
)

// CustomizedExplorer explore custom resource with webhook configuration.
type CustomizedExplorer struct {
	hookManager   configmanager.ConfigManager
	configManager *webhookutil.ClientManager
}

// NewCustomizedExplorer return a new CustomizedExplorer.
func NewCustomizedExplorer(kubeconfig string, informer informermanager.SingleClusterInformerManager) (*CustomizedExplorer, error) {
	cm, err := webhookutil.NewClientManager(
		[]schema.GroupVersion{configv1alpha1.SchemeGroupVersion},
		configv1alpha1.AddToScheme,
	)
	if err != nil {
		return nil, err
	}
	authInfoResolver, err := webhookutil.NewDefaultAuthenticationInfoResolver(kubeconfig)
	if err != nil {
		return nil, err
	}
	cm.SetAuthenticationInfoResolver(authInfoResolver)
	cm.SetServiceResolver(webhookutil.NewDefaultServiceResolver())

	return &CustomizedExplorer{
		hookManager:   configmanager.NewExploreConfigManager(informer),
		configManager: &cm,
	}, nil
}

// GetReplicas returns the desired replicas of the object as well as the requirements of each replica.
// return matched value to indicate whether there is a matching hook.
func (e *CustomizedExplorer) GetReplicas(ctx context.Context, operation configv1alpha1.OperationType,
	object runtime.Object) (replica int32, replicaRequires *workv1alpha2.ReplicaRequirements, matched bool, err error) {
	return 0, nil, false, err
}

// GetHealthy tells if the object in healthy state.
// return matched value to indicate whether there is a matching hook.
func (e *CustomizedExplorer) GetHealthy(ctx context.Context, operation configv1alpha1.OperationType,
	object runtime.Object) (healthy, matched bool, err error) {
	return false, false, err
}
