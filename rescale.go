package rescale

import (
	"bytes"
	"gopkg.in/gographics/imagick.v3/imagick"
	"image/jpeg"
	"image/png"
)

func Make(srcImageBytes []byte, countOfRescale int) ([]byte, error) {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	if err := mw.ReadImageBlob(srcImageBytes); err != nil { return nil, err }

	imageBuffer := bytes.NewBuffer(srcImageBytes)
	im, err := jpeg.Decode(imageBuffer)
	if err != nil {
		im, err = png.Decode(imageBuffer)
		if err != nil { return nil, err }
	}

	if err = mw.LiquidRescaleImage(
		uint(float64(im.Bounds().Size().X) / (1.2 * float64(countOfRescale))),
		uint(float64(im.Bounds().Size().Y) / (1.2 * float64(countOfRescale))),
		0, 0,
	); err != nil { return nil, err }
	if err = mw.ResizeImage(
		uint(im.Bounds().Size().X), uint(im.Bounds().Size().Y),
		imagick.FILTER_LANCZOS2,
	); err != nil { return nil, err }
	return mw.GetImageBlob(), nil
}
