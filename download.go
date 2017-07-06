package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Get urlFormat and save it to fileFormat
func Image(urlFormat string, fileFormat string, ins ...interface{}) error {
	imageUrl := fmt.Sprintf(urlFormat, ins...)
	filePath := fmt.Sprintf(fileFormat, ins...)

	fmt.Print("fetching: " + imageUrl + " to " + filePath + " ... ")

	// get url
	resp, err := http.Get(imageUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// make sure target dir exists
	err = os.MkdirAll(filepath.Dir(filePath), 0775)
	if err != nil {
		return err
	}

	// save response body to file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	size, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes fetched\n", size)

	return nil
}
