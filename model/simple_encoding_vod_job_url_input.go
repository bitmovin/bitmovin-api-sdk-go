package model

import (
	"encoding/json"
)

// SimpleEncodingVodJobUrlInput model
type SimpleEncodingVodJobUrlInput struct {
	// Define a URL pointing to the asset that should be encoded. The URL has to point to a file.  Currently the following protocols/storages systems are supported: HTTP(S), (S)FTP, S3, GCS, Azure Blob Storage, Akamai NetStorage. Note that most protocols will require `credentials` to access the asset. Check the description below which ones are applicable. See below how to construct the URLs for the individual protocols/storage systems.  ---  **HTTP** and **HTTPS**: * `http://<host>[:<port>]/path/file.mp4` * `https://<host>[:<port>]/path/file.mp4`  The port is defaulted to 80 if it's not provided. If the content is secured by Basic Authentication please provide corresponding credentials.  **FTP** and **SFTP**: * `ftp://<host>[:<port>]/path/file.mp4` * `sftp://<host>[:<port>]/path/file.mp4`  The port is defaulted to 21 (ftp) or  22 (sftp) if it's not provided. If the content is secured please provide corresponding credentials.  **S3**: * `s3://<my-bucket>/path/file.mp4`  Authentication can be done via accesskey/secretkey or role-based authentication. Generic S3 is currently NOT supported.  **GCS**: * `gcs://<my-bucket>/path/file.mp4`  Authentication can be done via accesskey/secretkey or a service account  **Azure Blob Storage (ABS)**: * `https://<account>.blob.core.windows.net/<container>/path/file.mp4`  It is required to provide the Azure key credentials for authentication.  **Akamai NetStorage**: * `https://<host>-nsu.akamaihd.net/<CP-code>/path/file.mp4`  It is required to provide username/password credentials for authentication. (required)
	Url *string `json:"url,omitempty"`
	// Credentials to be used for authentication and accessing the file. Check out the examples on how to define the credentials correctly.
	Credentials *SimpleEncodingVodJobCredentials `json:"credentials,omitempty"`
	// Defines the type of the input file, if no type is set it is assumed that the input file contains at least one video stream and optionally one or multiple audio streams.  Note that when defining video and audio inputs, you can either - add one single input without inputType, in which case that source file must contain a video stream and (if you want audio) one audio stream, or - add one single input with inputType=VIDEO and (if you want audio) one or more inputs with inputType=AUDIO (each containing one audio stream)  Other combinations are not valid.
	InputType SimpleEncodingVodJobInputType `json:"inputType,omitempty"`
	// The language of the audio track, the subtitles, or closed captions file. The language code  must be compliant with [BCP 47](https://datatracker.ietf.org/doc/html/rfc5646).  This property is mandatory if the input provided is of type SUBTITLES or CLOSED_CAPTIONS and  recommended for input type AUDIO and an input that does not specify an input type (combined  audio and video). If an audio input does not specify the language, it is defaulted to `und`  (undefined).
	Language *string `json:"language,omitempty"`
}

func (m SimpleEncodingVodJobUrlInput) SimpleEncodingVodJobInputSourceType() SimpleEncodingVodJobInputSourceType {
	return SimpleEncodingVodJobInputSourceType_URL
}
func (m SimpleEncodingVodJobUrlInput) MarshalJSON() ([]byte, error) {
	type M SimpleEncodingVodJobUrlInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "URL"

	return json.Marshal(x)
}
