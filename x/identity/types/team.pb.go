// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockchain/identity/team.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Team struct {
	Index   string         `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name    string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players []*Participant `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	Owner   *Participant   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Manager *Participant   `protobuf:"bytes,5,opt,name=manager,proto3" json:"manager,omitempty"`
	Trainer *Participant   `protobuf:"bytes,6,opt,name=trainer,proto3" json:"trainer,omitempty"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a7483b85c85ec1, []int{0}
}
func (m *Team) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Team.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return m.Size()
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetPlayers() []*Participant {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *Team) GetOwner() *Participant {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Team) GetManager() *Participant {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *Team) GetTrainer() *Participant {
	if m != nil {
		return m.Trainer
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "internetofsports.blockchain.identity.Team")
}

func init() { proto.RegisterFile("blockchain/identity/team.proto", fileDescriptor_07a7483b85c85ec1) }

var fileDescriptor_07a7483b85c85ec1 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0xe0, 0xb8, 0x4d, 0x8a, 0x30, 0x9b, 0xc5, 0x10, 0x31, 0x58, 0x11, 0x02, 0xa9, 0x93,
	0x23, 0x60, 0x60, 0x67, 0x41, 0x88, 0x01, 0x28, 0x4c, 0x6c, 0x4e, 0x7a, 0x6d, 0x2d, 0x1a, 0xdb,
	0x72, 0x0e, 0xd1, 0xbc, 0x05, 0x8f, 0xc5, 0xd8, 0x91, 0xb1, 0x4a, 0x5e, 0x04, 0x25, 0x69, 0xd4,
	0x0e, 0x0c, 0x55, 0x36, 0xdb, 0xd2, 0xff, 0xf9, 0x97, 0xee, 0x28, 0x4f, 0x96, 0x26, 0xfd, 0x48,
	0x17, 0x52, 0xe9, 0x58, 0x4d, 0x41, 0xa3, 0xc2, 0x22, 0x46, 0x90, 0x99, 0xb0, 0xce, 0xa0, 0x61,
	0x17, 0x4a, 0x23, 0x38, 0x0d, 0x68, 0x66, 0xb9, 0x35, 0x0e, 0x73, 0xb1, 0x0b, 0x88, 0x2e, 0x70,
	0x76, 0xf9, 0x9f, 0x62, 0xa5, 0x43, 0x95, 0x2a, 0x2b, 0x35, 0xb6, 0xd8, 0xf9, 0x66, 0x40, 0xfd,
	0x37, 0x90, 0x19, 0x3b, 0xa5, 0x81, 0xd2, 0x53, 0x58, 0x85, 0x24, 0x22, 0xe3, 0xe3, 0x49, 0x7b,
	0x61, 0x8c, 0xfa, 0x5a, 0x66, 0x10, 0x0e, 0x9a, 0xc7, 0xe6, 0xcc, 0x1e, 0xe9, 0x91, 0x5d, 0xca,
	0x02, 0x5c, 0x1e, 0x0e, 0xa3, 0xe1, 0xf8, 0xe4, 0xfa, 0x4a, 0x1c, 0xd2, 0x48, 0x3c, 0xef, 0x3e,
	0x9f, 0x74, 0x02, 0xbb, 0xa7, 0x81, 0xf9, 0xd2, 0xe0, 0x42, 0x3f, 0x22, 0xfd, 0xa8, 0x36, 0x5f,
	0xb7, 0xca, 0xa4, 0x96, 0x73, 0x70, 0x61, 0xd0, 0x97, 0xea, 0x84, 0x1a, 0x43, 0x27, 0x55, 0xdd,
	0x6b, 0xd4, 0x1b, 0xdb, 0x0a, 0x77, 0x2f, 0x3f, 0x25, 0x27, 0xeb, 0x92, 0x93, 0x4d, 0xc9, 0xc9,
	0x77, 0xc5, 0xbd, 0x75, 0xc5, 0xbd, 0xdf, 0x8a, 0x7b, 0xef, 0xb7, 0x73, 0x85, 0x8b, 0xcf, 0x44,
	0xa4, 0x26, 0x8b, 0x1f, 0xb6, 0xfe, 0xd3, 0xec, 0xb5, 0xf1, 0xe3, 0xbd, 0xf9, 0xad, 0xf6, 0xf6,
	0xa0, 0xb0, 0x90, 0x27, 0xa3, 0x66, 0x78, 0x37, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x4f,
	0xec, 0xfb, 0x2b, 0x02, 0x00, 0x00,
}

func (m *Team) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Team) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Team) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Trainer != nil {
		{
			size, err := m.Trainer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTeam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Manager != nil {
		{
			size, err := m.Manager.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTeam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Owner != nil {
		{
			size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTeam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Players) > 0 {
		for iNdEx := len(m.Players) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Players[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTeam(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTeam(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTeam(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTeam(dAtA []byte, offset int, v uint64) int {
	offset -= sovTeam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Team) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTeam(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTeam(uint64(l))
	}
	if len(m.Players) > 0 {
		for _, e := range m.Players {
			l = e.Size()
			n += 1 + l + sovTeam(uint64(l))
		}
	}
	if m.Owner != nil {
		l = m.Owner.Size()
		n += 1 + l + sovTeam(uint64(l))
	}
	if m.Manager != nil {
		l = m.Manager.Size()
		n += 1 + l + sovTeam(uint64(l))
	}
	if m.Trainer != nil {
		l = m.Trainer.Size()
		n += 1 + l + sovTeam(uint64(l))
	}
	return n
}

func sovTeam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTeam(x uint64) (n int) {
	return sovTeam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Team) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTeam
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
			return fmt.Errorf("proto: Team: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Team: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Players", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Players = append(m.Players, &Participant{})
			if err := m.Players[len(m.Players)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Owner == nil {
				m.Owner = &Participant{}
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manager", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Manager == nil {
				m.Manager = &Participant{}
			}
			if err := m.Manager.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trainer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
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
				return ErrInvalidLengthTeam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTeam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trainer == nil {
				m.Trainer = &Participant{}
			}
			if err := m.Trainer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTeam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTeam
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
func skipTeam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTeam
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
					return 0, ErrIntOverflowTeam
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
					return 0, ErrIntOverflowTeam
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
				return 0, ErrInvalidLengthTeam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTeam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTeam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTeam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTeam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTeam = fmt.Errorf("proto: unexpected end of group")
)