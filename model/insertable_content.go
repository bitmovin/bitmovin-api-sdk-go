package model
import (
	"time"
)

type InsertableContent struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Either a list of video files to be inserted in the live stream which have to match the codec, aspect ratio and frame rate of the live stream or a single image file. Supported image formats are: `.Y.U.V`, `Alias PIX`, `animated GIF`, `APNG`, `BMP`, `DPX`, `FITS`, `JPEG`, `JPEG 2000`, `JPEG-LS`, `PAM`, `PBM`, `PCX`, `PGM`, `PGMYUV`, `PNG`, `PPM`, `SGI`, `Sun Rasterfile`, `TIFF`, `Truevision Targa`, `WebP`, `XBM`, `XFace`, `XPM`, `XWD`
	Inputs []InsertableContentInput `json:"inputs,omitempty"`
	// Status of the insertable content.
	Status InsertableContentStatus `json:"status,omitempty"`
}

