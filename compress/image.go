package compress

import (
	"log"
	"os"
	"slices"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/presets"
)

var (
	supported_formats = []string{".jpeg", ".jpg"}
)

func SetCompressExtensions(format_list []string) (extensions []string) {
	files.LowerSlice(format_list)
	if format_list == nil {
		return supported_formats
	}
	if slices.Contains(format_list, "jpeg") || slices.Contains(format_list, "jpg") {
		extensions = append(extensions, ".jpg", ".jpeg")
	}
	return
}

func CompressImage(path string, overwrite bool) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf("An error occured during opening file at %s: %s \n", path, err)
		return
	}
	files.RemoveBloat(image)
	switch imgtype := image.Format(); imgtype {
	case vips.ImageTypeJPEG:
		bytes, _, _ := image.ExportJpeg(presets.JPEG(files.Quality))
		if !overwrite {
			path = files.UpdateFilename(path, "compressed")
		}
		err = os.WriteFile(path, bytes, 0644)
		if err != nil {
			log.Printf("An error occured during saving file at %s: %s \n", path, err)
			return
		}
	default:
		log.Printf("Passed unsupported filetype at: %s", path)
	}
}
