package interfaces

import (
	"log"
	"time"
	"webScraper/src/constants"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// This is the super interface that every object will implement
type Collectors interface {
	OnError(colly.ErrorCallback)
	OnRequest(colly.RequestCallback)
	OnResponse(colly.ResponseCallback)
	OnHTML(string, colly.HTMLCallback)
}

type Scraper struct {
	AllowedDomains []string
}

func (s Scraper) InitCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(s.AllowedDomains...), // Insert the domains
		colly.CacheDir(constants.CACHE),
	)
	collector.SetRequestTimeout(90 * time.Second)
	extensions.RandomUserAgent(collector) // Assign a random User Agent

	// Assign the Proxy
	// rp, err := proxy.RoundRobinProxySwitcher("http://20.206.106.192:8123", "http://65.109.228.231:8080")
	// if err != nil {
	// 	log.Println(err)
	// }
	// collector.SetProxyFunc(rp)
	return collector
}

func (s Scraper) OnRequest(r *colly.Request) {
	r.Headers.Set("cookie", "session-id=131-3218595-8184933; ubid-main=133-9263507-8038003; aws_lang=en; aws-target-data={'support':'1'}; aws-target-visitor-id=1678719079652-188289.34_0; AMCVS_7742037254C95E840A4C98A6@AdobeOrg=1; aws-mkto-trk=id:112-TZM-766&token:_mch-aws.amazon.com-1678719080108-37910; s_cc=true; aws-ubid-main=934-0131621-3781854; skin=noskin; x-main='qBJusB4Rt6m37PMgqk@v6KquZ7G?AI5o0liWirg7TPj6pjvcgFtIWx5NOTGwp@Rz'; at-main=Atza|IwEBICMk31CZyJJJmOmTAeWv3D2qHxu4HbCEACeg79J1hzAjj_tnnZOJD9xUqUIwTEQJhbMkF4CxHYuqezvCHWJvS0VNFgxINSdWsMz35jHYAz78fhVQW1p7RkgF5erWWrun3V2wuVZvazm_e9UzHxxcvTLNHEXyznvRA6b5eUzEkZRFJweZHXe2E6dsLi6Dq_pn_au0ZXvbijMYddm9-Qr6K-3K4Kh6Xw-rqCzjVQRvvRljnQ; sess-at-main='SmhsiaS5x3zj30ndvLj3krBdPom8CSuuTtXpLz3XYsU='; sst-main=Sst1|PQEEaqkhlRaB54Uucb6dumusCXbI1Q6_5cwTtTG0t_vzRbHljOv-Odh4NcFw2u1TYJhcckONcrPNGcIuy6Yo6Rn3sP_r71jTVS5L6bZHmbKsFvqY7EHT0KeiVeTwJysfhe5eGvuXIxaD4rGjtADjtJmjHPKPgSfGhhnZS43zrvf91BJAQj-mX4OEvCvVoGYFfVGjyGEC0WClIgg5fd8dPIpXO-_ngrzFrRyuEJ3-q9Hvpoo08SLPdrTJFeSi6TYLUOhkQHkYhCm_GqFrOLkUvLUPFj4SyJrcSAQ-8ILrYoUR-uU; lc-main=es_US; session-id-time=2082787201l; i18n-prefs=USD; regStatus=registered; awsc-color-theme=light; awsc-uh-opt-in=''; s_campaign=ps|ebd9fefd-469e-4e19-8d17-775ea6698635; s_eVar60=ebd9fefd-469e-4e19-8d17-775ea6698635; s_sq=[[B]]; noflush_awsccs_sid=3be02221654bfb2aa4ecc540c29a34d5afdf4ef8769f84175432ba8b67d77bcd; aws-userInfo={'arn':'arn:aws:iam::764539909266:root','alias':'','username':'andres.ramirez','keybase':'','issuer':'http://signin.aws.amazon.com/signin','signinType':'PUBLIC'}; AMCV_7742037254C95E840A4C98A6@AdobeOrg=1585540135|MCIDTS|19482|MCMID|54008824609818377872141446432288197497|MCAAMLH-1683664896|4|MCAAMB-1683212712|6G1ynYcLPuiQxYZrsz_pkqfLG9yMXBpb2zX5dvJdYQJzPXImdj0y|MCOPTOUT-1683067296s|NONE|MCAID|NONE|vVersion|4.4.0; sp-cdn='L5Z9:CO'; session-token='Byyiqfn3KRkfATn2800Pw0vELLN2EZ4TBQKUIaME9+Jw82/R2DtiNtfsjIvaCokNq9L2UENxWJV/LbdcBhz3eIbC1EfFxV38dsf0kVzL6kA2UZm0oxDaePgczWImFNl97fuWGM+HovcihMPALdxyeRNKo4/XS/tZ0Gzv04kT2Ak6L9jOIabEOtL+ApTCXStQNyGGHIQjdwYDRJoWiK/v080Ue0OI7efwxuzyb+9jQplRrD1ImmYWFA=='; csm-hit=tb:NRGXZSVVXPGAB0RDQW64+s-NRGXZSVVXPGAB0RDQW64|1683740996026&t:1683740996026&adb:adblk_yes")
	log.Println("Visiting", r.URL)
}

func (s Scraper) OnResponse(r *colly.Response) {
	log.Println("Response Code: ", r.StatusCode)
}

func (s Scraper) OnError(r *colly.Response, err error) {
	log.Println("Error making the scraping: ", err)
}

func (s Scraper) OnHTML(string, colly.HTMLCallback) {
	log.Println("Error making the scraping: ")
}
