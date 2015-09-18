package main



import (

	"net/http"
	"log"
	"os"
	"io"
	"os/exec"
)

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)

}

func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("req ok")
	io.WriteString(w,"ok")
	file, header, err := r.FormFile("file")


	if err != nil {
		log.Println("erreur lecture post, ", err)
		return
	}


	defer file.Close()

	os.MkdirAll("tmp",777)
	out, err := os.Create("tmp/" + header.Filename)
	if err != nil {
		log.Println("erreur création fichier temp, ", err)
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("erreur upload fichier temp, ", err)
	}

	log.Println("File uploaded successfully : ")
	io.WriteString(w, header.Filename)

	_=r.Close
	convert(out.Name(),"out/" + header.Filename)
}

func convert(sourcePath string, nameDestination string)  {
	cmd := exec.Command("ffmpeg -i uploadedfile -threads 16 test.webm")
	cmd.Start()
	log.Println(" processing..." )

	//cmd.Process.Kill()

	out,err := cmd.Output()
	if err != nil {
		log.Println("some error found",err)
	}
	log.Println("out",string(out))

}