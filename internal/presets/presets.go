package presets

import "github.com/davidbyttow/govips/v2/vips"

func JPEG(q int) *vips.JpegExportParams {
	return &vips.JpegExportParams{
		Quality:       q,
		SubsampleMode: vips.VipsForeignSubsampleOn,
		Interlace:     false,
		StripMetadata: true,
	}
}

func PNG() *vips.PngExportParams {
	return &vips.PngExportParams{
		Quality:       100,
		Compression:   6,
		Filter:        vips.PngFilterNone,
		Interlace:     false,
		StripMetadata: true,
	}
}

func SafePNG() *vips.ExportParams {
	return &vips.ExportParams{
		Format:        vips.ImageTypePNG,
		Compression:   6,
		Interlaced:    false,
		StripMetadata: true,
	}
}
