package prompb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MetricMetadata_MetricType int32

const (
	MetricMetadata_UNKNOWN        MetricMetadata_MetricType = 0
	MetricMetadata_COUNTER        MetricMetadata_MetricType = 1
	MetricMetadata_GAUGE          MetricMetadata_MetricType = 2
	MetricMetadata_HISTOGRAM      MetricMetadata_MetricType = 3
	MetricMetadata_GAUGEHISTOGRAM MetricMetadata_MetricType = 4
	MetricMetadata_SUMMARY        MetricMetadata_MetricType = 5
	MetricMetadata_INFO           MetricMetadata_MetricType = 6
	MetricMetadata_STATESET       MetricMetadata_MetricType = 7
)

var MetricMetadata_MetricType_name = map[int32]string{
	0: "UNKNOWN",
	1: "COUNTER",
	2: "GAUGE",
	3: "HISTOGRAM",
	4: "GAUGEHISTOGRAM",
	5: "SUMMARY",
	6: "INFO",
	7: "STATESET",
}

var MetricMetadata_MetricType_value = map[string]int32{
	"UNKNOWN":        0,
	"COUNTER":        1,
	"GAUGE":          2,
	"HISTOGRAM":      3,
	"GAUGEHISTOGRAM": 4,
	"SUMMARY":        5,
	"INFO":           6,
	"STATESET":       7,
}

func (x MetricMetadata_MetricType) String() string {
	return proto.EnumName(MetricMetadata_MetricType_name, int32(x))
}

func (MetricMetadata_MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0, 0}
}

type LabelMatcher_Type int32

const (
	LabelMatcher_EQ  LabelMatcher_Type = 0
	LabelMatcher_NEQ LabelMatcher_Type = 1
	LabelMatcher_RE  LabelMatcher_Type = 2
	LabelMatcher_NRE LabelMatcher_Type = 3
)

var LabelMatcher_Type_name = map[int32]string{
	0: "EQ",
	1: "NEQ",
	2: "RE",
	3: "NRE",
}

var LabelMatcher_Type_value = map[string]int32{
	"EQ":  0,
	"NEQ": 1,
	"RE":  2,
	"NRE": 3,
}

func (x LabelMatcher_Type) String() string {
	return proto.EnumName(LabelMatcher_Type_name, int32(x))
}

func (LabelMatcher_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6, 0}
}

// We require this to match chunkenc.Encoding.
type Chunk_Encoding int32

const (
	Chunk_UNKNOWN Chunk_Encoding = 0
	Chunk_XOR     Chunk_Encoding = 1
)

var Chunk_Encoding_name = map[int32]string{
	0: "UNKNOWN",
	1: "XOR",
}

var Chunk_Encoding_value = map[string]int32{
	"UNKNOWN": 0,
	"XOR":     1,
}

func (x Chunk_Encoding) String() string {
	return proto.EnumName(Chunk_Encoding_name, int32(x))
}

func (Chunk_Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{8, 0}
}

type MetricMetadata struct {
	// Represents the metric type, these match the set from Prometheus.
	// Refer to pkg/textparse/interface.go for details.
	Type                 MetricMetadata_MetricType `protobuf:"varint,1,opt,name=type,proto3,enum=prometheus.MetricMetadata_MetricType" json:"type,omitempty"`
	MetricFamilyName     string                    `protobuf:"bytes,2,opt,name=metric_family_name,json=metricFamilyName,proto3" json:"metric_family_name,omitempty"`
	Help                 string                    `protobuf:"bytes,4,opt,name=help,proto3" json:"help,omitempty"`
	Unit                 string                    `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MetricMetadata) Reset()         { *m = MetricMetadata{} }
func (m *MetricMetadata) String() string { return proto.CompactTextString(m) }
func (*MetricMetadata) ProtoMessage()    {}
func (*MetricMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}
func (m *MetricMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetricMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetricMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricMetadata.Merge(m, src)
}
func (m *MetricMetadata) XXX_Size() int {
	return m.Size()
}
func (m *MetricMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MetricMetadata proto.InternalMessageInfo

func (m *MetricMetadata) GetType() MetricMetadata_MetricType {
	if m != nil {
		return m.Type
	}
	return MetricMetadata_UNKNOWN
}

func (m *MetricMetadata) GetMetricFamilyName() string {
	if m != nil {
		return m.MetricFamilyName
	}
	return ""
}

func (m *MetricMetadata) GetHelp() string {
	if m != nil {
		return m.Help
	}
	return ""
}

func (m *MetricMetadata) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type Sample struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp is in ms format, see pkg/timestamp/timestamp.go for
	// conversion from time.Time to Prometheus timestamp.
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return m.Size()
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sample) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Exemplar struct {
	// Optional, can be empty.
	Labels []Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels"`
	Value  float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	// timestamp is in ms format, see pkg/timestamp/timestamp.go for
	// conversion from time.Time to Prometheus timestamp.
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Exemplar) Reset()         { *m = Exemplar{} }
func (m *Exemplar) String() string { return proto.CompactTextString(m) }
func (*Exemplar) ProtoMessage()    {}
func (*Exemplar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}
func (m *Exemplar) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Exemplar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Exemplar.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Exemplar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exemplar.Merge(m, src)
}
func (m *Exemplar) XXX_Size() int {
	return m.Size()
}
func (m *Exemplar) XXX_DiscardUnknown() {
	xxx_messageInfo_Exemplar.DiscardUnknown(m)
}

