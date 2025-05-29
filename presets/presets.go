package presets

import "github.com/davidbyttow/govips/v2/vips"

func JpegExportParams(q int) *vips.JpegExportParams {
	return &vips.JpegExportParams{
		Quality:       q,
		SubsampleMode: vips.VipsForeignSubsampleOn,
		Interlace:     false,
		StripMetadata: true,
	}
}

func PngExportParams() *vips.PngExportParams {
	return &vips.PngExportParams{
		Quality:       100,
		Compression:   6,
		Filter:        vips.PngFilterNone,
		Interlace:     false,
		StripMetadata: true,
	}
}
