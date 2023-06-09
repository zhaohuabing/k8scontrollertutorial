package main

import (
	"context"
	"fmt"
	api "github.com/zhaohuabing/k8scontrollertutorial/pkg/custom/apis/foo/v1alpha1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

type reconciler struct {
	client.Client
}

func (r *reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("foo", req.NamespacedName)
	log.V(1).Info("reconciling foo")

	var foo api.Foo
	if err := r.Get(ctx, req.NamespacedName, &foo); err != nil {
		log.Error(err, "unable to get foo")
		return ctrl.Result{}, err
	}

	fmt.Printf("Sync/Add/Update for foo %s\n", foo.GetName())
	return ctrl.Result{}, nil
}

func main() {
	ctrl.SetLogger(zap.New())

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:          true,
		LeaderElectionID:        "sample-controller",
		LeaderElectionNamespace: "kube-system",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// in a real controller, we'd create a new scheme for this
	err = api.AddToScheme(mgr.GetScheme())
	if err != nil {
		setupLog.Error(err, "unable to add scheme")
		os.Exit(1)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&api.Foo{}).
		Complete(&reconciler{
			Client: mgr.GetClient(),
		})
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	err = ctrl.NewWebhookManagedBy(mgr).
		For(&api.Foo{}).
		Complete()
	if err != nil {
		setupLog.Error(err, "unable to create webhook")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
