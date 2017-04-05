// +build example
//
// Do not build by default.

package main

import (
	"bytes"
	"flag"
	"image"
	"image/draw"
	"image/png"
	"log"

	. "github.com/peterhellberg/microview"
)

func main() {
	name := flag.String("name", "/dev/cu.usbserial-DA00SSM3", "name of the serial port")

	flag.Parse()

	mv, err := OpenMicroView(*name)
	if err != nil {
		log.Fatal(err)
	}

	gopher, err := png.Decode(bytes.NewReader(gopherPNGBytes))
	if err != nil {
		log.Fatal(err)
	}

	draw.Draw(mv, mv.Bounds(), gopher, image.ZP, draw.Src)
}

var gopherPNGBytes = []byte{
	137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 64,
	0, 0, 0, 48, 8, 6, 0, 0, 0, 161, 75, 124, 31, 0, 0, 0, 1, 115, 82, 71, 66, 0,
	174, 206, 28, 233, 0, 0, 1, 89, 105, 84, 88, 116, 88, 77, 76, 58, 99, 111, 109,
	46, 97, 100, 111, 98, 101, 46, 120, 109, 112, 0, 0, 0, 0, 0, 60, 120, 58, 120,
	109, 112, 109, 101, 116, 97, 32, 120, 109, 108, 110, 115, 58, 120, 61, 34, 97,
	100, 111, 98, 101, 58, 110, 115, 58, 109, 101, 116, 97, 47, 34, 32, 120, 58,
	120, 109, 112, 116, 107, 61, 34, 88, 77, 80, 32, 67, 111, 114, 101, 32, 53, 46,
	52, 46, 48, 34, 62, 10, 32, 32, 32, 60, 114, 100, 102, 58, 82, 68, 70, 32,
	120, 109, 108, 110, 115, 58, 114, 100, 102, 61, 34, 104, 116, 116, 112, 58, 47,
	47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 49, 57, 57, 57, 47, 48,
	50, 47, 50, 50, 45, 114, 100, 102, 45, 115, 121, 110, 116, 97, 120, 45, 110,
	115, 35, 34, 62, 10, 32, 32, 32, 32, 32, 32, 60, 114, 100, 102, 58, 68, 101,
	115, 99, 114, 105, 112, 116, 105, 111, 110, 32, 114, 100, 102, 58, 97, 98, 111,
	117, 116, 61, 34, 34, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 120,
	109, 108, 110, 115, 58, 116, 105, 102, 102, 61, 34, 104, 116, 116, 112, 58, 47,
	47, 110, 115, 46, 97, 100, 111, 98, 101, 46, 99, 111, 109, 47, 116, 105, 102,
	102, 47, 49, 46, 48, 47, 34, 62, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 60,
	116, 105, 102, 102, 58, 79, 114, 105, 101, 110, 116, 97, 116, 105, 111, 110,
	62, 49, 60, 47, 116, 105, 102, 102, 58, 79, 114, 105, 101, 110, 116, 97, 116,
	105, 111, 110, 62, 10, 32, 32, 32, 32, 32, 32, 60, 47, 114, 100, 102, 58, 68,
	101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 62, 10, 32, 32, 32, 60, 47,
	114, 100, 102, 58, 82, 68, 70, 62, 10, 60, 47, 120, 58, 120, 109, 112, 109,
	101, 116, 97, 62, 10, 76, 194, 39, 89, 0, 0, 1, 99, 73, 68, 65, 84, 104, 222,
	237, 154, 217, 14, 132, 48, 8, 69, 141, 154, 56, 255, 255, 195, 53, 62, 152,
	24, 66, 129, 118, 186, 0, 165, 73, 51, 47, 90, 185, 103, 128, 110, 108, 91,
	180, 104, 110, 218, 239, 60, 211, 219, 151, 16, 9, 5, 195, 95, 151, 98, 185,
	103, 205, 3, 184, 142, 35, 181, 112, 101, 115, 161, 208, 43, 126, 171, 199, 28,
	65, 112, 100, 210, 42, 250, 198, 247, 225, 156, 129, 185, 100, 244, 244, 199,
	141, 115, 46, 61, 51, 83, 139, 190, 41, 21, 43, 137, 229, 183, 107, 10, 51,
	214, 158, 111, 38, 253, 10, 213, 36, 164, 171, 23, 184, 155, 75, 107, 195, 192,
	43, 0, 81, 88, 142, 22, 207, 37, 210, 105, 97, 48, 211, 37, 123, 218, 16, 0,
	168, 177, 103, 185, 255, 168, 28, 196, 78, 229, 179, 22, 40, 35, 226, 31, 46,
	240, 84, 136, 159, 53, 5, 6, 128, 0, 16, 0, 210, 82, 226, 217, 165, 254, 42, 0,
	138, 221, 223, 227, 73, 171, 8, 0, 182, 29, 246, 0, 66, 180, 0, 162, 132, 190,
	32, 220, 3, 104, 178, 175, 182, 8, 160, 228, 31, 182, 8, 129, 5, 80, 34, 202,
	21, 128, 218, 93, 152, 37, 8, 162, 115, 128, 0, 176, 50, 0, 12, 2, 183, 71,
	183, 2, 64, 124, 71, 1, 111, 89, 185, 11, 17, 43, 0, 138, 175, 198, 48, 47, 88,
	14, 0, 230, 9, 212, 179, 90, 97, 84, 217, 150, 59, 171, 167, 4, 187, 2, 0, 95,
	196, 54, 67, 112, 208, 86, 197, 13, 61, 196, 87, 219, 67, 189, 44, 41, 91, 81,
	123, 0, 210, 34, 121, 96, 116, 53, 93, 174, 54, 169, 27, 146, 2, 200, 93, 171,
	107, 112, 253, 191, 108, 226, 118, 135, 26, 42, 63, 154, 37, 61, 73, 28, 213,
	148, 204, 72, 250, 181, 239, 73, 157, 112, 56, 56, 53, 13, 98, 51, 66, 73, 153,
	140, 24, 212, 204, 154, 35, 202, 176, 45, 90, 52, 147, 237, 6, 172, 193, 89,
	224, 237, 243, 250, 142, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130,
}