var xxx_messageInfo_Exemplar proto.InternalMessageInfo

func (m *Exemplar) GetLabels() []Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Exemplar) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Exemplar) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// TimeSeries represents samples and labels for a single time series.
type TimeSeries struct {
	// For a timeseries to be valid, and for the samples and exemplars
	// to be ingested by the remote system properly, the labels field is required.
	Labels               []Label    `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels"`
	Samples              []Sample   `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples"`
	Exemplars            []Exemplar `protobuf:"bytes,3,rep,name=exemplars,proto3" json:"exemplars"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}
func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return m.Size()
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetLabels() []Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

func (m *TimeSeries) GetExemplars() []Exemplar {
	if m != nil {
		return m.Exemplars
	}
	return nil
}

type Label struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}
func (m *Label) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Label.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return m.Size()
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Labels struct {
	Labels               []Label  `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Labels) Reset()         { *m = Labels{} }
func (m *Labels) String() string { return proto.CompactTextString(m) }
func (*Labels) ProtoMessage()    {}
func (*Labels) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}
func (m *Labels) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Labels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Labels.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Labels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Labels.Merge(m, src)
}
func (m *Labels) XXX_Size() int {
	return m.Size()
}
func (m *Labels) XXX_DiscardUnknown() {
	xxx_messageInfo_Labels.DiscardUnknown(m)
}

var xxx_messageInfo_Labels proto.InternalMessageInfo

func (m *Labels) GetLabels() []Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Matcher specifies a rule, which can match or set of labels or not.
type LabelMatcher struct {
	Type                 LabelMatcher_Type `protobuf:"varint,1,opt,name=type,proto3,enum=prometheus.LabelMatcher_Type" json:"type,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LabelMatcher) Reset()         { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()    {}
func (*LabelMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6}
}
func (m *LabelMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelMatcher.Merge(m, src)
}
func (m *LabelMatcher) XXX_Size() int {
	return m.Size()
}
func (m *LabelMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_LabelMatcher proto.InternalMessageInfo

func (m *LabelMatcher) GetType() LabelMatcher_Type {
	if m != nil {
		return m.Type
	}
	return LabelMatcher_EQ
}

func (m *LabelMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReadHints struct {
	StepMs               int64    `protobuf:"varint,1,opt,name=step_ms,json=stepMs,proto3" json:"step_ms,omitempty"`
	Func                 string   `protobuf:"bytes,2,opt,name=func,proto3" json:"func,omitempty"`
	StartMs              int64    `protobuf:"varint,3,opt,name=start_ms,json=startMs,proto3" json:"start_ms,omitempty"`
	EndMs                int64    `protobuf:"varint,4,opt,name=end_ms,json=endMs,proto3" json:"end_ms,omitempty"`
	Grouping             []string `protobuf:"bytes,5,rep,name=grouping,proto3" json:"grouping,omitempty"`
	By                   bool     `protobuf:"varint,6,opt,name=by,proto3" json:"by,omitempty"`
	RangeMs              int64    `protobuf:"varint,7,opt,name=range_ms,json=rangeMs,proto3" json:"range_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadHints) Reset()         { *m = ReadHints{} }
func (m *ReadHints) String() string { return proto.CompactTextString(m) }
func (*ReadHints) ProtoMessage()    {}
func (*ReadHints) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{7}
}
func (m *ReadHints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadHints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadHints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadHints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadHints.Merge(m, src)
}
func (m *ReadHints) XXX_Size() int {
	return m.Size()
}
func (m *ReadHints) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadHints.DiscardUnknown(m)
}

