// Code generated by protoc-gen-go.
// source: fileformat.proto
// DO NOT EDIT!

/*
Package OSMPBF is a generated protocol buffer package.

It is generated from these files:
    fileformat.proto
    osmformat.proto

It has these top-level messages:
    Blob
    BlobHeader
*/
package OSMPBF

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Blob struct {
    Raw     []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
    RawSize *int32 `protobuf:"varint,2,opt,name=raw_size" json:"raw_size,omitempty"`
    // Possible compressed versions of the data.
    ZlibData []byte `protobuf:"bytes,3,opt,name=zlib_data" json:"zlib_data,omitempty"`
    // PROPOSED feature for LZMA compressed data. SUPPORT IS NOT REQUIRED.
    LzmaData []byte `protobuf:"bytes,4,opt,name=lzma_data" json:"lzma_data,omitempty"`
    // Formerly used for bzip2 compressed data. Depreciated in 2010.
    OBSOLETEBzip2Data []byte `protobuf:"bytes,5,opt,name=OBSOLETE_bzip2_data" json:"OBSOLETE_bzip2_data,omitempty"`
    XXX_unrecognized  []byte `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}

func (m *Blob) GetRaw() []byte {
    if m != nil {
        return m.Raw
    }
    return nil
}

func (m *Blob) GetRawSize() int32 {
    if m != nil && m.RawSize != nil {
        return *m.RawSize
    }
    return 0
}

func (m *Blob) GetZlibData() []byte {
    if m != nil {
        return m.ZlibData
    }
    return nil
}

func (m *Blob) GetLzmaData() []byte {
    if m != nil {
        return m.LzmaData
    }
    return nil
}

func (m *Blob) GetOBSOLETEBzip2Data() []byte {
    if m != nil {
        return m.OBSOLETEBzip2Data
    }
    return nil
}

type BlobHeader struct {
    Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
    Indexdata        []byte  `protobuf:"bytes,2,opt,name=indexdata" json:"indexdata,omitempty"`
    Datasize         *int32  `protobuf:"varint,3,req,name=datasize" json:"datasize,omitempty"`
    XXX_unrecognized []byte  `json:"-"`
}

func (m *BlobHeader) Reset()         { *m = BlobHeader{} }
func (m *BlobHeader) String() string { return proto.CompactTextString(m) }
func (*BlobHeader) ProtoMessage()    {}

func (m *BlobHeader) GetType() string {
    if m != nil && m.Type != nil {
        return *m.Type
    }
    return ""
}

func (m *BlobHeader) GetIndexdata() []byte {
    if m != nil {
        return m.Indexdata
    }
    return nil
}

func (m *BlobHeader) GetDatasize() int32 {
    if m != nil && m.Datasize != nil {
        return *m.Datasize
    }
    return 0
}

func init() {
}
