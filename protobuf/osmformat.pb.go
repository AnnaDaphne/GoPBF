// Code generated by protoc-gen-go.
// source: osmformat.proto
// DO NOT EDIT!

package OSMPBF

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Relation_MemberType int32

const (
    Relation_NODE     Relation_MemberType = 0
    Relation_WAY      Relation_MemberType = 1
    Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
    0:  "NODE",
    1:  "WAY",
    2:  "RELATION",
}
var Relation_MemberType_value = map[string]int32{
    "NODE":     0,
    "WAY":      1,
    "RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
    p := new(Relation_MemberType)
    *p = x
    return p
}
func (x Relation_MemberType) String() string {
    return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
    value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
    if err != nil {
        return err
    }
    *x = Relation_MemberType(value)
    return nil
}

type HeaderBlock struct {
    Bbox *HeaderBBox `protobuf:"bytes,1,opt,name=bbox" json:"bbox,omitempty"`
    // Additional tags to aid in parsing this dataset
    RequiredFeatures []string `protobuf:"bytes,4,rep,name=required_features" json:"required_features,omitempty"`
    OptionalFeatures []string `protobuf:"bytes,5,rep,name=optional_features" json:"optional_features,omitempty"`
    Writingprogram   *string  `protobuf:"bytes,16,opt,name=writingprogram" json:"writingprogram,omitempty"`
    Source           *string  `protobuf:"bytes,17,opt,name=source" json:"source,omitempty"`
    // replication timestamp, expressed in seconds since the epoch,
    // otherwise the same value as in the "timestamp=..." field
    // in the state.txt file used by Osmosis
    OsmosisReplicationTimestamp *int64 `protobuf:"varint,32,opt,name=osmosis_replication_timestamp" json:"osmosis_replication_timestamp,omitempty"`
    // replication sequence number (sequenceNumber in state.txt)
    OsmosisReplicationSequenceNumber *int64 `protobuf:"varint,33,opt,name=osmosis_replication_sequence_number" json:"osmosis_replication_sequence_number,omitempty"`
    // replication base URL (from Osmosis' configuration.txt file)
    OsmosisReplicationBaseUrl *string `protobuf:"bytes,34,opt,name=osmosis_replication_base_url" json:"osmosis_replication_base_url,omitempty"`
    XXX_unrecognized          []byte  `json:"-"`
}

func (m *HeaderBlock) Reset()         { *m = HeaderBlock{} }
func (m *HeaderBlock) String() string { return proto.CompactTextString(m) }
func (*HeaderBlock) ProtoMessage()    {}

func (m *HeaderBlock) GetBbox() *HeaderBBox {
    if m != nil {
        return m.Bbox
    }
    return nil
}

func (m *HeaderBlock) GetRequiredFeatures() []string {
    if m != nil {
        return m.RequiredFeatures
    }
    return nil
}

func (m *HeaderBlock) GetOptionalFeatures() []string {
    if m != nil {
        return m.OptionalFeatures
    }
    return nil
}

func (m *HeaderBlock) GetWritingprogram() string {
    if m != nil && m.Writingprogram != nil {
        return *m.Writingprogram
    }
    return ""
}

func (m *HeaderBlock) GetSource() string {
    if m != nil && m.Source != nil {
        return *m.Source
    }
    return ""
}

func (m *HeaderBlock) GetOsmosisReplicationTimestamp() int64 {
    if m != nil && m.OsmosisReplicationTimestamp != nil {
        return *m.OsmosisReplicationTimestamp
    }
    return 0
}

func (m *HeaderBlock) GetOsmosisReplicationSequenceNumber() int64 {
    if m != nil && m.OsmosisReplicationSequenceNumber != nil {
        return *m.OsmosisReplicationSequenceNumber
    }
    return 0
}

func (m *HeaderBlock) GetOsmosisReplicationBaseUrl() string {
    if m != nil && m.OsmosisReplicationBaseUrl != nil {
        return *m.OsmosisReplicationBaseUrl
    }
    return ""
}

type HeaderBBox struct {
    Left             *int64 `protobuf:"zigzag64,1,req,name=left" json:"left,omitempty"`
    Right            *int64 `protobuf:"zigzag64,2,req,name=right" json:"right,omitempty"`
    Top              *int64 `protobuf:"zigzag64,3,req,name=top" json:"top,omitempty"`
    Bottom           *int64 `protobuf:"zigzag64,4,req,name=bottom" json:"bottom,omitempty"`
    XXX_unrecognized []byte `json:"-"`
}

func (m *HeaderBBox) Reset()         { *m = HeaderBBox{} }
func (m *HeaderBBox) String() string { return proto.CompactTextString(m) }
func (*HeaderBBox) ProtoMessage()    {}

