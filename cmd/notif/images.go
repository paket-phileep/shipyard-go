package notif

import (
	logger "github.com/charmbracelet/log"
)

func CompletedReadingImageFile(data any) {
	logger.Info("Sucessfully retrieved images", data)
}

func ReadingImageFile() {
	logger.Info("Reading image.yml file...")
}

// func UnmarshalingImages(images any) {
// 	logger.Info("Unmarshaling.", images)
// }

func CompletedUnmarshalingImages(image any) {
	logger.Info("Unmarshaled image...", image)
}
