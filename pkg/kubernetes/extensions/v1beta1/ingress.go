// Code generated by helmit-generate. DO NOT EDIT.

package v1beta1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var IngressKind = resource.Kind{
	Group:   "extensions",
	Version: "v1beta1",
	Kind:    "Ingress",
	Scoped:  true,
}

var IngressResource = resource.Type{
	Kind: IngressKind,
	Name: "ingresses",
}

func NewIngress(ingress *extensionsv1beta1.Ingress, client resource.Client) *Ingress {
	return &Ingress{
		Resource: resource.NewResource(ingress.ObjectMeta, IngressKind, client),
		Object:   ingress,
	}
}

type Ingress struct {
	*resource.Resource
	Object *extensionsv1beta1.Ingress
}

func (r *Ingress) Delete() error {
	return r.Clientset().
		ExtensionsV1beta1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, IngressKind.Scoped).
		Resource(IngressResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
