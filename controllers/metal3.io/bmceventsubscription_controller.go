/*

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

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	metal3v1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
)

type BMCEventSubscriptionReconciler struct {
	client.Client
	Log logr.Logger
}

func (r *BMCEventSubscriptionReconciler) Reconcile(ctx context.Context, request ctrl.Request) (result ctrl.Result, err error) {
	reconcileCounters.With(hostMetricLabels(request)).Inc()
	defer func() {
		if err != nil {
			reconcileErrorCounter.Inc()
		}
	}()

	reqLogger := r.Log.WithValues("bmceventsubscription", request.NamespacedName)
	reqLogger.Info("start")

	// Fetch the BMCEventSubscription
	subscription := &metal3v1alpha1.BMCEventSubscription{}
	err = r.Get(ctx, request.NamespacedName, subscription)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// Request object not found, could have been deleted after
			// reconcile request.  Owned objects are automatically
			// garbage collected. For additional cleanup logic use
			// finalizers.  Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, errors.Wrap(err, "could not load host data")
	}

	return
}

// SetupWithManager registers the reconciler to be run by the manager
func (r *BMCEventSubscriptionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&metal3v1alpha1.BMCEventSubscription{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}
