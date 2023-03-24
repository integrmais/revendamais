package revendamais

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type PostService Client

type Post struct {
	XMLName xml.Name `xml:"AD"`
	ID string `xml:"ID"`
	Title string `xml:"TITLE"`
	Category string `xml:"CATEGORY"`
	Description string `xml:"DESCRIPTION"`
	Accessories string `xml:"ACCESSORIES"`
	PromotionPrice string `xml:"PROMOTION_PRICE"`
	HP string `xml:"HP"`
	Make string `xml:"MAKE"`
	Model string `xml:"MODEL"`
	BaseModel string `xml:"BASE_MODEL"`
	Year string `xml:"YEAR"`
	FabricYear string `xml:"FABRIC_YEAR"`
	Condition string `xml:"CONDITION"`
	Mileage string `xml:"MILEAGE"`
	Fuel string `xml:"FUEL"`
	Gear string `xml:"GEAR"`
	MOTOR string `xml:"MOTOR"`
	Doors string `xml:"DOORS"`
	Color string `xml:"COLOR"`
	Price string `xml:"PRICE"`
	Phone string `xml:"PHONE"`
	BodyType string `xml:"BODY_TYPE"`
	LargeImages []Image `xml:"IMAGES_LARGE"`
	FIPE string `xml:"FIPE"`
	Now string `xml:"DATE"`
	LastUpdated string `xml:"LAST_UPDATE"`
}

type Posts struct {
	XMLName xml.Name `xml:"ADS"`
	Posts []Post `xml:"AD"`
}

type Image struct {
	Url string `xml:"IMAGE_URL_LARGE"`
}

func (s *PostService) List() (Posts, error) {
	url := fmt.Sprintf("%s/application/index.php/apiGeneratorXml/generator/sitedaloja/%s.xml", s.BaseUrl, s.StoreId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Posts{}, err
	}

	return s.Get(req)
}


func (s *PostService) Get(req *http.Request) (Posts, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Posts{}, err
	}

	defer res.Body.Close()

	byteValue, err := io.ReadAll(res.Body)
	if err != nil {
		return Posts{}, err
	}

	var posts Posts
	err = xml.Unmarshal(byteValue, &posts)
	if err != nil {
		return Posts{}, err
	}

	return posts, nil
}

