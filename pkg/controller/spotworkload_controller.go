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

package controllers

import (
	"context"
	"fmt"

	"github.com/Technion-SpotOS/SpotWorkload/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SpotWorkloadReconciler reconciles a SpotWorkload object
type SpotWorkloadReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpotWorkloadReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	SpotWorkload := &v1alpha1.SpotWorkload{}
	_ = r.Get(ctx, req.NamespacedName, SpotWorkload)

	// TODO: create controllers for:
	// 1. Creating a spot instance
	// 2. Deploying the instance as a node
	// 3. Updating CR Status with the node name

	return ctrl.Result{}, nil
}

// setupSpotWorkloadController sets up the controller with the Manager.
func setupSpotWorkloadController(mgr ctrl.Manager) error {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.SpotWorkload{}).
		Complete(&SpotWorkloadReconciler{mgr.GetClient(), mgr.GetScheme()}); err != nil {
		return fmt.Errorf("failed to add spot-workload controller to the manager: %w", err)
	}

	return nil
}
