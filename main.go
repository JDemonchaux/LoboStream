package main



import (

	"net/http"
	"log"
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

func main() {
	os.MkdirAll("tmp", 777)
	os.MkdirAll("out", 777)
	os.MkdirAll("media", 777)

	q := NewQueue()
	go q.Start()
	http.Handle("/", http.FileServer(http.Dir("./Views/")))
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		upload(w, r, q)
	})
	http.HandleFunc("/tosort", tosort)

	http.ListenAndServe(":80", nil)

}
func tosort(w http.ResponseWriter, r *http.Request) {

	mediaPath, _ := filepath.Abs("out")
	dir, _ := ioutil.ReadDir(mediaPath)
	var tabNameFile[]string

	for _, d := range dir {
		tabNameFile = append(tabNameFile, d.Name())
	}

	wJson, _ := json.Marshal(tabNameFile)

	io.WriteString(w, string(wJson))
	_ = r.Close
}

func upload(w http.ResponseWriter, r *http.Request, q *Queue) {
	log.Println("req ok")
	io.WriteString(w, "ok")

	err := r.ParseMultipartForm(800000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["files[]"]
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println("Erreur lecture video")
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
		log.Println(files[i].Filename)
		q.Push(&Node{files[i].Filename})

	}
	log.Println("Upload OK")

	_ = r.Close
}

