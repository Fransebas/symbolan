// Code generated from Symbolan.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 382,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 6, 28, 181, 10, 28, 13, 28, 14,
	28, 182, 5, 28, 185, 10, 28, 3, 29, 3, 29, 3, 29, 6, 29, 190, 10, 29, 13,
	29, 14, 29, 191, 5, 29, 194, 10, 29, 3, 30, 3, 30, 3, 30, 6, 30, 199, 10,
	30, 13, 30, 14, 30, 200, 5, 30, 203, 10, 30, 3, 31, 3, 31, 3, 32, 5, 32,
	208, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 6, 34, 224, 10, 34, 13, 34, 14,
	34, 225, 5, 34, 228, 10, 34, 3, 35, 5, 35, 231, 10, 35, 3, 35, 3, 35, 3,
	35, 6, 35, 236, 10, 35, 13, 35, 14, 35, 237, 5, 35, 240, 10, 35, 3, 36,
	5, 36, 243, 10, 36, 3, 36, 3, 36, 3, 36, 6, 36, 248, 10, 36, 13, 36, 14,
	36, 249, 5, 36, 252, 10, 36, 3, 37, 5, 37, 255, 10, 37, 3, 37, 3, 37, 3,
	37, 6, 37, 260, 10, 37, 13, 37, 14, 37, 261, 5, 37, 264, 10, 37, 3, 38,
	5, 38, 267, 10, 38, 3, 38, 3, 38, 3, 38, 6, 38, 272, 10, 38, 13, 38, 14,
	38, 273, 5, 38, 276, 10, 38, 3, 39, 3, 39, 5, 39, 280, 10, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 7, 40, 286, 10, 40, 12, 40, 14, 40, 289, 11, 40, 3, 41,
	3, 41, 7, 41, 293, 10, 41, 12, 41, 14, 41, 296, 11, 41, 3, 42, 3, 42, 3,
	42, 6, 42, 301, 10, 42, 13, 42, 14, 42, 302, 3, 43, 3, 43, 3, 43, 5, 43,
	308, 10, 43, 3, 43, 5, 43, 311, 10, 43, 3, 43, 5, 43, 314, 10, 43, 3, 43,
	3, 43, 3, 43, 5, 43, 319, 10, 43, 5, 43, 321, 10, 43, 3, 44, 3, 44, 5,
	44, 325, 10, 44, 3, 44, 3, 44, 3, 45, 3, 45, 7, 45, 331, 10, 45, 12, 45,
	14, 45, 334, 11, 45, 3, 46, 5, 46, 337, 10, 46, 3, 46, 6, 46, 340, 10,
	46, 13, 46, 14, 46, 341, 3, 47, 5, 47, 345, 10, 47, 3, 47, 3, 47, 3, 48,
	5, 48, 350, 10, 48, 3, 48, 3, 48, 3, 49, 3, 49, 5, 49, 356, 10, 49, 3,
	49, 3, 49, 3, 50, 5, 50, 361, 10, 50, 3, 50, 3, 50, 3, 51, 6, 51, 366,
	10, 51, 13, 51, 14, 51, 367, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52,
	7, 52, 376, 10, 52, 12, 52, 14, 52, 379, 11, 52, 3, 52, 3, 52, 2, 2, 53,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77,
	40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 2, 93, 2, 95, 2,
	97, 2, 99, 47, 101, 48, 103, 49, 3, 2, 20, 3, 2, 122, 124, 3, 2, 50, 59,
	3, 2, 99, 101, 3, 2, 103, 103, 4, 2, 45, 45, 47, 47, 3, 2, 88, 88, 3, 2,
	80, 80, 3, 2, 69, 69, 3, 2, 72, 72, 3, 2, 51, 59, 4, 2, 90, 90, 122, 122,
	3, 2, 67, 92, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 57, 5, 2,
	50, 59, 67, 72, 99, 104, 4, 2, 71, 71, 103, 103, 4, 2, 11, 11, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 417, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53,
	3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2,
	61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2,
	2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2,
	2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2,
	2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 99, 3,
	2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 3, 105, 3, 2, 2, 2, 5,
	108, 3, 2, 2, 2, 7, 111, 3, 2, 2, 2, 9, 113, 3, 2, 2, 2, 11, 116, 3, 2,
	2, 2, 13, 118, 3, 2, 2, 2, 15, 120, 3, 2, 2, 2, 17, 122, 3, 2, 2, 2, 19,
	124, 3, 2, 2, 2, 21, 126, 3, 2, 2, 2, 23, 128, 3, 2, 2, 2, 25, 130, 3,
	2, 2, 2, 27, 132, 3, 2, 2, 2, 29, 134, 3, 2, 2, 2, 31, 136, 3, 2, 2, 2,
	33, 138, 3, 2, 2, 2, 35, 140, 3, 2, 2, 2, 37, 142, 3, 2, 2, 2, 39, 146,
	3, 2, 2, 2, 41, 150, 3, 2, 2, 2, 43, 154, 3, 2, 2, 2, 45, 158, 3, 2, 2,
	2, 47, 162, 3, 2, 2, 2, 49, 166, 3, 2, 2, 2, 51, 171, 3, 2, 2, 2, 53, 173,
	3, 2, 2, 2, 55, 176, 3, 2, 2, 2, 57, 186, 3, 2, 2, 2, 59, 195, 3, 2, 2,
	2, 61, 204, 3, 2, 2, 2, 63, 207, 3, 2, 2, 2, 65, 213, 3, 2, 2, 2, 67, 219,
	3, 2, 2, 2, 69, 230, 3, 2, 2, 2, 71, 242, 3, 2, 2, 2, 73, 254, 3, 2, 2,
	2, 75, 266, 3, 2, 2, 2, 77, 279, 3, 2, 2, 2, 79, 283, 3, 2, 2, 2, 81, 290,
	3, 2, 2, 2, 83, 297, 3, 2, 2, 2, 85, 320, 3, 2, 2, 2, 87, 324, 3, 2, 2,
	2, 89, 328, 3, 2, 2, 2, 91, 336, 3, 2, 2, 2, 93, 344, 3, 2, 2, 2, 95, 349,
	3, 2, 2, 2, 97, 353, 3, 2, 2, 2, 99, 360, 3, 2, 2, 2, 101, 365, 3, 2, 2,
	2, 103, 371, 3, 2, 2, 2, 105, 106, 7, 42, 2, 2, 106, 107, 7, 43, 2, 2,
	107, 4, 3, 2, 2, 2, 108, 109, 7, 63, 2, 2, 109, 110, 7, 64, 2, 2, 110,
	6, 3, 2, 2, 2, 111, 112, 7, 96, 2, 2, 112, 8, 3, 2, 2, 2, 113, 114, 7,
	44, 2, 2, 114, 115, 7, 44, 2, 2, 115, 10, 3, 2, 2, 2, 116, 117, 7, 44,
	2, 2, 117, 12, 3, 2, 2, 2, 118, 119, 7, 49, 2, 2, 119, 14, 3, 2, 2, 2,
	120, 121, 7, 45, 2, 2, 121, 16, 3, 2, 2, 2, 122, 123, 7, 47, 2, 2, 123,
	18, 3, 2, 2, 2, 124, 125, 7, 42, 2, 2, 125, 20, 3, 2, 2, 2, 126, 127, 7,
	43, 2, 2, 127, 22, 3, 2, 2, 2, 128, 129, 7, 125, 2, 2, 129, 24, 3, 2, 2,
	2, 130, 131, 7, 127, 2, 2, 131, 26, 3, 2, 2, 2, 132, 133, 7, 93, 2, 2,
	133, 28, 3, 2, 2, 2, 134, 135, 7, 95, 2, 2, 135, 30, 3, 2, 2, 2, 136, 137,
	7, 63, 2, 2, 137, 32, 3, 2, 2, 2, 138, 139, 7, 46, 2, 2, 139, 34, 3, 2,
	2, 2, 140, 141, 7, 35, 2, 2, 141, 36, 3, 2, 2, 2, 142, 143, 7, 111, 2,
	2, 143, 144, 7, 107, 2, 2, 144, 145, 7, 112, 2, 2, 145, 38, 3, 2, 2, 2,
	146, 147, 7, 111, 2, 2, 147, 148, 7, 99, 2, 2, 148, 149, 7, 122, 2, 2,
	149, 40, 3, 2, 2, 2, 150, 151, 7, 117, 2, 2, 151, 152, 7, 107, 2, 2, 152,
	153, 7, 112, 2, 2, 153, 42, 3, 2, 2, 2, 154, 155, 7, 101, 2, 2, 155, 156,
	7, 113, 2, 2, 156, 157, 7, 117, 2, 2, 157, 44, 3, 2, 2, 2, 158, 159, 7,
	118, 2, 2, 159, 160, 7, 99, 2, 2, 160, 161, 7, 112, 2, 2, 161, 46, 3, 2,
	2, 2, 162, 163, 7, 110, 2, 2, 163, 164, 7, 113, 2, 2, 164, 165, 7, 105,
	2, 2, 165, 48, 3, 2, 2, 2, 166, 167, 7, 117, 2, 2, 167, 168, 7, 115, 2,
	2, 168, 169, 7, 116, 2, 2, 169, 170, 7, 118, 2, 2, 170, 50, 3, 2, 2, 2,
	171, 172, 7, 70, 2, 2, 172, 52, 3, 2, 2, 2, 173, 174, 7, 114, 2, 2, 174,
	175, 7, 107, 2, 2, 175, 54, 3, 2, 2, 2, 176, 177, 7, 102, 2, 2, 177, 184,
	9, 2, 2, 2, 178, 180, 7, 97, 2, 2, 179, 181, 9, 3, 2, 2, 180, 179, 3, 2,
	2, 2, 181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2,
	183, 185, 3, 2, 2, 2, 184, 178, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185,
	56, 3, 2, 2, 2, 186, 193, 9, 2, 2, 2, 187, 189, 7, 97, 2, 2, 188, 190,
	9, 3, 2, 2, 189, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 189, 3, 2,
	2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2, 193, 187, 3, 2, 2, 2,
	193, 194, 3, 2, 2, 2, 194, 58, 3, 2, 2, 2, 195, 202, 9, 4, 2, 2, 196, 198,
	7, 97, 2, 2, 197, 199, 9, 3, 2, 2, 198, 197, 3, 2, 2, 2, 199, 200, 3, 2,
	2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 3, 2, 2, 2,
	202, 196, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 60, 3, 2, 2, 2, 204, 205,
	9, 5, 2, 2, 205, 62, 3, 2, 2, 2, 206, 208, 9, 6, 2, 2, 207, 206, 3, 2,
	2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 7, 107, 2,
	2, 210, 211, 7, 112, 2, 2, 211, 212, 7, 104, 2, 2, 212, 64, 3, 2, 2, 2,
	213, 214, 7, 119, 2, 2, 214, 215, 7, 112, 2, 2, 215, 216, 7, 102, 2, 2,
	216, 217, 7, 103, 2, 2, 217, 218, 7, 104, 2, 2, 218, 66, 3, 2, 2, 2, 219,
	220, 7, 102, 2, 2, 220, 227, 9, 7, 2, 2, 221, 223, 7, 97, 2, 2, 222, 224,
	9, 3, 2, 2, 223, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 223, 3, 2,
	2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 221, 3, 2, 2, 2,
	227, 228, 3, 2, 2, 2, 228, 68, 3, 2, 2, 2, 229, 231, 9, 6, 2, 2, 230, 229,
	3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 239, 9, 8,
	2, 2, 233, 235, 7, 97, 2, 2, 234, 236, 9, 3, 2, 2, 235, 234, 3, 2, 2, 2,
	236, 237, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	240, 3, 2, 2, 2, 239, 233, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 70, 3,
	2, 2, 2, 241, 243, 9, 6, 2, 2, 242, 241, 3, 2, 2, 2, 242, 243, 3, 2, 2,
	2, 243, 244, 3, 2, 2, 2, 244, 251, 9, 9, 2, 2, 245, 247, 7, 97, 2, 2, 246,
	248, 9, 3, 2, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 247,
	3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251, 245, 3, 2,
	2, 2, 251, 252, 3, 2, 2, 2, 252, 72, 3, 2, 2, 2, 253, 255, 9, 6, 2, 2,
	254, 253, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	263, 9, 7, 2, 2, 257, 259, 7, 97, 2, 2, 258, 260, 9, 3, 2, 2, 259, 258,
	3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2,
	2, 2, 262, 264, 3, 2, 2, 2, 263, 257, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2,
	264, 74, 3, 2, 2, 2, 265, 267, 9, 6, 2, 2, 266, 265, 3, 2, 2, 2, 266, 267,
	3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 275, 9, 10, 2, 2, 269, 271, 7, 97,
	2, 2, 270, 272, 9, 3, 2, 2, 271, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2,
	273, 271, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 276, 3, 2, 2, 2, 275,
	269, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 76, 3, 2, 2, 2, 277, 280, 5,
	91, 46, 2, 278, 280, 5, 85, 43, 2, 279, 277, 3, 2, 2, 2, 279, 278, 3, 2,
	2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 7, 107, 2, 2, 282, 78, 3, 2, 2, 2,
	283, 287, 9, 11, 2, 2, 284, 286, 9, 3, 2, 2, 285, 284, 3, 2, 2, 2, 286,
	289, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 80, 3,
	2, 2, 2, 289, 287, 3, 2, 2, 2, 290, 294, 7, 50, 2, 2, 291, 293, 5, 93,
	47, 2, 292, 291, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2,
	294, 295, 3, 2, 2, 2, 295, 82, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297, 298,
	7, 50, 2, 2, 298, 300, 9, 12, 2, 2, 299, 301, 5, 95, 48, 2, 300, 299, 3,
	2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2,
	2, 303, 84, 3, 2, 2, 2, 304, 313, 5, 91, 46, 2, 305, 307, 7, 48, 2, 2,
	306, 308, 5, 91, 46, 2, 307, 306, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308,
	310, 3, 2, 2, 2, 309, 311, 5, 97, 49, 2, 310, 309, 3, 2, 2, 2, 310, 311,
	3, 2, 2, 2, 311, 314, 3, 2, 2, 2, 312, 314, 5, 97, 49, 2, 313, 305, 3,
	2, 2, 2, 313, 312, 3, 2, 2, 2, 314, 321, 3, 2, 2, 2, 315, 316, 7, 48, 2,
	2, 316, 318, 5, 91, 46, 2, 317, 319, 5, 97, 49, 2, 318, 317, 3, 2, 2, 2,
	318, 319, 3, 2, 2, 2, 319, 321, 3, 2, 2, 2, 320, 304, 3, 2, 2, 2, 320,
	315, 3, 2, 2, 2, 321, 86, 3, 2, 2, 2, 322, 325, 5, 91, 46, 2, 323, 325,
	5, 85, 43, 2, 324, 322, 3, 2, 2, 2, 324, 323, 3, 2, 2, 2, 325, 326, 3,
	2, 2, 2, 326, 327, 7, 107, 2, 2, 327, 88, 3, 2, 2, 2, 328, 332, 9, 13,
	2, 2, 329, 331, 9, 14, 2, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2,
	332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 90, 3, 2, 2, 2, 334, 332,
	3, 2, 2, 2, 335, 337, 9, 6, 2, 2, 336, 335, 3, 2, 2, 2, 336, 337, 3, 2,
	2, 2, 337, 339, 3, 2, 2, 2, 338, 340, 9, 3, 2, 2, 339, 338, 3, 2, 2, 2,
	340, 341, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342,
	92, 3, 2, 2, 2, 343, 345, 9, 6, 2, 2, 344, 343, 3, 2, 2, 2, 344, 345, 3,
	2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 347, 9, 15, 2, 2, 347, 94, 3, 2, 2,
	2, 348, 350, 9, 6, 2, 2, 349, 348, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350,
	351, 3, 2, 2, 2, 351, 352, 9, 16, 2, 2, 352, 96, 3, 2, 2, 2, 353, 355,
	9, 17, 2, 2, 354, 356, 9, 6, 2, 2, 355, 354, 3, 2, 2, 2, 355, 356, 3, 2,
	2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 5, 91, 46, 2, 358, 98, 3, 2, 2, 2,
	359, 361, 7, 15, 2, 2, 360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361,
	362, 3, 2, 2, 2, 362, 363, 7, 12, 2, 2, 363, 100, 3, 2, 2, 2, 364, 366,
	9, 18, 2, 2, 365, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 365, 3, 2,
	2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 8, 51, 2, 2,
	370, 102, 3, 2, 2, 2, 371, 372, 7, 49, 2, 2, 372, 373, 7, 49, 2, 2, 373,
	377, 3, 2, 2, 2, 374, 376, 10, 19, 2, 2, 375, 374, 3, 2, 2, 2, 376, 379,
	3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 380, 3, 2,
	2, 2, 379, 377, 3, 2, 2, 2, 380, 381, 8, 52, 3, 2, 381, 104, 3, 2, 2, 2,
	43, 2, 182, 184, 191, 193, 200, 202, 207, 225, 227, 230, 237, 239, 242,
	249, 251, 254, 261, 263, 266, 273, 275, 279, 287, 294, 302, 307, 310, 313,
	318, 320, 324, 332, 336, 341, 344, 349, 355, 360, 367, 377, 4, 8, 2, 2,
	2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'()'", "'=>'", "'^'", "'**'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "'='", "','", "'!'", "'min'", "'max'", "'sin'",
	"'cos'", "'tan'", "'log'", "'sqrt'", "'D'", "'pi'", "", "", "", "", "",
	"'undef'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
	"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "FACTORIAL", "MIN", "MAX",
	"SIN", "COS", "TAN", "LOG", "SQRT", "DERIVATIVE_OPERATOR", "PI", "DERIVATIVE",
	"VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", "DERIVATE_RULE", "NUMERIC_RULE",
	"CONST_RULE", "VAR_RULE", "EXPR_RULE", "IMAGINARY", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "IDENTIFIER", "NEWLINE", "WS",
	"LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "L_PAREN",
	"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA",
	"FACTORIAL", "MIN", "MAX", "SIN", "COS", "TAN", "LOG", "SQRT", "DERIVATIVE_OPERATOR",
	"PI", "DERIVATIVE", "VAR", "CONST", "EULER", "INFINITE", "UNDEFINED", "DERIVATE_RULE",
	"NUMERIC_RULE", "CONST_RULE", "VAR_RULE", "EXPR_RULE", "IMAGINARY", "DECIMAL_LIT",
	"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "IDENTIFIER", "DECIMALS",
	"OCTAL_DIGIT", "HEX_DIGIT", "EXPONENT", "NEWLINE", "WS", "LINE_COMMENT",
}

type SymbolanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSymbolanLexer(input antlr.CharStream) *SymbolanLexer {

	l := new(SymbolanLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Symbolan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SymbolanLexer tokens.
const (
	SymbolanLexerT__0                = 1
	SymbolanLexerT__1                = 2
	SymbolanLexerT__2                = 3
	SymbolanLexerT__3                = 4
	SymbolanLexerT__4                = 5
	SymbolanLexerT__5                = 6
	SymbolanLexerT__6                = 7
	SymbolanLexerT__7                = 8
	SymbolanLexerL_PAREN             = 9
	SymbolanLexerR_PAREN             = 10
	SymbolanLexerL_CURLY             = 11
	SymbolanLexerR_CURLY             = 12
	SymbolanLexerL_BRACKET           = 13
	SymbolanLexerR_BRACKET           = 14
	SymbolanLexerASSIGN              = 15
	SymbolanLexerCOMMA               = 16
	SymbolanLexerFACTORIAL           = 17
	SymbolanLexerMIN                 = 18
	SymbolanLexerMAX                 = 19
	SymbolanLexerSIN                 = 20
	SymbolanLexerCOS                 = 21
	SymbolanLexerTAN                 = 22
	SymbolanLexerLOG                 = 23
	SymbolanLexerSQRT                = 24
	SymbolanLexerDERIVATIVE_OPERATOR = 25
	SymbolanLexerPI                  = 26
	SymbolanLexerDERIVATIVE          = 27
	SymbolanLexerVAR                 = 28
	SymbolanLexerCONST               = 29
	SymbolanLexerEULER               = 30
	SymbolanLexerINFINITE            = 31
	SymbolanLexerUNDEFINED           = 32
	SymbolanLexerDERIVATE_RULE       = 33
	SymbolanLexerNUMERIC_RULE        = 34
	SymbolanLexerCONST_RULE          = 35
	SymbolanLexerVAR_RULE            = 36
	SymbolanLexerEXPR_RULE           = 37
	SymbolanLexerIMAGINARY           = 38
	SymbolanLexerDECIMAL_LIT         = 39
	SymbolanLexerOCTAL_LIT           = 40
	SymbolanLexerHEX_LIT             = 41
	SymbolanLexerFLOAT_LIT           = 42
	SymbolanLexerIMAGINARY_LIT       = 43
	SymbolanLexerIDENTIFIER          = 44
	SymbolanLexerNEWLINE             = 45
	SymbolanLexerWS                  = 46
	SymbolanLexerLINE_COMMENT        = 47
)
