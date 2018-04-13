package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//打开原始图片
	imaB ,_ := os.Open("/Users/zhangyalei/Desktop/jj.png")
	//获取jpeg图片io对象
	img,_ := jpeg.Decode(imaB)
	//调用defer 在io文件结束时关闭
	defer imaB.Close()

	//打开水印出片
	wmb,_ := os.Open("/Users/zhangyalei/Desktop/ll.png")
	watermark,_ := png.Decode(wmb)

	defer wmb.Close()

	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10,img.Bounds().Dy()-watermark.Bounds().Dy()-10)

	b := img.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m,b,img,image.ZP,draw.Src)
	draw.Draw(m,watermark.Bounds().Add(offset),watermark,image.ZP,draw.Over)
	imgw,_:= os.Create("new.jpg")
	jpeg.Encode(imgw,m,&jpeg.Options{100})
	defer imgw.Close()
	fmt.Println("水印添加结束，请查看new.jpg图片")



}