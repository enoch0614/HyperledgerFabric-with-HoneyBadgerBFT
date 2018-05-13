// Code generated by protoc-gen-go. DO NOT EDIT.
// source: honeybadgerbft.proto

/*
Package orderer is a generated protocol buffer package.

It is generated from these files:
	honeybadgerbft.proto

It has these top-level messages:
	HoneyBadgerBFTMessage
	HoneyBadgerBFTMessageNewHeight
	HoneyBadgerBFTMessageCommonCoin
	HoneyBadgerBFTMessageBinaryAgreement
	HoneyBadgerBFTMessageBinaryAgreementEST
	HoneyBadgerBFTMessageBinaryAgreementAUX
	HoneyBadgerBFTMessageReliableBroadcast
	HoneyBadgerBFTMessageReliableBroadcastVAL
	HoneyBadgerBFTMessageReliableBroadcastECHO
	HoneyBadgerBFTMessageReliableBroadcastREADY
	HoneyBadgerBFTMessageThresholdEncryption
	HoneyBadgerBFTMessageThresholdEncryptionShare
	HoneyBadgerBFTMessageRequireBlock
	HoneyBadgerBFTMessageResponseBlock
*/
package orderer

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

// HoneyBadgerBFTMessage is a wrapper type for the messages
// that the Kafka-based orderer deals with.
type HoneyBadgerBFTMessage struct {
	Sender   uint64 `protobuf:"varint,1,opt,name=sender" json:"sender,omitempty"`
	Receiver uint64 `protobuf:"varint,2,opt,name=receiver" json:"receiver,omitempty"`
	ChainId  string `protobuf:"bytes,3,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	Height   uint64 `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Instance uint64 `protobuf:"varint,5,opt,name=instance" json:"instance,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*HoneyBadgerBFTMessage_New_Height
	//	*HoneyBadgerBFTMessage_CommonCoin
	//	*HoneyBadgerBFTMessage_BinaryAgreement
	//	*HoneyBadgerBFTMessage_ReliableBroadcast
	//	*HoneyBadgerBFTMessage_ThresholdEncryption
	//	*HoneyBadgerBFTMessage_RequireBlock
	//	*HoneyBadgerBFTMessage_ResponseBlock
	Type isHoneyBadgerBFTMessage_Type `protobuf_oneof:"Type"`
}

func (m *HoneyBadgerBFTMessage) Reset()                    { *m = HoneyBadgerBFTMessage{} }
func (m *HoneyBadgerBFTMessage) String() string            { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessage) ProtoMessage()               {}
func (*HoneyBadgerBFTMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isHoneyBadgerBFTMessage_Type interface{ isHoneyBadgerBFTMessage_Type() }

type HoneyBadgerBFTMessage_New_Height struct {
	New_Height *HoneyBadgerBFTMessageNewHeight `protobuf:"bytes,6,opt,name=new_Height,json=newHeight,oneof"`
}
type HoneyBadgerBFTMessage_CommonCoin struct {
	CommonCoin *HoneyBadgerBFTMessageCommonCoin `protobuf:"bytes,7,opt,name=common_coin,json=commonCoin,oneof"`
}
type HoneyBadgerBFTMessage_BinaryAgreement struct {
	BinaryAgreement *HoneyBadgerBFTMessageBinaryAgreement `protobuf:"bytes,8,opt,name=binary_agreement,json=binaryAgreement,oneof"`
}
type HoneyBadgerBFTMessage_ReliableBroadcast struct {
	ReliableBroadcast *HoneyBadgerBFTMessageReliableBroadcast `protobuf:"bytes,9,opt,name=reliable_broadcast,json=reliableBroadcast,oneof"`
}
type HoneyBadgerBFTMessage_ThresholdEncryption struct {
	ThresholdEncryption *HoneyBadgerBFTMessageThresholdEncryption `protobuf:"bytes,10,opt,name=threshold_encryption,json=thresholdEncryption,oneof"`
}
type HoneyBadgerBFTMessage_RequireBlock struct {
	RequireBlock *HoneyBadgerBFTMessageRequireBlock `protobuf:"bytes,11,opt,name=require_block,json=requireBlock,oneof"`
}
type HoneyBadgerBFTMessage_ResponseBlock struct {
	ResponseBlock *HoneyBadgerBFTMessageResponseBlock `protobuf:"bytes,12,opt,name=response_block,json=responseBlock,oneof"`
}

func (*HoneyBadgerBFTMessage_New_Height) isHoneyBadgerBFTMessage_Type()          {}
func (*HoneyBadgerBFTMessage_CommonCoin) isHoneyBadgerBFTMessage_Type()          {}
func (*HoneyBadgerBFTMessage_BinaryAgreement) isHoneyBadgerBFTMessage_Type()     {}
func (*HoneyBadgerBFTMessage_ReliableBroadcast) isHoneyBadgerBFTMessage_Type()   {}
func (*HoneyBadgerBFTMessage_ThresholdEncryption) isHoneyBadgerBFTMessage_Type() {}
func (*HoneyBadgerBFTMessage_RequireBlock) isHoneyBadgerBFTMessage_Type()        {}
func (*HoneyBadgerBFTMessage_ResponseBlock) isHoneyBadgerBFTMessage_Type()       {}

func (m *HoneyBadgerBFTMessage) GetType() isHoneyBadgerBFTMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetSender() uint64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *HoneyBadgerBFTMessage) GetReceiver() uint64 {
	if m != nil {
		return m.Receiver
	}
	return 0
}

func (m *HoneyBadgerBFTMessage) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *HoneyBadgerBFTMessage) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *HoneyBadgerBFTMessage) GetInstance() uint64 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *HoneyBadgerBFTMessage) GetNew_Height() *HoneyBadgerBFTMessageNewHeight {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_New_Height); ok {
		return x.New_Height
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetCommonCoin() *HoneyBadgerBFTMessageCommonCoin {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_CommonCoin); ok {
		return x.CommonCoin
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetBinaryAgreement() *HoneyBadgerBFTMessageBinaryAgreement {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_BinaryAgreement); ok {
		return x.BinaryAgreement
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetReliableBroadcast() *HoneyBadgerBFTMessageReliableBroadcast {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_ReliableBroadcast); ok {
		return x.ReliableBroadcast
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetThresholdEncryption() *HoneyBadgerBFTMessageThresholdEncryption {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_ThresholdEncryption); ok {
		return x.ThresholdEncryption
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetRequireBlock() *HoneyBadgerBFTMessageRequireBlock {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_RequireBlock); ok {
		return x.RequireBlock
	}
	return nil
}

func (m *HoneyBadgerBFTMessage) GetResponseBlock() *HoneyBadgerBFTMessageResponseBlock {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessage_ResponseBlock); ok {
		return x.ResponseBlock
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HoneyBadgerBFTMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HoneyBadgerBFTMessage_OneofMarshaler, _HoneyBadgerBFTMessage_OneofUnmarshaler, _HoneyBadgerBFTMessage_OneofSizer, []interface{}{
		(*HoneyBadgerBFTMessage_New_Height)(nil),
		(*HoneyBadgerBFTMessage_CommonCoin)(nil),
		(*HoneyBadgerBFTMessage_BinaryAgreement)(nil),
		(*HoneyBadgerBFTMessage_ReliableBroadcast)(nil),
		(*HoneyBadgerBFTMessage_ThresholdEncryption)(nil),
		(*HoneyBadgerBFTMessage_RequireBlock)(nil),
		(*HoneyBadgerBFTMessage_ResponseBlock)(nil),
	}
}

func _HoneyBadgerBFTMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HoneyBadgerBFTMessage)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessage_New_Height:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.New_Height); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_CommonCoin:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CommonCoin); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_BinaryAgreement:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BinaryAgreement); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_ReliableBroadcast:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReliableBroadcast); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_ThresholdEncryption:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ThresholdEncryption); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_RequireBlock:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequireBlock); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessage_ResponseBlock:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponseBlock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HoneyBadgerBFTMessage.Type has unexpected type %T", x)
	}
	return nil
}

func _HoneyBadgerBFTMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HoneyBadgerBFTMessage)
	switch tag {
	case 6: // Type.new_Height
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageNewHeight)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_New_Height{msg}
		return true, err
	case 7: // Type.common_coin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageCommonCoin)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_CommonCoin{msg}
		return true, err
	case 8: // Type.binary_agreement
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageBinaryAgreement)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_BinaryAgreement{msg}
		return true, err
	case 9: // Type.reliable_broadcast
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageReliableBroadcast)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_ReliableBroadcast{msg}
		return true, err
	case 10: // Type.threshold_encryption
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageThresholdEncryption)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_ThresholdEncryption{msg}
		return true, err
	case 11: // Type.require_block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageRequireBlock)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_RequireBlock{msg}
		return true, err
	case 12: // Type.response_block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageResponseBlock)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessage_ResponseBlock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HoneyBadgerBFTMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HoneyBadgerBFTMessage)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessage_New_Height:
		s := proto.Size(x.New_Height)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_CommonCoin:
		s := proto.Size(x.CommonCoin)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_BinaryAgreement:
		s := proto.Size(x.BinaryAgreement)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_ReliableBroadcast:
		s := proto.Size(x.ReliableBroadcast)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_ThresholdEncryption:
		s := proto.Size(x.ThresholdEncryption)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_RequireBlock:
		s := proto.Size(x.RequireBlock)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessage_ResponseBlock:
		s := proto.Size(x.ResponseBlock)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HoneyBadgerBFTMessageNewHeight struct {
}

func (m *HoneyBadgerBFTMessageNewHeight) Reset()                    { *m = HoneyBadgerBFTMessageNewHeight{} }
func (m *HoneyBadgerBFTMessageNewHeight) String() string            { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageNewHeight) ProtoMessage()               {}
func (*HoneyBadgerBFTMessageNewHeight) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type HoneyBadgerBFTMessageCommonCoin struct {
	Round   uint64 `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *HoneyBadgerBFTMessageCommonCoin) Reset()                    { *m = HoneyBadgerBFTMessageCommonCoin{} }
