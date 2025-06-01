package scale

import (
	"log"
	"os"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/presets"
)

type Search_options struct {
	All, Jpeg, Png bool
}

func SetScaleExtensions(so Search_options) (extensions []string) {
	if so.All {
		extensions = append(extensions, ".jpg", ".jpeg", ".png")
		return
	}
	if so.Jpeg {
		extensions = append(extensions, ".jpg", ".jpeg")
	}
	if so.Png {
		extensions = append(extensions, ".png")
	}
	return
}

func ScaleImage(path string, overwrite bool) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf("An error occured during opening file at %s: %s \n", path, err)
		return
	}
	image.Resize(0.5, vips.KernelAuto)
	switch imgtype := image.Format(); imgtype {
	case vips.ImageTypeJPEG:
		bytes, _, _ := image.ExportJpeg(presets.JPEG(100))
		if !overwrite {
			path = files.NewFilename(path, "downscaled")
		}
		err = os.WriteFile(path, bytes, 0644)
		if err != nil {
			log.Printf("An error occured during saving file at %s: %s \n", path, err)
			return
		}
	case vips.ImageTypePNG:
		newImage, _ := image.ToImage(presets.SafePNG())
		if !overwrite {
			path = files.NewFilename(path, "downscaled")
		}
		err = imgio.Save(path, newImage, imgio.PNGEncoder())
		if err != nil {
			log.Printf("An error occured during saving file at %s: %s \n", path, err)
			return
		}
	}
}
