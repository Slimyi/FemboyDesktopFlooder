package r34dl

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Posts struct {
	XMLName xml.Name `xml:"posts"`
	Text    string   `xml:",chardata"`
	Count   string   `xml:"count,attr"`
	Offset  string   `xml:"offset,attr"`
	Post    []struct {
		Text          string `xml:",chardata"`
		Height        string `xml:"height,attr"`
		Score         string `xml:"score,attr"`
		FileURL       string `xml:"file_url,attr"`
		ParentID      string `xml:"parent_id,attr"`
		SampleURL     string `xml:"sample_url,attr"`
		SampleWidth   string `xml:"sample_width,attr"`
		SampleHeight  string `xml:"sample_height,attr"`
		PreviewURL    string `xml:"preview_url,attr"`
		Rating        string `xml:"rating,attr"`
		Tags          string `xml:"tags,attr"`
		ID            string `xml:"id,attr"`
		Width         string `xml:"width,attr"`
		Change        string `xml:"change,attr"`
		Md5           string `xml:"md5,attr"`
		CreatorID     string `xml:"creator_id,attr"`
		HasChildren   string `xml:"has_children,attr"`
		CreatedAt     string `xml:"created_at,attr"`
		Status        string `xml:"status,attr"`
		Source        string `xml:"source,attr"`
		HasNotes      string `xml:"has_notes,attr"`
		HasComments   string `xml:"has_comments,attr"`
		PreviewWidth  string `xml:"preview_width,attr"`
		PreviewHeight string `xml:"preview_height,attr"`
	} `xml:"post"`
}

func Fetch(tags string, limit int, page int) Posts {
	var result Posts
	if limit > 50 || limit < 1 {
		fmt.Println("Limit either too high or too low. Try: r34-dl help get")
		return Posts{}
	}
	var url string
	if tags != "" {
		url = "https://api.rule34.xxx/index.php?page=dapi&s=post&q=index&pid=" + fmt.Sprintf("%d", page) + "&limit=" + fmt.Sprintf("%d", limit) + "&tags=" + tags
	} else {
		url = "https://api.rule34.xxx/index.php?page=dapi&s=post&q=index&pid=" + fmt.Sprintf("%d", page) + "&limit=" + fmt.Sprintf("%d", limit)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	xml.NewDecoder(res.Body).Decode(&result)
	return result
}
