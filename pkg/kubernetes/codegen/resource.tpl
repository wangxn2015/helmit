// Code generated by helmit-generate. DO NOT EDIT.

{{- $resource := .Resource }}
{{- $name := (.Resource.Names.Singular | toLowerCamel) }}
{{- $kind := (printf "%s.%s" .Resource.Kind.Package.Alias .Resource.Kind.Kind) }}

package {{ $resource.Package.Name }}

import (
    "github.com/onosproject/helmit/pkg/kubernetes/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	{{ .Resource.Kind.Package.Alias }} {{ .Resource.Kind.Package.Path | quote }}
    {{- range $ref := $resource.References }}
    {{- if not (eq $ref.Reference.Package.Path $resource.Package.Path) }}
    {{ $ref.Reference.Package.Alias }} {{ $ref.Reference.Package.Path | quote }}
    {{- end }}
    {{- end }}
    "time"
)

var {{ $resource.Types.Kind }} = resource.Kind{
	Group:   {{ $resource.Kind.Group | quote }},
	Version: {{ $resource.Kind.Version | quote }},
	Kind:    {{ $resource.Kind.Kind | quote }},
	{{- if $resource.Kind.Scoped }}
	Scoped: true,
	{{- else }}
	Scoped: false,
	{{- end }}
}

var {{ $resource.Types.Resource }} = resource.Type{
	Kind: {{ $resource.Types.Kind }},
	Name: {{ $resource.Names.Plural | lower | quote }},
}

func New{{ $resource.Types.Struct }}({{ $name }} *{{ $kind }}, client resource.Client) *{{ $resource.Types.Struct }} {
	return &{{ $resource.Types.Struct }}{
		Resource: resource.NewResource({{ $name }}.ObjectMeta, {{ .Resource.Types.Kind }}, client),
		Object: {{ $name }},
        {{- range $ref := $resource.References }}
        {{- if eq $ref.Resource.Package.Path $resource.Package.Path }}
        {{ $ref.Reference.Types.Interface }}: New{{ $ref.Reference.Types.Interface }}(client, resource.NewUIDFilter({{ $name }}.UID)),
        {{- else }}
        {{ $ref.Reference.Types.Interface }}: {{ $ref.Reference.Package.Alias }}.New{{ $ref.Reference.Types.Interface }}(client, resource.NewUIDFilter({{ $name }}.UID)),
        {{- end }}
        {{- end }}
	}
}

type {{ $resource.Types.Struct }} struct {
	*resource.Resource
	Object *{{ $kind }}
    {{- range $ref := .Resource.References }}
    {{- if eq $ref.Resource.Package.Path $resource.Package.Path }}
    {{ $ref.Reference.Types.Interface }}
    {{- else }}
    {{ $ref.Reference.Package.Alias }}.{{ $ref.Reference.Types.Interface }}
    {{- end }}
    {{- end }}
}

func (r *{{ $resource.Types.Struct }}) Delete() error {
	return r.Clientset().
        {{ .Group.Names.Proper }}().
        RESTClient().
	    Delete().
	    NamespaceIfScoped(r.Namespace, {{ .Resource.Types.Kind }}.Scoped).
		Resource({{ .Resource.Types.Resource }}.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