var xxx_messageInfo_ReadHints proto.InternalMessageInfo

func (m *ReadHints) GetStepMs() int64 {
	if m != nil {
		return m.StepMs
	}
	return 0
}

func (m *ReadHints) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *ReadHints) GetStartMs() int64 {
	if m != nil {
		return m.StartMs
	}
	return 0
}

func (m *ReadHints) GetEndMs() int64 {
	if m != nil {
		return m.EndMs
	}
	return 0
}

func (m *ReadHints) GetGrouping() []string {
	if m != nil {
		return m.Grouping
	}
	return nil
}

func (m *ReadHints) GetBy() bool {
	if m != nil {
		return m.By
	}
	return false
}

func (m *ReadHints) GetRangeMs() int64 {
	if m != nil {
		return m.RangeMs
	}
	return 0
}

// Chunk represents a TSDB chunk.
// Time range [min, max] is inclusive.
type Chunk struct {
	MinTimeMs            int64          `protobuf:"varint,1,opt,name=min_time_ms,json=minTimeMs,proto3" json:"min_time_ms,omitempty"`
	MaxTimeMs            int64          `protobuf:"varint,2,opt,name=max_time_ms,json=maxTimeMs,proto3" json:"max_time_ms,omitempty"`
	Type                 Chunk_Encoding `protobuf:"varint,3,opt,name=type,proto3,enum=prometheus.Chunk_Encoding" json:"type,omitempty"`
	Data                 []byte         `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{8}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return m.Size()
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetMinTimeMs() int64 {
	if m != nil {
		return m.MinTimeMs
	}
	return 0
}

func (m *Chunk) GetMaxTimeMs() int64 {
	if m != nil {
		return m.MaxTimeMs
	}
	return 0
}

func (m *Chunk) GetType() Chunk_Encoding {
	if m != nil {
		return m.Type
	}
	return Chunk_UNKNOWN
}

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ChunkedSeries represents single, encoded time series.
type ChunkedSeries struct {
	// Labels should be sorted.
	Labels []Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels"`
	// Chunks will be in start time order and may overlap.
	Chunks               []Chunk  `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkedSeries) Reset()         { *m = ChunkedSeries{} }
func (m *ChunkedSeries) String() string { return proto.CompactTextString(m) }
func (*ChunkedSeries) ProtoMessage()    {}
func (*ChunkedSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{9}
}
func (m *ChunkedSeries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChunkedSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChunkedSeries.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChunkedSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkedSeries.Merge(m, src)
}
func (m *ChunkedSeries) XXX_Size() int {
	return m.Size()
}
func (m *ChunkedSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkedSeries.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkedSeries proto.InternalMessageInfo

func (m *ChunkedSeries) GetLabels() []Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ChunkedSeries) GetChunks() []Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
	proto.RegisterEnum("prometheus.MetricMetadata_MetricType", MetricMetadata_MetricType_name, MetricMetadata_MetricType_value)
	proto.RegisterEnum("prometheus.LabelMatcher_Type", LabelMatcher_Type_name, LabelMatcher_Type_value)
	proto.RegisterEnum("prometheus.Chunk_Encoding", Chunk_Encoding_name, Chunk_Encoding_value)
	proto.RegisterType((*MetricMetadata)(nil), "prometheus.MetricMetadata")
	proto.RegisterType((*Sample)(nil), "prometheus.Sample")
	proto.RegisterType((*Exemplar)(nil), "prometheus.Exemplar")
	proto.RegisterType((*TimeSeries)(nil), "prometheus.TimeSeries")
	proto.RegisterType((*Label)(nil), "prometheus.Label")
	proto.RegisterType((*Labels)(nil), "prometheus.Labels")
	proto.RegisterType((*LabelMatcher)(nil), "prometheus.LabelMatcher")
	proto.RegisterType((*ReadHints)(nil), "prometheus.ReadHints")
	proto.RegisterType((*Chunk)(nil), "prometheus.Chunk")
	proto.RegisterType((*ChunkedSeries)(nil), "prometheus.ChunkedSeries")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0xf5, 0xf0, 0x29, 0x5e, 0xd9, 0x02, 0x3d, 0x50, 0x51, 0xd6, 0x68, 0x55, 0x81, 0x40, 0x01,
	0x2d, 0x0a, 0x19, 0x76, 0x37, 0x35, 0xd0, 0x8d, 0x6c, 0xd0, 0x0f, 0xd4, 0x94, 0xe0, 0x91, 0x84,
	0x3e, 0x36, 0xc2, 0x48, 0x1a, 0x4b, 0x44, 0xc4, 0x47, 0x38, 0x54, 0x60, 0x7d, 0x48, 0x76, 0xf9,
	0x83, 0x20, 0x8b, 0xfc, 0x85, 0x97, 0xf9, 0x82, 0x20, 0xf0, 0x97, 0x04, 0x33, 0xa4, 0x4c, 0x29,
	0x4e, 0x16, 0xce, 0xee, 0xde, 0x7b, 0xce, 0xb9, 0x8f, 0xb9, 0x97, 0x84, 0x6a, 0xb6, 0x4a, 0x18,
	0x6f, 0x27, 0x69, 0x9c, 0xc5, 0x18, 0x92, 0x34, 0x0e, 0x59, 0x36, 0x67, 0x4b, 0x7e, 0x50, 0x9f,
	0xc5, 0xb3, 0x58, 0x86, 0x0f, 0x85, 0x95, 0x33, 0xdc, 0x37, 0x0a, 0xd4, 0x7c, 0x96, 0xa5, 0xc1,
	0xc4, 0x67, 0x19, 0x9d, 0xd2, 0x8c, 0xe2, 0x13, 0xd0, 0x44, 0x0e, 0x07, 0x35, 0x51, 0xab, 0x76,
	0xfc, 0x5b, 0xbb, 0xcc, 0xd1, 0xde, 0x66, 0x16, 0xee, 0x60, 0x95, 0x30, 0x22, 0x25, 0xf8, 0x77,
	0xc0, 0xa1, 0x8c, 0x8d, 0x6e, 0x69, 0x18, 0x2c, 0x56, 0xa3, 0x88, 0x86, 0xcc, 0x51, 0x9a, 0xa8,
	0x65, 0x11, 0x3b, 0x47, 0xce, 0x25, 0xd0, 0xa5, 0x21, 0xc3, 0x18, 0xb4, 0x39, 0x5b, 0x24, 0x8e,
	0x26, 0x71, 0x69, 0x8b, 0xd8, 0x32, 0x0a, 0x32, 0x47, 0xcf, 0x63, 0xc2, 0x76, 0x57, 0x00, 0x65,
	0x25, 0x5c, 0x05, 0x73, 0xd8, 0xfd, 0xbb, 0xdb, 0xfb, 0xa7, 0x6b, 0xef, 0x08, 0xe7, 0xac, 0x37,
	0xec, 0x0e, 0x3c, 0x62, 0x23, 0x6c, 0x81, 0x7e, 0xd1, 0x19, 0x5e, 0x78, 0xb6, 0x82, 0xf7, 0xc0,
	0xba, 0xbc, 0xea, 0x0f, 0x7a, 0x17, 0xa4, 0xe3, 0xdb, 0x2a, 0xc6, 0x50, 0x93, 0x48, 0x19, 0xd3,
	0x84, 0xb4, 0x3f, 0xf4, 0xfd, 0x0e, 0xf9, 0xcf, 0xd6, 0x71, 0x05, 0xb4, 0xab, 0xee, 0x79, 0xcf,
	0x36, 0xf0, 0x2e, 0x54, 0xfa, 0x83, 0xce, 0xc0, 0xeb, 0x7b, 0x03, 0xdb, 0x74, 0xff, 0x02, 0xa3,
	0x4f, 0xc3, 0x64, 0xc1, 0x70, 0x1d, 0xf4, 0x57, 0x74, 0xb1, 0xcc, 0x9f, 0x05, 0x91, 0xdc, 0xc1,
	0x3f, 0x83, 0x95, 0x05, 0x21, 0xe3, 0x19, 0x0d, 0x13, 0x39, 0xa7, 0x4a, 0xca, 0x80, 0x1b, 0x43,
	0xc5, 0xbb, 0x63, 0x61, 0xb2, 0xa0, 0x29, 0x3e, 0x04, 0x63, 0x41, 0xc7, 0x6c, 0xc1, 0x1d, 0xd4,
	0x54, 0x5b, 0xd5, 0xe3, 0xfd, 0xcd, 0x77, 0xbd, 0x16, 0xc8, 0xa9, 0x76, 0xff, 0xf1, 0xd7, 0x1d,
	0x52, 0xd0, 0xca, 0x82, 0xca, 0x37, 0x0b, 0xaa, 0x5f, 0x16, 0x7c, 0x8b, 0x00, 0x06, 0x41, 0xc8,
	0xfa, 0x2c, 0x0d, 0x18, 0x7f, 0x7e, 0xcd, 0x63, 0x30, 0xb9, 0x1c, 0x97, 0x3b, 0x8a, 0x54, 0xe0,
	0x4d, 0x45, 0xfe, 0x12, 0x85, 0x64, 0x4d, 0xc4, 0x7f, 0x82, 0xc5, 0x8a, 0x21, 0xb9, 0xa3, 0x4a,
	0x55, 0x7d, 0x53, 0xb5, 0x7e, 0x81, 0x42, 0x57, 0x92, 0xdd, 0x23, 0xd0, 0x65, 0x13, 0x62, 0xe9,
	0xf2, 0x50, 0x50, 0xbe, 0x74, 0x61, 0x6f, 0x8f, 0x6f, 0x15, 0xe3, 0xbb, 0x27, 0x60, 0x5c, 0xe7,
	0xad, 0x3e, 0x77, 0x36, 0xf7, 0x35, 0x82, 0x5d, 0x19, 0xf7, 0x69, 0x36, 0x99, 0xb3, 0x14, 0x1f,
	0x6d, 0xdd, 0xf9, 0x2f, 0x4f, 0xf4, 0x05, 0xaf, 0xbd, 0x71, 0xdf, 0xeb, 0x46, 0x95, 0xaf, 0x35,
	0xaa, 0x6e, 0x36, 0xda, 0x02, 0x4d, 0x5e, 0xab, 0x01, 0x8a, 0x77, 0x63, 0xef, 0x60, 0x13, 0xd4,
	0xae, 0x77, 0x63, 0x23, 0x11, 0x20, 0xe2, 0x42, 0x45, 0x80, 0x78, 0xb6, 0xea, 0xbe, 0x47, 0x60,
	0x11, 0x46, 0xa7, 0x97, 0x41, 0x94, 0x71, 0xfc, 0x23, 0x98, 0x3c, 0x63, 0xc9, 0x28, 0xe4, 0xb2,
	0x2f, 0x95, 0x18, 0xc2, 0xf5, 0xb9, 0x28, 0x7d, 0xbb, 0x8c, 0x26, 0xeb, 0xd2, 0xc2, 0xc6, 0x3f,
	0x41, 0x85, 0x67, 0x34, 0xcd, 0x04, 0x3b, 0xbf, 0x05, 0x53, 0xfa, 0x3e, 0xc7, 0x3f, 0x80, 0xc1,
	0xa2, 0xa9, 0x00, 0x34, 0x09, 0xe8, 0x2c, 0x9a, 0xfa, 0x1c, 0x1f, 0x40, 0x65, 0x96, 0xc6, 0xcb,
	0x24, 0x88, 0x66, 0x8e, 0xde, 0x54, 0x5b, 0x16, 0x79, 0xf4, 0x71, 0x0d, 0x94, 0xf1, 0xca, 0x31,
	0x9a, 0xa8, 0x55, 0x21, 0xca, 0x78, 0x25, 0xb2, 0xa7, 0x34, 0x9a, 0x31, 0x91, 0xc4, 0xcc, 0xb3,
	0x4b, 0xdf, 0xe7, 0xee, 0x3b, 0x04, 0xfa, 0xd9, 0x7c, 0x19, 0xbd, 0xc0, 0x0d, 0xa8, 0x86, 0x41,
	0x34, 0x12, 0x27, 0x58, 0xf6, 0x6c, 0x85, 0x41, 0x24, 0xce, 0xd0, 0xe7, 0x12, 0xa7, 0x77, 0x8f,
	0x78, 0xf1, 0x89, 0x84, 0xf4, 0xae, 0xc0, 0xdb, 0xc5, 0x12, 0x54, 0xb9, 0x84, 0x83, 0xcd, 0x25,
	0xc8, 0x02, 0x6d, 0x2f, 0x9a, 0xc4, 0xd3, 0x20, 0x9a, 0x95, 0x1b, 0x10, 0xbf, 0x1e, 0x39, 0xd5,
	0x2e, 0x91, 0xb6, 0xdb, 0x84, 0xca, 0x9a, 0xb5, 0xfd, 0x77, 0x30, 0x41, 0xfd, 0xb7, 0x47, 0x6c,
	0xe4, 0xbe, 0x84, 0x3d, 0x99, 0x8d, 0x4d, 0xbf, 0xf7, 0xcb, 0x38, 0x04, 0x63, 0x22, 0x32, 0xac,
	0x3f, 0x8c, 0xfd, 0x27, 0x9d, 0xae, 0x05, 0x39, 0xed, 0xb4, 0x7e, 0xff, 0xd0, 0x40, 0x1f, 0x1e,
	0x1a, 0xe8, 0xd3, 0x43, 0x03, 0xfd, 0x6f, 0x08, 0x76, 0x32, 0x1e, 0x1b, 0xf2, 0xaf, 0xfb, 0xc7,
	0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa3, 0x6c, 0x23, 0xa6, 0x05, 0x00, 0x00,
}

func (m *MetricMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Unit) > 0 {
		i -= len(m.Unit)
		copy(dAtA[i:], m.Unit)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Unit)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Help) > 0 {
		i -= len(m.Help)
		copy(dAtA[i:], m.Help)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Help)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MetricFamilyName) > 0 {
		i -= len(m.MetricFamilyName)
		copy(dAtA[i:], m.MetricFamilyName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MetricFamilyName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Sample) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sample) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sample) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *Exemplar) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exemplar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Exemplar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TimeSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeSeries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeSeries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Exemplars) > 0 {
		for iNdEx := len(m.Exemplars) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Exemplars[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Samples) > 0 {
		for iNdEx := len(m.Samples) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Samples[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Label) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Label) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Label) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Labels) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Labels) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Labels) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LabelMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReadHints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadHints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadHints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RangeMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RangeMs))
		i--
		dAtA[i] = 0x38
	}
	if m.By {
		i--
		if m.By {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Grouping) > 0 {
		for iNdEx := len(m.Grouping) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Grouping[iNdEx])
			copy(dAtA[i:], m.Grouping[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Grouping[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.EndMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EndMs))
		i--
		dAtA[i] = 0x20
	}
	if m.StartMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.StartMs))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Func) > 0 {
		i -= len(m.Func)
		copy(dAtA[i:], m.Func)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Func)))
		i--
		dAtA[i] = 0x12
	}
	if m.StepMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.StepMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxTimeMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxTimeMs))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTimeMs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinTimeMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChunkedSeries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunkedSeries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChunkedSeries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chunks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetricMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.MetricFamilyName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Help)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Sample) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sovTypes(uint64(m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Exemplar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sovTypes(uint64(m.Timestamp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TimeSeries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Exemplars) > 0 {
		for _, e := range m.Exemplars {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Label) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Labels) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LabelMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReadHints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StepMs != 0 {
		n += 1 + sovTypes(uint64(m.StepMs))
	}
	l = len(m.Func)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartMs != 0 {
		n += 1 + sovTypes(uint64(m.StartMs))
	}
	if m.EndMs != 0 {
		n += 1 + sovTypes(uint64(m.EndMs))
	}
	if len(m.Grouping) > 0 {
		for _, s := range m.Grouping {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.By {
		n += 2
	}
	if m.RangeMs != 0 {
		n += 1 + sovTypes(uint64(m.RangeMs))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTimeMs != 0 {
		n += 1 + sovTypes(uint64(m.MinTimeMs))
	}
	if m.MaxTimeMs != 0 {
		n += 1 + sovTypes(uint64(m.MaxTimeMs))
	}
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChunkedSeries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MetricMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MetricMetadata_MetricType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricFamilyName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricFamilyName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Help", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Help = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Sample) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Sample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Exemplar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Exemplar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Exemplar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, Label{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *TimeSeries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TimeSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, Label{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Samples = append(m.Samples, Sample{})
			if err := m.Samples[len(m.Samples)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exemplars", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Exemplars = append(m.Exemplars, Exemplar{})
			if err := m.Exemplars[len(m.Exemplars)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Label) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Label: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Label: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Labels) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Labels: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Labels: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, Label{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LabelMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LabelMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= LabelMatcher_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ReadHints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ReadHints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadHints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StepMs", wireType)
			}
			m.StepMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StepMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Func", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Func = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartMs", wireType)
			}
			m.StartMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndMs", wireType)
			}
			m.EndMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grouping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grouping = append(m.Grouping, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field By", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.By = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeMs", wireType)
			}
			m.RangeMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTimeMs", wireType)
			}
			m.MinTimeMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTimeMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTimeMs", wireType)
			}
			m.MaxTimeMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTimeMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Chunk_Encoding(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ChunkedSeries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ChunkedSeries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChunkedSeries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, Label{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, Chunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