func (m *HoneyBadgerBFTMessageCommonCoin) String() string            { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageCommonCoin) ProtoMessage()               {}
func (*HoneyBadgerBFTMessageCommonCoin) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *HoneyBadgerBFTMessageCommonCoin) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *HoneyBadgerBFTMessageCommonCoin) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type HoneyBadgerBFTMessageBinaryAgreement struct {
	// Types that are valid to be assigned to Type:
	//	*HoneyBadgerBFTMessageBinaryAgreement_Est
	//	*HoneyBadgerBFTMessageBinaryAgreement_Aux
	Type  isHoneyBadgerBFTMessageBinaryAgreement_Type `protobuf_oneof:"Type"`
	Round uint64                                      `protobuf:"varint,3,opt,name=round" json:"round,omitempty"`
	Value bool                                        `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
}

func (m *HoneyBadgerBFTMessageBinaryAgreement) Reset()         { *m = HoneyBadgerBFTMessageBinaryAgreement{} }
func (m *HoneyBadgerBFTMessageBinaryAgreement) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageBinaryAgreement) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageBinaryAgreement) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{3}
}

type isHoneyBadgerBFTMessageBinaryAgreement_Type interface{ isHoneyBadgerBFTMessageBinaryAgreement_Type() }

type HoneyBadgerBFTMessageBinaryAgreement_Est struct {
	Est *HoneyBadgerBFTMessageBinaryAgreementEST `protobuf:"bytes,1,opt,name=est,oneof"`
}
type HoneyBadgerBFTMessageBinaryAgreement_Aux struct {
	Aux *HoneyBadgerBFTMessageBinaryAgreementAUX `protobuf:"bytes,2,opt,name=aux,oneof"`
}

func (*HoneyBadgerBFTMessageBinaryAgreement_Est) isHoneyBadgerBFTMessageBinaryAgreement_Type() {}
func (*HoneyBadgerBFTMessageBinaryAgreement_Aux) isHoneyBadgerBFTMessageBinaryAgreement_Type() {}

func (m *HoneyBadgerBFTMessageBinaryAgreement) GetType() isHoneyBadgerBFTMessageBinaryAgreement_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HoneyBadgerBFTMessageBinaryAgreement) GetEst() *HoneyBadgerBFTMessageBinaryAgreementEST {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessageBinaryAgreement_Est); ok {
		return x.Est
	}
	return nil
}

func (m *HoneyBadgerBFTMessageBinaryAgreement) GetAux() *HoneyBadgerBFTMessageBinaryAgreementAUX {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessageBinaryAgreement_Aux); ok {
		return x.Aux
	}
	return nil
}

func (m *HoneyBadgerBFTMessageBinaryAgreement) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *HoneyBadgerBFTMessageBinaryAgreement) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HoneyBadgerBFTMessageBinaryAgreement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HoneyBadgerBFTMessageBinaryAgreement_OneofMarshaler, _HoneyBadgerBFTMessageBinaryAgreement_OneofUnmarshaler, _HoneyBadgerBFTMessageBinaryAgreement_OneofSizer, []interface{}{
		(*HoneyBadgerBFTMessageBinaryAgreement_Est)(nil),
		(*HoneyBadgerBFTMessageBinaryAgreement_Aux)(nil),
	}
}

func _HoneyBadgerBFTMessageBinaryAgreement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HoneyBadgerBFTMessageBinaryAgreement)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessageBinaryAgreement_Est:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Est); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessageBinaryAgreement_Aux:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aux); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HoneyBadgerBFTMessageBinaryAgreement.Type has unexpected type %T", x)
	}
	return nil
}

func _HoneyBadgerBFTMessageBinaryAgreement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HoneyBadgerBFTMessageBinaryAgreement)
	switch tag {
	case 1: // Type.est
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageBinaryAgreementEST)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessageBinaryAgreement_Est{msg}
		return true, err
	case 2: // Type.aux
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageBinaryAgreementAUX)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessageBinaryAgreement_Aux{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HoneyBadgerBFTMessageBinaryAgreement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HoneyBadgerBFTMessageBinaryAgreement)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessageBinaryAgreement_Est:
		s := proto.Size(x.Est)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessageBinaryAgreement_Aux:
		s := proto.Size(x.Aux)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HoneyBadgerBFTMessageBinaryAgreementEST struct {
}

func (m *HoneyBadgerBFTMessageBinaryAgreementEST) Reset() {
	*m = HoneyBadgerBFTMessageBinaryAgreementEST{}
}
func (m *HoneyBadgerBFTMessageBinaryAgreementEST) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageBinaryAgreementEST) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageBinaryAgreementEST) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{4}
}

type HoneyBadgerBFTMessageBinaryAgreementAUX struct {
}

func (m *HoneyBadgerBFTMessageBinaryAgreementAUX) Reset() {
	*m = HoneyBadgerBFTMessageBinaryAgreementAUX{}
}
func (m *HoneyBadgerBFTMessageBinaryAgreementAUX) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageBinaryAgreementAUX) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageBinaryAgreementAUX) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{5}
}

type HoneyBadgerBFTMessageReliableBroadcast struct {
	// Types that are valid to be assigned to Type:
	//	*HoneyBadgerBFTMessageReliableBroadcast_Val
	//	*HoneyBadgerBFTMessageReliableBroadcast_Echo
	//	*HoneyBadgerBFTMessageReliableBroadcast_Ready
	Type      isHoneyBadgerBFTMessageReliableBroadcast_Type `protobuf_oneof:"Type"`
	PadLength uint64                                        `protobuf:"varint,4,opt,name=pad_length,json=padLength" json:"pad_length,omitempty"`
	Block     []byte                                        `protobuf:"bytes,5,opt,name=block,proto3" json:"block,omitempty"`
	RootHash  []byte                                        `protobuf:"bytes,6,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	Branch    [][]byte                                      `protobuf:"bytes,7,rep,name=branch,proto3" json:"branch,omitempty"`
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) Reset() {
	*m = HoneyBadgerBFTMessageReliableBroadcast{}
}
func (m *HoneyBadgerBFTMessageReliableBroadcast) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageReliableBroadcast) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageReliableBroadcast) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{6}
}

