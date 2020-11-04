package model

// InsertableContent model
type InsertableContent struct {
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Either a list of movie files to be inserted in the live stream or a single image file. The movie files must have a video stream at stream position 0, which matches the codec, resolution and framerate of the livestream. The number of audio streams must also be the same and they have to match the sample format, number of channels and sample rate of the audio streams of the livestream. Supported image formats are: `.Y.U.V`, `Alias PIX`, `animated GIF`, `APNG`, `BMP`, `DPX`, `FITS`, `JPEG`, `JPEG 2000`, `JPEG-LS`, `PAM`, `PBM`, `PCX`, `PGM`, `PGMYUV`, `PNG`, `PPM`, `SGI`, `Sun Rasterfile`, `TIFF`, `Truevision Targa`, `WebP`, `XBM`, `XFace`, `XPM`, `XWD`
	Inputs []InsertableContentInput `json:"inputs,omitempty"`
	// Status of the insertable content.
	Status InsertableContentStatus `json:"status,omitempty"`
}
