// Copyright Red Hat

package v1alpha1

import (
	"fmt"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	cl "sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var (
	authrealmlog = logf.Log.WithName("authrealm-resource")
	Client       cl.Client
)

func (r *AuthRealm) SetupWebhookWithManager(mgr ctrl.Manager) error {
	Client = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:verbs=create;update,path=/validate-identitatem-io-identitatem-io-v1alpha1-authrealm,mutating=false,failurePolicy=fail,sideEffects=None,groups=identitatem.io.identitatem.io,resources=authrealms,versions=v1alpha1,name=vauthrealm.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &AuthRealm{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AuthRealm) ValidateCreate() error {
	authrealmlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	//var allErrs field.ErrorList

	// authRealmList := &AuthRealmList{}

	// if err := Client.List(context.TODO(), authRealmList); err != nil {
	// 	return fmt.Errorf("unable to list AuthRealms: %s", err)
	// }

	// for _, authRealm := range authRealmList.Items {
	// 	if authRealm.Name == r.Name {
	// 		return fmt.Errorf("AuthRealm CR already exists with that name")
	// 	}
	// }

	// MUST follow rules for internet subdomain name
	//domainRegex, _ := regexp.Compile(`^(?!-)[A-Za-z0-9-]{1, 63}(?<!-)$`)
	//domainRegex, _ := regexp.Compile(`^[A-Za-z0-9-](?:[A-Za-z0-9-]{0, 61}[a-z0-9])$`)
	//domainRegex, _ := regexp.Compile(`^([a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]?[^-_])$`)
	domainRegex, _ := regexp.Compile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`) // DNS-1123 subdomain

	if r.Spec.RouteSubDomain == "" {
		return fmt.Errorf("RouteSubDomain MUST be specified")
	} else if !domainRegex.MatchString(r.Spec.RouteSubDomain) {
		return fmt.Errorf(
			"RouteSubDomain \"%s\" is invalid: a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example', regex used for validation is \"%s\"",
			r.Spec.RouteSubDomain,
			domainRegex.String())
	}

	return nil
	// if len(allErrs) == 0 {
	//     return nil
	// }

	// return apierrors.NewInvalid(
	//     schema.GroupKind{Group: r.TypeMeta.GroupVersionKind().Group(), Kind: r.TypeMeta.Kind()},
	//     r.Name,
	// 	allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AuthRealm) ValidateUpdate(old runtime.Object) error {
	authrealmlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	oldAuthRealm := old.(*AuthRealm)
	if r.Spec.RouteSubDomain != oldAuthRealm.Spec.RouteSubDomain {
		return fmt.Errorf("RouteSubDomain is immutable and cannot be changed")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AuthRealm) ValidateDelete() error {
	authrealmlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
