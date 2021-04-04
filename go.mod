module github.com/c3sr/dlframework

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/jaegertracing/jaeger => github.com/uber/jaeger v1.22.0
	github.com/uber/jaeger => github.com/jaegertracing/jaeger v1.22.0
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)

go 1.15

require (
	github.com/DataDog/go-python3 v0.0.0-20191126174558-6ed25e33b3c4
	github.com/GeertJohan/go-sourcepath v0.0.0-20150925135350-83e8b8723a9b
	github.com/Masterminds/semver v1.5.0
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/VividCortex/robustly v0.0.0-20210119222408-48da771af5f6
	github.com/aarondl/tpl v0.0.0-20180717141031-b5afe9b3122c
	github.com/apache/arrow/go/arrow v0.0.0-20201229220542-30ce2eb5d4dc // indirect
	github.com/c3sr/archive v1.0.0
	github.com/c3sr/cmd v1.0.0
	github.com/c3sr/config v1.0.1
	github.com/c3sr/database v1.0.0
	github.com/c3sr/downloadmanager v1.0.0
	github.com/c3sr/grpc v1.0.0
	github.com/c3sr/image v1.0.0
	github.com/c3sr/libkv v1.0.0
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/machine v1.0.0
	github.com/c3sr/monitoring v1.0.0
	github.com/c3sr/nvidia-smi v1.0.0
	github.com/c3sr/parallel v1.0.1
	github.com/c3sr/pipeline v1.0.0
	github.com/c3sr/registry v1.0.0
	github.com/c3sr/serializer v1.0.0
	github.com/c3sr/tracer v1.0.0
	github.com/c3sr/utils v1.0.0
	github.com/c3sr/uuid v1.0.1
	github.com/c3sr/vipertags v1.0.0
	github.com/c3sr/web v1.0.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cheggaaa/pb v1.0.29
	github.com/cockroachdb/cmux v0.0.0-20170110192607-30d10be49292
	github.com/davecgh/go-spew v1.1.1
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052
	github.com/fatih/color v1.10.0
	github.com/glendc/go-external-ip v0.0.0-20200601212049-c872357d968e
	github.com/go-openapi/errors v0.20.0
	github.com/go-openapi/loads v0.20.2
	github.com/go-openapi/runtime v0.19.26
	github.com/go-openapi/spec v0.20.3
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.20.2
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.1
	github.com/golang/snappy v0.0.3
	github.com/gorilla/schema v1.2.0
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/h2non/filetype v1.1.1
	github.com/jessevdk/go-flags v1.4.0
	github.com/jinzhu/copier v0.2.8
	github.com/justinas/nosurf v1.1.1
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/klauspost/shutdown2 v1.1.0
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/echo/v4 v4.2.1
	github.com/levigross/grequests v0.0.0-20190908174114-253788527a1a
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/oliamb/cutter v0.2.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/unknwon/com v1.0.1
	github.com/volatiletech/authboss v2.4.1+incompatible // indirect
	github.com/volatiletech/authboss-clientstate v0.0.0-20200826024349-8d4e74078241
	github.com/volatiletech/authboss/v3 v3.0.3
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20201222180813-1025295fd063 // indirect
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4
	golang.org/x/oauth2 v0.0.0-20210313182246-cd4f82c27b84 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210318145829-90b20ab00860
	google.golang.org/grpc v1.36.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.4.0
	gorgonia.org/tensor v0.9.14
)
