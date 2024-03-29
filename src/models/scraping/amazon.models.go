package scraping

import (
	"errors"
	"log"
	"strings"
	"webScraper/src/constants"

	"github.com/gocolly/colly"
)

type AmazonColector struct {
	Id             string
	Name           string
	CurrentPrice   string
	Disccount      string
	HighPrice      string
	Brand          string
	Description    map[int]string
	ImageURL       string
	ProductURL     string
	AllowedDomains []string
}

// Implement the Collectors Interface
func (a AmazonColector) OnRequest(r *colly.Request) {
	r.Headers.Set("cookie", "session-id=131-3218595-8184933; ubid-main=133-9263507-8038003; aws_lang=en; aws-target-data={'support':'1'}; aws-target-visitor-id=1678719079652-188289.34_0; AMCVS_7742037254C95E840A4C98A6@AdobeOrg=1; aws-mkto-trk=id:112-TZM-766&token:_mch-aws.amazon.com-1678719080108-37910; s_cc=true; aws-ubid-main=934-0131621-3781854; skin=noskin; x-main='qBJusB4Rt6m37PMgqk@v6KquZ7G?AI5o0liWirg7TPj6pjvcgFtIWx5NOTGwp@Rz'; at-main=Atza|IwEBICMk31CZyJJJmOmTAeWv3D2qHxu4HbCEACeg79J1hzAjj_tnnZOJD9xUqUIwTEQJhbMkF4CxHYuqezvCHWJvS0VNFgxINSdWsMz35jHYAz78fhVQW1p7RkgF5erWWrun3V2wuVZvazm_e9UzHxxcvTLNHEXyznvRA6b5eUzEkZRFJweZHXe2E6dsLi6Dq_pn_au0ZXvbijMYddm9-Qr6K-3K4Kh6Xw-rqCzjVQRvvRljnQ; sess-at-main='SmhsiaS5x3zj30ndvLj3krBdPom8CSuuTtXpLz3XYsU='; sst-main=Sst1|PQEEaqkhlRaB54Uucb6dumusCXbI1Q6_5cwTtTG0t_vzRbHljOv-Odh4NcFw2u1TYJhcckONcrPNGcIuy6Yo6Rn3sP_r71jTVS5L6bZHmbKsFvqY7EHT0KeiVeTwJysfhe5eGvuXIxaD4rGjtADjtJmjHPKPgSfGhhnZS43zrvf91BJAQj-mX4OEvCvVoGYFfVGjyGEC0WClIgg5fd8dPIpXO-_ngrzFrRyuEJ3-q9Hvpoo08SLPdrTJFeSi6TYLUOhkQHkYhCm_GqFrOLkUvLUPFj4SyJrcSAQ-8ILrYoUR-uU; lc-main=es_US; session-id-time=2082787201l; i18n-prefs=USD; regStatus=registered; awsc-color-theme=light; awsc-uh-opt-in=''; s_campaign=ps|ebd9fefd-469e-4e19-8d17-775ea6698635; s_eVar60=ebd9fefd-469e-4e19-8d17-775ea6698635; s_sq=[[B]]; noflush_awsccs_sid=3be02221654bfb2aa4ecc540c29a34d5afdf4ef8769f84175432ba8b67d77bcd; aws-userInfo={'arn':'arn:aws:iam::764539909266:root','alias':'','username':'andres.ramirez','keybase':'','issuer':'http://signin.aws.amazon.com/signin','signinType':'PUBLIC'}; AMCV_7742037254C95E840A4C98A6@AdobeOrg=1585540135|MCIDTS|19482|MCMID|54008824609818377872141446432288197497|MCAAMLH-1683664896|4|MCAAMB-1683212712|6G1ynYcLPuiQxYZrsz_pkqfLG9yMXBpb2zX5dvJdYQJzPXImdj0y|MCOPTOUT-1683067296s|NONE|MCAID|NONE|vVersion|4.4.0; sp-cdn='L5Z9:CO'; session-token='Byyiqfn3KRkfATn2800Pw0vELLN2EZ4TBQKUIaME9+Jw82/R2DtiNtfsjIvaCokNq9L2UENxWJV/LbdcBhz3eIbC1EfFxV38dsf0kVzL6kA2UZm0oxDaePgczWImFNl97fuWGM+HovcihMPALdxyeRNKo4/XS/tZ0Gzv04kT2Ak6L9jOIabEOtL+ApTCXStQNyGGHIQjdwYDRJoWiK/v080Ue0OI7efwxuzyb+9jQplRrD1ImmYWFA=='; csm-hit=tb:NRGXZSVVXPGAB0RDQW64+s-NRGXZSVVXPGAB0RDQW64|1683740996026&t:1683740996026&adb:adblk_yes")
	log.Println("Visiting", r.URL)
}

