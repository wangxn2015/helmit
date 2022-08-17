// Code generated by helmit-generate. DO NOT EDIT.

package v1beta1

import (
	"context"
	"github.com/wangxn2015/helmit/pkg/kubernetes/resource"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

type CustomResourceDefinitionsReader interface {
	Get(ctx context.Context, name string) (*CustomResourceDefinition, error)
	List(ctx context.Context) ([]*CustomResourceDefinition, error)
}

func NewCustomResourceDefinitionsReader(client resource.Client, filter resource.Filter) CustomResourceDefinitionsReader {
	return &customResourceDefinitionsReader{
		Client: client,
		filter: filter,
	}
}

type customResourceDefinitionsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *customResourceDefinitionsReader) Get(ctx context.Context, name string) (*CustomResourceDefinition, error) {
	customResourceDefinition := &apiextensionsv1beta1.CustomResourceDefinition{}
	client, err := clientset.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.ApiextensionsV1beta1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), CustomResourceDefinitionKind.Scoped).
		Resource(CustomResourceDefinitionResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(customResourceDefinition)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   CustomResourceDefinitionKind.Group,
			Version: CustomResourceDefinitionKind.Version,
			Kind:    CustomResourceDefinitionKind.Kind,
		}, customResourceDefinition.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    CustomResourceDefinitionKind.Group,
				Resource: CustomResourceDefinitionResource.Name,
			}, name)
		}
	}
	return NewCustomResourceDefinition(customResourceDefinition, c.Client), nil
}

func (c *customResourceDefinitionsReader) List(ctx context.Context) ([]*CustomResourceDefinition, error) {
	list := &apiextensionsv1beta1.CustomResourceDefinitionList{}
	client, err := clientset.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.ApiextensionsV1beta1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), CustomResourceDefinitionKind.Scoped).
		Resource(CustomResourceDefinitionResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*CustomResourceDefinition, 0, len(list.Items))
	for _, customResourceDefinition := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   CustomResourceDefinitionKind.Group,
			Version: CustomResourceDefinitionKind.Version,
			Kind:    CustomResourceDefinitionKind.Kind,
		}, customResourceDefinition.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := customResourceDefinition
			results = append(results, NewCustomResourceDefinition(&copy, c.Client))
		}
	}
	return results, nil
}
