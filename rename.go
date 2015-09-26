package main
import (
	"log"
	"os/exec"
	"path/filepath"
	"os"
)



func rename(fileName string){
	log.Println("Rename")
	filePath, err := filepath.Abs("out" + string(os.PathSeparator) + fileName)
	bin, err := filepath.Abs("bin" + string(os.PathSeparator) + "filebot" + string(os.PathSeparator) + "filebot.exe")
	log.Println("Rename", filePath)

	out, err := exec.Command(bin, "-rename", filePath, "--lang", "fr", "--db", "TheTVDB", "-non-strict").CombinedOutput()
	if err != nil {
		log.Println("some error found",err)
	}
	log.Println("out",string(out))

}