func (m *HeaderBBox) GetLeft() int64 {
    if m != nil && m.Left != nil {
        return *m.Left
    }
    return 0
}

func (m *HeaderBBox) GetRight() int64 {
    if m != nil && m.Right != nil {
        return *m.Right
    }
    return 0
}

func (m *HeaderBBox) GetTop() int64 {
    if m != nil && m.Top != nil {
        return *m.Top
    }
    return 0
}

func (m *HeaderBBox) GetBottom() int64 {
    if m != nil && m.Bottom != nil {
        return *m.Bottom
    }
    return 0
}

type PrimitiveBlock struct {
    Stringtable    *StringTable      `protobuf:"bytes,1,req,name=stringtable" json:"stringtable,omitempty"`
    Primitivegroup []*PrimitiveGroup `protobuf:"bytes,2,rep,name=primitivegroup" json:"primitivegroup,omitempty"`
    // Granularity, units of nanodegrees, used to store coordinates in this block
    Granularity *int32 `protobuf:"varint,17,opt,name=granularity,def=100" json:"granularity,omitempty"`
    // Offset value between the output coordinates coordinates and the granularity grid in unites of nanodegrees.
    LatOffset *int64 `protobuf:"varint,19,opt,name=lat_offset,def=0" json:"lat_offset,omitempty"`
    LonOffset *int64 `protobuf:"varint,20,opt,name=lon_offset,def=0" json:"lon_offset,omitempty"`
    // Granularity of dates, normally represented in units of milliseconds since the 1970 epoch.
    DateGranularity  *int32 `protobuf:"varint,18,opt,name=date_granularity,def=1000" json:"date_granularity,omitempty"`
    XXX_unrecognized []byte `json:"-"`
}

func (m *PrimitiveBlock) Reset()         { *m = PrimitiveBlock{} }
func (m *PrimitiveBlock) String() string { return proto.CompactTextString(m) }
func (*PrimitiveBlock) ProtoMessage()    {}

const Default_PrimitiveBlock_Granularity int32 = 100
const Default_PrimitiveBlock_LatOffset int64 = 0
const Default_PrimitiveBlock_LonOffset int64 = 0
const Default_PrimitiveBlock_DateGranularity int32 = 1000

func (m *PrimitiveBlock) GetStringtable() *StringTable {
    if m != nil {
        return m.Stringtable
    }
    return nil
}

func (m *PrimitiveBlock) GetPrimitivegroup() []*PrimitiveGroup {
    if m != nil {
        return m.Primitivegroup
    }
    return nil
}

func (m *PrimitiveBlock) GetGranularity() int32 {
    if m != nil && m.Granularity != nil {
        return *m.Granularity
    }
    return Default_PrimitiveBlock_Granularity
}

func (m *PrimitiveBlock) GetLatOffset() int64 {
    if m != nil && m.LatOffset != nil {
        return *m.LatOffset
    }
    return Default_PrimitiveBlock_LatOffset
}

func (m *PrimitiveBlock) GetLonOffset() int64 {
    if m != nil && m.LonOffset != nil {
        return *m.LonOffset
    }
    return Default_PrimitiveBlock_LonOffset
}

func (m *PrimitiveBlock) GetDateGranularity() int32 {
    if m != nil && m.DateGranularity != nil {
        return *m.DateGranularity
    }
    return Default_PrimitiveBlock_DateGranularity
}

