package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
)

// Pass this as an arg to the program
const filename = "image2.bmp"

type bmp struct {
	filename       string
	bmpType        string
	size           uint32
	pixelDataStart uint32
	dibHeaderSize  uint32
	headerName     string
	width          uint32
	height         uint32
	colorPlanes    uint16
	bitsPerPixel   uint16
	data           []byte
}

var headerMap = map[uint32]string{
	12:  "BITMAPCOREHEADER",
	64:  "OS22XBITMAPHEADER",
	16:  "OS22XBITMAPHEADER",
	40:  "BITMAPINFOHEADER",
	52:  "BITMAPV2INFOHEADER",
	56:  "BITMAPV3INFOHEADER",
	108: "BITMAPV4HEADER",
	124: "BITMAPV5HEADER",
}

func newBmpFromFile(filename string) (*bmp, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	bmpType := string(dat[0]) + string(dat[1])
	size := binary.LittleEndian.Uint32(dat[2:6])
	pixelDataStart := binary.LittleEndian.Uint32(dat[10:14])
	data := dat[pixelDataStart:]
	dibHeaderSize := binary.LittleEndian.Uint32(dat[14:18])

	headerName, ok := headerMap[dibHeaderSize]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Header name not found given dibHeaderSize %d", dibHeaderSize))
	}

	width := binary.LittleEndian.Uint32(dat[18:22])
	height := binary.LittleEndian.Uint32(dat[22:26])
	colorPlanes := binary.LittleEndian.Uint16(dat[26:28])
	bitsPerPixel := binary.LittleEndian.Uint16(dat[28:30])

	newBmp := bmp{
		filename,
		bmpType,
		size,
		pixelDataStart,
		dibHeaderSize,
		headerName,
		width,
		height,
		colorPlanes,
		bitsPerPixel,
		data,
	}

	return &newBmp, nil
}

type pixel struct {
	red   byte
	blue  byte
	green byte
}

func (p pixel) String() string {
	return fmt.Sprintf("{Red: %d Blue: %d Green: %d}", p.red, p.blue, p.green)
}

func main() {
	// pass filename as an arg to the program
	b, err := newBmpFromFile(filename)
	if err != nil {
		fmt.Printf("Error building bmp %s %s\n", filename, err)
	}
	output := []pixel{}
	for i := 3; i < len(b.data); i += 3 {
		output = append(output, pixel{blue: b.data[i-3], green: b.data[i-2], red: b.data[i-1]})
	}

	fmt.Println(output)
}