type isHoneyBadgerBFTMessageReliableBroadcast_Type interface{ isHoneyBadgerBFTMessageReliableBroadcast_Type() }

type HoneyBadgerBFTMessageReliableBroadcast_Val struct {
	Val *HoneyBadgerBFTMessageReliableBroadcastVAL `protobuf:"bytes,1,opt,name=val,oneof"`
}
type HoneyBadgerBFTMessageReliableBroadcast_Echo struct {
	Echo *HoneyBadgerBFTMessageReliableBroadcastECHO `protobuf:"bytes,2,opt,name=echo,oneof"`
}
type HoneyBadgerBFTMessageReliableBroadcast_Ready struct {
	Ready *HoneyBadgerBFTMessageReliableBroadcastREADY `protobuf:"bytes,3,opt,name=ready,oneof"`
}

func (*HoneyBadgerBFTMessageReliableBroadcast_Val) isHoneyBadgerBFTMessageReliableBroadcast_Type()   {}
func (*HoneyBadgerBFTMessageReliableBroadcast_Echo) isHoneyBadgerBFTMessageReliableBroadcast_Type()  {}
func (*HoneyBadgerBFTMessageReliableBroadcast_Ready) isHoneyBadgerBFTMessageReliableBroadcast_Type() {}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetType() isHoneyBadgerBFTMessageReliableBroadcast_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetVal() *HoneyBadgerBFTMessageReliableBroadcastVAL {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessageReliableBroadcast_Val); ok {
		return x.Val
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetEcho() *HoneyBadgerBFTMessageReliableBroadcastECHO {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessageReliableBroadcast_Echo); ok {
		return x.Echo
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetReady() *HoneyBadgerBFTMessageReliableBroadcastREADY {
	if x, ok := m.GetType().(*HoneyBadgerBFTMessageReliableBroadcast_Ready); ok {
		return x.Ready
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetPadLength() uint64 {
	if m != nil {
		return m.PadLength
	}
	return 0
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *HoneyBadgerBFTMessageReliableBroadcast) GetBranch() [][]byte {
	if m != nil {
		return m.Branch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HoneyBadgerBFTMessageReliableBroadcast) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HoneyBadgerBFTMessageReliableBroadcast_OneofMarshaler, _HoneyBadgerBFTMessageReliableBroadcast_OneofUnmarshaler, _HoneyBadgerBFTMessageReliableBroadcast_OneofSizer, []interface{}{
		(*HoneyBadgerBFTMessageReliableBroadcast_Val)(nil),
		(*HoneyBadgerBFTMessageReliableBroadcast_Echo)(nil),
		(*HoneyBadgerBFTMessageReliableBroadcast_Ready)(nil),
	}
}

func _HoneyBadgerBFTMessageReliableBroadcast_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HoneyBadgerBFTMessageReliableBroadcast)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessageReliableBroadcast_Val:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Val); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessageReliableBroadcast_Echo:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Echo); err != nil {
			return err
		}
	case *HoneyBadgerBFTMessageReliableBroadcast_Ready:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ready); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HoneyBadgerBFTMessageReliableBroadcast.Type has unexpected type %T", x)
	}
	return nil
}

