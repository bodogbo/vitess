// Code generated by protoc-gen-go.
// source: throttlerdata.proto
// DO NOT EDIT!

/*
Package throttlerdata is a generated protocol buffer package.

It is generated from these files:
	throttlerdata.proto

It has these top-level messages:
	MaxRatesRequest
	MaxRatesResponse
	SetMaxRateRequest
	SetMaxRateResponse
	MaxReplicationLagModuleConfig
*/
package throttlerdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MaxRatesRequest is the payload for the MaxRates RPC.
type MaxRatesRequest struct {
}

func (m *MaxRatesRequest) Reset()                    { *m = MaxRatesRequest{} }
func (m *MaxRatesRequest) String() string            { return proto.CompactTextString(m) }
func (*MaxRatesRequest) ProtoMessage()               {}
func (*MaxRatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// MaxRatesResponse is returned by the MaxRates RPC.
type MaxRatesResponse struct {
	// max_rates returns the max rate for each throttler. It's keyed by the
	// throttler name.
	Rates map[string]int64 `protobuf:"bytes,1,rep,name=rates" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *MaxRatesResponse) Reset()                    { *m = MaxRatesResponse{} }
func (m *MaxRatesResponse) String() string            { return proto.CompactTextString(m) }
func (*MaxRatesResponse) ProtoMessage()               {}
func (*MaxRatesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MaxRatesResponse) GetRates() map[string]int64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

// SetMaxRateRequest is the payload for the SetMaxRate RPC.
type SetMaxRateRequest struct {
	Rate int64 `protobuf:"varint,1,opt,name=rate" json:"rate,omitempty"`
}

func (m *SetMaxRateRequest) Reset()                    { *m = SetMaxRateRequest{} }
func (m *SetMaxRateRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMaxRateRequest) ProtoMessage()               {}
func (*SetMaxRateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// SetMaxRateResponse is returned by the SetMaxRate RPC.
type SetMaxRateResponse struct {
	// names is the list of throttler names which were updated.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *SetMaxRateResponse) Reset()                    { *m = SetMaxRateResponse{} }
func (m *SetMaxRateResponse) String() string            { return proto.CompactTextString(m) }
func (*SetMaxRateResponse) ProtoMessage()               {}
func (*SetMaxRateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// MaxReplicationLagModuleConfig holds the configuration parameters
// for the MaxReplicationLagModule which allows to adaptively
// adjust the throttling rate based on the observed replication lag
// across all replicas.
type MaxReplicationLagModuleConfig struct {
	// target_replication_lag_sec is the replication lag (in seconds) the
	// MaxReplicationLagModule tries to aim for.
	// If it is within the target, it tries to increase the throttler
	// rate, otherwise it will lower it based on an educated guess of the
	// slave throughput.
	TargetReplicationLagSec int64 `protobuf:"varint,1,opt,name=target_replication_lag_sec,json=targetReplicationLagSec" json:"target_replication_lag_sec,omitempty"`
	// max_replication_lag_sec is meant as a last resort.
	// By default, the module tries to find out the system maximum capacity while
	// trying to keep the replication lag around "target_replication_lag_sec".
	// Usually, we'll wait min_duration_between_changes_sec to see the effect of a
	// throttler rate change on the replication lag.
	// But if the lag goes above this field's value we will go into an "emergency"
	// state and throttle more aggressively (see "emergency_decrease" below).
	// This is the only way to ensure that the system will recover.
	MaxReplicationLagSec int64 `protobuf:"varint,2,opt,name=max_replication_lag_sec,json=maxReplicationLagSec" json:"max_replication_lag_sec,omitempty"`
	// initial_rate is the rate at which the module will start.
	InitialRate int64 `protobuf:"varint,3,opt,name=initial_rate,json=initialRate" json:"initial_rate,omitempty"`
	// max_increase defines by how much we will increase the rate
	// e.g. 0.05 increases the rate by 5% while 1.0 by 100%.
	// Note that any increase will let the system wait for at least
	// (1 / MaxIncrease) seconds. If we wait for shorter periods of time, we
	// won't notice if the rate increase also increases the replication lag.
	// (If the system was already at its maximum capacity (e.g. 1k QPS) and we
	// increase the rate by e.g. 5% to 1050 QPS, it will take 20 seconds until
	// 1000 extra queries are buffered and the lag increases by 1 second.)
	MaxIncrease float64 `protobuf:"fixed64,4,opt,name=max_increase,json=maxIncrease" json:"max_increase,omitempty"`
	// emergency_decrease defines by how much we will decrease the current rate
	// if the observed replication lag is above "max_replication_lag_sec".
	// E.g. 0.50 decreases the current rate by 50%.
	EmergencyDecrease float64 `protobuf:"fixed64,5,opt,name=emergency_decrease,json=emergencyDecrease" json:"emergency_decrease,omitempty"`
	// min_duration_between_changes_sec specifies how long we'll wait for the last
	// rate increase or decrease to have an effect on the system.
	MinDurationBetweenChangesSec int64 `protobuf:"varint,6,opt,name=min_duration_between_changes_sec,json=minDurationBetweenChangesSec" json:"min_duration_between_changes_sec,omitempty"`
	// max_duration_between_increases_sec specifies how long we'll wait at most
	// for the last rate increase to have an effect on the system.
	MaxDurationBetweenIncreasesSec int64 `protobuf:"varint,7,opt,name=max_duration_between_increases_sec,json=maxDurationBetweenIncreasesSec" json:"max_duration_between_increases_sec,omitempty"`
	// ignore_n_slowest_replicas will ignore replication lag updates from the
	// N slowest replicas. Under certain circumstances, replicas are still
	// considered e.g. a) if the lag is at most max_replication_lag_sec, b) there
	// are less than N+1 replicas or c) the lag increased on each replica such
	// that all replicas were ignored in a row.
	IgnoreNSlowestReplicas int32 `protobuf:"varint,8,opt,name=ignore_n_slowest_replicas,json=ignoreNSlowestReplicas" json:"ignore_n_slowest_replicas,omitempty"`
}

func (m *MaxReplicationLagModuleConfig) Reset()                    { *m = MaxReplicationLagModuleConfig{} }
func (m *MaxReplicationLagModuleConfig) String() string            { return proto.CompactTextString(m) }
func (*MaxReplicationLagModuleConfig) ProtoMessage()               {}
func (*MaxReplicationLagModuleConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*MaxRatesRequest)(nil), "throttlerdata.MaxRatesRequest")
	proto.RegisterType((*MaxRatesResponse)(nil), "throttlerdata.MaxRatesResponse")
	proto.RegisterType((*SetMaxRateRequest)(nil), "throttlerdata.SetMaxRateRequest")
	proto.RegisterType((*SetMaxRateResponse)(nil), "throttlerdata.SetMaxRateResponse")
	proto.RegisterType((*MaxReplicationLagModuleConfig)(nil), "throttlerdata.MaxReplicationLagModuleConfig")
}

func init() { proto.RegisterFile("throttlerdata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0x65, 0x19, 0xec, 0x14, 0xc4, 0x6a, 0x2a, 0x16, 0x26, 0x40, 0x21, 0x37, 0x44,
	0x93, 0xe8, 0x05, 0x08, 0x69, 0xc0, 0x0d, 0x62, 0x03, 0x09, 0xc4, 0xb8, 0x70, 0x1f, 0xc0, 0xf2,
	0xd2, 0x43, 0x66, 0x91, 0xd8, 0xc5, 0x76, 0x59, 0xfa, 0x12, 0xbc, 0x1e, 0xaf, 0x83, 0xfc, 0x67,
	0x2b, 0x2d, 0xdc, 0xf9, 0x7c, 0xe7, 0x77, 0xbe, 0xf3, 0x9d, 0x36, 0x70, 0xdf, 0x5d, 0x1a, 0xed,
	0x5c, 0x87, 0x66, 0x2e, 0x9c, 0x98, 0x2e, 0x8c, 0x76, 0x9a, 0xde, 0xdd, 0x10, 0xab, 0x31, 0xdc,
	0x3b, 0x17, 0x03, 0x13, 0x0e, 0x2d, 0xc3, 0x1f, 0x4b, 0xb4, 0xae, 0xfa, 0x45, 0xe0, 0x60, 0xad,
	0xd9, 0x85, 0x56, 0x16, 0xe9, 0x3b, 0xc8, 0x8d, 0x17, 0x0a, 0x52, 0x66, 0xf5, 0xe8, 0xc5, 0xf1,
	0x74, 0xd3, 0x7b, 0x9b, 0x9f, 0x86, 0xea, 0x83, 0x72, 0x66, 0xc5, 0xe2, 0xe0, 0xd1, 0x09, 0xc0,
	0x5a, 0xa4, 0x07, 0x90, 0x7d, 0xc7, 0x55, 0x41, 0x4a, 0x52, 0xef, 0x33, 0xff, 0xa4, 0x13, 0xc8,
	0x7f, 0x8a, 0x6e, 0x89, 0xc5, 0x4e, 0x49, 0xea, 0x8c, 0xc5, 0xe2, 0xcd, 0xce, 0x09, 0xa9, 0x9e,
	0xc1, 0x78, 0x86, 0x2e, 0xad, 0x48, 0x29, 0x29, 0x85, 0x5d, 0xef, 0x1b, 0x1c, 0x32, 0x16, 0xde,
	0xd5, 0x31, 0xd0, 0xbf, 0xc1, 0x14, 0x7d, 0x02, 0xb9, 0x12, 0x7d, 0x8a, 0xbe, 0xcf, 0x62, 0x51,
	0xfd, 0xce, 0xe0, 0xb1, 0x27, 0x71, 0xd1, 0xc9, 0x46, 0x38, 0xa9, 0xd5, 0x17, 0xd1, 0x9e, 0xeb,
	0xf9, 0xb2, 0xc3, 0x53, 0xad, 0xbe, 0xc9, 0x96, 0xbe, 0x85, 0x23, 0x27, 0x4c, 0x8b, 0x8e, 0x9b,
	0x35, 0xc4, 0x3b, 0xd1, 0x72, 0x8b, 0x4d, 0xda, 0x7b, 0x18, 0x89, 0x4d, 0x97, 0x19, 0x36, 0xf4,
	0x15, 0x1c, 0xf6, 0x62, 0xf8, 0xef, 0x64, 0xbc, 0x6f, 0xd2, 0x6f, 0x2f, 0xf7, 0x63, 0x4f, 0xe1,
	0x8e, 0x54, 0xd2, 0x49, 0xd1, 0xf1, 0x70, 0x5d, 0x16, 0xd8, 0x51, 0xd2, 0xfc, 0x59, 0x1e, 0xf1,
	0xce, 0x52, 0x35, 0x06, 0x85, 0xc5, 0x62, 0xb7, 0x24, 0x35, 0x61, 0xa3, 0x5e, 0x0c, 0x9f, 0x92,
	0x44, 0x9f, 0x03, 0xc5, 0x1e, 0x4d, 0x8b, 0xaa, 0x59, 0xf1, 0x39, 0x26, 0x30, 0x0f, 0xe0, 0xf8,
	0xa6, 0x73, 0x96, 0x1a, 0xf4, 0x23, 0x94, 0xbd, 0x54, 0x7c, 0xbe, 0x34, 0x31, 0xe8, 0x05, 0xba,
	0x2b, 0x44, 0xc5, 0x9b, 0x4b, 0xa1, 0x5a, 0xb4, 0x21, 0xf4, 0x5e, 0x08, 0xf2, 0xa8, 0x97, 0xea,
	0x2c, 0x61, 0xef, 0x23, 0x75, 0x1a, 0x21, 0x1f, 0xfe, 0x33, 0x54, 0x3e, 0xd9, 0x3f, 0x3e, 0xd7,
	0x51, 0xa3, 0xd3, 0xad, 0xe0, 0xf4, 0xa4, 0x17, 0xc3, 0x96, 0xd3, 0x75, 0xfc, 0xe0, 0xf5, 0x1a,
	0x1e, 0xca, 0x56, 0x69, 0x83, 0x5c, 0x71, 0xdb, 0xe9, 0x2b, 0xb4, 0x37, 0x7f, 0x83, 0x2d, 0x6e,
	0x97, 0xa4, 0xce, 0xd9, 0x83, 0x08, 0x7c, 0x9d, 0xc5, 0x76, 0xfa, 0x31, 0xed, 0xc5, 0x5e, 0xf8,
	0xd0, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3d, 0x0c, 0xc2, 0xff, 0x02, 0x00, 0x00,
}