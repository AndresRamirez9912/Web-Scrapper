package scraping

import (
	"encoding/json"
	"log"
	"strconv"
	"webScraper/src/constants"

	"github.com/gocolly/colly"
)

type ExitoCollector struct {
	AllowedDomains []string

	Id          string
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	Discount    string
	ProductURL  string `json:"@id"`
	Offers      offers `json:"offers"`
	Brand       struct {
		Name string `json:"name"`
	} `json:"brand"`
}

type offers struct {
	LowPrice  int    `json:"lowPrice"`
	HighPrice int    `json:"highPrice"`
	Currency  string `json:"priceCurrency"`
}

// Implement the Collectors Interface
func (e ExitoCollector) OnRequest(r *colly.Request) {
	r.Headers.Set("cookie", "session-id=131-3218595-8184933; ubid-main=133-9263507-8038003; aws_lang=en; aws-target-data={'support':'1'}; aws-target-visitor-id=1678719079652-188289.34_0; AMCVS_7742037254C95E840A4C98A6@AdobeOrg=1; aws-mkto-trk=id:112-TZM-766&token:_mch-aws.amazon.com-1678719080108-37910; s_cc=true; aws-ubid-main=934-0131621-3781854; skin=noskin; x-main='qBJusB4Rt6m37PMgqk@v6KquZ7G?AI5o0liWirg7TPj6pjvcgFtIWx5NOTGwp@Rz'; at-main=Atza|IwEBICMk31CZyJJJmOmTAeWv3D2qHxu4HbCEACeg79J1hzAjj_tnnZOJD9xUqUIwTEQJhbMkF4CxHYuqezvCHWJvS0VNFgxINSdWsMz35jHYAz78fhVQW1p7RkgF5erWWrun3V2wuVZvazm_e9UzHxxcvTLNHEXyznvRA6b5eUzEkZRFJweZHXe2E6dsLi6Dq_pn_au0ZXvbijMYddm9-Qr6K-3K4Kh6Xw-rqCzjVQRvvRljnQ; sess-at-main='SmhsiaS5x3zj30ndvLj3krBdPom8CSuuTtXpLz3XYsU='; sst-main=Sst1|PQEEaqkhlRaB54Uucb6dumusCXbI1Q6_5cwTtTG0t_vzRbHljOv-Odh4NcFw2u1TYJhcckONcrPNGcIuy6Yo6Rn3sP_r71jTVS5L6bZHmbKsFvqY7EHT0KeiVeTwJysfhe5eGvuXIxaD4rGjtADjtJmjHPKPgSfGhhnZS43zrvf91BJAQj-mX4OEvCvVoGYFfVGjyGEC0WClIgg5fd8dPIpXO-_ngrzFrRyuEJ3-q9Hvpoo08SLPdrTJFeSi6TYLUOhkQHkYhCm_GqFrOLkUvLUPFj4SyJrcSAQ-8ILrYoUR-uU; lc-main=es_US; session-id-time=2082787201l; i18n-prefs=USD; regStatus=registered; awsc-color-theme=light; awsc-uh-opt-in=''; s_campaign=ps|ebd9fefd-469e-4e19-8d17-775ea6698635; s_eVar60=ebd9fefd-469e-4e19-8d17-775ea6698635; s_sq=[[B]]; noflush_awsccs_sid=3be02221654bfb2aa4ecc540c29a34d5afdf4ef8769f84175432ba8b67d77bcd; aws-userInfo={'arn':'arn:aws:iam::764539909266:root','alias':'','username':'andres.ramirez','keybase':'','issuer':'http://signin.aws.amazon.com/signin','signinType':'PUBLIC'}; AMCV_7742037254C95E840A4C98A6@AdobeOrg=1585540135|MCIDTS|19482|MCMID|54008824609818377872141446432288197497|MCAAMLH-1683664896|4|MCAAMB-1683212712|6G1ynYcLPuiQxYZrsz_pkqfLG9yMXBpb2zX5dvJdYQJzPXImdj0y|MCOPTOUT-1683067296s|NONE|MCAID|NONE|vVersion|4.4.0; sp-cdn='L5Z9:CO'; session-token='Byyiqfn3KRkfATn2800Pw0vELLN2EZ4TBQKUIaME9+Jw82/R2DtiNtfsjIvaCokNq9L2UENxWJV/LbdcBhz3eIbC1EfFxV38dsf0kVzL6kA2UZm0oxDaePgczWImFNl97fuWGM+HovcihMPALdxyeRNKo4/XS/tZ0Gzv04kT2Ak6L9jOIabEOtL+ApTCXStQNyGGHIQjdwYDRJoWiK/v080Ue0OI7efwxuzyb+9jQplRrD1ImmYWFA=='; csm-hit=tb:NRGXZSVVXPGAB0RDQW64+s-NRGXZSVVXPGAB0RDQW64|1683740996026&t:1683740996026&adb:adblk_yes")
	log.Println("Visiting", r.URL)
}

func (e ExitoCollector) OnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func (e ExitoCollector) OnError(r *colly.Response, err error) {
	log.Println("Error making the scraping: ", err)
}

func (e *ExitoCollector) OnHTML(h *colly.HTMLElement) {
	e.Id = h.ChildText("span.vtex-product-identifier-0-x-product-identifier__value")
	exitoData := h.ChildText("script[type='application/ld+json']") // Send the response
	err := json.Unmarshal([]byte(exitoData), &e)                   // Unmarshall and store on the object
	if err != nil {
		log.Println("Error unmarshaling scraping response ", err)
	}
}

func (e ExitoCollector) GetQuerySelector() string {
	return constants.EXITO_QUERY_SELECTOR
}

func (e *ExitoCollector) SetURL(URL string) {
	e.ProductURL = URL
}

func (e ExitoCollector) GetDomains() []string {
	return e.AllowedDomains
}

// Implement the Product interface
func (e *ExitoCollector) CreateProductStructure(userId string) Product {
	current := strconv.Itoa(e.Offers.LowPrice)
	high := strconv.Itoa(e.Offers.HighPrice)
	return Product{
		Product_id:      e.Id,
		User_product_id: userId + e.Id,
		Name:            e.Name,
		Brand:           e.Brand.Name,
		Description:     e.Description,
		ImageURL:        e.ImageURL,
		ProductURL:      e.ProductURL,
		Current_price:   current,
		Discount:        e.Discount,
		High_price:      high,
	}
}
