// Code generated from TreeLan.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 332,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 6, 25, 301, 10, 25, 13, 25, 14, 25,
	302, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 7, 28, 313,
	10, 28, 12, 28, 14, 28, 316, 11, 28, 5, 28, 318, 10, 28, 3, 29, 3, 29,
	5, 29, 322, 10, 29, 3, 29, 3, 29, 3, 30, 6, 30, 327, 10, 30, 13, 30, 14,
	30, 328, 3, 30, 3, 30, 2, 2, 31, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 2, 57, 2, 59, 29, 3, 2, 8, 4, 2, 67, 92, 99, 124, 3,
	2, 51, 59, 3, 2, 50, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 334, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 63, 3, 2,
	2, 2, 7, 65, 3, 2, 2, 2, 9, 67, 3, 2, 2, 2, 11, 76, 3, 2, 2, 2, 13, 89,
	3, 2, 2, 2, 15, 104, 3, 2, 2, 2, 17, 113, 3, 2, 2, 2, 19, 125, 3, 2, 2,
	2, 21, 135, 3, 2, 2, 2, 23, 139, 3, 2, 2, 2, 25, 143, 3, 2, 2, 2, 27, 147,
	3, 2, 2, 2, 29, 162, 3, 2, 2, 2, 31, 171, 3, 2, 2, 2, 33, 176, 3, 2, 2,
	2, 35, 180, 3, 2, 2, 2, 37, 197, 3, 2, 2, 2, 39, 206, 3, 2, 2, 2, 41, 215,
	3, 2, 2, 2, 43, 235, 3, 2, 2, 2, 45, 255, 3, 2, 2, 2, 47, 283, 3, 2, 2,
	2, 49, 300, 3, 2, 2, 2, 51, 304, 3, 2, 2, 2, 53, 307, 3, 2, 2, 2, 55, 317,
	3, 2, 2, 2, 57, 319, 3, 2, 2, 2, 59, 326, 3, 2, 2, 2, 61, 62, 7, 42, 2,
	2, 62, 4, 3, 2, 2, 2, 63, 64, 7, 43, 2, 2, 64, 6, 3, 2, 2, 2, 65, 66, 7,
	97, 2, 2, 66, 8, 3, 2, 2, 2, 67, 68, 7, 67, 2, 2, 68, 69, 7, 70, 2, 2,
	69, 70, 7, 70, 2, 2, 70, 71, 7, 75, 2, 2, 71, 72, 7, 86, 2, 2, 72, 73,
	7, 75, 2, 2, 73, 74, 7, 81, 2, 2, 74, 75, 7, 80, 2, 2, 75, 10, 3, 2, 2,
	2, 76, 77, 7, 85, 2, 2, 77, 78, 7, 87, 2, 2, 78, 79, 7, 68, 2, 2, 79, 80,
	7, 85, 2, 2, 80, 81, 7, 86, 2, 2, 81, 82, 7, 84, 2, 2, 82, 83, 7, 67, 2,
	2, 83, 84, 7, 69, 2, 2, 84, 85, 7, 86, 2, 2, 85, 86, 7, 75, 2, 2, 86, 87,
	7, 81, 2, 2, 87, 88, 7, 80, 2, 2, 88, 12, 3, 2, 2, 2, 89, 90, 7, 79, 2,
	2, 90, 91, 7, 87, 2, 2, 91, 92, 7, 78, 2, 2, 92, 93, 7, 86, 2, 2, 93, 94,
	7, 75, 2, 2, 94, 95, 7, 82, 2, 2, 95, 96, 7, 78, 2, 2, 96, 97, 7, 75, 2,
	2, 97, 98, 7, 69, 2, 2, 98, 99, 7, 67, 2, 2, 99, 100, 7, 86, 2, 2, 100,
	101, 7, 75, 2, 2, 101, 102, 7, 81, 2, 2, 102, 103, 7, 80, 2, 2, 103, 14,
	3, 2, 2, 2, 104, 105, 7, 70, 2, 2, 105, 106, 7, 75, 2, 2, 106, 107, 7,
	88, 2, 2, 107, 108, 7, 75, 2, 2, 108, 109, 7, 85, 2, 2, 109, 110, 7, 75,
	2, 2, 110, 111, 7, 81, 2, 2, 111, 112, 7, 80, 2, 2, 112, 16, 3, 2, 2, 2,
	113, 114, 7, 71, 2, 2, 114, 115, 7, 90, 2, 2, 115, 116, 7, 82, 2, 2, 116,
	117, 7, 81, 2, 2, 117, 118, 7, 80, 2, 2, 118, 119, 7, 71, 2, 2, 119, 120,
	7, 80, 2, 2, 120, 121, 7, 86, 2, 2, 121, 122, 7, 75, 2, 2, 122, 123, 7,
	67, 2, 2, 123, 124, 7, 78, 2, 2, 124, 18, 3, 2, 2, 2, 125, 126, 7, 78,
	2, 2, 126, 127, 7, 81, 2, 2, 127, 128, 7, 73, 2, 2, 128, 129, 7, 67, 2,
	2, 129, 130, 7, 84, 2, 2, 130, 131, 7, 75, 2, 2, 131, 132, 7, 86, 2, 2,
	132, 133, 7, 74, 2, 2, 133, 134, 7, 79, 2, 2, 134, 20, 3, 2, 2, 2, 135,
	136, 7, 85, 2, 2, 136, 137, 7, 75, 2, 2, 137, 138, 7, 80, 2, 2, 138, 22,
	3, 2, 2, 2, 139, 140, 7, 69, 2, 2, 140, 141, 7, 81, 2, 2, 141, 142, 7,
	85, 2, 2, 142, 24, 3, 2, 2, 2, 143, 144, 7, 86, 2, 2, 144, 145, 7, 67,
	2, 2, 145, 146, 7, 80, 2, 2, 146, 26, 3, 2, 2, 2, 147, 148, 7, 86, 2, 2,
	148, 149, 7, 84, 2, 2, 149, 150, 7, 67, 2, 2, 150, 151, 7, 80, 2, 2, 151,
	152, 7, 85, 2, 2, 152, 153, 7, 72, 2, 2, 153, 154, 7, 81, 2, 2, 154, 155,
	7, 84, 2, 2, 155, 156, 7, 79, 2, 2, 156, 157, 7, 67, 2, 2, 157, 158, 7,
	86, 2, 2, 158, 159, 7, 75, 2, 2, 159, 160, 7, 81, 2, 2, 160, 161, 7, 80,
	2, 2, 161, 28, 3, 2, 2, 2, 162, 163, 7, 71, 2, 2, 163, 164, 7, 83, 2, 2,
	164, 165, 7, 87, 2, 2, 165, 166, 7, 67, 2, 2, 166, 167, 7, 86, 2, 2, 167,
	168, 7, 75, 2, 2, 168, 169, 7, 81, 2, 2, 169, 170, 7, 80, 2, 2, 170, 30,
	3, 2, 2, 2, 171, 172, 7, 78, 2, 2, 172, 173, 7, 71, 2, 2, 173, 174, 7,
	67, 2, 2, 174, 175, 7, 72, 2, 2, 175, 32, 3, 2, 2, 2, 176, 177, 7, 67,
	2, 2, 177, 178, 7, 80, 2, 2, 178, 179, 7, 91, 2, 2, 179, 34, 3, 2, 2, 2,
	180, 181, 7, 80, 2, 2, 181, 182, 7, 87, 2, 2, 182, 183, 7, 79, 2, 2, 183,
	184, 7, 71, 2, 2, 184, 185, 7, 84, 2, 2, 185, 186, 7, 75, 2, 2, 186, 187,
	7, 69, 2, 2, 187, 188, 7, 97, 2, 2, 188, 189, 7, 69, 2, 2, 189, 190, 7,
	81, 2, 2, 190, 191, 7, 80, 2, 2, 191, 192, 7, 85, 2, 2, 192, 193, 7, 86,
	2, 2, 193, 194, 7, 67, 2, 2, 194, 195, 7, 80, 2, 2, 195, 196, 7, 86, 2,
	2, 196, 36, 3, 2, 2, 2, 197, 198, 7, 69, 2, 2, 198, 199, 7, 81, 2, 2, 199,
	200, 7, 80, 2, 2, 200, 201, 7, 85, 2, 2, 201, 202, 7, 86, 2, 2, 202, 203,
	7, 67, 2, 2, 203, 204, 7, 80, 2, 2, 204, 205, 7, 86, 2, 2, 205, 38, 3,
	2, 2, 2, 206, 207, 7, 88, 2, 2, 207, 208, 7, 67, 2, 2, 208, 209, 7, 84,
	2, 2, 209, 210, 7, 75, 2, 2, 210, 211, 7, 67, 2, 2, 211, 212, 7, 68, 2,
	2, 212, 213, 7, 78, 2, 2, 213, 214, 7, 71, 2, 2, 214, 40, 3, 2, 2, 2, 215,
	216, 7, 69, 2, 2, 216, 217, 7, 81, 2, 2, 217, 218, 7, 80, 2, 2, 218, 219,
	7, 85, 2, 2, 219, 220, 7, 86, 2, 2, 220, 221, 7, 67, 2, 2, 221, 222, 7,
	80, 2, 2, 222, 223, 7, 86, 2, 2, 223, 224, 7, 97, 2, 2, 224, 225, 7, 71,
	2, 2, 225, 226, 7, 90, 2, 2, 226, 227, 7, 82, 2, 2, 227, 228, 7, 84, 2,
	2, 228, 229, 7, 71, 2, 2, 229, 230, 7, 85, 2, 2, 230, 231, 7, 85, 2, 2,
	231, 232, 7, 75, 2, 2, 232, 233, 7, 81, 2, 2, 233, 234, 7, 80, 2, 2, 234,
	42, 3, 2, 2, 2, 235, 236, 7, 88, 2, 2, 236, 237, 7, 67, 2, 2, 237, 238,
	7, 84, 2, 2, 238, 239, 7, 75, 2, 2, 239, 240, 7, 67, 2, 2, 240, 241, 7,
	68, 2, 2, 241, 242, 7, 78, 2, 2, 242, 243, 7, 71, 2, 2, 243, 244, 7, 97,
	2, 2, 244, 245, 7, 71, 2, 2, 245, 246, 7, 90, 2, 2, 246, 247, 7, 82, 2,
	2, 247, 248, 7, 84, 2, 2, 248, 249, 7, 71, 2, 2, 249, 250, 7, 85, 2, 2,
	250, 251, 7, 85, 2, 2, 251, 252, 7, 75, 2, 2, 252, 253, 7, 81, 2, 2, 253,
	254, 7, 80, 2, 2, 254, 44, 3, 2, 2, 2, 255, 256, 7, 80, 2, 2, 256, 257,
	7, 87, 2, 2, 257, 258, 7, 79, 2, 2, 258, 259, 7, 71, 2, 2, 259, 260, 7,
	84, 2, 2, 260, 261, 7, 75, 2, 2, 261, 262, 7, 69, 2, 2, 262, 263, 7, 97,
	2, 2, 263, 264, 7, 69, 2, 2, 264, 265, 7, 81, 2, 2, 265, 266, 7, 80, 2,
	2, 266, 267, 7, 85, 2, 2, 267, 268, 7, 86, 2, 2, 268, 269, 7, 67, 2, 2,
	269, 270, 7, 80, 2, 2, 270, 271, 7, 86, 2, 2, 271, 272, 7, 97, 2, 2, 272,
	273, 7, 71, 2, 2, 273, 274, 7, 90, 2, 2, 274, 275, 7, 82, 2, 2, 275, 276,
	7, 84, 2, 2, 276, 277, 7, 71, 2, 2, 277, 278, 7, 85, 2, 2, 278, 279, 7,
	85, 2, 2, 279, 280, 7, 75, 2, 2, 280, 281, 7, 81, 2, 2, 281, 282, 7, 80,
	2, 2, 282, 46, 3, 2, 2, 2, 283, 284, 7, 85, 2, 2, 284, 285, 7, 91, 2, 2,
	285, 286, 7, 85, 2, 2, 286, 287, 7, 86, 2, 2, 287, 288, 7, 71, 2, 2, 288,
	289, 7, 79, 2, 2, 289, 290, 7, 97, 2, 2, 290, 291, 7, 72, 2, 2, 291, 292,
	7, 87, 2, 2, 292, 293, 7, 80, 2, 2, 293, 294, 7, 69, 2, 2, 294, 295, 7,
	86, 2, 2, 295, 296, 7, 75, 2, 2, 296, 297, 7, 81, 2, 2, 297, 298, 7, 80,
	2, 2, 298, 48, 3, 2, 2, 2, 299, 301, 9, 2, 2, 2, 300, 299, 3, 2, 2, 2,
	301, 302, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	50, 3, 2, 2, 2, 304, 305, 7, 63, 2, 2, 305, 306, 7, 64, 2, 2, 306, 52,
	3, 2, 2, 2, 307, 308, 5, 55, 28, 2, 308, 54, 3, 2, 2, 2, 309, 318, 7, 50,
	2, 2, 310, 314, 9, 3, 2, 2, 311, 313, 9, 4, 2, 2, 312, 311, 3, 2, 2, 2,
	313, 316, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315,
	318, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 317, 309, 3, 2, 2, 2, 317, 310,
	3, 2, 2, 2, 318, 56, 3, 2, 2, 2, 319, 321, 9, 5, 2, 2, 320, 322, 9, 6,
	2, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2,
	323, 324, 5, 55, 28, 2, 324, 58, 3, 2, 2, 2, 325, 327, 9, 7, 2, 2, 326,
	325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329,
	3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 8, 30, 2, 2, 331, 60, 3, 2,
	2, 2, 8, 2, 302, 314, 317, 321, 328, 3, 8, 2, 2,
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
	"", "'('", "')'", "'_'", "'ADDITION'", "'SUBSTRACTION'", "'MULTIPLICATION'",
	"'DIVISION'", "'EXPONENTIAL'", "'LOGARITHM'", "'SIN'", "'COS'", "'TAN'",
	"'TRANSFORMATION'", "'EQUATION'", "'LEAF'", "'ANY'", "'NUMERIC_CONSTANT'",
	"'CONSTANT'", "'VARIABLE'", "'CONSTANT_EXPRESSION'", "'VARIABLE_EXPRESSION'",
	"'NUMERIC_CONSTANT_EXPRESSION'", "'SYSTEM_FUNCTION'", "", "'=>'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "ADDITION", "SUBSTRACTION", "MULTIPLICATION", "DIVISION",
	"EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", "EQUATION",
	"LEAF", "ANY", "NUMERIC_CONSTANT", "CONSTANT", "VARIABLE", "CONSTANT_EXPRESSION",
	"VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", "SYSTEM_FUNCTION",
	"ID", "ASSIGN", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "ADDITION", "SUBSTRACTION", "MULTIPLICATION", "DIVISION",
	"EXPONENTIAL", "LOGARITHM", "SIN", "COS", "TAN", "TRANSFORMATION", "EQUATION",
	"LEAF", "ANY", "NUMERIC_CONSTANT", "CONSTANT", "VARIABLE", "CONSTANT_EXPRESSION",
	"VARIABLE_EXPRESSION", "NUMERIC_CONSTANT_EXPRESSION", "SYSTEM_FUNCTION",
	"ID", "ASSIGN", "NUMBER", "INT", "EXP", "WS",
}

