package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"time"
)

// this is the function used to help in the processing of image in submit job function

func DownloadAndProcessImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	_, _, err = image.Decode(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	width := rand.Intn(100) + 1
	height := rand.Intn(100) + 1
	perimeter := 2 * (width + height)

	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	fmt.Printf("Perimeter of image: %d\n", perimeter)

	return nil
}
