// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/wangxn2015/helmit/pkg/kubernetes/resource"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var CustomResourceDefinitionKind = resource.Kind{
	Group:   "apiextensions.k8s.io",
	Version: "v1",
	Kind:    "CustomResourceDefinition",
	Scoped:  false,
}

var CustomResourceDefinitionResource = resource.Type{
	Kind: CustomResourceDefinitionKind,
	Name: "customresourcedefinitions",
}

func NewCustomResourceDefinition(customResourceDefinition *apiextensionsv1.CustomResourceDefinition, client resource.Client) *CustomResourceDefinition {
	return &CustomResourceDefinition{
		Resource: resource.NewResource(customResourceDefinition.ObjectMeta, CustomResourceDefinitionKind, client),
		Object:   customResourceDefinition,
	}
}

type CustomResourceDefinition struct {
	*resource.Resource
	Object *apiextensionsv1.CustomResourceDefinition
}

func (r *CustomResourceDefinition) Delete(ctx context.Context) error {
	client, err := clientset.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.ApiextensionsV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, CustomResourceDefinitionKind.Scoped).
		Resource(CustomResourceDefinitionResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Error()
}
