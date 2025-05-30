package scale

import (
	"os"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/presets"
)

func ScaleImage(path string, ep *vips.ExportParams, overwrite bool) {
	image, err := vips.NewImageFromFile(path)
	if !files.CheckError(err, false) {
		image.Resize(0.5, vips.KernelAuto)
	}
	switch imgtype := image.Format(); imgtype {
	case vips.ImageTypeJPEG:
		bytes, _, _ := image.ExportJpeg(presets.JpegExportParams(100))
		if !overwrite {
			path = files.NewFilename(path, "downscaled")
		}
		err = os.WriteFile(path, bytes, 0644)
	case vips.ImageTypePNG:
		newImage, _ := image.ToImage(presets.PngExportParamsSafe())
		if !overwrite {
			path = files.NewFilename(path, "downscaled")
		}
		imgio.Save(path, newImage, imgio.PNGEncoder())
	}
	// imagebytes := setExport(image)
	// if !overwrite {
	// 	path = files.NewFilename(path, "downscaled")
	// }
	// err = os.WriteFile(path, imagebytes, 0644)
	// files.CheckError(err, false)
}
