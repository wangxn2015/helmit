// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var SecretKind = resource.Kind{
	Group:   "core",
	Version: "v1",
	Kind:    "Secret",
	Scoped:  true,
}

var SecretResource = resource.Type{
	Kind: SecretKind,
	Name: "secrets",
}

func NewSecret(secret *corev1.Secret, client resource.Client) *Secret {
	return &Secret{
		Resource: resource.NewResource(secret.ObjectMeta, SecretKind, client),
		Object:   secret,
	}
}

type Secret struct {
	*resource.Resource
	Object *corev1.Secret
}

func (r *Secret) Delete() error {
	return r.Clientset().
		CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, SecretKind.Scoped).
		Resource(SecretResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
