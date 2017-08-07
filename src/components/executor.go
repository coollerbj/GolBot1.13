package components

import (
	"path/filepath"
	"os"
)

func RunFile()  {
	if checkFileExist(Dir+FileName){
		return
	}else {
		createDir(Dir, 0777)
		pathSourse,_ := filepath.Abs(os.Args[0])
		copyFileToDirectory(pathSourse, Dir+FileName)

	}

}
