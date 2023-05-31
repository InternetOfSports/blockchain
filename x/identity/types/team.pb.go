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
	Index           string                  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name            string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players         map[string]*Participant `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Owner           *Participant            `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Manager         *Participant            `protobuf:"bytes,5,opt,name=manager,proto3" json:"manager,omitempty"`
	Trainer         *Participant            `protobuf:"bytes,6,opt,name=trainer,proto3" json:"trainer,omitempty"`
	JoiningRequests map[string]*Participant `protobuf:"bytes,7,rep,name=joiningRequests,proto3" json:"joiningRequests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	JoiningInvites  map[string]*Participant `protobuf:"bytes,8,rep,name=joiningInvites,proto3" json:"joiningInvites,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *Team) GetPlayers() map[string]*Participant {
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

func (m *Team) GetJoiningRequests() map[string]*Participant {
	if m != nil {
		return m.JoiningRequests
	}
	return nil
}

func (m *Team) GetJoiningInvites() map[string]*Participant {
	if m != nil {
		return m.JoiningInvites
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "internetofsports.blockchain.identity.Team")
	proto.RegisterMapType((map[string]*Participant)(nil), "internetofsports.blockchain.identity.Team.JoiningInvitesEntry")
	proto.RegisterMapType((map[string]*Participant)(nil), "internetofsports.blockchain.identity.Team.JoiningRequestsEntry")
	proto.RegisterMapType((map[string]*Participant)(nil), "internetofsports.blockchain.identity.Team.PlayersEntry")
}

func init() { proto.RegisterFile("blockchain/identity/team.proto", fileDescriptor_07a7483b85c85ec1) }

