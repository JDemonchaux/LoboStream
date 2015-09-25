package main
import (
	"net/http"
	"log"
	"io/ioutil"
"encoding/xml"
)


type MirrorsType struct {
	Mirror []MirrorType

}
type MirrorType struct {
	Id string `xml:"id"`
	Mirrorpath string `xml:"mirrorpath"`
	Typemask string `xml:"typemask"`
}

type LanguagesType struct {
	Language []LanguageType

}
type LanguageType struct {
	Name string `xml:"name"`
	Abbreviation string `xml:"abbreviation"`
	Id string `xml:"id"`
}

type DataType struct {
	Series []SeriesType
}
type SeriesType struct {
	Id string `xml:"seriesid"`
	Language string `xml:"language"`
	Name string `xml:"SeriesName"`
	Overview string `xml:"Overview"`
	FirstAired string `xml:"FirstAired"`
	Network string `xml:"Network"`

}


func initSerieBot()  {
	var apikey = "A889B887A9B8BC82"
	respmir,err := http.Get("http://thetvdb.com/api/"+ apikey +"/mirrors.xml")
	if err != nil {
		log.Println("error seriebot get miror thetvdb ", err)
	}
	defer respmir.Body.Close()

	body, err := ioutil.ReadAll(respmir.Body)
	if err != nil {
		log.Println("error seriebot get miror thetvdb read body ",err)
	}

	log.Println(string(body))
	var mirrors MirrorsType
	err = xml.Unmarshal(body, &mirrors)
	if err != nil {
		log.Println("error seriebot get miror thetvdb read XML ",err)
	}


	log.Println("mirrors ", mirrors)// end miror


	resplan,err := http.Get(mirrors.Mirror[0].Mirrorpath + "/api/"+ apikey +"/languages.xml")
	if err != nil {
		log.Println("error seriebot get miror thetvdb ", err)
	}
	defer resplan.Body.Close()

	body, err = ioutil.ReadAll(resplan.Body)
	if err != nil {
		log.Println("error seriebot get miror thetvdb read body ",err)
	}

	log.Println(string(body))
	var languages LanguagesType
	err = xml.Unmarshal(body, &languages)
	if err != nil {
		log.Println("error seriebot get miror thetvdb read XML ",err)
	}

	log.Println("languages ", languages)// end miror
	GetSeriesPerName("South%20Park")

}

func GetSeriesPerName(name string){
	lang := "fr"
	log.Println("Get http://thetvdb.com/api/GetSeries.php?seriesname=" + name +"&language=" +lang)
	resp,err := http.Get("http://thetvdb.com/api/GetSeries.php?seriesname=" + name + "&language=" + lang)
	if err != nil {
		log.Println("error seriebot get series ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error seriebot get series read body ",err)
	}

	//log.Println(string(body))
	var series DataType
	err = xml.Unmarshal(body, &series)
	if err != nil {
		log.Println("error seriebot get series read XML ",err)
	}

	log.Println("series ", series)// end miror
	log.Println(series.Series[0].Name)
	log.Println(series.Series[0].Overview)
}

func GetSeriesPerId(name string){
	lang := "fr"
	log.Println("Get http://thetvdb.com/api/A889B887A9B8BC82/series/75897/all/fr.zip")
	resp,err := http.Get("http://thetvdb.com/api/GetSeries.php?seriesname=" + name + "&language=" + lang)
	if err != nil {
		log.Println("error seriebot get series ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error seriebot get series read body ",err)
	}

	//log.Println(string(body))
	var series DataType
	err = xml.Unmarshal(body, &series)
	if err != nil {
		log.Println("error seriebot get series read XML ",err)
	}

	log.Println("series ", series)// end miror
	log.Println(series.Series[0].Name)
	log.Println(series.Series[0].Overview)

}