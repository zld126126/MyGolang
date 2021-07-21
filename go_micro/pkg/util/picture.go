package util

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func testCreateImg() {
	fileAddress1 := "/Users/dong/Desktop/1.jpg"
	fileAddress2 := "/Users/dong/Desktop/2.png"
	createAddress := "/Users/dong/Desktop/3.jpg"
	err := CreateImg(fileAddress1, fileAddress2, createAddress)
	if err != nil {
		logrus.WithError(err).Println("create img error")
	}
}

func CreateImg(fileAddress1, fileAddress2, createAddress string) error {
	//背景图
	//如果是windows 换成c:/1.jpg
	backgroudImgFile, err := os.Open(fileAddress1)
	if err != nil {
		logrus.WithError(err).WithField("address", fileAddress1).Println("open file error")
		return errors.WithStack(err)
	}
	backgroudImg, err := jpeg.Decode(backgroudImgFile)
	if err != nil {
		logrus.WithError(err).WithField("address", fileAddress1).Println("decode file error")
		return errors.WithStack(err)
	}
	defer backgroudImgFile.Close()
	backgroudBound := backgroudImg.Bounds()
	//x轴坐标总数
	backgroudX := backgroudBound.Size().X
	//y轴坐标总数
	backgroudY := backgroudBound.Size().Y
	//添加图
	//如果是windows 换成c:/1.jpg
	centerImgFile, err := os.Open(fileAddress2)
	if err != nil {
		logrus.WithError(err).WithField("address", fileAddress2).Println("open file error")
		return errors.WithStack(err)
	}
	centerImg, err := png.Decode(centerImgFile)
	if err != nil {
		logrus.WithError(err).WithField("address", fileAddress2).Println("decode file error")
		return errors.WithStack(err)
	}
	defer centerImgFile.Close()
	centerBound := centerImg.Bounds()
	//x轴坐标总数
	centerX := centerBound.Size().X
	//y轴坐标总数
	centerY := centerBound.Size().Y

	//坐标偏差，x轴y轴 计算
	newImgX := (backgroudX - centerX) / 2
	newImgY := (backgroudY - centerY) / 2
	offset := image.Pt(newImgX, newImgY)
	//x轴坐标总数
	m := image.NewRGBA(backgroudBound)
	draw.Draw(m, backgroudBound, backgroudImg, image.ZP, draw.Src)
	draw.Draw(m, centerImg.Bounds().Add(offset), centerImg, image.ZP, draw.Over)
	//如果是windows 换成c:/1.jpg
	imgw, _ := os.Create(createAddress)
	err = jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	if err != nil {
		logrus.WithError(err).WithField("address", createAddress).Println("create file error")
		return errors.WithStack(err)
	}
	defer imgw.Close()
	return nil
}
