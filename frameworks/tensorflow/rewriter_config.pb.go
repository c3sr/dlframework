// Code generated by protoc-gen-gogo.
// source: rewriter_config.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RewriterConfig_MemOptType int32

const (
	// Fully disabled
	RewriterConfig_NO_MEM_OPT RewriterConfig_MemOptType = 0
	// Driven by manual annotations
	RewriterConfig_MANUAL RewriterConfig_MemOptType = 1
)

var RewriterConfig_MemOptType_name = map[int32]string{
	0: "NO_MEM_OPT",
	1: "MANUAL",
}
var RewriterConfig_MemOptType_value = map[string]int32{
	"NO_MEM_OPT": 0,
	"MANUAL":     1,
}

func (x RewriterConfig_MemOptType) String() string {
	return proto.EnumName(RewriterConfig_MemOptType_name, int32(x))
}
func (RewriterConfig_MemOptType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorRewriterConfig, []int{1, 0}
}

type AutoParallelOptions struct {
	Enable      bool  `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	NumReplicas int32 `protobuf:"varint,2,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
}

func (m *AutoParallelOptions) Reset()         { *m = AutoParallelOptions{} }
func (m *AutoParallelOptions) String() string { return proto.CompactTextString(m) }
func (*AutoParallelOptions) ProtoMessage()    {}
func (*AutoParallelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorRewriterConfig, []int{0}
}

func (m *AutoParallelOptions) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *AutoParallelOptions) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

type RewriterConfig struct {
	OptimizeTensorLayout bool                      `protobuf:"varint,1,opt,name=optimize_tensor_layout,json=optimizeTensorLayout,proto3" json:"optimize_tensor_layout,omitempty"`
	DisableModelPruning  bool                      `protobuf:"varint,2,opt,name=disable_model_pruning,json=disableModelPruning,proto3" json:"disable_model_pruning,omitempty"`
	ConstantFolding      bool                      `protobuf:"varint,3,opt,name=constant_folding,json=constantFolding,proto3" json:"constant_folding,omitempty"`
	MemoryOptimization   RewriterConfig_MemOptType `protobuf:"varint,4,opt,name=memory_optimization,json=memoryOptimization,proto3,enum=tensorflow.RewriterConfig_MemOptType" json:"memory_optimization,omitempty"`
	AutoParallel         *AutoParallelOptions      `protobuf:"bytes,5,opt,name=auto_parallel,json=autoParallel" json:"auto_parallel,omitempty"`
	// If non-empty, will use this as an alternative way to specify a list of
	// optimizations to turn on and the order of the optimizations.
	Optimizers []string `protobuf:"bytes,100,rep,name=optimizers" json:"optimizers,omitempty"`
}

func (m *RewriterConfig) Reset()                    { *m = RewriterConfig{} }
func (m *RewriterConfig) String() string            { return proto.CompactTextString(m) }
func (*RewriterConfig) ProtoMessage()               {}
func (*RewriterConfig) Descriptor() ([]byte, []int) { return fileDescriptorRewriterConfig, []int{1} }

func (m *RewriterConfig) GetOptimizeTensorLayout() bool {
	if m != nil {
		return m.OptimizeTensorLayout
	}
	return false
}

func (m *RewriterConfig) GetDisableModelPruning() bool {
	if m != nil {
		return m.DisableModelPruning
	}
	return false
}

func (m *RewriterConfig) GetConstantFolding() bool {
	if m != nil {
		return m.ConstantFolding
	}
	return false
}

func (m *RewriterConfig) GetMemoryOptimization() RewriterConfig_MemOptType {
	if m != nil {
		return m.MemoryOptimization
	}
	return RewriterConfig_NO_MEM_OPT
}

func (m *RewriterConfig) GetAutoParallel() *AutoParallelOptions {
	if m != nil {
		return m.AutoParallel
	}
	return nil
}

func (m *RewriterConfig) GetOptimizers() []string {
	if m != nil {
		return m.Optimizers
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoParallelOptions)(nil), "tensorflow.AutoParallelOptions")
	proto.RegisterType((*RewriterConfig)(nil), "tensorflow.RewriterConfig")
	proto.RegisterEnum("tensorflow.RewriterConfig_MemOptType", RewriterConfig_MemOptType_name, RewriterConfig_MemOptType_value)
}
func (m *AutoParallelOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AutoParallelOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Enable {
		dAtA[i] = 0x8
		i++
		if m.Enable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NumReplicas != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRewriterConfig(dAtA, i, uint64(m.NumReplicas))
	}
	return i, nil
}

