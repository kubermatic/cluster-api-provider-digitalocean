/*
Copyright 2021.

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

package controllers

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	clusterv1 "sigs.k8s.io/cluster-api-provider-digitalocean/api/v1alpha4"
	controlplanev1alpha4 "sigs.k8s.io/cluster-api-provider-digitalocean/controlplane/doks/api/v1alpha4"
)

// DOKSControlPlaneReconciler reconciles a DOKSControlPlane object
type DOKSControlPlaneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=dokscontrolplanes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=dokscontrolplanes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=dokscontrolplanes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DOKSControlPlane object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DOKSControlPlaneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DOKSControlPlaneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	c, err := ctrl.NewControllerManagedBy(mgr).
		For(&controlplanev1alpha4.DOKSControlPlane{}).
		Build(r)
	if err != nil {
		return errors.Wrap(err, "failed setting up with a controller manager")
	}

	err = c.Watch(
		&source.Kind{Type: &clusterv1.Cluster{}},
		handler.EnqueueRequestsFromMapFunc(r.ClusterToKubeadmControlPlane),
		predicates.ClusterUnpausedAndInfrastructureReady(ctrl.LoggerFrom(ctx)),
	)
	if err != nil {
		return errors.Wrap(err, "failed adding Watch for Clusters to controller manager")
	}

	return err
}
