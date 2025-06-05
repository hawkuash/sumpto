package convert

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
	Supported_formats_jpeg = []string{".png"}
	Supported_formats_png  = []string{".jpeg", ".jpg"}
)

func SetPNGConvertExtensions(format_list []string) (extensions []string) {
	files.LowerSlice(format_list)
	if format_list == nil {
		return Supported_formats_png
	}
	if slices.Contains(format_list, "jpeg") || slices.Contains(format_list, "jpg") {
		extensions = append(extensions, ".jpg", ".jpeg")
	}
	return
}

func SetJPEGConvertExtensions(format_list []string) (extensions []string) {
	files.LowerSlice(format_list)
	if format_list == nil {
		return Supported_formats_jpeg
	}
	if slices.Contains(format_list, "png") {
		extensions = append(extensions, ".png")
	}
	return
}

func ConvertToPNG(path string) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf("An error occured during opening file at %s: %s \n", path, err)
		return
	}
	files.RemoveBloat(image)
	newImage, _ := image.ToImage(presets.SafePNG())
	path = files.UpdateExtension(path, ".png")
	err = imgio.Save(path, newImage, imgio.PNGEncoder())
	if err != nil {
		log.Printf("An error occured during saving file at %s: %s \n", path, err)
		return
	}
}

func ConvertToJPEG(path string) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf("An error occured during opening file at %s: %s \n", path, err)
		return
	}
	files.RemoveBloat(image)
	bytes, _, _ := image.ExportJpeg(presets.JPEG(100))
	path = files.UpdateExtension(path, ".jpg")
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Printf("An error occured during saving file at %s: %s \n", path, err)
		return
	}
}
