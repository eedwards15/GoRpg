package ui

import "golang.org/x/image/font"

func CenterTextXPos(value string, fontUsed font.Face, screenWidth int) int {
	return (screenWidth - (font.MeasureString(fontUsed, value).Floor())) / 2
}

func RightAligned(value string, fontUsed font.Face, screenWidth int) int {
	return screenWidth - (font.MeasureString(fontUsed, value).Floor())
}