func _HoneyBadgerBFTMessageReliableBroadcast_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HoneyBadgerBFTMessageReliableBroadcast)
	switch tag {
	case 1: // Type.val
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageReliableBroadcastVAL)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessageReliableBroadcast_Val{msg}
		return true, err
	case 2: // Type.echo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageReliableBroadcastECHO)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessageReliableBroadcast_Echo{msg}
		return true, err
	case 3: // Type.ready
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HoneyBadgerBFTMessageReliableBroadcastREADY)
		err := b.DecodeMessage(msg)
		m.Type = &HoneyBadgerBFTMessageReliableBroadcast_Ready{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HoneyBadgerBFTMessageReliableBroadcast_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HoneyBadgerBFTMessageReliableBroadcast)
	// Type
	switch x := m.Type.(type) {
	case *HoneyBadgerBFTMessageReliableBroadcast_Val:
		s := proto.Size(x.Val)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessageReliableBroadcast_Echo:
		s := proto.Size(x.Echo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HoneyBadgerBFTMessageReliableBroadcast_Ready:
		s := proto.Size(x.Ready)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HoneyBadgerBFTMessageReliableBroadcastVAL struct {
}

func (m *HoneyBadgerBFTMessageReliableBroadcastVAL) Reset() {
	*m = HoneyBadgerBFTMessageReliableBroadcastVAL{}
}
func (m *HoneyBadgerBFTMessageReliableBroadcastVAL) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageReliableBroadcastVAL) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageReliableBroadcastVAL) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{7}
}

