/*
Copyright 2022.

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
	"sigs.k8s.io/controller-runtime/pkg/log"

	dellv1 "github.com/king-jam/kubecon2022/controller/api/v1"
)

// ServerReconciler reconciles a Server object
type ServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=dell.kubecon.dell.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dell.kubecon.dell.com,resources=servers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dell.kubecon.dell.com,resources=servers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Server object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var server dellv1.Server
	if err := r.Get(ctx, req.NamespacedName, &server); err != nil {
		log.Error(err, "unable to get it")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("object", "desired", server.Spec.PowerState, "current", server.Status.PowerState)

	if server.Spec.PowerState == dellv1.PoweredOff {
		// do some client call to power it off
		log.Info("would power it off")
		server.Status.PowerState = dellv1.PoweredOff
	} else {
		// do some client call to power it on
		log.Info("would power it on")
		server.Status.PowerState = dellv1.PoweredOn
	}

	if err := r.Status().Update(ctx, &server); err != nil {
		log.Error(err, "unable to update")

		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dellv1.Server{}).
		Complete(r)
}
