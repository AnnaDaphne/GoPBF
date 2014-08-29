package gopbf

type BlockReader struct {
    cursor  uint32
    raw     mmap.MMap
}

func pError(e error) {
    if e != nil {
        panic(e)
    }
}

/*
 * Get size of BlobHeader (aka block header)
 */
func (br *BlockReader) ReadU32() uint32 { // 4
    b = br.Read(4)
    return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

/*
 * Get actual BlobHeader and Blob (aka block header and block data)
 * Not to be confused with OSMHeader and OSMData (blob types)
 */
func (br *BlockReader) Read(size uint32) (bytes []byte) {
    bytes = br.raw[br.cursor:br.cursor + size]
    cursor += size
    return
}