// Group of OSMPrimitives. All primitives in a group must be the same type.
type PrimitiveGroup struct {
    Nodes            []*Node      `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
    Dense            *DenseNodes  `protobuf:"bytes,2,opt,name=dense" json:"dense,omitempty"`
    Ways             []*Way       `protobuf:"bytes,3,rep,name=ways" json:"ways,omitempty"`
    Relations        []*Relation  `protobuf:"bytes,4,rep,name=relations" json:"relations,omitempty"`
    Changesets       []*ChangeSet `protobuf:"bytes,5,rep,name=changesets" json:"changesets,omitempty"`
    XXX_unrecognized []byte       `json:"-"`
}

func (m *PrimitiveGroup) Reset()         { *m = PrimitiveGroup{} }
func (m *PrimitiveGroup) String() string { return proto.CompactTextString(m) }
func (*PrimitiveGroup) ProtoMessage()    {}

func (m *PrimitiveGroup) GetNodes() []*Node {
    if m != nil {
        return m.Nodes
    }
    return nil
}

func (m *PrimitiveGroup) GetDense() *DenseNodes {
    if m != nil {
        return m.Dense
    }
    return nil
}

func (m *PrimitiveGroup) GetWays() []*Way {
    if m != nil {
        return m.Ways
    }
    return nil
}

func (m *PrimitiveGroup) GetRelations() []*Relation {
    if m != nil {
        return m.Relations
    }
    return nil
}

func (m *PrimitiveGroup) GetChangesets() []*ChangeSet {
    if m != nil {
        return m.Changesets
    }
    return nil
}

// * String table, contains the common strings in each block.
//
// Note that we reserve index '0' as a delimiter, so the entry at that
// index in the table is ALWAYS blank and unused.
//
type StringTable struct {
    S                [][]byte `protobuf:"bytes,1,rep,name=s" json:"s,omitempty"`
    XXX_unrecognized []byte   `json:"-"`
}

func (m *StringTable) Reset()         { *m = StringTable{} }
func (m *StringTable) String() string { return proto.CompactTextString(m) }
func (*StringTable) ProtoMessage()    {}

func (m *StringTable) GetS() [][]byte {
    if m != nil {
        return m.S
    }
    return nil
}

// Optional metadata that may be included into each primitive.
type Info struct {
    Version   *int32  `protobuf:"varint,1,opt,name=version,def=-1" json:"version,omitempty"`
    Timestamp *int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
    Changeset *int64  `protobuf:"varint,3,opt,name=changeset" json:"changeset,omitempty"`
    Uid       *int32  `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
    UserSid   *uint32 `protobuf:"varint,5,opt,name=user_sid" json:"user_sid,omitempty"`
    // The visible flag is used to store history information. It indicates that
    // the current object version has been created by a delete operation on the
    // OSM API.
    // When a writer sets this flag, it MUST add a required_features tag with
    // value "HistoricalInformation" to the HeaderBlock.
    // If this flag is not available for some object it MUST be assumed to be
    // true if the file has the required_features tag "HistoricalInformation"
    // set.
    Visible          *bool  `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
    XXX_unrecognized []byte `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}

const Default_Info_Version int32 = -1

func (m *Info) GetVersion() int32 {
    if m != nil && m.Version != nil {
        return *m.Version
    }
    return Default_Info_Version
}

func (m *Info) GetTimestamp() int64 {
    if m != nil && m.Timestamp != nil {
        return *m.Timestamp
    }
    return 0
}

func (m *Info) GetChangeset() int64 {
    if m != nil && m.Changeset != nil {
        return *m.Changeset
    }
    return 0
}

func (m *Info) GetUid() int32 {
    if m != nil && m.Uid != nil {
        return *m.Uid
    }
    return 0
}

func (m *Info) GetUserSid() uint32 {
    if m != nil && m.UserSid != nil {
        return *m.UserSid
    }
    return 0
}

func (m *Info) GetVisible() bool {
    if m != nil && m.Visible != nil {
        return *m.Visible
    }
    return false
}

// * Optional metadata that may be included into each primitive. Special dense format used in DenseNodes.
type DenseInfo struct {
    Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
    Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
    Changeset []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset" json:"changeset,omitempty"`
    Uid       []int32 `protobuf:"zigzag32,4,rep,packed,name=uid" json:"uid,omitempty"`
    UserSid   []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid" json:"user_sid,omitempty"`
    // The visible flag is used to store history information. It indicates that
    // the current object version has been created by a delete operation on the
    // OSM API.
    // When a writer sets this flag, it MUST add a required_features tag with
    // value "HistoricalInformation" to the HeaderBlock.
    // If this flag is not available for some object it MUST be assumed to be
    // true if the file has the required_features tag "HistoricalInformation"
    // set.
    Visible          []bool `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
    XXX_unrecognized []byte `json:"-"`
}

func (m *DenseInfo) Reset()         { *m = DenseInfo{} }
func (m *DenseInfo) String() string { return proto.CompactTextString(m) }
func (*DenseInfo) ProtoMessage()    {}

func (m *DenseInfo) GetVersion() []int32 {
    if m != nil {
        return m.Version
    }
    return nil
}

func (m *DenseInfo) GetTimestamp() []int64 {
    if m != nil {
        return m.Timestamp
    }
    return nil
}

func (m *DenseInfo) GetChangeset() []int64 {
    if m != nil {
        return m.Changeset
    }
    return nil
}

func (m *DenseInfo) GetUid() []int32 {
    if m != nil {
        return m.Uid
    }
    return nil
}

func (m *DenseInfo) GetUserSid() []int32 {
    if m != nil {
        return m.UserSid
    }
    return nil
}

func (m *DenseInfo) GetVisible() []bool {
    if m != nil {
        return m.Visible
    }
    return nil
}

// THIS IS STUB DESIGN FOR CHANGESETS. NOT USED RIGHT NOW.
// TODO:    REMOVE THIS?
type ChangeSet struct {
    Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
    XXX_unrecognized []byte `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}

func (m *ChangeSet) GetId() int64 {
    if m != nil && m.Id != nil {
        return *m.Id
    }
    return 0
}

