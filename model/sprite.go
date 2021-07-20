package model

// Sprite model
type Sprite struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
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
	// Height of one thumbnail, either height or width are required fields. If only one is given the encoder will calculate the other way value based on the aspect ratio of the video file. If the encoder version is below 2.83.0 both are required
	Height *int32 `json:"height,omitempty"`
	// Width of one thumbnail, either height or width are required fields. If only one is given the encoder will calculate the other way value based on the aspect ratio of the video file. If the encoder version is below 2.83.0 both are required
	Width *int32     `json:"width,omitempty"`
	Unit  SpriteUnit `json:"unit,omitempty"`
	// Distance in the given unit between a screenshot
	Distance *float64 `json:"distance,omitempty"`
	// Name of the sprite image. File extension \".jpg\"/\".jpeg\" or \".png\" is required. (required)
	SpriteName *string `json:"spriteName,omitempty"`
	// Filename of the sprite image. If not set, spriteName will be used, but without an extension.
	Filename *string `json:"filename,omitempty"`
	// Filename of the vtt-file. The file-extension \".vtt\" is required.
	VttName *string          `json:"vttName,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Number of images per file. If more images are generated than specified in this value, multiple sprites will be created. You can use the placeholder '%number%' in the spriteName to specify the naming policy. Either this property must be set or hTiles and vTiles.
	ImagesPerFile *int32 `json:"imagesPerFile,omitempty"`
	// Number of rows of images per file.  Has to be set together with vTiles. If this property and vTiles are set, the imagesPerFile property must not be set.  It is recommended to use the placeholder '%number%' in the spriteName to allow the generation of multiple sprites.  Only supported starting with encoder version `2.76.0`.
	HTiles *int32 `json:"hTiles,omitempty"`
	// Number of columns of images per file.  Has to be set together with hTiles. If this property and hTiles are set, the imagesPerFile property must not be set.  It is recommended to use the placeholder '%number%' in the spriteName to allow the generation of multiple sprites.  Only supported starting with encoder version `2.76.0`.
	VTiles *int32 `json:"vTiles,omitempty"`
	// Additional configuration for JPEG sprite generation.  If this property is set the extension of the file must be '.jpg.' or '.jpeg'  Only supported starting with encoder version `2.76.0`
	JpegConfig *SpriteJpegConfig `json:"jpegConfig,omitempty"`
	// The creation mode for the thumbnails in the Sprite.  Two possible creation modes exist: generate thumbnails starting with the beginning of the video or after the first configured period.  When using distance=10 and unit=SECONDS and INTERVAL_END, the first image of the sprite is from the second 10 of the video. When using distance=10 and unit=SECONDS and INTERVAL_START, the first image of the sprite is from the very start of the video, while the second image is from second 10 of the video.  It is recommended to use 'INTERVAL_START' when using the sprites for trick play so that there is an additional thumbnail from the beginning of the video.  Only supported starting with encoder version `2.76.0`.
	CreationMode SpriteCreationMode `json:"creationMode,omitempty"`
	// Specifies the aspect mode that is used when both height and width are specified Only supported starting with encoder version `2.85.0`.
	AspectMode ThumbnailAspectMode `json:"aspectMode,omitempty"`
}
