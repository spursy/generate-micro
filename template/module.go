package template

var (
	ModuleSRV = `module {{.Dir}}

go 1.16

// github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
replace github.com/nsqio/go-nsq => git.5th.im/long-bridge-core-system/go-nsq v0.0.2-0.20200327020757-79a8c85c399e

require (
	git.5th.im/lb-public/gear v1.12.4
	github.com/demdxx/gocast v1.0.1 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.7.0
	github.com/prometheus/client_golang v1.8.0 // indirect
)
`
)



