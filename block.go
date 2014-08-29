package gopbf

/*
 * File structure: sequence of blocks
 * Starting with a block containing OSMHeader blob
 */
type FileBlocks []*Block

type Block struct {
	*BlockReader
    BlobHeader *OSMPBF.BlobHeader
    Blob       *OSMPBF.Blob
}

/*
 * 1. Length of BlobHeader
 * 2. BlobHeader
 * 3. Blob (length in BlobHeader)
 */
func (b *Block) ParseBlock() *Block {

}

/*
 * BlobHeader: block header
 */
func (b *Block) ParseBlobHeader(size uint32) *Block {
	
}

/*
 * Blob: block data/body
 * Serialized and possibly zlib/deflate compressed:
 * - OSMHeader (serialized HeaderBlock) or...
 * - OSMData (serialized PrimitiveBlock)
 */
func (b *Block) ParseBlob(size uint32) *Block {
	
}