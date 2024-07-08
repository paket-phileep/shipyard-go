package notif

import (
	logger "github.com/charmbracelet/log"
)

func CompletedReadingBundleFile(data any) {
	logger.Info("Sucessfully retrieved bundles", data)
}

func ReadingBundleFile() {
	logger.Info("Reading bundle.yml file...")
}

func UnmarshalingBundles() {
	logger.Info("Unmarshaling.")
}

func CompletedUnmarshalingBundles(bundle any) {
	logger.Info("Unmarshaled bundle...", bundle)
}
