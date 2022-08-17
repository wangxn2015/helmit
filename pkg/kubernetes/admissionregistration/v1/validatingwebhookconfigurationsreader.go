// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/wangxn2015/helmit/pkg/kubernetes/resource"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type ValidatingWebhookConfigurationsReader interface {
	Get(ctx context.Context, name string) (*ValidatingWebhookConfiguration, error)
	List(ctx context.Context) ([]*ValidatingWebhookConfiguration, error)
}

func NewValidatingWebhookConfigurationsReader(client resource.Client, filter resource.Filter) ValidatingWebhookConfigurationsReader {
	return &validatingWebhookConfigurationsReader{
		Client: client,
		filter: filter,
	}
}

type validatingWebhookConfigurationsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *validatingWebhookConfigurationsReader) Get(ctx context.Context, name string) (*ValidatingWebhookConfiguration, error) {
	validatingWebhookConfiguration := &admissionregistrationv1.ValidatingWebhookConfiguration{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.AdmissionregistrationV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), ValidatingWebhookConfigurationKind.Scoped).
		Resource(ValidatingWebhookConfigurationResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(validatingWebhookConfiguration)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   ValidatingWebhookConfigurationKind.Group,
			Version: ValidatingWebhookConfigurationKind.Version,
			Kind:    ValidatingWebhookConfigurationKind.Kind,
		}, validatingWebhookConfiguration.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    ValidatingWebhookConfigurationKind.Group,
				Resource: ValidatingWebhookConfigurationResource.Name,
			}, name)
		}
	}
	return NewValidatingWebhookConfiguration(validatingWebhookConfiguration, c.Client), nil
}

func (c *validatingWebhookConfigurationsReader) List(ctx context.Context) ([]*ValidatingWebhookConfiguration, error) {
	list := &admissionregistrationv1.ValidatingWebhookConfigurationList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.AdmissionregistrationV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), ValidatingWebhookConfigurationKind.Scoped).
		Resource(ValidatingWebhookConfigurationResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*ValidatingWebhookConfiguration, 0, len(list.Items))
	for _, validatingWebhookConfiguration := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   ValidatingWebhookConfigurationKind.Group,
			Version: ValidatingWebhookConfigurationKind.Version,
			Kind:    ValidatingWebhookConfigurationKind.Kind,
		}, validatingWebhookConfiguration.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := validatingWebhookConfiguration
			results = append(results, NewValidatingWebhookConfiguration(&copy, c.Client))
		}
	}
	return results, nil
}