type TreeLanLexer struct {
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

func NewTreeLanLexer(input antlr.CharStream) *TreeLanLexer {

	l := new(TreeLanLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TreeLan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TreeLanLexer tokens.
const (
	TreeLanLexerT__0                        = 1
	TreeLanLexerT__1                        = 2
	TreeLanLexerT__2                        = 3
	TreeLanLexerADDITION                    = 4
	TreeLanLexerSUBSTRACTION                = 5
	TreeLanLexerMULTIPLICATION              = 6
	TreeLanLexerDIVISION                    = 7
	TreeLanLexerEXPONENTIAL                 = 8
	TreeLanLexerLOGARITHM                   = 9
	TreeLanLexerSIN                         = 10
	TreeLanLexerCOS                         = 11
	TreeLanLexerTAN                         = 12
	TreeLanLexerTRANSFORMATION              = 13
	TreeLanLexerEQUATION                    = 14
	TreeLanLexerLEAF                        = 15
	TreeLanLexerANY                         = 16
	TreeLanLexerNUMERIC_CONSTANT            = 17
	TreeLanLexerCONSTANT                    = 18
	TreeLanLexerVARIABLE                    = 19
	TreeLanLexerCONSTANT_EXPRESSION         = 20
	TreeLanLexerVARIABLE_EXPRESSION         = 21
	TreeLanLexerNUMERIC_CONSTANT_EXPRESSION = 22
	TreeLanLexerSYSTEM_FUNCTION             = 23
	TreeLanLexerID                          = 24
	TreeLanLexerASSIGN                      = 25
	TreeLanLexerNUMBER                      = 26
	TreeLanLexerWS                          = 27
)