type HoneyBadgerBFTMessageReliableBroadcastECHO struct {
}

func (m *HoneyBadgerBFTMessageReliableBroadcastECHO) Reset() {
	*m = HoneyBadgerBFTMessageReliableBroadcastECHO{}
}
func (m *HoneyBadgerBFTMessageReliableBroadcastECHO) String() string {
	return proto.CompactTextString(m)
}
func (*HoneyBadgerBFTMessageReliableBroadcastECHO) ProtoMessage() {}
func (*HoneyBadgerBFTMessageReliableBroadcastECHO) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{8}
}

type HoneyBadgerBFTMessageReliableBroadcastREADY struct {
}

func (m *HoneyBadgerBFTMessageReliableBroadcastREADY) Reset() {
	*m = HoneyBadgerBFTMessageReliableBroadcastREADY{}
}
func (m *HoneyBadgerBFTMessageReliableBroadcastREADY) String() string {
	return proto.CompactTextString(m)
}
func (*HoneyBadgerBFTMessageReliableBroadcastREADY) ProtoMessage() {}
func (*HoneyBadgerBFTMessageReliableBroadcastREADY) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{9}
}

type HoneyBadgerBFTMessageThresholdEncryption struct {
	Shares []*HoneyBadgerBFTMessageThresholdEncryptionShare `protobuf:"bytes,1,rep,name=shares" json:"shares,omitempty"`
}

func (m *HoneyBadgerBFTMessageThresholdEncryption) Reset() {
	*m = HoneyBadgerBFTMessageThresholdEncryption{}
}
func (m *HoneyBadgerBFTMessageThresholdEncryption) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageThresholdEncryption) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageThresholdEncryption) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{10}
}

func (m *HoneyBadgerBFTMessageThresholdEncryption) GetShares() []*HoneyBadgerBFTMessageThresholdEncryptionShare {
	if m != nil {
		return m.Shares
	}
	return nil
}

type HoneyBadgerBFTMessageThresholdEncryptionShare struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *HoneyBadgerBFTMessageThresholdEncryptionShare) Reset() {
	*m = HoneyBadgerBFTMessageThresholdEncryptionShare{}
}
func (m *HoneyBadgerBFTMessageThresholdEncryptionShare) String() string {
	return proto.CompactTextString(m)
}
func (*HoneyBadgerBFTMessageThresholdEncryptionShare) ProtoMessage() {}
func (*HoneyBadgerBFTMessageThresholdEncryptionShare) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{11}
}

