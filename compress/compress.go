package compress

import (
	"bytes"
	"compress/flate"
	"io"
)

// Compress returns a compressed byte slice.
func Compress(src []byte) []byte {
	compressedData := new(bytes.Buffer)
	compress(src, compressedData, -2)
	return compressedData.Bytes()
}

// Decompress returns a decompressed byte slice.
func Decompress(src []byte) []byte {
	compressedData := bytes.NewBuffer(src)
	deCompressedData := new(bytes.Buffer)
	decompress(compressedData, deCompressedData)
	return deCompressedData.Bytes()
}

// compress uses flate to compress a byte slice to a corresponding level
func compress(src []byte, dest io.Writer, level int) {
	compressor, err := flate.NewWriter(dest, level)
	if err != nil {
		return
	}
	_, _ = compressor.Write(src)
	compressor.Close()
}

// compress uses flate to decompress an io.Reader
func decompress(src io.Reader, dest io.Writer) {
	decompressor := flate.NewReader(src)
	_, _ = io.Copy(dest, decompressor)
	decompressor.Close()
}
