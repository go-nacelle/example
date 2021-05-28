module example

go 1.16

require (
	github.com/apex/log v1.9.0 // indirect
	github.com/derision-test/glock v0.0.0-20210316032053-f5b74334bb29 // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/go-nacelle/config v1.2.2-0.20210104042626-7a6afdea884f
	github.com/go-nacelle/grpcbase v1.0.1
	github.com/go-nacelle/httpbase v1.0.1
	github.com/go-nacelle/nacelle v1.2.0
	github.com/go-nacelle/process v1.1.1-0.20210415190733-22a6a9b328b3
	github.com/go-nacelle/workerbase v1.0.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/kr/pretty v0.2.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.21.1
)

replace (
	github.com/go-nacelle/config => ../config
	github.com/go-nacelle/log => ../log
	github.com/go-nacelle/nacelle => ../nacelle
	github.com/go-nacelle/process => ../process
	github.com/go-nacelle/service => ../service
	github.com/go-nacelle/httpbase => ../httpbase
	github.com/go-nacelle/workerbase => ../workerbase
	github.com/go-nacelle/grpcbase => ../grpcbase
)
