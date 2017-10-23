package distro

import (
	"bytes"
	"compress/gzip"
)

type gzipTransformer struct {
}

func (gzt gzipTransformer) Transform(data bytes.Buffer) bytes.Buffer {
	var buf bytes.Buffer
	// TODO: reuse the gzip writer could improve performace
	zw := gzip.NewWriter(&buf)
	zw.Write(data.Bytes())
	zw.Close()
	return buf
}
