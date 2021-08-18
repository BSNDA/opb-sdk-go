module github.com/bianjieai/opb-sdk-go

go 1.16

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

require (
	github.com/bianjieai/irita-sdk-go v1.1.1-0.20210818060217-87f48face9f8
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/grpc v1.40.0 // indirect
)