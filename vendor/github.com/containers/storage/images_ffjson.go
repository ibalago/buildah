// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: ./images.go

package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/opencontainers/go-digest"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Image) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Image) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteByte(',')
	if len(j.Digest) != 0 {
		buf.WriteString(`"digest":`)
		fflib.WriteJsonString(buf, string(j.Digest))
		buf.WriteByte(',')
	}
	if len(j.Names) != 0 {
		buf.WriteString(`"names":`)
		if j.Names != nil {
			buf.WriteString(`[`)
			for i, v := range j.Names {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.NamesHistory) != 0 {
		buf.WriteString(`"names-history":`)
		if j.NamesHistory != nil {
			buf.WriteString(`[`)
			for i, v := range j.NamesHistory {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.TopLayer) != 0 {
		buf.WriteString(`"layer":`)
		fflib.WriteJsonString(buf, string(j.TopLayer))
		buf.WriteByte(',')
	}
	if len(j.MappedTopLayers) != 0 {
		buf.WriteString(`"mapped-layers":`)
		if j.MappedTopLayers != nil {
			buf.WriteString(`[`)
			for i, v := range j.MappedTopLayers {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.Metadata) != 0 {
		buf.WriteString(`"metadata":`)
		fflib.WriteJsonString(buf, string(j.Metadata))
		buf.WriteByte(',')
	}
	if len(j.BigDataNames) != 0 {
		buf.WriteString(`"big-data-names":`)
		if j.BigDataNames != nil {
			buf.WriteString(`[`)
			for i, v := range j.BigDataNames {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BigDataSizes) != 0 {
		if j.BigDataSizes == nil {
			buf.WriteString(`"big-data-sizes":null`)
		} else {
			buf.WriteString(`"big-data-sizes":{ `)
			for key, value := range j.BigDataSizes {
				fflib.WriteJsonString(buf, key)
				buf.WriteString(`:`)
				fflib.FormatBits2(buf, uint64(value), 10, value < 0)
				buf.WriteByte(',')
			}
			buf.Rewind(1)
			buf.WriteByte('}')
		}
		buf.WriteByte(',')
	}
	if len(j.BigDataDigests) != 0 {
		if j.BigDataDigests == nil {
			buf.WriteString(`"big-data-digests":null`)
		} else {
			buf.WriteString(`"big-data-digests":{ `)
			for key, value := range j.BigDataDigests {
				fflib.WriteJsonString(buf, key)
				buf.WriteString(`:`)
				fflib.WriteJsonString(buf, string(value))
				buf.WriteByte(',')
			}
			buf.Rewind(1)
			buf.WriteByte('}')
		}
		buf.WriteByte(',')
	}
	if true {
		buf.WriteString(`"created":`)

		{

			obj, err = j.Created.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
		buf.WriteByte(',')
	}
	if len(j.Flags) != 0 {
		buf.WriteString(`"flags":`)
		/* Falling back. type=map[string]interface {} kind=map */
		err = buf.Encode(j.Flags)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtImagebase = iota
	ffjtImagenosuchkey

	ffjtImageID

	ffjtImageDigest

	ffjtImageNames

	ffjtImageNamesHistory

	ffjtImageTopLayer

	ffjtImageMappedTopLayers

	ffjtImageMetadata

	ffjtImageBigDataNames

	ffjtImageBigDataSizes

	ffjtImageBigDataDigests

	ffjtImageCreated

	ffjtImageFlags
)

var ffjKeyImageID = []byte("id")

var ffjKeyImageDigest = []byte("digest")

var ffjKeyImageNames = []byte("names")

var ffjKeyImageNamesHistory = []byte("names-history")

var ffjKeyImageTopLayer = []byte("layer")

var ffjKeyImageMappedTopLayers = []byte("mapped-layers")

var ffjKeyImageMetadata = []byte("metadata")

var ffjKeyImageBigDataNames = []byte("big-data-names")

var ffjKeyImageBigDataSizes = []byte("big-data-sizes")

var ffjKeyImageBigDataDigests = []byte("big-data-digests")

var ffjKeyImageCreated = []byte("created")

var ffjKeyImageFlags = []byte("flags")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Image) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Image) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtImagebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtImagenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyImageBigDataNames, kn) {
						currentKey = ffjtImageBigDataNames
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyImageBigDataSizes, kn) {
						currentKey = ffjtImageBigDataSizes
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyImageBigDataDigests, kn) {
						currentKey = ffjtImageBigDataDigests
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyImageCreated, kn) {
						currentKey = ffjtImageCreated
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyImageDigest, kn) {
						currentKey = ffjtImageDigest
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyImageFlags, kn) {
						currentKey = ffjtImageFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyImageID, kn) {
						currentKey = ffjtImageID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyImageTopLayer, kn) {
						currentKey = ffjtImageTopLayer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyImageMappedTopLayers, kn) {
						currentKey = ffjtImageMappedTopLayers
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyImageMetadata, kn) {
						currentKey = ffjtImageMetadata
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyImageNames, kn) {
						currentKey = ffjtImageNames
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyImageNamesHistory, kn) {
						currentKey = ffjtImageNamesHistory
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyImageFlags, kn) {
					currentKey = ffjtImageFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyImageCreated, kn) {
					currentKey = ffjtImageCreated
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageBigDataDigests, kn) {
					currentKey = ffjtImageBigDataDigests
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageBigDataSizes, kn) {
					currentKey = ffjtImageBigDataSizes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageBigDataNames, kn) {
					currentKey = ffjtImageBigDataNames
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyImageMetadata, kn) {
					currentKey = ffjtImageMetadata
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageMappedTopLayers, kn) {
					currentKey = ffjtImageMappedTopLayers
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyImageTopLayer, kn) {
					currentKey = ffjtImageTopLayer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageNamesHistory, kn) {
					currentKey = ffjtImageNamesHistory
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageNames, kn) {
					currentKey = ffjtImageNames
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyImageDigest, kn) {
					currentKey = ffjtImageDigest
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyImageID, kn) {
					currentKey = ffjtImageID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtImagenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtImageID:
					goto handle_ID

				case ffjtImageDigest:
					goto handle_Digest

				case ffjtImageNames:
					goto handle_Names

				case ffjtImageNamesHistory:
					goto handle_NamesHistory

				case ffjtImageTopLayer:
					goto handle_TopLayer

				case ffjtImageMappedTopLayers:
					goto handle_MappedTopLayers

				case ffjtImageMetadata:
					goto handle_Metadata

				case ffjtImageBigDataNames:
					goto handle_BigDataNames

				case ffjtImageBigDataSizes:
					goto handle_BigDataSizes

				case ffjtImageBigDataDigests:
					goto handle_BigDataDigests

				case ffjtImageCreated:
					goto handle_Created

				case ffjtImageFlags:
					goto handle_Flags

				case ffjtImagenosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Digest:

	/* handler: j.Digest type=digest.Digest kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Digest", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Digest = digest.Digest(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Names:

	/* handler: j.Names type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Names = nil
		} else {

			j.Names = []string{}

			wantVal := true

			for {

				var tmpJNames string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJNames type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJNames = string(string(outBuf))

					}
				}

				j.Names = append(j.Names, tmpJNames)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NamesHistory:

	/* handler: j.NamesHistory type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.NamesHistory = nil
		} else {

			j.NamesHistory = []string{}

			wantVal := true

			for {

				var tmpJNamesHistory string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJNamesHistory type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJNamesHistory = string(string(outBuf))

					}
				}

				j.NamesHistory = append(j.NamesHistory, tmpJNamesHistory)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TopLayer:

	/* handler: j.TopLayer type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.TopLayer = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MappedTopLayers:

	/* handler: j.MappedTopLayers type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.MappedTopLayers = nil
		} else {

			j.MappedTopLayers = []string{}

			wantVal := true

			for {

				var tmpJMappedTopLayers string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJMappedTopLayers type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJMappedTopLayers = string(string(outBuf))

					}
				}

				j.MappedTopLayers = append(j.MappedTopLayers, tmpJMappedTopLayers)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Metadata:

	/* handler: j.Metadata type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Metadata = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BigDataNames:

	/* handler: j.BigDataNames type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BigDataNames = nil
		} else {

			j.BigDataNames = []string{}

			wantVal := true

			for {

				var tmpJBigDataNames string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBigDataNames type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBigDataNames = string(string(outBuf))

					}
				}

				j.BigDataNames = append(j.BigDataNames, tmpJBigDataNames)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BigDataSizes:

	/* handler: j.BigDataSizes type=map[string]int64 kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BigDataSizes = nil
		} else {

			j.BigDataSizes = make(map[string]int64, 0)

			wantVal := true

			for {

				var k string

				var tmpJBigDataSizes int64

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJBigDataSizes type=int64 kind=int64 quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmpJBigDataSizes = int64(tval)

					}
				}

				j.BigDataSizes[k] = tmpJBigDataSizes

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BigDataDigests:

	/* handler: j.BigDataDigests type=map[string]digest.Digest kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BigDataDigests = nil
		} else {

			j.BigDataDigests = make(map[string]digest.Digest, 0)

			wantVal := true

			for {

				var k string

				var tmpJBigDataDigests digest.Digest

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJBigDataDigests type=digest.Digest kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Digest", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBigDataDigests = digest.Digest(string(outBuf))

					}
				}

				j.BigDataDigests[k] = tmpJBigDataDigests

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Created:

	/* handler: j.Created type=time.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Created.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Flags:

	/* handler: j.Flags type=map[string]interface {} kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Flags = nil
		} else {

			j.Flags = make(map[string]interface{}, 0)

			wantVal := true

			for {

				var k string

				var tmpJFlags interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJFlags type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJFlags)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Flags[k] = tmpJFlags

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *imageStore) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *imageStore) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{}`)
	return nil
}

const (
	ffjtimageStorebase = iota
	ffjtimageStorenosuchkey
)

// UnmarshalJSON umarshall json - template of ffjson
func (j *imageStore) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *imageStore) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtimageStorebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtimageStorenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				}

				currentKey = ffjtimageStorenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtimageStorenosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
