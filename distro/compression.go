package distro

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
)

type gzipTransformer struct {
	writer *gzip.Writer
}

func (gzt gzipTransformer) Transform(data bytes.Buffer) bytes.Buffer {
	var buf bytes.Buffer

	if gzt.writer == nil {
		gzt.writer = gzip.NewWriter(&buf)
	} else {
		gzt.writer.Reset(&buf)
	}

	gzt.writer.Write(data.Bytes())
	gzt.writer.Close()
	return buf
}

type zlibTransformer struct {
	writer *zlib.Writer
}

func (zlt zlibTransformer) Transform(data bytes.Buffer) bytes.Buffer {
	var buf bytes.Buffer

	if zlt.writer == nil {
		zlt.writer = zlib.NewWriter(&buf)
	} else {
		zlt.writer.Reset(&buf)
	}

	zlt.writer.Write(data.Bytes())
	zlt.writer.Close()
	return buf
}