func (m *HoneyBadgerBFTMessageThresholdEncryptionShare) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type HoneyBadgerBFTMessageRequireBlock struct {
}

func (m *HoneyBadgerBFTMessageRequireBlock) Reset()         { *m = HoneyBadgerBFTMessageRequireBlock{} }
func (m *HoneyBadgerBFTMessageRequireBlock) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageRequireBlock) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageRequireBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{12}
}

type HoneyBadgerBFTMessageResponseBlock struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *HoneyBadgerBFTMessageResponseBlock) Reset()         { *m = HoneyBadgerBFTMessageResponseBlock{} }
func (m *HoneyBadgerBFTMessageResponseBlock) String() string { return proto.CompactTextString(m) }
func (*HoneyBadgerBFTMessageResponseBlock) ProtoMessage()    {}
func (*HoneyBadgerBFTMessageResponseBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{13}
}

func (m *HoneyBadgerBFTMessageResponseBlock) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*HoneyBadgerBFTMessage)(nil), "orderer.HoneyBadgerBFTMessage")
	proto.RegisterType((*HoneyBadgerBFTMessageNewHeight)(nil), "orderer.HoneyBadgerBFTMessageNewHeight")
	proto.RegisterType((*HoneyBadgerBFTMessageCommonCoin)(nil), "orderer.HoneyBadgerBFTMessageCommonCoin")
	proto.RegisterType((*HoneyBadgerBFTMessageBinaryAgreement)(nil), "orderer.HoneyBadgerBFTMessageBinaryAgreement")
	proto.RegisterType((*HoneyBadgerBFTMessageBinaryAgreementEST)(nil), "orderer.HoneyBadgerBFTMessageBinaryAgreementEST")
	proto.RegisterType((*HoneyBadgerBFTMessageBinaryAgreementAUX)(nil), "orderer.HoneyBadgerBFTMessageBinaryAgreementAUX")
	proto.RegisterType((*HoneyBadgerBFTMessageReliableBroadcast)(nil), "orderer.HoneyBadgerBFTMessageReliableBroadcast")
	proto.RegisterType((*HoneyBadgerBFTMessageReliableBroadcastVAL)(nil), "orderer.HoneyBadgerBFTMessageReliableBroadcastVAL")
	proto.RegisterType((*HoneyBadgerBFTMessageReliableBroadcastECHO)(nil), "orderer.HoneyBadgerBFTMessageReliableBroadcastECHO")
	proto.RegisterType((*HoneyBadgerBFTMessageReliableBroadcastREADY)(nil), "orderer.HoneyBadgerBFTMessageReliableBroadcastREADY")
	proto.RegisterType((*HoneyBadgerBFTMessageThresholdEncryption)(nil), "orderer.HoneyBadgerBFTMessageThresholdEncryption")
	proto.RegisterType((*HoneyBadgerBFTMessageThresholdEncryptionShare)(nil), "orderer.HoneyBadgerBFTMessageThresholdEncryptionShare")
	proto.RegisterType((*HoneyBadgerBFTMessageRequireBlock)(nil), "orderer.HoneyBadgerBFTMessageRequireBlock")
	proto.RegisterType((*HoneyBadgerBFTMessageResponseBlock)(nil), "orderer.HoneyBadgerBFTMessageResponseBlock")
}

