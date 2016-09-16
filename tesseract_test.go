package tesseract

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/GeertJohan/go.leptonica.v1"
)

func TestClear2(ts *testing.T) {
	// create new tess instance and point it to the tessdata location. Set language to english.
	tessdataprefix := os.Getenv("TESSDATA_PREFIX")
	if tessdataprefix == "" {
		tessdataprefix = "/usr/local/share"
	}
	t, err := NewTess(filepath.Join(tessdataprefix, "tessdata"), "eng")
	if err != nil {
		ts.Errorf("Error while initializing")
		// 	log.Fatalf("Error while initializing Tess: %s\n", err)
	}
	defer t.Close()
	// open a new Pix from file with leptonica
	pix, err := leptonica.NewPixFromFile("/Users/yorhe/work/gocode/src/github.com/yourhe/go.tesseract/tessexample/FelixScan.jpg")
	if err != nil {
		ts.Fatalf("Error while getting pix from file: %s\n", err)
	}
	defer pix.Close() // remember to cleanup
	t.SetImagePix(pix)
	t.GetRegions(nil)
	// set the page seg mode to autodetect
	// t.SetPageSegMode(PSM_AUTO_OSD)

}
