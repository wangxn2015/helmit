// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"context"
	"github.com/wangxn2015/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

var PersistentVolumeClaimKind = resource.Kind{
	Group:   "",
	Version: "v1",
	Kind:    "PersistentVolumeClaim",
	Scoped:  true,
}

var PersistentVolumeClaimResource = resource.Type{
	Kind: PersistentVolumeClaimKind,
	Name: "persistentvolumeclaims",
}

func NewPersistentVolumeClaim(persistentVolumeClaim *corev1.PersistentVolumeClaim, client resource.Client) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		Resource: resource.NewResource(persistentVolumeClaim.ObjectMeta, PersistentVolumeClaimKind, client),
		Object:   persistentVolumeClaim,
	}
}

type PersistentVolumeClaim struct {
	*resource.Resource
	Object *corev1.PersistentVolumeClaim
}

func (r *PersistentVolumeClaim) Delete(ctx context.Context) error {
	client, err := kubernetes.NewForConfig(r.Config())
	if err != nil {
		return err
	}
	return client.CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, PersistentVolumeClaimKind.Scoped).
		Resource(PersistentVolumeClaimResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Error()
}
