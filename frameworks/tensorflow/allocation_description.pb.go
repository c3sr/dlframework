// Code generated by protoc-gen-gogo.
// source: allocation_description.proto
// DO NOT EDIT!

/*
	Package tensorflow is a generated protocol buffer package.

	It is generated from these files:
		allocation_description.proto
		attr_value.proto
		cluster.proto
		config.proto
		control_flow.proto
		cost_graph.proto
		debug.proto
		device_attributes.proto
		device_properties.proto
		error_codes.proto
		function.proto
		graph.proto
		graph_transfer_info.proto
		kernel_def.proto
		log_memory.proto
		master.proto
		master_service.proto
		meta_graph.proto
		named_tensor.proto
		node_def.proto
		op_def.proto
		op_gen_overrides.proto
		queue_runner.proto
		reader_base.proto
		remote_fused_graph_execute_info.proto
		resource_handle.proto
		rewriter_config.proto
		saved_model.proto
		saver.proto
		step_stats.proto
		summary.proto
		tensor_bundle.proto
		tensor_description.proto
		tensorflow_server.proto
		tensor.proto
		tensor_shape.proto
		tensor_slice.proto
		types.proto
		variable.proto
		versions.proto
		worker.proto
		worker_service.proto

	It has these top-level messages:
		AllocationDescription
		AttrValue
		NameAttrList
		JobDef
		ClusterDef
		GPUOptions
		OptimizerOptions
		GraphOptions
		ThreadPoolOptionProto
		RPCOptions
		ConfigProto
		RunOptions
		RunMetadata
		ValuesDef
		CondContextDef
		WhileContextDef
		CostGraphDef
		DebugTensorWatch
		DebugOptions
		DeviceLocality
		DeviceAttributes
		DeviceProperties
		Error
		FunctionDefLibrary
		FunctionDef
		GradientDef
		GraphDef
		GraphTransferInfo
		KernelDef
		MemoryLogStep
		MemoryLogTensorAllocation
		MemoryLogTensorDeallocation
		MemoryLogTensorOutput
		MemoryLogRawAllocation
		MemoryLogRawDeallocation
		CreateSessionRequest
		CreateSessionResponse
		ExtendSessionRequest
		ExtendSessionResponse
		RunStepRequest
		RunStepResponse
		PartialRunSetupRequest
		PartialRunSetupResponse
		CloseSessionRequest
		CloseSessionResponse
		ResetRequest
		ResetResponse
		ListDevicesRequest
		ListDevicesResponse
		MetaGraphDef
		CollectionDef
		TensorInfo
		SignatureDef
		AssetFileDef
		NamedTensorProto
		NodeDef
		OpDef
		OpDeprecation
		OpList
		OpGenOverride
		OpGenOverrides
		QueueRunnerDef
		ReaderBaseState
		RemoteFusedGraphExecuteInfo
		ResourceHandle
		AutoParallelOptions
		RewriterConfig
		SavedModel
		SaverDef
		AllocatorMemoryUsed
		NodeOutput
		MemoryStats
		NodeExecStats
		DeviceStepStats
		StepStats
		SummaryDescription
		HistogramProto
		Summary
		BundleHeaderProto
		BundleEntryProto
		TensorDescription
		ServerDef
		TensorProto
		TensorShapeProto
		TensorSliceProto
		VariableDef
		SaveSliceInfoDef
		VersionDef
		GetStatusRequest
		GetStatusResponse
		CreateWorkerSessionRequest
		CreateWorkerSessionResponse
		RegisterGraphRequest
		RegisterGraphResponse
		DeregisterGraphRequest
		DeregisterGraphResponse
		CleanupAllRequest
		CleanupAllResponse
		ExecutorOpts
		RunGraphRequest
		RunGraphResponse
		CleanupGraphRequest
		CleanupGraphResponse
		RecvTensorRequest
		RecvTensorResponse
		LoggingRequest
		LabeledStepStats
		LoggingResponse
		TraceOpts
		TracingRequest
		TracingResponse
*/
package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
}

func (m *AllocationDescription) Reset()         { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()    {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptorAllocationDescription, []int{0}
}

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}
func (m *AllocationDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocationDescription) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestedBytes != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.RequestedBytes))
	}
	if m.AllocatedBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.AllocatedBytes))
	}
	if len(m.AllocatorName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(len(m.AllocatorName)))
		i += copy(dAtA[i:], m.AllocatorName)
	}
	if m.AllocationId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.AllocationId))
	}
	if m.HasSingleReference {
		dAtA[i] = 0x28
		i++
		if m.HasSingleReference {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Ptr != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAllocationDescription(dAtA, i, uint64(m.Ptr))
	}
	return i, nil
}

