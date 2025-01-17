// Code generated by helmit-generate. DO NOT EDIT.

package {{ .Reader.Package.Name }}

import (
    "github.com/onosproject/helmit/pkg/kubernetes/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	{{ .Resource.Client.Package.Alias }} {{ .Resource.Client.Package.Path | quote }}
	{{ .Resource.Kind.Package.Alias }} {{ .Resource.Kind.Package.Path | quote }}
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
	"context"
)

type {{ .Reader.Types.Interface }} interface {
	Get(ctx context.Context, name string) (*{{ .Resource.Types.Struct }}, error)
	List(ctx context.Context) ([]*{{ .Resource.Types.Struct }}, error)
}

func New{{ .Reader.Types.Interface }}(client resource.Client, filter resource.Filter) {{ .Reader.Types.Interface }} {
	return &{{ .Reader.Types.Struct }}{
		Client: client,
		filter: filter,
	}
}

type {{ .Reader.Types.Struct }} struct {
	resource.Client
	filter resource.Filter
}

{{- $singular := (.Resource.Names.Singular | toLowerCamel) }}
{{- $kind := (printf "%s.%s" .Resource.Kind.Package.Alias .Resource.Kind.Kind) }}
{{- $listKind := (printf "%s.%s" .Resource.Kind.Package.Alias .Resource.Kind.ListKind) }}

func (c *{{ .Reader.Types.Struct }}) Get(ctx context.Context, name string) (*{{ .Resource.Types.Struct }}, error) {
    {{ $singular }} := &{{ $kind }}{}
    client, err := {{ .Resource.Client.Package.Alias }}.NewForConfig(c.Config())
    if err != nil {
        return nil, err
    }
	err = client.{{ .Group.Names.Proper }}().
        RESTClient().
	    Get().
	    NamespaceIfScoped(c.Namespace(), {{ .Resource.Types.Kind }}.Scoped).
		Resource({{ .Resource.Types.Resource }}.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into({{ $singular }})
	if err != nil {
		return nil, err
	} else {
        ok, err := c.filter(metav1.GroupVersionKind{
            Group:   {{ .Resource.Types.Kind }}.Group,
            Version: {{ .Resource.Types.Kind }}.Version,
            Kind:    {{ .Resource.Types.Kind }}.Kind,
        }, {{ $singular }}.ObjectMeta)
        if err != nil {
            return nil, err
        } else if !ok {
            return nil, errors.NewNotFound(schema.GroupResource{
                Group:    {{ .Resource.Types.Kind }}.Group,
                Resource: {{ .Resource.Types.Resource }}.Name,
            }, name)
        }
    }
	return New{{ .Resource.Types.Struct }}({{ $singular }}, c.Client), nil
}

func (c *{{ .Reader.Types.Struct }}) List(ctx context.Context) ([]*{{ .Resource.Types.Struct }}, error) {
    list := &{{ $listKind }}{}
    client, err := {{ .Resource.Client.Package.Alias }}.NewForConfig(c.Config())
    if err != nil {
        return nil, err
    }
	err = client.{{ .Group.Names.Proper }}().
        RESTClient().
	    Get().
	    NamespaceIfScoped(c.Namespace(), {{ .Resource.Types.Kind }}.Scoped).
		Resource({{ .Resource.Types.Resource }}.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do(ctx).
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*{{ .Resource.Types.Struct }}, 0, len(list.Items))
	for _, {{ $singular }} := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   {{ .Resource.Types.Kind }}.Group,
			Version: {{ .Resource.Types.Kind }}.Version,
			Kind:    {{ .Resource.Types.Kind }}.Kind,
		}, {{ $singular }}.ObjectMeta)
        if err != nil {
            return nil, err
        } else if ok {
            copy := {{ $singular }}
    	    results = append(results, New{{ .Resource.Types.Struct }}(&copy, c.Client))
        }
	}
	return results, nil
}
