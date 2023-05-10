package uploader

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func InsertImage(tmpPrevFullPath string, vidPath string, genVidPth string) error {
	fmt.Println("-------------------------")
	fmt.Println(tmpPrevFullPath)
	fmt.Println(vidPath)
	fmt.Println("-------------------------")
	//ffmpeg -i /home/geek/Videos/in_ok.mp4 -i /home/geek/Videos/in_ok.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic out3.mp4
	app := "/usr/bin/ffmpeg"

	par0 := "-y"

	par1Key := "-i"
	par1Val := vidPath //"in.mp4" inVideoFileName
	//par1Val := vidPath //"in.mp4" inVideoFileName

	par2Key := "-i"
	par2Val := tmpPrevFullPath //"in.png" //inThumbFileName generatingPreview("video.png", "00:03:00")

	par3Key := "-map"
	par3Val := "1"

	par4Key := "-map"
	par4Val := "0"

	par5Key := "-c"
	par5Val := "copy"

	par6 := "-disposition:0"

	par7Key := "attached_pic"
	par7Val := genVidPth
	//par7Val := "/var/tmp/outVideo1.mp4"

	par8Key := "-frames:v"
	par8Val := "1"

	cmd := exec.Command(app, par0, par1Key, par1Val, par2Key, par2Val, par3Key, par3Val, par4Key, par4Val, par5Key, par5Val, par6, par7Key, par7Val, par8Key, par8Val)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	fmt.Println(string(stdout))
	return nil
}

func GenTempThumb(file string, hms string) (string, error) {
	//fmt.Println("file: " + file)
	//uts := time.Now().Unix()
	// ffmpeg -y -i in.mp4 -ss 00:00:03 -frames:v 1 thumb.jpg
	// ffmpeg -y -i myvideo.mp4 -ss 00:00:03 -frames:v 1 thumb.jpeg
	app := "/usr/bin/ffmpeg"
	par1 := "-y"
	par2Key := "-i"
	par2Val := file // myvideo.avi
	par3Key := "-ss"
	par3Val := hms // 00:00:30

	par4Key := "-frames:v"
	par4Val := "1"
	par5 := "/var/tmp/thumb_" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"

	cmd := exec.Command(app, par1, par2Key, par2Val, par3Key, par3Val, par4Key, par4Val, par5)
	_, err := cmd.Output()
	if err != nil {
		//  Возвращает количество записанных байтов и любую возникшую ошибку записи.
		_, _ = fmt.Fprintf(os.Stderr, "%s", err)
		return "", err
	} else {
		//onlyName := file[:len(file)-len(filepath.Ext(file))]
		fmt.Println("Миниатюра сгенерирована!!! " + par5)
		return par5, nil
		//os.Remove(par5)
	}
}

func GenPreview(vidPath string, genVidPth string) (string, error) {
	// "/home/geek/Videos/in.mp4"
	prevPath, err := GenTempThumb(vidPath, "00:00:30")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return "", err
	}
	// tmpPrevFullPath string, vidPath string, genVidPth string
	err = InsertImage(prevPath, vidPath, genVidPth)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err)
		return "", err
	}
	return prevPath, nil
}
