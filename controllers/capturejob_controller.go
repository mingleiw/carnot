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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	batchv1 "entropie.ai/carnot/api/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"entropie.ai/carnot/pkg/capture"
)

const (
	addPodNameLabelAnnotation = "entropie.ai/add-pod-name-label"
	podNameLabel              = "entropie.ai/pod-name"
)

// CaptureJobReconciler reconciles a CaptureJob object
type CaptureJobReconciler struct {
	capture.Capture
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch.entropie.ai,resources=capturejobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch.entropie.ai,resources=capturejobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=batch.entropie.ai,resources=capturejobs/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CaptureJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *CaptureJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	r.Capture.WithPort("5678").WithIface("eth0").Start()

	// TODO(user): your logic here
	// Load the CronJob by name
	var captureJob batchv1.CaptureJob

	if req.Name == "capturejob-sample1" {
		log.V(1).Info(req.Namespace)
		if err := r.Get(ctx, req.NamespacedName, &captureJob); err != nil {
			log.Error(err, "unable to fetch CaptureJob")
			if apierrors.IsNotFound(err) {
				// we'll ignore not-found errors, since we can get them on deleted requests.
				return ctrl.Result{}, nil
			}

			return ctrl.Result{}, err
		}

		log.V(1).Info("found the capture job", "port", captureJob.Spec.ListeningPort)
	}

	// Load the pod
	var service corev1.Service
	if err := r.Get(ctx, req.NamespacedName, &service); err != nil {
		if apierrors.IsNotFound(err) {
			// we'll ignore not-found errors, since we can get them on deleted requests.
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch service")
		return ctrl.Result{}, err
	}

	// labelShouldBePresent := pod.Annotations[addPodNameLabelAnnotation] == "true"
	// labelIsPresent := pod.Labels[podNameLabel] == pod.Name

	log.V(1).Info(service.Name)
	log.V(1).Info(service.Spec.ClusterIP)
	log.V(1).Info(service.Spec.Ports[0].String())

	// if labelShouldBePresent == labelIsPresent {
	// 	// The desired state and actual state of the Pod are the same.
	// 	// No further action is required by the operator at this moment.
	// 	log.Info("no update required")
	// 	return ctrl.Result{}, nil
	// }

	// if labelShouldBePresent {
	// 	// If the label should be set but is not, set it.
	// 	if pod.Labels == nil {
	// 		pod.Labels = make(map[string]string)
	// 	}
	// 	pod.Labels[podNameLabel] = pod.Name
	// 	log.Info("adding label")
	// } else {
	// 	// If the label should not be set but is, remove it.
	// 	delete(pod.Labels, podNameLabel)
	// 	log.Info("removing label")
	// }

	// if err := r.Update(ctx, &pod); err != nil {
	// 	if apierrors.IsConflict(err) {
	// 		// The Pod has been updated since we read it.
	// 		// Requeue the Pod to try to reconciliate again.
	// 		return ctrl.Result{Requeue: true}, nil
	// 	}
	// 	if apierrors.IsNotFound(err) {
	// 		// The Pod has been deleted since we read it.
	// 		// Requeue the Pod to try to reconciliate again.
	// 		return ctrl.Result{Requeue: true}, nil
	// 	}
	// 	log.Error(err, "unable to update Pod")

	// 	return ctrl.Result{}, err
	// }

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CaptureJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.CaptureJob{}).
		Watches(&source.Kind{Type: &corev1.Service{}},
			&handler.EnqueueRequestForObject{}).
		Complete(r)
}
