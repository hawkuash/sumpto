package scale

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
)

func ScaleImage(path string, ep *vips.ExportParams, overwrite bool) {
	image, err := vips.NewImageFromFile(path)
	if !files.CheckError(err, false) {
		image.Resize(0.5, vips.KernelAuto)
	}
	imagebytes, _, err := image.ExportNative()
	if !overwrite {
		path = files.NewFilename(path, "downscaled")
	}
	err = os.WriteFile(path, imagebytes, 0644)
	files.CheckError(err, false)
}