func init() { proto.RegisterFile("honeybadgerbft.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0xeb, 0x44,
	0x10, 0xc7, 0x9d, 0x93, 0xef, 0x49, 0x0e, 0x1f, 0x4b, 0x40, 0x06, 0x04, 0x04, 0xf3, 0x71, 0x72,
	0x4e, 0xdb, 0x04, 0x52, 0xc4, 0x25, 0x52, 0xd2, 0xa6, 0x72, 0x45, 0x29, 0xea, 0x36, 0x45, 0xa5,
	0x37, 0x66, 0x6d, 0x4f, 0x63, 0x0b, 0x67, 0xd7, 0xac, 0x9d, 0xb4, 0xe1, 0x09, 0x78, 0x52, 0x5e,
	0x82, 0x1b, 0xe4, 0xb5, 0x1d, 0x52, 0xe0, 0x24, 0xf1, 0xe5, 0x7f, 0x32, 0xff, 0xdf, 0xee, 0x8e,
	0x67, 0x46, 0x81, 0x8e, 0x27, 0x38, 0xae, 0x6c, 0xe6, 0xce, 0x50, 0xda, 0xf7, 0x71, 0x3f, 0x94,
	0x22, 0x16, 0xa4, 0x2e, 0xa4, 0x8b, 0x12, 0xa5, 0xf1, 0x47, 0x0d, 0xde, 0x35, 0x93, 0x8c, 0xb1,
	0xca, 0x18, 0x9f, 0x4d, 0x7f, 0xc0, 0x28, 0x62, 0x33, 0x24, 0xef, 0x41, 0x2d, 0x42, 0xee, 0xa2,
	0xd4, 0x4b, 0xdd, 0x52, 0xaf, 0x42, 0x33, 0x45, 0x3e, 0x80, 0x86, 0x44, 0x07, 0xfd, 0x25, 0x4a,
	0xfd, 0x99, 0xfa, 0x65, 0xad, 0xc9, 0xfb, 0xd0, 0x70, 0x3c, 0xe6, 0x73, 0xcb, 0x77, 0xf5, 0x72,
	0xb7, 0xd4, 0x6b, 0xd2, 0xba, 0xd2, 0xe7, 0x6e, 0x82, 0xf3, 0xd0, 0x9f, 0x79, 0xb1, 0x5e, 0x49,
	0x71, 0xa9, 0x4a, 0x70, 0x3e, 0x8f, 0x62, 0xc6, 0x1d, 0xd4, 0xab, 0x29, 0x2e, 0xd7, 0xc4, 0x04,
	0xe0, 0xf8, 0x60, 0x99, 0xa9, 0xaf, 0xd6, 0x2d, 0xf5, 0x5a, 0xc3, 0x17, 0xfd, 0xec, 0xea, 0xfd,
	0xff, 0xbd, 0xf6, 0x25, 0x3e, 0xa4, 0xe9, 0xa6, 0x46, 0x9b, 0x3c, 0x17, 0xe4, 0x7b, 0x68, 0x39,
	0x62, 0x3e, 0x17, 0xdc, 0x72, 0x84, 0xcf, 0xf5, 0xba, 0x42, 0xf5, 0xb6, 0xa3, 0x4e, 0x94, 0xe1,
	0x44, 0xf8, 0xdc, 0xd4, 0x28, 0x38, 0x6b, 0x45, 0xee, 0xe0, 0x2d, 0xdb, 0xe7, 0x4c, 0xae, 0x2c,
	0x36, 0x93, 0x88, 0x73, 0xe4, 0xb1, 0xde, 0x50, 0xc4, 0xa3, 0xed, 0xc4, 0xb1, 0x72, 0x8d, 0x72,
	0x93, 0xa9, 0xd1, 0x37, 0xed, 0xa7, 0x21, 0xf2, 0x0b, 0x10, 0x89, 0x81, 0xcf, 0xec, 0x00, 0x2d,
	0x5b, 0x0a, 0xe6, 0x3a, 0x2c, 0x8a, 0xf5, 0xa6, 0xa2, 0x0f, 0xb6, 0xd3, 0x69, 0xe6, 0x1b, 0xe7,
	0x36, 0x53, 0xa3, 0x6f, 0xcb, 0x7f, 0x07, 0xc9, 0x3d, 0x74, 0x62, 0x4f, 0x62, 0xe4, 0x89, 0xc0,
	0xb5, 0x90, 0x3b, 0x72, 0x15, 0xc6, 0xbe, 0xe0, 0x3a, 0xa8, 0x33, 0xbe, 0xde, 0x7e, 0xc6, 0x34,
	0x77, 0x4e, 0xd6, 0x46, 0x53, 0xa3, 0xef, 0xc4, 0xff, 0x0d, 0x93, 0x2b, 0x78, 0x2e, 0xf1, 0xb7,
	0x85, 0x2f, 0xd1, 0xb2, 0x03, 0xe1, 0xfc, 0xaa, 0xb7, 0xd4, 0x01, 0xaf, 0x76, 0x3d, 0x42, 0x59,
	0xc6, 0x89, 0xc3, 0xd4, 0x68, 0x5b, 0x6e, 0x68, 0x32, 0x85, 0x37, 0x24, 0x46, 0xa1, 0xe0, 0x51,
	0xce, 0x6c, 0x2b, 0xe6, 0xc1, 0x2e, 0x66, 0xea, 0xc9, 0xa1, 0xcf, 0xe5, 0x66, 0x60, 0x5c, 0x83,
	0xca, 0x74, 0x15, 0xa2, 0xd1, 0x85, 0x8f, 0xb7, 0xb7, 0x94, 0x71, 0x05, 0x9f, 0xec, 0xe8, 0x14,
	0xd2, 0x81, 0xaa, 0x14, 0x0b, 0xee, 0x66, 0x43, 0x93, 0x0a, 0xa2, 0x43, 0x3d, 0x64, 0xab, 0x40,
	0x30, 0x57, 0x8d, 0x4c, 0x9b, 0xe6, 0xd2, 0xf8, 0xb3, 0x04, 0x9f, 0xef, 0xd3, 0x2b, 0xe4, 0x14,
	0xca, 0x18, 0xc5, 0x0a, 0xdb, 0x1a, 0x7e, 0x55, 0xa8, 0xcf, 0x26, 0xd7, 0x53, 0x53, 0xa3, 0x89,
	0x3d, 0xa1, 0xb0, 0xc5, 0xa3, 0xba, 0x44, 0x51, 0xca, 0xe8, 0xe6, 0x36, 0xa1, 0xb0, 0xc5, 0xe3,
	0x3f, 0x8f, 0x2c, 0x6f, 0x3e, 0xb2, 0x03, 0xd5, 0x25, 0x0b, 0x16, 0xa8, 0x06, 0xbc, 0x41, 0x53,
	0xb1, 0xae, 0xee, 0x4b, 0x78, 0xb1, 0xe7, 0x5d, 0xf7, 0x4d, 0x1d, 0xdd, 0xdc, 0x1a, 0x7f, 0x3d,
	0x83, 0x2f, 0xf7, 0x1b, 0x06, 0x72, 0x06, 0xe5, 0x25, 0x0b, 0xb2, 0x02, 0x0e, 0x0b, 0x8e, 0xd2,
	0x4f, 0xa3, 0x8b, 0xe4, 0xf1, 0x4b, 0x16, 0x90, 0x73, 0xa8, 0xa0, 0xe3, 0x89, 0xac, 0x86, 0xc7,
	0x05, 0x41, 0x93, 0x13, 0xf3, 0x47, 0x53, 0xa3, 0x0a, 0x41, 0x2e, 0xa0, 0x2a, 0x91, 0xb9, 0x2b,
	0x55, 0xc7, 0xd6, 0xf0, 0x9b, 0x82, 0x2c, 0x3a, 0x19, 0x9d, 0xfe, 0x6c, 0x6a, 0x34, 0x85, 0x90,
	0x8f, 0x00, 0x42, 0xe6, 0x5a, 0x01, 0xf2, 0x59, 0xec, 0x65, 0x5b, 0xb6, 0x19, 0x32, 0xf7, 0x42,
	0x05, 0x92, 0xcf, 0x93, 0xce, 0x4c, 0x55, 0x75, 0x60, 0x2a, 0xc8, 0x87, 0xd0, 0x94, 0x42, 0xc4,
	0x96, 0xc7, 0x22, 0x4f, 0x6d, 0xd8, 0x36, 0x6d, 0x24, 0x01, 0x93, 0x45, 0x5e, 0xb2, 0xb3, 0x6d,
	0xc9, 0xb8, 0xe3, 0xe9, 0xf5, 0x6e, 0xb9, 0xd7, 0xa6, 0x99, 0x5a, 0x7f, 0xd3, 0x03, 0x78, 0xb9,
	0x77, 0xf9, 0x8c, 0x43, 0x78, 0xb5, 0x7f, 0x89, 0x8c, 0x23, 0x38, 0x28, 0x50, 0x04, 0xe3, 0x77,
	0xe8, 0xed, 0xbb, 0xaf, 0xc8, 0x25, 0xd4, 0x22, 0x8f, 0x49, 0x8c, 0xf4, 0x52, 0xb7, 0xdc, 0x6b,
	0x0d, 0xbf, 0x2d, 0xbc, 0xf2, 0xae, 0x13, 0x3b, 0xcd, 0x28, 0xc6, 0x39, 0x1c, 0x15, 0x32, 0x6e,
	0x6e, 0x83, 0xd2, 0xd3, 0x6d, 0xf0, 0x19, 0x7c, 0xba, 0x73, 0x2b, 0x1a, 0xdf, 0x81, 0xb1, 0x7b,
	0xcd, 0xbd, 0xfe, 0x90, 0xf1, 0x0d, 0x7c, 0x21, 0xe4, 0xac, 0xef, 0xad, 0x42, 0x94, 0x01, 0x26,
	0x8c, 0xfe, 0x3d, 0xb3, 0xa5, 0xef, 0xa4, 0xff, 0x0d, 0xa2, 0xbc, 0x1c, 0x77, 0x87, 0x33, 0x3f,
	0xf6, 0x16, 0x76, 0xdf, 0x11, 0xf3, 0xc1, 0x46, 0xf6, 0x20, 0xcd, 0x1e, 0xa4, 0xd9, 0x83, 0x2c,
	0xdb, 0xae, 0x29, 0x7d, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x3b, 0x87, 0x7e, 0x71,
	0x08, 0x00, 0x00,
}
