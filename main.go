package main



import (

"net/http"
"log"
"os"
"io"
"os/exec"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Views/")))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)

}

func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("req ok")
	io.WriteString(w,"ok")

	err := r.ParseMultipartForm(800000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

		//get the *fileheaders
	files := m.File["videos"]
	log.Println(files)
	for i, _ := range files {
			//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			//create destination file making sure the path is writeable.
		dst, err := os.Create("tmp/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			log.Println("Erreur creation destination")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			log.Println("Erreur copie dans destination")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		convert(dst.Name(),"out/" + files[i].Filename)

	}
	log.Println("Upload OK")

	_=r.Close

		// file, header, err := r.FormFile("videos[]")


		// if err != nil {
		// 	log.Println("erreur lecture post, ", err)
		// 	return
		// }


		// defer file.Close()

		// os.MkdirAll("tmp",777)
		// out, err := os.Create("tmp/" + header.Filename)
		// if err != nil {
		// 	log.Println("erreur création fichier temp, ", err)
		// 	return
		// }

		// defer out.Close()

		// _, err = io.Copy(out, file)
		// if err != nil {
		// 	log.Println("erreur upload fichier temp, ", err)
		// }

		// log.Println("File uploaded successfully : ")
		// io.WriteString(w, header.Filename)

		// _=r.Close
		// convert(out.Name(),"out/" + header.Filename)
}

func convert(sourcePath string, nameDestination string)  {

   out, err := exec.Command("ffmpeg.exe", "-i", sourcePath, "-codec:a", "aac", "-strict", "-2", nameDestination+".mp4").CombinedOutput()
   if err != nil {
      log.Println("some error found",err)
   }

   log.Println("out",string(out))
}