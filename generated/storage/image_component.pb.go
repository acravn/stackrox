// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/image_component.proto

package storage

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImageComponent struct {
	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Component ID,store" sql:"pk"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Component,store" sql:"pk"`
	Version   string     `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" search:"Component Version,store" sql:"pk"`
	License   *License   `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
	Priority  int64      `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Source    SourceType `protobuf:"varint,6,opt,name=source,proto3,enum=storage.SourceType" json:"source,omitempty" search:"Component Source,store"`
	RiskScore float32    `protobuf:"fixed32,7,opt,name=risk_score,json=riskScore,proto3" json:"risk_score,omitempty" search:"Component Risk Score,hidden"`
	// Types that are valid to be assigned to SetTopCvss:
	//	*ImageComponent_TopCvss
	SetTopCvss isImageComponent_SetTopCvss `protobuf_oneof:"set_top_cvss"`
	// Component version that fixes all the fixable vulnerabilities in this component.
	FixedBy              string   `protobuf:"bytes,9,opt,name=fixed_by,json=fixedBy,proto3" json:"fixed_by,omitempty"`
	OperatingSystem      string   `protobuf:"bytes,10,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty" sql:"pk"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageComponent) Reset()         { *m = ImageComponent{} }
func (m *ImageComponent) String() string { return proto.CompactTextString(m) }
func (*ImageComponent) ProtoMessage()    {}
func (*ImageComponent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f72cea254a8774ea, []int{0}
}
func (m *ImageComponent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImageComponent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImageComponent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImageComponent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageComponent.Merge(m, src)
}
func (m *ImageComponent) XXX_Size() int {
	return m.Size()
}
func (m *ImageComponent) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageComponent.DiscardUnknown(m)
}

var xxx_messageInfo_ImageComponent proto.InternalMessageInfo

type isImageComponent_SetTopCvss interface {
	isImageComponent_SetTopCvss()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isImageComponent_SetTopCvss
}

type ImageComponent_TopCvss struct {
	TopCvss float32 `protobuf:"fixed32,8,opt,name=top_cvss,json=topCvss,proto3,oneof" json:"top_cvss,omitempty" search:"Component Top CVSS,store"`
}

func (*ImageComponent_TopCvss) isImageComponent_SetTopCvss() {}
func (m *ImageComponent_TopCvss) Clone() isImageComponent_SetTopCvss {
	if m == nil {
		return nil
	}
	cloned := new(ImageComponent_TopCvss)
	*cloned = *m

	return cloned
}

func (m *ImageComponent) GetSetTopCvss() isImageComponent_SetTopCvss {
	if m != nil {
		return m.SetTopCvss
	}
	return nil
}

func (m *ImageComponent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImageComponent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageComponent) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ImageComponent) GetLicense() *License {
	if m != nil {
		return m.License
	}
	return nil
}

func (m *ImageComponent) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *ImageComponent) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_OS
}

func (m *ImageComponent) GetRiskScore() float32 {
	if m != nil {
		return m.RiskScore
	}
	return 0
}

func (m *ImageComponent) GetTopCvss() float32 {
	if x, ok := m.GetSetTopCvss().(*ImageComponent_TopCvss); ok {
		return x.TopCvss
	}
	return 0
}

func (m *ImageComponent) GetFixedBy() string {
	if m != nil {
		return m.FixedBy
	}
	return ""
}

func (m *ImageComponent) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ImageComponent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ImageComponent_TopCvss)(nil),
	}
}

func (m *ImageComponent) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ImageComponent) Clone() *ImageComponent {
	if m == nil {
		return nil
	}
	cloned := new(ImageComponent)
	*cloned = *m

	cloned.License = m.License.Clone()
	if m.SetTopCvss != nil {
		cloned.SetTopCvss = m.SetTopCvss.Clone()
	}
	return cloned
}

func init() {
	proto.RegisterType((*ImageComponent)(nil), "storage.ImageComponent")
}

func init() { proto.RegisterFile("storage/image_component.proto", fileDescriptor_f72cea254a8774ea) }

var fileDescriptor_f72cea254a8774ea = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x6f, 0xd3, 0x30,
	0x1c, 0xc6, 0xe7, 0xae, 0x34, 0xad, 0x99, 0xca, 0xe4, 0x71, 0xf0, 0x2a, 0xd1, 0x98, 0x00, 0x52,
	0x40, 0xa3, 0x93, 0x06, 0x12, 0xb0, 0x63, 0x8a, 0x34, 0x26, 0x71, 0x40, 0xc9, 0xb4, 0x03, 0x97,
	0x28, 0x4b, 0x4c, 0x66, 0x75, 0x8d, 0x8d, 0x6d, 0xaa, 0xe5, 0x9b, 0xf0, 0x91, 0x38, 0x72, 0xe5,
	0x12, 0xa1, 0xf2, 0x0d, 0xf2, 0x09, 0x50, 0x9c, 0x17, 0x09, 0x85, 0xdd, 0xec, 0xc7, 0xbf, 0xe7,
	0xf1, 0x5f, 0x8f, 0x0d, 0x1f, 0x29, 0xcd, 0x65, 0x94, 0xd2, 0x63, 0xb6, 0x8e, 0x52, 0x1a, 0xc6,
	0x7c, 0x2d, 0x78, 0x46, 0x33, 0xbd, 0x10, 0x92, 0x6b, 0x8e, 0xac, 0xe6, 0x78, 0x76, 0xf0, 0x0f,
	0x57, 0x9f, 0xce, 0x1e, 0xa6, 0x3c, 0xe5, 0x66, 0x79, 0x5c, 0xad, 0x6a, 0xd5, 0xf9, 0x35, 0x84,
	0xd3, 0xf3, 0x8a, 0x5a, 0xb6, 0x61, 0xe8, 0x2d, 0x1c, 0xb0, 0x04, 0x03, 0x02, 0xdc, 0x89, 0xe7,
	0x96, 0x85, 0xfd, 0x54, 0xd1, 0x48, 0xc6, 0xd7, 0xa7, 0x4e, 0x87, 0x90, 0xf3, 0xf7, 0x47, 0xd5,
	0x15, 0xd4, 0x21, 0xea, 0xeb, 0xcd, 0xa9, 0x23, 0x56, 0x8e, 0x3f, 0x60, 0x09, 0x7a, 0x07, 0x87,
	0x59, 0xb4, 0xa6, 0x78, 0x60, 0xbc, 0xcf, 0xca, 0xc2, 0x7e, 0xdc, 0xf3, 0xf6, 0x8c, 0xc6, 0x82,
	0xce, 0xa0, 0xb5, 0xa1, 0x52, 0x31, 0x9e, 0xe1, 0x5d, 0xe3, 0x7e, 0x59, 0x16, 0xf6, 0xf3, 0xfe,
	0xcd, 0x97, 0x35, 0xd4, 0x4b, 0x69, 0xdd, 0xe8, 0x05, 0xb4, 0x6e, 0x58, 0x4c, 0x33, 0x45, 0xf1,
	0x90, 0x00, 0xf7, 0xfe, 0xc9, 0xfe, 0xa2, 0x69, 0x63, 0xf1, 0xb1, 0xd6, 0xfd, 0x16, 0x40, 0x33,
	0x38, 0x16, 0x92, 0x71, 0xc9, 0x74, 0x8e, 0xef, 0x11, 0xe0, 0xee, 0xfa, 0xdd, 0x1e, 0x7d, 0x82,
	0x23, 0xc5, 0xbf, 0xc9, 0x98, 0xe2, 0x11, 0x01, 0xee, 0xf4, 0xe4, 0xa0, 0x8b, 0x09, 0x8c, 0x7c,
	0x91, 0x0b, 0xea, 0x3d, 0x29, 0x0b, 0xdb, 0xee, 0x0f, 0x59, 0x13, 0xcd, 0x8c, 0x7e, 0x93, 0x83,
	0xce, 0x20, 0x94, 0x4c, 0xad, 0x42, 0x15, 0x73, 0x49, 0xb1, 0x45, 0x80, 0x3b, 0xb8, 0xab, 0x5f,
	0x9f, 0xa9, 0x15, 0x09, 0x2a, 0xee, 0xe8, 0x9a, 0x25, 0x09, 0xcd, 0x1c, 0x7f, 0x52, 0x79, 0x8d,
	0x84, 0x3c, 0x38, 0xd6, 0x5c, 0x84, 0xf1, 0x46, 0x29, 0x3c, 0x36, 0x31, 0xff, 0xaf, 0x9a, 0x5c,
	0x70, 0x41, 0x96, 0x97, 0x41, 0xd0, 0x4c, 0xf2, 0x61, 0xc7, 0xb7, 0x34, 0x17, 0xcb, 0x8d, 0x52,
	0xe8, 0x10, 0x8e, 0xbf, 0xb0, 0x5b, 0x9a, 0x84, 0x57, 0x39, 0x9e, 0x54, 0x85, 0xfb, 0x96, 0xd9,
	0x7b, 0x39, 0x7a, 0x03, 0xf7, 0xb9, 0xa0, 0x32, 0xd2, 0x2c, 0x4b, 0x43, 0x95, 0x2b, 0x4d, 0xd7,
	0x18, 0x9a, 0x37, 0xd9, 0x2b, 0x0b, 0x7b, 0xdc, 0x55, 0xfe, 0xa0, 0xa3, 0x02, 0x03, 0x79, 0x53,
	0xb8, 0xa7, 0xa8, 0x0e, 0xdb, 0xd9, 0xbc, 0xd7, 0x3f, 0xb6, 0x73, 0xf0, 0x73, 0x3b, 0x07, 0xbf,
	0xb7, 0x73, 0xf0, 0xfd, 0xcf, 0x7c, 0x07, 0x1e, 0x32, 0xbe, 0x50, 0x3a, 0x8a, 0x57, 0x92, 0xdf,
	0xd6, 0x1f, 0xb0, 0x6d, 0xf5, 0x73, 0xfb, 0x79, 0xaf, 0x46, 0x46, 0x7f, 0xf5, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x71, 0xd7, 0xab, 0xed, 0x02, 0x00, 0x00,
}

func (m *ImageComponent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImageComponent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImageComponent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.OperatingSystem) > 0 {
		i -= len(m.OperatingSystem)
		copy(dAtA[i:], m.OperatingSystem)
		i = encodeVarintImageComponent(dAtA, i, uint64(len(m.OperatingSystem)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.FixedBy) > 0 {
		i -= len(m.FixedBy)
		copy(dAtA[i:], m.FixedBy)
		i = encodeVarintImageComponent(dAtA, i, uint64(len(m.FixedBy)))
		i--
		dAtA[i] = 0x4a
	}
	if m.SetTopCvss != nil {
		{
			size := m.SetTopCvss.Size()
			i -= size
			if _, err := m.SetTopCvss.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.RiskScore != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.RiskScore))))
		i--
		dAtA[i] = 0x3d
	}
	if m.Source != 0 {
		i = encodeVarintImageComponent(dAtA, i, uint64(m.Source))
		i--
		dAtA[i] = 0x30
	}
	if m.Priority != 0 {
		i = encodeVarintImageComponent(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x28
	}
	if m.License != nil {
		{
			size, err := m.License.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintImageComponent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintImageComponent(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintImageComponent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintImageComponent(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ImageComponent_TopCvss) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImageComponent_TopCvss) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.TopCvss))))
	i--
	dAtA[i] = 0x45
	return len(dAtA) - i, nil
}
func encodeVarintImageComponent(dAtA []byte, offset int, v uint64) int {
	offset -= sovImageComponent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ImageComponent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovImageComponent(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovImageComponent(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovImageComponent(uint64(l))
	}
	if m.License != nil {
		l = m.License.Size()
		n += 1 + l + sovImageComponent(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovImageComponent(uint64(m.Priority))
	}
	if m.Source != 0 {
		n += 1 + sovImageComponent(uint64(m.Source))
	}
	if m.RiskScore != 0 {
		n += 5
	}
	if m.SetTopCvss != nil {
		n += m.SetTopCvss.Size()
	}
	l = len(m.FixedBy)
	if l > 0 {
		n += 1 + l + sovImageComponent(uint64(l))
	}
	l = len(m.OperatingSystem)
	if l > 0 {
		n += 1 + l + sovImageComponent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImageComponent_TopCvss) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 5
	return n
}

func sovImageComponent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozImageComponent(x uint64) (n int) {
	return sovImageComponent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ImageComponent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageComponent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImageComponent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageComponent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field License", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.License == nil {
				m.License = &License{}
			}
			if err := m.License.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			m.Source = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Source |= SourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field RiskScore", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.RiskScore = float32(math.Float32frombits(v))
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopCvss", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.SetTopCvss = &ImageComponent_TopCvss{float32(math.Float32frombits(v))}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FixedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatingSystem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthImageComponent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageComponent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatingSystem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImageComponent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageComponent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipImageComponent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImageComponent
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
					return 0, ErrIntOverflowImageComponent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowImageComponent
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
			if length < 0 {
				return 0, ErrInvalidLengthImageComponent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupImageComponent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthImageComponent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthImageComponent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImageComponent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupImageComponent = fmt.Errorf("proto: unexpected end of group")
)