func encodeFixed64AllocationDescription(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32AllocationDescription(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAllocationDescription(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AllocationDescription) Size() (n int) {
	var l int
	_ = l
	if m.RequestedBytes != 0 {
		n += 1 + sovAllocationDescription(uint64(m.RequestedBytes))
	}
	if m.AllocatedBytes != 0 {
		n += 1 + sovAllocationDescription(uint64(m.AllocatedBytes))
	}
	l = len(m.AllocatorName)
	if l > 0 {
		n += 1 + l + sovAllocationDescription(uint64(l))
	}
	if m.AllocationId != 0 {
		n += 1 + sovAllocationDescription(uint64(m.AllocationId))
	}
	if m.HasSingleReference {
		n += 2
	}
	if m.Ptr != 0 {
		n += 1 + sovAllocationDescription(uint64(m.Ptr))
	}
	return n
}

func sovAllocationDescription(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAllocationDescription(x uint64) (n int) {
	return sovAllocationDescription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllocationDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocationDescription
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllocationDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocationDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedBytes", wireType)
			}
			m.RequestedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatedBytes", wireType)
			}
			m.AllocatedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocatedBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAllocationDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationId", wireType)
			}
			m.AllocationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocationId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasSingleReference", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasSingleReference = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ptr", wireType)
			}
			m.Ptr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ptr |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAllocationDescription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAllocationDescription
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAllocationDescription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllocationDescription
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAllocationDescription
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAllocationDescription
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAllocationDescription
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAllocationDescription(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAllocationDescription = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllocationDescription   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("allocation_description.proto", fileDescriptorAllocationDescription) }

var fileDescriptorAllocationDescription = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x40, 0xfd, 0x76, 0x1c, 0x34, 0x38, 0x2a, 0x41, 0x21, 0xa0, 0x94, 0xa2, 0x88, 0x5d, 0x15,
	0xc1, 0x13, 0x58, 0xdc, 0x08, 0x22, 0x43, 0x3d, 0x40, 0xc9, 0xb4, 0xbf, 0x33, 0xc5, 0xb6, 0xa9,
	0x3f, 0x91, 0xc1, 0x5b, 0x78, 0x2c, 0x97, 0x1e, 0x41, 0xea, 0x25, 0x74, 0x27, 0xe9, 0x38, 0xa9,
	0x0b, 0x77, 0x9f, 0xf7, 0x1f, 0xc9, 0xe7, 0xb1, 0x13, 0x59, 0x55, 0x2a, 0x93, 0xa6, 0x54, 0x4d,
	0x9a, 0xa3, 0xce, 0xa8, 0x6c, 0xed, 0x1c, 0xb5, 0xa4, 0x8c, 0xe2, 0xcc, 0x60, 0xa3, 0x15, 0x15,
	0x95, 0x5a, 0x9e, 0x7e, 0x03, 0x3b, 0xba, 0x76, 0xf2, 0xcd, 0xe0, 0xf2, 0x0b, 0xb6, 0x4f, 0xf8,
	0xf4, 0x8c, 0xda, 0x60, 0x9e, 0xce, 0x5e, 0x0c, 0x6a, 0x01, 0x01, 0x84, 0x5e, 0xb2, 0xe7, 0x70,
	0x6c, 0xa9, 0x15, 0x7f, 0xbf, 0x73, 0xe2, 0xe6, 0x4a, 0x74, 0x78, 0x25, 0x9e, 0xb3, 0x35, 0x51,
	0x94, 0x36, 0xb2, 0x46, 0xe1, 0x05, 0x10, 0xee, 0x24, 0x13, 0x47, 0xef, 0x65, 0x8d, 0xfc, 0x8c,
	0x4d, 0xfe, 0x9c, 0x5f, 0xe6, 0x62, 0xd4, 0xbf, 0xb6, 0x3b, 0xc0, 0xdb, 0x9c, 0x5f, 0xb2, 0xc3,
	0x85, 0xd4, 0xa9, 0x2e, 0x9b, 0x79, 0x85, 0x29, 0x61, 0x81, 0x84, 0x4d, 0x86, 0x62, 0x2b, 0x80,
	0x70, 0x3b, 0xe1, 0x0b, 0xa9, 0x1f, 0xfa, 0x55, 0xb2, 0xde, 0xf0, 0x03, 0xe6, 0xb5, 0x86, 0xc4,
	0x38, 0x80, 0x70, 0x94, 0xd8, 0x31, 0xbe, 0x7b, 0xeb, 0x7c, 0x78, 0xef, 0x7c, 0xf8, 0xe8, 0x7c,
	0x78, 0xfd, 0xf4, 0x37, 0x98, 0x50, 0x34, 0x8f, 0x86, 0x3a, 0x51, 0x41, 0xb2, 0xc6, 0xa5, 0xa2,
	0xc7, 0xf8, 0xf8, 0xdf, 0x48, 0x53, 0xdb, 0x53, 0x4f, 0xe1, 0x0b, 0x60, 0x36, 0xee, 0xe3, 0x5e,
	0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x44, 0x96, 0xea, 0x7c, 0x01, 0x00, 0x00,
}
