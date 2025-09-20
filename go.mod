module bigchain

go 1.18

replace git.schwanenlied.me/yawning/bsaes.git => github.com/Yawning/bsaes v0.0.0-20180720073208-c0276d75487e

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace github.com/bigchain/bigchaind => ./

require google.golang.org/grpc v1.44.0

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/beorn7/perks v1.0.0 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.1.1 // indirect
	github.com/juju/clock v1.0.2 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/juju/loggo/v2 v2.0.0 // indirect
	github.com/juju/utils/v3 v3.1.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4 // indirect
	github.com/prometheus/common v0.4.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190507164030-5867b95ac084 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.0.3 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require (
	github.com/NebulousLabs/fastrand v0.0.0-20181203155948-6fb6489aac4e // indirect
	github.com/NebulousLabs/go-upnp v0.0.0-20180202185039-29b680b06c82
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da
	github.com/aead/siphash v1.0.1
	github.com/arl/statsviz v0.2.2-0.20201115121518-5ea9f0cf1bd1
	github.com/bigchain/bigchaind v0.0.0-00010101000000-000000000000
	github.com/btcsuite/winsvc v1.0.0
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/dchest/blake2b v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/emirpasic/gods v1.12.1-0.20200630092735-7e2349589531
	github.com/fsnotify/fsnotify v1.4.10-0.20200417215612-7f4cf4dd2b52 // indirect
	github.com/go-errors/errors v1.0.1
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/golang/protobuf v1.5.4
	github.com/golang/snappy v0.0.2
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackpal/gateway v1.0.5
	github.com/jackpal/go-nat-pmp v0.0.0-20170405195558-28a68d0c24ad
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/jessevdk/go-flags v1.4.1-0.20200711081900-c17162fe8fd7
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.11-0.20200806011408-6821bec9fa5c
	github.com/juju/loggo v0.0.0-20210728185423-eebad3a902c4 // indirect
	github.com/kkdai/bstream v1.0.0
	github.com/lightninglabs/protobuf-hex-display v1.4.3-hex-display
	github.com/ltcsuite/ltcd v0.0.0-20190101042124-f37f8bf35796
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/miekg/dns v0.0.0-20171125082028-79bfde677fa8
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1
	github.com/nxadm/tail v1.4.6-0.20201001195649-edf6bc2dfc36 // indirect
	github.com/onsi/ginkgo v1.14.3-0.20201013214636-dfe369837f25
	github.com/onsi/gomega v1.10.3
	github.com/prometheus/client_golang v0.9.3
	github.com/sethgrid/pester v1.1.1-0.20200617174401-d2ad9ec9a8b6
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/urfave/cli v1.18.0
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.6-0.20200807205753-f6be82302843
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.3.0
	golang.org/x/net v0.7.0
	golang.org/x/sys v0.16.0
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20220204002441-d6cc3cc0770e // indirect
	google.golang.org/protobuf v1.33.0
	gopkg.in/errgo.v1 v1.0.1 // indirect
	gopkg.in/macaroon-bakery.v2 v2.0.1
	gopkg.in/macaroon.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)
