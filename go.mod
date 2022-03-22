module github.com/aliyun/aliyun-cli

go 1.17

replace github.com/alibabacloud-go/tea => ./local/tea

replace github.com/aliyun/credentials-go => ./local/credentials-go

replace github.com/syndtr/goleveldb => ./local/goleveldb

replace github.com/onsi/ginkgo => ./local/ginkgo

replace github.com/onsi/ginkgo/v2 => ./local/ginkgo_v2

replace github.com/smartystreets/goconvey => ./local/goconvey

replace golang.org/x/mod => ./local/mod

require (
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1529
	github.com/aliyun/aliyun-oss-go-sdk v2.2.1+incompatible
	github.com/aliyun/credentials-go v1.2.1
	github.com/alyu/configparser v0.0.0-20191103060215-744e9a66e7bc
	github.com/droundy/goopt v0.0.0-20170604162106-0b8effe182da
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/ini.v1 v1.66.2
)

require (
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