type Node struct {
    Id  *int64 `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
    // Parallel arrays.
    Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
    Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
    Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
    Lat              *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
    Lon              *int64   `protobuf:"zigzag64,9,req,name=lon" json:"lon,omitempty"`
    XXX_unrecognized []byte   `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

func (m *Node) GetId() int64 {
    if m != nil && m.Id != nil {
        return *m.Id
    }
    return 0
}

func (m *Node) GetKeys() []uint32 {
    if m != nil {
        return m.Keys
    }
    return nil
}

func (m *Node) GetVals() []uint32 {
    if m != nil {
        return m.Vals
    }
    return nil
}

func (m *Node) GetInfo() *Info {
    if m != nil {
        return m.Info
    }
    return nil
}

func (m *Node) GetLat() int64 {
    if m != nil && m.Lat != nil {
        return *m.Lat
    }
    return 0
}

func (m *Node) GetLon() int64 {
    if m != nil && m.Lon != nil {
        return *m.Lon
    }
    return 0
}

type DenseNodes struct {
    Id  []int64 `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
    // repeated Info info = 4;
    Denseinfo *DenseInfo `protobuf:"bytes,5,opt,name=denseinfo" json:"denseinfo,omitempty"`
    Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
    Lon       []int64    `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"`
    // Special packing of keys and vals into one array. May be empty if all nodes in this block are tagless.
    KeysVals         []int32 `protobuf:"varint,10,rep,packed,name=keys_vals" json:"keys_vals,omitempty"`
    XXX_unrecognized []byte  `json:"-"`
}

func (m *DenseNodes) Reset()         { *m = DenseNodes{} }
func (m *DenseNodes) String() string { return proto.CompactTextString(m) }
func (*DenseNodes) ProtoMessage()    {}

func (m *DenseNodes) GetId() []int64 {
    if m != nil {
        return m.Id
    }
    return nil
}

func (m *DenseNodes) GetDenseinfo() *DenseInfo {
    if m != nil {
        return m.Denseinfo
    }
    return nil
}

func (m *DenseNodes) GetLat() []int64 {
    if m != nil {
        return m.Lat
    }
    return nil
}

func (m *DenseNodes) GetLon() []int64 {
    if m != nil {
        return m.Lon
    }
    return nil
}

func (m *DenseNodes) GetKeysVals() []int32 {
    if m != nil {
        return m.KeysVals
    }
    return nil
}

type Way struct {
    Id  *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
    // Parallel arrays.
    Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
    Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
    Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
    Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
    XXX_unrecognized []byte   `json:"-"`
}

func (m *Way) Reset()         { *m = Way{} }
func (m *Way) String() string { return proto.CompactTextString(m) }
func (*Way) ProtoMessage()    {}

func (m *Way) GetId() int64 {
    if m != nil && m.Id != nil {
        return *m.Id
    }
    return 0
}

func (m *Way) GetKeys() []uint32 {
    if m != nil {
        return m.Keys
    }
    return nil
}

func (m *Way) GetVals() []uint32 {
    if m != nil {
        return m.Vals
    }
    return nil
}

func (m *Way) GetInfo() *Info {
    if m != nil {
        return m.Info
    }
    return nil
}

func (m *Way) GetRefs() []int64 {
    if m != nil {
        return m.Refs
    }
    return nil
}

type Relation struct {
    Id  *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
    // Parallel arrays.
    Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
    Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
    Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
    // Parallel arrays
    RolesSid         []int32               `protobuf:"varint,8,rep,packed,name=roles_sid" json:"roles_sid,omitempty"`
    Memids           []int64               `protobuf:"zigzag64,9,rep,packed,name=memids" json:"memids,omitempty"`
    Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=OSMPBF.Relation_MemberType" json:"types,omitempty"`
    XXX_unrecognized []byte                `json:"-"`
}

func (m *Relation) Reset()         { *m = Relation{} }
func (m *Relation) String() string { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()    {}

func (m *Relation) GetId() int64 {
    if m != nil && m.Id != nil {
        return *m.Id
    }
    return 0
}

func (m *Relation) GetKeys() []uint32 {
    if m != nil {
        return m.Keys
    }
    return nil
}

func (m *Relation) GetVals() []uint32 {
    if m != nil {
        return m.Vals
    }
    return nil
}

func (m *Relation) GetInfo() *Info {
    if m != nil {
        return m.Info
    }
    return nil
}

func (m *Relation) GetRolesSid() []int32 {
    if m != nil {
        return m.RolesSid
    }
    return nil
}

func (m *Relation) GetMemids() []int64 {
    if m != nil {
        return m.Memids
    }
    return nil
}

func (m *Relation) GetTypes() []Relation_MemberType {
    if m != nil {
        return m.Types
    }
    return nil
}

func init() {
    proto.RegisterEnum("OSMPBF.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}
