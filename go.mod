module github.com/wanxn2015/helmit

go 1.16

require (
	github.com/dustinkirkland/golang-petname v0.0.0-20191129215211-8e5a1ed0cff0
	github.com/fatih/color v1.9.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/iancoleman/strcase v0.1.2
	github.com/joncalhoun/pipe v0.0.0-20170510025636-72505674a733
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/wangxn2015/helmit v0.0.0-00010101000000-000000000000
	github.com/wangxn2015/onos-lib-go v1.8.15
	google.golang.org/grpc v1.41.0
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.7.2
	k8s.io/api v0.22.4
	k8s.io/apiextensions-apiserver v0.22.4
	k8s.io/apimachinery v0.22.4
	k8s.io/client-go v0.22.4
)

//replace github.com/onosproject/onos-lib-go v0.8.1 => github.com/wangxn2015/onos-lib-go v1.8.15

replace github.com/wangxn2015/onos-lib-go => /home/baicells/go_project/modified-onos-module/onos-lib-go

replace github.com/wangxn2015/helmit => /home/baicells/go_project/modified-onos-module/helmit
