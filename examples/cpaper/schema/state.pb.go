// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommercialPaper_State int32

const (
	CommercialPaper_ISSUED   CommercialPaper_State = 0
	CommercialPaper_TRADING  CommercialPaper_State = 1
	CommercialPaper_REDEEMED CommercialPaper_State = 2
)

var CommercialPaper_State_name = map[int32]string{
	0: "ISSUED",
	1: "TRADING",
	2: "REDEEMED",
}
var CommercialPaper_State_value = map[string]int32{
	"ISSUED":   0,
	"TRADING":  1,
	"REDEEMED": 2,
}

func (x CommercialPaper_State) String() string {
	return proto.EnumName(CommercialPaper_State_name, int32(x))
}
func (CommercialPaper_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type CommercialPaper struct {
	Issuer       string                     `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	PaperNumber  string                     `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber" json:"paper_number,omitempty"`
	Owner        string                     `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	IssueDate    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_date,json=issueDate" json:"issue_date,omitempty"`
	MaturityDate *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=maturity_date,json=maturityDate" json:"maturity_date,omitempty"`
	FaceValue    int32                      `protobuf:"varint,6,opt,name=face_value,json=faceValue" json:"face_value,omitempty"`
	State        CommercialPaper_State      `protobuf:"varint,7,opt,name=state,enum=schema.CommercialPaper_State" json:"state,omitempty"`
}

func (m *CommercialPaper) Reset()                    { *m = CommercialPaper{} }
func (m *CommercialPaper) String() string            { return proto.CompactTextString(m) }
func (*CommercialPaper) ProtoMessage()               {}
func (*CommercialPaper) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommercialPaper) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *CommercialPaper) GetPaperNumber() string {
	if m != nil {
		return m.PaperNumber
	}
	return ""
}

func (m *CommercialPaper) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CommercialPaper) GetIssueDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueDate
	}
	return nil
}

func (m *CommercialPaper) GetMaturityDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.MaturityDate
	}
	return nil
}

func (m *CommercialPaper) GetFaceValue() int32 {
	if m != nil {
		return m.FaceValue
	}
	return 0
}

func (m *CommercialPaper) GetState() CommercialPaper_State {
	if m != nil {
		return m.State
	}
	return CommercialPaper_ISSUED
}

func init() {
	proto.RegisterType((*CommercialPaper)(nil), "schema.CommercialPaper")
	proto.RegisterEnum("schema.CommercialPaper_State", CommercialPaper_State_name, CommercialPaper_State_value)
}

func init() { proto.RegisterFile("state.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xcd, 0xb4, 0x9d, 0xfb, 0x3a, 0x75, 0x04, 0x91, 0x30, 0x18, 0xd6, 0x9d, 0x7a, 0xca,
	0x60, 0x3b, 0x79, 0x12, 0xb1, 0x45, 0x76, 0x70, 0x48, 0x36, 0xbd, 0x8e, 0x6c, 0x7e, 0x9b, 0x85,
	0x66, 0x29, 0x69, 0xaa, 0xf8, 0xef, 0xfa, 0x97, 0x48, 0x13, 0x7b, 0xf1, 0xe2, 0xf1, 0xbd, 0xef,
	0xf7, 0x3e, 0x1e, 0x0f, 0xa2, 0xca, 0x4a, 0x8b, 0xbc, 0x34, 0xda, 0x6a, 0x1a, 0x56, 0xdb, 0x77,
	0x54, 0x72, 0x78, 0xbd, 0xd7, 0x7a, 0x5f, 0xe0, 0xc4, 0xb9, 0x9b, 0x7a, 0x37, 0xb1, 0xb9, 0xc2,
	0xca, 0x4a, 0x55, 0x7a, 0x70, 0xfc, 0xdd, 0x81, 0x8b, 0x07, 0xad, 0x14, 0x9a, 0x6d, 0x2e, 0x8b,
	0x67, 0x59, 0xa2, 0xa1, 0x57, 0x10, 0xe6, 0x55, 0x55, 0xa3, 0x61, 0x24, 0x26, 0x49, 0x4f, 0xfc,
	0x2a, 0x7a, 0x03, 0xfd, 0xb2, 0x01, 0xd6, 0x87, 0x5a, 0x6d, 0xd0, 0xb0, 0x8e, 0xbb, 0x46, 0xce,
	0x5b, 0x38, 0x8b, 0x5e, 0x42, 0xa0, 0x3f, 0x0f, 0x68, 0xd8, 0xb1, 0xbb, 0x79, 0x41, 0x6f, 0x01,
	0xdc, 0x8b, 0xf5, 0x9b, 0xb4, 0xc8, 0x4e, 0x62, 0x92, 0x44, 0xd3, 0x21, 0xf7, 0xd5, 0x78, 0x5b,
	0x8d, 0xaf, 0xda, 0x6a, 0xa2, 0xe7, 0xe8, 0x54, 0x5a, 0xa4, 0x77, 0x70, 0xa6, 0xa4, 0xad, 0x4d,
	0x6e, 0xbf, 0x7c, 0x3a, 0xf8, 0x37, 0xdd, 0x6f, 0x03, 0xee, 0xc1, 0x08, 0x60, 0x27, 0xb7, 0xb8,
	0xfe, 0x90, 0x45, 0x8d, 0x2c, 0x8c, 0x49, 0x12, 0x88, 0x5e, 0xe3, 0xbc, 0x36, 0x06, 0x9d, 0x41,
	0xe0, 0x76, 0x63, 0xdd, 0x98, 0x24, 0xe7, 0xd3, 0x11, 0xf7, 0xc3, 0xf1, 0x3f, 0x9b, 0xf0, 0x65,
	0x03, 0x09, 0xcf, 0x8e, 0x39, 0x04, 0x4e, 0x53, 0x80, 0x70, 0xbe, 0x5c, 0xbe, 0x64, 0xe9, 0xe0,
	0x88, 0x46, 0xd0, 0x5d, 0x89, 0xfb, 0x74, 0xbe, 0x78, 0x1c, 0x10, 0xda, 0x87, 0x53, 0x91, 0xa5,
	0x59, 0xf6, 0x94, 0xa5, 0x83, 0xce, 0x26, 0x74, 0x2d, 0x67, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x69, 0xd2, 0x59, 0xb0, 0xa3, 0x01, 0x00, 0x00,
}