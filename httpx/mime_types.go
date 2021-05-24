package httpx

import "strings"

// mime types
const (
	MimetypeJson        = "application/json"
	MimetypeFormEncoded = "application/x-www-form-urlencoded"
	MimeTypeOctetStream = "application/octet-stream"

	MimeTypePlainText  = "text/plain"
	MimeTypeCSS        = "text/css"
	MimeTypeHTML       = "text/html"
	MimeTypeJavascript = "text/javascript"

	MimeTypeMultipart = "multipart/form-data"

	// images.

	MimeTypeAPNG = "image/apng"
	MimeTypeBMP  = "image/bmp"
	MimeTypeGIF  = "image/gif"
	MimeTypeICO  = "image/x-icon"
	MimeTypeJPG  = "image/jpeg"
	MimeTypePNG  = "image/png"
	MimeTypeSVG  = "image/svg+xml"
	MimeTypeTIFF = "image/tiff"
	MimeTypeWEBP = "image/webp"
)

// MimeType is for http mime type.
type MimeType struct {
	Type    string // the type
	SubType string // the subtype
}

func (m *MimeType) String() string {
	return m.Type + "/" + m.SubType
}

// Unpack is for convenient unpack mime type values.
func (m *MimeType) Unpack() (_type, subType string) {
	return m.Type, m.SubType
}

// ParseMimeType parse a mimetype string to type and subtype.
// This function always return with no error, but results may be undefined and vary for invalid input mime-type strings.
func ParseMimeType(mimeType string) *MimeType {
	idx := strings.IndexByte(mimeType, '/')
	if idx < 0 {
		return &MimeType{mimeType, ""}
	}
	return &MimeType{mimeType[0:idx], mimeType[idx+1:]}
}
