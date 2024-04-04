package gb2312

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"unicode/utf8"
)

//按大老师说的做 https://en.wikipedia.org/wiki/HZ_(character_encoding)

func NewEncoder() *encoding.Encoder {
	e := new(Encoder)
	return &encoding.Encoder{Transformer: e}
}

const (
	asciiState = iota
	gbState
)

type Encoder int

func (e *Encoder) Reset() {
	*e = asciiState
}

func (_ *Encoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	r, size := rune(0), 0
	for ; nSrc < len(src); nSrc += size {
		r = rune(src[nSrc])

		// Decode a 1-byte rune.
		if r < utf8.RuneSelf {
			size = 1
			if nDst >= len(dst) {
				err = transform.ErrShortDst
				break
			}
			dst[nDst] = uint8(r)
			nDst += 1
			continue
		}

		// Decode a multi-byte rune.
		r, size = utf8.DecodeRune(src[nSrc:])
		if size == 1 {
			// All valid runes of size 1 (those below utf8.RuneSelf) were
			// handled above. We have invalid UTF-8 or we haven't seen the
			// full character yet.
			if !atEOF && !utf8.FullRune(src[nSrc:]) {
				err = transform.ErrShortSrc
				break
			}
		}

		// func init checks that the switch covers all tables.
		switch {
		case encode0Low <= r && r < encode0High:
			if r = rune(encode0[r-encode0Low]); r != 0 {
				goto writeGB
			}
		case encode1Low <= r && r < encode1High:
			if r = rune(encode1[r-encode1Low]); r != 0 {
				goto writeGB
			}
		case encode2Low <= r && r < encode2High:
			if r = rune(encode2[r-encode2Low]); r != 0 {
				goto writeGB
			}
		case encode3Low <= r && r < encode3High:
			if r = rune(encode3[r-encode3Low]); r != 0 {
				goto writeGB
			}
		case encode4Low <= r && r < encode4High:
			if r = rune(encode4[r-encode4Low]); r != 0 {
				goto writeGB
			}
		}

	writeGB:
		c0 := uint8(r >> 8)
		c1 := uint8(r)
		if nDst+2 > len(dst) {
			err = transform.ErrShortDst
			break
		}
		dst[nDst+0] = c0
		dst[nDst+1] = c1
		nDst += 2
		continue
	}
	// TODO: should one always terminate in ASCII state to make it safe to
	// concatenate two HZ-GB2312-encoded strings?
	return nDst, nSrc, err
}
