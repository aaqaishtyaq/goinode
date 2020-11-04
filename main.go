package main

import (
	"fmt"
	"strings"
	"os/exec"
	"runtime"
)

func main() {
	df := "df"
	inode := "-i"
	disk := "/"
	cmd := exec.Command(df, inode, disk)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err.Error())
		return
	}

	inodeData := strings.Split(string(stdout),"\n")
	header := strings.Fields(inodeData[0])
	listing := strings.Fields(inodeData[1])

	i := 0
	dfData := make(map[string]string)
	for _, v := range listing{
		v := strings.Trim(v, "%")
		dfData[header[i]] = v
		i++
	}

	os := runtime.GOOS
	switch os {
	case "darwin":
		inodePercent := dfData["%iused"]
		fmt.Println(inodePercent)
	case "linux":
		inodePercent := dfData["Use%"]
		fmt.Println(inodePercent)
	default:
		inodePercent := 0
		fmt.Println(inodePercent)
	}

}
