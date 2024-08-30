package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	edgelevelcomv1alpha1 "github.com/edgelevel/lastpass-operator/api/v1alpha1"
)

// LastPassCustomReconciler reconciles a LastPassCustom object
type LastPassCustomReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=edgelevel.com,resources=lastpasscustoms,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=edgelevel.com,resources=lastpasscustoms/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=edgelevel.com,resources=lastpasscustoms/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LastPassCustom object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *LastPassCustomReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LastPassCustomReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&edgelevelcomv1alpha1.LastPassCustom{}).
		Complete(r)
}
