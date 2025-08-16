package convert

import (
	"log"
	"os"
	"slices"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/internal/files"
	"github.com/hawkuash/sumpto/internal/presets"
)

var (
	supportedFormatsJPEG = []string{".png"}
	supportedFormatsPNG  = []string{".jpeg", ".jpg"}
)

func SetPNGConvertExtensions(formatList []string) (extensions []string) {
	files.LowerSlice(formatList)

	if formatList == nil {
		return supportedFormatsPNG
	}

	if slices.Contains(formatList, "jpeg") ||
		slices.Contains(formatList, "jpg") {
		extensions = append(extensions, ".jpg", ".jpeg")
	}

	return
}

func SetJPEGConvertExtensions(formatList []string) (extensions []string) {
	files.LowerSlice(formatList)

	if formatList == nil {
		return supportedFormatsJPEG
	}

	if slices.Contains(formatList, "png") {
		extensions = append(extensions, ".png")
	}

	return
}

func ConvertToPNG(path string) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf(
			"An error occured during opening file at %s: %s \n",
			path,
			err,
		)

		return
	}

	files.RemoveBloat(image)
	newImage, _ := image.ToImage(presets.SafePNG())
	path = files.UpdateExtension(path, ".png")

	err = imgio.Save(path, newImage, imgio.PNGEncoder())
	if err != nil {
		log.Printf(
			"An error occured during saving file at %s: %s \n",
			path,
			err,
		)

		return
	}
}

func ConvertToJPEG(path string) {
	image, err := vips.NewImageFromFile(path)
	if err != nil {
		log.Printf(
			"An error occured during opening file at %s: %s \n",
			path,
			err,
		)

		return
	}

	files.RemoveBloat(image)
	bytes, _, _ := image.ExportJpeg(presets.JPEG(files.Quality))
	path = files.UpdateExtension(path, ".jpg")

	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Printf(
			"An error occured during saving file at %s: %s \n",
			path,
			err,
		)

		return
	}
}