func (a AmazonColector) OnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func (a AmazonColector) OnError(r *colly.Response, err error) {
	log.Println("Error making the scraping: ", err)
}

func (a *AmazonColector) OnHTML(h *colly.HTMLElement) {
	// Get the Id of the product
	err := a.getProductId(a.ProductURL)
	if err != nil {
		return
	}

	a.Name = h.ChildText(constants.AMAZON_QUERY_NAME)   // Product Tittle
	a.Brand = h.ChildText(constants.AMAZON_QUERY_BRAND) // Brand

	// Description
	var description = make(map[int]string)
	h.ForEach(constants.AMAZON_QUERY_DESCRIPTION, func(i int, h *colly.HTMLElement) {
		description[i] = h.Text
	})
	a.Description = description

	if h.ChildAttr(constants.AMAZON_QUERY_IMAGE_URL, constants.AMAZON_QUERY_IMAGE_URL_ATTR) != "" {
		a.ImageURL = h.ChildAttr(constants.AMAZON_QUERY_IMAGE_URL, constants.AMAZON_QUERY_IMAGE_URL_ATTR) // Image URL
	}

	// Discount Form
	a.Disccount = h.ChildText(constants.AMAZON_QUERY_DISCOUNT_DISCOUNT)        // Product Discount
	a.CurrentPrice = h.ChildText(constants.AMAZON_QUERY_CURRENTPRICE_DISCOUNT) // Product Lower Price
	a.HighPrice = h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_DISCOUNT)      // Original Price, withou Discount

	// Current Form - No discount
	if a.HighPrice == "" {
		a.HighPrice = h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_CURRENT)
		a.CurrentPrice = h.ChildText(constants.AMAZON_QUERY_CURRENTPRICE_CURRENT)
		prices := strings.Split(h.ChildText(constants.AMAZON_QUERY_HIGHTPRICE_CURRENT), "US")
		if len(prices) > 2 {
			a.HighPrice = prices[2]
			a.CurrentPrice = prices[2]
		}
	}

	// Table Form
	if a.HighPrice == "" {
		var prices = make(map[int]string)
		h.ForEach(constants.AMAZON_QUERY_PRICES_TABLE, func(i int, h *colly.HTMLElement) {
			prices[i] = h.ChildText(constants.AMAZON_QUERY_PRICE_ELEMENT_TABLE)
		})
		a.HighPrice = prices[0]
		a.CurrentPrice = prices[1]
		a.Disccount = prices[2]
	}
}

func (a AmazonColector) GetDomains() []string {
	return a.AllowedDomains
}

// Owner Functions
func (a *AmazonColector) getProductId(productURL string) error {
	productWithoutQuery := strings.Replace(productURL, "?", "/", -1)
	urlElements := strings.Split(productWithoutQuery, "/")
	for index, value := range urlElements {
		if value == "dp" {
			a.Id = urlElements[index+1] // Assign the Id
			return nil
		}
	}
	return errors.New("Product Id not found")
}

func (a AmazonColector) GetQuerySelector() string {
	return constants.AMAZON_QUERY_SELECTOR
}

func (a *AmazonColector) SetURL(URL string) {
	a.ProductURL = URL
}

func (a *AmazonColector) CreateProductStructure(userId string) Product {
	description := "&"
	for _, v := range a.Description {
		description = v + description
	}
	return Product{
		Product_id:      a.Id,
		User_product_id: userId + a.Id,
		Name:            a.Name,
		Brand:           a.Brand,
		Description:     description,
		ImageURL:        a.ImageURL,
		ProductURL:      a.ProductURL,
		Current_price:   a.CurrentPrice,
		Discount:        a.Disccount,
		High_price:      a.HighPrice,
	}
}
