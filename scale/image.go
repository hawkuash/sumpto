package scale

import (
	"log"
	"os"
	"slices"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/presets"
)

var (
	supported_formats = []string{".jpg", ".jpeg", ".png"}
)

func SetScaleExtensions(format_list []string) (extensions []string) {
	files.LowerSlice(format_list)
	if format_list == nil {
		return supported_formats
	}
	if slices.Contains(format_list, "jpeg") || slices.Contains(format_list, "jpg") {
		extensions = append(extensions, ".jpg", ".jpeg")
	}
	if slices.Contains(format_list, "png") {
		extensions = append(extensions, ".png")
	}
	return
}

func ScaleImage(path string, overwrite bool, limit int) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf("An error occured during opening file at %s: %s \n", path, err)
		return
	}
	if min(image.Height(), image.Width()) < limit {
		log.Printf("Image dimensions exceeded size limit: %s", path)
		return
	}
	image.Resize(0.5, vips.KernelAuto)
	switch imgtype := image.Format(); imgtype {
	case vips.ImageTypeJPEG:
		bytes, _, _ := image.ExportJpeg(presets.JPEG(100))
		if !overwrite {
			path = files.UpdateFilename(path, "downscaled")
		}
		err = os.WriteFile(path, bytes, 0644)
		if err != nil {
			log.Printf("An error occured during saving file at %s: %s \n", path, err)
			return
		}
	case vips.ImageTypePNG:
		newImage, _ := image.ToImage(presets.SafePNG())
		if !overwrite {
			path = files.UpdateFilename(path, "downscaled")
		}
		err = imgio.Save(path, newImage, imgio.PNGEncoder())
		if err != nil {
			log.Printf("An error occured during saving file at %s: %s \n", path, err)
			return
		}
	}
}
