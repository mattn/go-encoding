package encoding

import (
	"strings"

	enc "golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

var encodingMap = map[string]*enc.Encoding{
	"CP437":             &charmap.CodePage437,
	"CP866":             &charmap.CodePage866,
	"ISO-2022-JP":       &japanese.ISO2022JP,
	"ISO-8859-10":       &charmap.ISO8859_10,
	"ISO-8859-13":       &charmap.ISO8859_13,
	"ISO-8859-14":       &charmap.ISO8859_14,
	"ISO-8859-15":       &charmap.ISO8859_15,
	"ISO-8859-16":       &charmap.ISO8859_16,
	"ISO-8859-2":        &charmap.ISO8859_2,
	"ISO-8859-3":        &charmap.ISO8859_3,
	"ISO-8859-4":        &charmap.ISO8859_4,
	"ISO-8859-5":        &charmap.ISO8859_5,
	"ISO-8859-6":        &charmap.ISO8859_6,
	"ISO-8859-7":        &charmap.ISO8859_7,
	"ISO-8859-8":        &charmap.ISO8859_8,
	"KOI8R":             &charmap.KOI8R,
	"KOI8U":             &charmap.KOI8U,
	"Macintosh":         &charmap.Macintosh,
	"MacintoshCyrillic": &charmap.MacintoshCyrillic,
	"Windows1250":       &charmap.Windows1250,
	"Windows1251":       &charmap.Windows1251,
	"Windows1252":       &charmap.Windows1252,
	"Windows1253":       &charmap.Windows1253,
	"Windows1254":       &charmap.Windows1254,
	"Windows1255":       &charmap.Windows1255,
	"Windows1256":       &charmap.Windows1256,
	"Windows1257":       &charmap.Windows1257,
	"Windows1258":       &charmap.Windows1258,
	"Windows874":        &charmap.Windows874,
	"XUserDefined":      &charmap.XUserDefined,
	"Big5":              &traditionalchinese.Big5,
	"EUC-KR":            &korean.EUCKR,
	"HZ-GB2312":         &simplifiedchinese.HZGB2312,
	"sjis":              &japanese.ShiftJIS,
	"CP932":             &japanese.ShiftJIS,
	"Shift_JIS":         &japanese.ShiftJIS,
	"EUC-JP":            &japanese.EUCJP,
}

var encodingNames []string

func init() {
	for k := range encodingMap {
		encodingNames = append(encodingNames, k)
	}
}

func EncodingNames() []string {
	return encodingNames
}

func GetEncoding(name string) enc.Encoding {
	name = strings.ToLower(name)
	for k, v := range encodingMap {
		if strings.ToLower(k) == name {
			return *v
		}
	}
	return nil
}
