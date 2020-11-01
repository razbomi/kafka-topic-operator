package main

import (
	"os"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr    manager.Manager
	logger = ctrl.Log.WithName("pod-controller")
)

func init() {
	// +kubebuilder:scaffold:scheme
}

// func podController() controller.Controller {
// 	c, err := controller.New("pod-controller", mgr, controller.Options{
// 		Reconciler: reconcile.Func(func(o reconcile.Request) (reconcile.Result, error) {
// 			logger.Info("creating, updating, deleting objects here")
// 			return reconcile.Result{}, nil
// 		}),
// 	})
// 	if err != nil {
// 		logger.Error(err, "unable to create pod-controller")
// 		os.Exit(1)
// 	}
// 	return c
// }

func main() {
	logger.Info("starting manager")
	// c := podController()
	// err := c.Watch(&source.Kind{Type: &v1.Pod{}}, &handler.EnqueueRequestForObject{})

	err := builder.
		ControllerManagedBy(mgr).  // Create the ControllerManagedBy
		For(&appsv1.ReplicaSet{}). // ReplicaSet is the Application API
		Owns(&v1.Pod{}).           // ReplicaSet owns Pods created by it
		Complete(&ReplicaSetReconciler{})
	if err != nil {
		logger.Error(err, "unable to watch pods")
		os.Exit(1)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Error(err, "could not start manager")
		os.Exit(1)
	}
}

// ReplicaSetReconciler HelloWorld
type ReplicaSetReconciler struct {
	client.Client
}

// Reconcile HelloWorld
func (a *ReplicaSetReconciler) Reconcile(req reconcile.Request) (reconcile.Result, error) {
	logger.Info("reconcile")
	return reconcile.Result{}, nil
}

// InjectClient HelloWorld
func (a *ReplicaSetReconciler) InjectClient(c client.Client) error {
	a.Client = c
	return nil
}