var fileDescriptor_07a7483b85c85ec1 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x3f, 0x4f, 0xfa, 0x40,
	0x18, 0xc7, 0x29, 0xff, 0xfa, 0xfb, 0x1d, 0x46, 0xcd, 0xc9, 0xd0, 0x30, 0x34, 0xc4, 0x68, 0xc2,
	0xd4, 0x46, 0x1c, 0x30, 0x0e, 0x9a, 0x98, 0x18, 0x83, 0x0e, 0x42, 0x75, 0x72, 0x3b, 0xca, 0x01,
	0x27, 0xf4, 0xae, 0x5e, 0x9f, 0x22, 0x7d, 0x17, 0x2e, 0xbe, 0x27, 0x47, 0x46, 0x47, 0x03, 0x6f,
	0xc4, 0xf4, 0x0f, 0x01, 0x09, 0x03, 0x90, 0xb0, 0x5d, 0x2f, 0xf9, 0x7e, 0x3e, 0x77, 0xcf, 0x73,
	0x4f, 0x91, 0xde, 0x1a, 0x08, 0xbb, 0x6f, 0xf7, 0x08, 0xe3, 0x26, 0x6b, 0x53, 0x0e, 0x0c, 0x02,
	0x13, 0x28, 0x71, 0x0c, 0x57, 0x0a, 0x10, 0xf8, 0x84, 0x71, 0xa0, 0x92, 0x53, 0x10, 0x1d, 0xcf,
	0x15, 0x12, 0x3c, 0x63, 0x1e, 0x30, 0x66, 0x81, 0xd2, 0xe9, 0x2a, 0x8a, 0x4b, 0x24, 0x30, 0x9b,
	0xb9, 0x84, 0x43, 0x0c, 0x3b, 0xfe, 0x54, 0x51, 0xf6, 0x99, 0x12, 0x07, 0x17, 0x51, 0x8e, 0xf1,
	0x36, 0x1d, 0x69, 0x4a, 0x59, 0xa9, 0xfc, 0xb7, 0xe2, 0x0f, 0x8c, 0x51, 0x96, 0x13, 0x87, 0x6a,
	0xe9, 0x68, 0x33, 0x5a, 0xe3, 0x26, 0x52, 0xdd, 0x01, 0x09, 0xa8, 0xf4, 0xb4, 0x4c, 0x39, 0x53,
	0x29, 0x54, 0x6b, 0xc6, 0x3a, 0x27, 0x32, 0x42, 0x8d, 0xd1, 0x88, 0x93, 0xb7, 0x1c, 0x64, 0x60,
	0xcd, 0x38, 0xf8, 0x0e, 0xe5, 0xc4, 0x3b, 0xa7, 0x52, 0xcb, 0x96, 0x95, 0x4a, 0xa1, 0x7a, 0xb6,
	0x1e, 0xb0, 0x31, 0xbf, 0x8d, 0x15, 0xe7, 0xf1, 0x03, 0x52, 0x1d, 0xc2, 0x49, 0x97, 0x4a, 0x2d,
	0xb7, 0x2d, 0x6a, 0x46, 0x08, 0x61, 0x20, 0x09, 0x0b, 0xcf, 0x95, 0xdf, 0x1a, 0x96, 0x10, 0x30,
	0x43, 0x07, 0xaf, 0x82, 0x71, 0xc6, 0xbb, 0x16, 0x7d, 0xf3, 0xa9, 0x07, 0x9e, 0xa6, 0x46, 0xd5,
	0xbb, 0xde, 0xa0, 0x7a, 0xf7, 0x7f, 0x09, 0x71, 0x15, 0x97, 0xb9, 0xb8, 0x83, 0xf6, 0x93, 0xad,
	0x3a, 0x1f, 0x32, 0xa0, 0x9e, 0xf6, 0x2f, 0x32, 0x5d, 0x6d, 0x6e, 0x4a, 0x00, 0xb1, 0x68, 0x89,
	0x5a, 0x72, 0xd0, 0xde, 0x62, 0x3b, 0xf1, 0x21, 0xca, 0xf4, 0x69, 0x90, 0x3c, 0xa0, 0x70, 0x19,
	0xf6, 0x75, 0x48, 0x06, 0x7e, 0xfc, 0x7e, 0xb6, 0xeb, 0x6b, 0x94, 0xbf, 0x4c, 0x5f, 0x28, 0x25,
	0x1f, 0x15, 0x57, 0xdd, 0x7f, 0xd7, 0x5a, 0x40, 0x47, 0x2b, 0x8a, 0xb1, 0x63, 0xeb, 0x4d, 0xf3,
	0x6b, 0xa2, 0x2b, 0xe3, 0x89, 0xae, 0xfc, 0x4c, 0x74, 0xe5, 0x63, 0xaa, 0xa7, 0xc6, 0x53, 0x3d,
	0xf5, 0x3d, 0xd5, 0x53, 0x2f, 0xb5, 0x2e, 0x83, 0x9e, 0xdf, 0x32, 0x6c, 0xe1, 0x98, 0xf5, 0xc4,
	0xf0, 0xd8, 0x79, 0x8a, 0x0c, 0xe6, 0xc2, 0xd0, 0x8f, 0x16, 0x7e, 0x1e, 0x81, 0x4b, 0xbd, 0x56,
	0x3e, 0x9a, 0xf8, 0xf3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x21, 0x99, 0x56, 0x60, 0x04,
	0x00, 0x00,
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
	if len(m.JoiningInvites) > 0 {
		for k := range m.JoiningInvites {
			v := m.JoiningInvites[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTeam(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTeam(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTeam(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.JoiningRequests) > 0 {
		for k := range m.JoiningRequests {
			v := m.JoiningRequests[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTeam(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTeam(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTeam(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
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
		for k := range m.Players {
			v := m.Players[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTeam(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTeam(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTeam(dAtA, i, uint64(baseI-i))
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
		for k, v := range m.Players {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTeam(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTeam(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTeam(uint64(mapEntrySize))
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
	if len(m.JoiningRequests) > 0 {
		for k, v := range m.JoiningRequests {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTeam(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTeam(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTeam(uint64(mapEntrySize))
		}
	}
	if len(m.JoiningInvites) > 0 {
		for k, v := range m.JoiningInvites {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTeam(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTeam(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTeam(uint64(mapEntrySize))
		}
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
			if m.Players == nil {
				m.Players = make(map[string]*Participant)
			}
			var mapkey string
			var mapvalue *Participant
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTeam
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTeam
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Participant{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTeam(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTeam
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Players[mapkey] = mapvalue
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoiningRequests", wireType)
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
			if m.JoiningRequests == nil {
				m.JoiningRequests = make(map[string]*Participant)
			}
			var mapkey string
			var mapvalue *Participant
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTeam
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTeam
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Participant{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTeam(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTeam
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.JoiningRequests[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoiningInvites", wireType)
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
			if m.JoiningInvites == nil {
				m.JoiningInvites = make(map[string]*Participant)
			}
			var mapkey string
			var mapvalue *Participant
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTeam
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTeam
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTeam
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTeam
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Participant{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTeam(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTeam
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.JoiningInvites[mapkey] = mapvalue
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
