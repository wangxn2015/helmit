// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/wangxn2015/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type PersistentVolumeClaimsReader interface {
	Get(ctx context.Context, name string) (*PersistentVolumeClaim, error)
	List(ctx context.Context) ([]*PersistentVolumeClaim, error)
}

func NewPersistentVolumeClaimsReader(client resource.Client, filter resource.Filter) PersistentVolumeClaimsReader {
	return &persistentVolumeClaimsReader{
		Client: client,
		filter: filter,
	}
}

type persistentVolumeClaimsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *persistentVolumeClaimsReader) Get(ctx context.Context, name string) (*PersistentVolumeClaim, error) {
	persistentVolumeClaim := &corev1.PersistentVolumeClaim{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), PersistentVolumeClaimKind.Scoped).
		Resource(PersistentVolumeClaimResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(persistentVolumeClaim)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   PersistentVolumeClaimKind.Group,
			Version: PersistentVolumeClaimKind.Version,
			Kind:    PersistentVolumeClaimKind.Kind,
		}, persistentVolumeClaim.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    PersistentVolumeClaimKind.Group,
				Resource: PersistentVolumeClaimResource.Name,
			}, name)
		}
	}
	return NewPersistentVolumeClaim(persistentVolumeClaim, c.Client), nil
}

func (c *persistentVolumeClaimsReader) List(ctx context.Context) ([]*PersistentVolumeClaim, error) {
	list := &corev1.PersistentVolumeClaimList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), PersistentVolumeClaimKind.Scoped).
		Resource(PersistentVolumeClaimResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*PersistentVolumeClaim, 0, len(list.Items))
	for _, persistentVolumeClaim := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   PersistentVolumeClaimKind.Group,
			Version: PersistentVolumeClaimKind.Version,
			Kind:    PersistentVolumeClaimKind.Kind,
		}, persistentVolumeClaim.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := persistentVolumeClaim
			results = append(results, NewPersistentVolumeClaim(&copy, c.Client))
		}
	}
	return results, nil
}
