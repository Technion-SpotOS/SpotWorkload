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

	"github.com/Technion-SpotOS/SpotWorkload/pkg/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SpotWorkloadReconciler reconciles a SpotWorkload object.
type SpotWorkloadReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpotWorkloadReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	spotWorkload := &v1alpha1.SpotWorkload{}
	if err := r.Get(ctx, req.NamespacedName, spotWorkload); err != nil {
		return ctrl.Result{Requeue: true}, client.IgnoreNotFound(err)
	}

	// Update SpotWorkload status
	// FOR DEMO, TODO: Update
	// Assumed CR is at stage "scheduled" (pending effective scheduling)
	componentStatuses := map[string]v1alpha1.ComponentStatus{}
	for name := range spotWorkload.Spec.Components {
		if err := r.scheduleWorkload(ctx, req.NamespacedName,
			spotWorkload.Status.Components[name].InstanceName); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to update deployment %v: %w", req.NamespacedName, err)
		}

		componentStatuses[name] = v1alpha1.ComponentStatus{
			Stage:        "scheduled",
			InstanceName: spotWorkload.Status.Components[name].InstanceName,
		}
	}

	if err := r.Status().Update(ctx, spotWorkload); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to update SpotWorkload status: %w", err)
	}

	return ctrl.Result{}, nil
}

func (r *SpotWorkloadReconciler) scheduleWorkload(ctx context.Context, deploymentKey types.NamespacedName,
	targetInstance string,
) error {
	// fetch the Deployment
	var deployment appsv1.Deployment
	err := r.Client.Get(ctx, deploymentKey, &deployment)
	if err != nil {
		return err
	}

	// set the toleration for the Deployment
	toleration := corev1.Toleration{
		Key:      "instance-type",
		Operator: corev1.TolerationOpEqual,
		Value:    "spot",
		Effect:   corev1.TaintEffectNoSchedule,
	}

	deployment.Spec.Template.Spec.Tolerations = []corev1.Toleration{toleration}

	nodeSelectorRequirementNodeName := corev1.NodeSelectorRequirement{
		Key:      "kubernetes.io/hostname",
		Operator: corev1.NodeSelectorOpIn,
		Values:   []string{targetInstance},
	}

	nodeSelectorTerm := corev1.NodeSelectorTerm{
		MatchExpressions: []corev1.NodeSelectorRequirement{nodeSelectorRequirementNodeName},
	}

	nodeAffinity := &corev1.NodeAffinity{
		RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
			NodeSelectorTerms: []corev1.NodeSelectorTerm{nodeSelectorTerm},
		},
	}

	deployment.Spec.Template.Spec.Affinity = &corev1.Affinity{
		NodeAffinity: nodeAffinity,
	}

	// Update the Deployment with the modified toleration and node affinity
	return r.Client.Update(ctx, &deployment)
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
