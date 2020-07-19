package common

import(
	"fmt"
	"github.com/gucchisk/bytestring"
)

type Format int
const (
	Unknown Format = iota
	Ascii
	Hex
	Base64
)

func NewFormat(format string) (Format, error) {
	switch format {
	case "ascii":
		return Ascii, nil
	case "hex":
		return Hex, nil
	case "base64":
		return Base64, nil
	default:
		return Unknown, fmt.Errorf("unknown format: %s", format)
	}
}

func ByteToStr(b []byte, out Format) string {
	bytes := bytestring.NewBytes(b)
	switch out {
	case Ascii:
		return bytes.String()
	case Hex:
		return bytes.HexString()
	case Base64:
		return bytes.Base64()
	default:
		return ""
	}
}

func StrToByte(s string, in Format) ([]byte, error) {
	var b bytestring.Bytes
	var err error
	switch in {
	case Ascii:
		b, err = bytestring.NewBytesFrom(s, bytestring.Normal)
	case Hex:
		b, err = bytestring.NewBytesFrom(s, bytestring.Hex)
	case Base64:
		b, err = bytestring.NewBytesFrom(s, bytestring.Base64)
	default:
		return []byte{}, fmt.Errorf("unknown format")
	}
	return b.ByteArray(), err
}