func (m *RewriterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewriterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OptimizeTensorLayout {
		dAtA[i] = 0x8
		i++
		if m.OptimizeTensorLayout {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DisableModelPruning {
		dAtA[i] = 0x10
		i++
		if m.DisableModelPruning {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ConstantFolding {
		dAtA[i] = 0x18
		i++
		if m.ConstantFolding {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MemoryOptimization != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRewriterConfig(dAtA, i, uint64(m.MemoryOptimization))
	}
	if m.AutoParallel != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRewriterConfig(dAtA, i, uint64(m.AutoParallel.Size()))
		n1, err := m.AutoParallel.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Optimizers) > 0 {
		for _, s := range m.Optimizers {
			dAtA[i] = 0xa2
			i++
			dAtA[i] = 0x6
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64RewriterConfig(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32RewriterConfig(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRewriterConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AutoParallelOptions) Size() (n int) {
	var l int
	_ = l
	if m.Enable {
		n += 2
	}
	if m.NumReplicas != 0 {
		n += 1 + sovRewriterConfig(uint64(m.NumReplicas))
	}
	return n
}

func (m *RewriterConfig) Size() (n int) {
	var l int
	_ = l
	if m.OptimizeTensorLayout {
		n += 2
	}
	if m.DisableModelPruning {
		n += 2
	}
	if m.ConstantFolding {
		n += 2
	}
	if m.MemoryOptimization != 0 {
		n += 1 + sovRewriterConfig(uint64(m.MemoryOptimization))
	}
	if m.AutoParallel != nil {
		l = m.AutoParallel.Size()
		n += 1 + l + sovRewriterConfig(uint64(l))
	}
	if len(m.Optimizers) > 0 {
		for _, s := range m.Optimizers {
			l = len(s)
			n += 2 + l + sovRewriterConfig(uint64(l))
		}
	}
	return n
}

func sovRewriterConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRewriterConfig(x uint64) (n int) {
	return sovRewriterConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AutoParallelOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewriterConfig
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
			return fmt.Errorf("proto: AutoParallelOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AutoParallelOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
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
			m.Enable = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumReplicas", wireType)
			}
			m.NumReplicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumReplicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRewriterConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRewriterConfig
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
func (m *RewriterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewriterConfig
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
			return fmt.Errorf("proto: RewriterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewriterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptimizeTensorLayout", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
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
			m.OptimizeTensorLayout = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableModelPruning", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
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
			m.DisableModelPruning = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantFolding", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
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
			m.ConstantFolding = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryOptimization", wireType)
			}
			m.MemoryOptimization = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryOptimization |= (RewriterConfig_MemOptType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoParallel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRewriterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AutoParallel == nil {
				m.AutoParallel = &AutoParallelOptions{}
			}
			if err := m.AutoParallel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Optimizers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewriterConfig
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
				return ErrInvalidLengthRewriterConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Optimizers = append(m.Optimizers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewriterConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRewriterConfig
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
func skipRewriterConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewriterConfig
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
					return 0, ErrIntOverflowRewriterConfig
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
					return 0, ErrIntOverflowRewriterConfig
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
				return 0, ErrInvalidLengthRewriterConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRewriterConfig
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
				next, err := skipRewriterConfig(dAtA[start:])
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
	ErrInvalidLengthRewriterConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewriterConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rewriter_config.proto", fileDescriptorRewriterConfig) }

var fileDescriptorRewriterConfig = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x67, 0xca, 0xaa, 0x71, 0x36, 0x4a, 0xe5, 0x6e, 0x53, 0xae, 0x42, 0xa8, 0x84, 0x14,
	0x6e, 0x7a, 0x51, 0x78, 0x81, 0x0e, 0xc6, 0xd5, 0xb2, 0x44, 0x56, 0xe1, 0xd6, 0xf2, 0x5a, 0xb7,
	0xb2, 0xb0, 0x7d, 0x2c, 0xc7, 0x51, 0x55, 0xde, 0x01, 0x89, 0xc7, 0xe2, 0x92, 0x47, 0x40, 0xe5,
	0x25, 0xb8, 0x44, 0x49, 0x53, 0x2d, 0x95, 0x76, 0xe9, 0xff, 0xb7, 0xfe, 0xef, 0x9c, 0x5f, 0x07,
	0xae, 0xbc, 0xdc, 0x78, 0x15, 0xa4, 0xe7, 0x0b, 0xb4, 0x2b, 0xb5, 0x9e, 0x38, 0x8f, 0x01, 0x29,
	0x04, 0x69, 0x4b, 0xf4, 0x2b, 0x8d, 0x9b, 0x71, 0x01, 0xa3, 0x59, 0x15, 0xb0, 0x10, 0x5e, 0x68,
	0x2d, 0x75, 0xee, 0x82, 0x42, 0x5b, 0xd2, 0x6b, 0xe8, 0x4b, 0x2b, 0x1e, 0xb4, 0x8c, 0x48, 0x42,
	0xd2, 0x33, 0xd6, 0xbe, 0xe8, 0x1b, 0xb8, 0xb0, 0x95, 0xe1, 0x5e, 0x3a, 0xad, 0x16, 0xa2, 0x8c,
	0x9e, 0x25, 0x24, 0x3d, 0x65, 0xe7, 0xb6, 0x32, 0xac, 0x95, 0xc6, 0x3f, 0x7a, 0x30, 0x60, 0x2d,
	0xf7, 0x63, 0x83, 0xa5, 0x1f, 0xe0, 0x1a, 0x5d, 0x50, 0x46, 0x7d, 0x97, 0x7c, 0xcf, 0xe6, 0x5a,
	0x6c, 0xb1, 0x0a, 0x6d, 0xfa, 0xe5, 0xc1, 0x9d, 0x37, 0xe6, 0x5d, 0xe3, 0xd1, 0x29, 0x5c, 0x2d,
	0x55, 0x59, 0x63, 0xb9, 0xc1, 0xa5, 0xd4, 0xdc, 0xf9, 0xca, 0x2a, 0xbb, 0x6e, 0xa0, 0x67, 0x6c,
	0xd4, 0x9a, 0x59, 0xed, 0x15, 0x7b, 0x8b, 0xbe, 0x83, 0xe1, 0x02, 0x6d, 0x19, 0x84, 0x0d, 0x7c,
	0x85, 0x7a, 0x59, 0x7f, 0xef, 0x35, 0xdf, 0x5f, 0x1d, 0xf4, 0xcf, 0x7b, 0x99, 0x7e, 0x85, 0x91,
	0x91, 0x06, 0xfd, 0x96, 0xb7, 0x74, 0x51, 0xaf, 0x1e, 0x3d, 0x4f, 0x48, 0x3a, 0x98, 0xbe, 0x9d,
	0x3c, 0x76, 0x34, 0x39, 0xde, 0x66, 0x92, 0x49, 0x93, 0xbb, 0x30, 0xdf, 0x3a, 0xc9, 0xe8, 0x3e,
	0x21, 0xef, 0x04, 0xd0, 0x4f, 0xf0, 0x52, 0x54, 0x01, 0xb9, 0x6b, 0x2b, 0x8d, 0x4e, 0x13, 0x92,
	0x9e, 0x4f, 0x5f, 0x77, 0x13, 0x9f, 0xa8, 0x9c, 0x5d, 0x88, 0x8e, 0x48, 0x63, 0x80, 0x43, 0x29,
	0xbe, 0x8c, 0x96, 0x49, 0x2f, 0x7d, 0xc1, 0x3a, 0xca, 0x38, 0x05, 0x78, 0x9c, 0x83, 0x0e, 0x00,
	0xee, 0x73, 0x9e, 0xdd, 0x66, 0x3c, 0x2f, 0xe6, 0xc3, 0x13, 0x0a, 0xd0, 0xcf, 0x66, 0xf7, 0x5f,
	0x66, 0x77, 0x43, 0x72, 0x73, 0xfb, 0x6b, 0x17, 0x93, 0xdf, 0xbb, 0x98, 0xfc, 0xd9, 0xc5, 0xe4,
	0xe7, 0xdf, 0xf8, 0x04, 0x22, 0xf4, 0xeb, 0xee, 0x34, 0x2b, 0x2f, 0x8c, 0xdc, 0xa0, 0xff, 0x76,
	0x73, 0x79, 0xbc, 0x6a, 0x51, 0x9f, 0x4b, 0x59, 0x90, 0x7f, 0x84, 0x3c, 0xf4, 0x9b, 0xdb, 0x79,
	0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x95, 0x24, 0x2a, 0x54, 0x02, 0x00, 0x00,
}
