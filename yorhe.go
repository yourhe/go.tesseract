package tesseract

// #cgo LDFLAGS: -L /usr/local/lib -ltesseract
// #include "/usr/local/Cellar/tesseract/3.02.02_3/include/tesseract/capi.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"

	"gopkg.in/GeertJohan/go.leptonica.v1"
)

// BOXA* TessBaseAPIGetRegions( TessBaseAPI* handle, PIXA** pixa);
func (t *Tess) GetRegions(pix *leptonica.Pix) {
	box := C.TessBaseAPIGetRegions(t.tba, nil)
	fmt.Print(box)
}
