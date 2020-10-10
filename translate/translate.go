package translate

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/pkg/errors"

	"github.com/Hiroshii8/GoTrans/util"
)

type Translate struct{}

const (
	basicURL = "http://translate.google.com/m"
)

var (
	languageSupport = []string{
		"auto", "af", "sq", "am", "ar", "hy", "az", "eu", "be", "bn", "bs",
		"bg", "ca", "ceb", "ny", "zh-CN", "zh-TW", "co", "hr", "cs", "da", "nl", "en",
		"kk", "km", "ko", "ku", "ky", "lo", "la", "lv", "lt", "lb", "mk", "mg", "ms",
		"eo", "et", "tl", "fi", "fr", "fy", "gl", "ka", "de", "el", "gu", "ht", "ha",
		"ml", "mt", "mi", "mr", "mn", "my", "ne", "no", "ps", "fa", "pl", "pt", "ma",
		"ro", "ru", "sm", "gd", "sr", "st", "sn", "sd", "si", "sk", "sl", "so", "es",
		"haw", "iw", "hi", "hmn", "hu", "is", "ig", "id", "ga", "it", "ja", "jw", "kn",
		"su", "sw", "sv", "tg", "ta", "te", "th", "tr", "uk", "ur", "uz", "vi", "cy",
		"xh", "yi", "yo", "zu",
	}
)

// New for init Translate Struct
func New() *Translate {
	return &Translate{}
}

// Request will call function for request the translation
func (t *Translate) Request(from, to, text string) (string, error) {

	// validate the input before formatting
	if !util.ValidateInput(from, to, languageSupport) {
		return "", errors.New("input not valid!")
	}

	// parse URL
	baseUrl, err := url.Parse(basicURL)
	if err != nil {
		return "", errors.Wrap(err, "url.Parse")
	}

	params := url.Values{}
	params.Add("sl", from)
	params.Add("tl", to)
	params.Add("q", text)
	params.Add("ie", "UTF-8")
	params.Add("oe", "UTF-8")
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", errors.Wrap(err, "[http.Get]")
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "[ioutil.ReadAll]")
	}

	resultText := string(content)
	regex := regexp.MustCompile(`<div dir="ltr" class="t0">(.*?)</div>`)
	res := regex.FindAllStringSubmatch(resultText, -1)

	if len(res) > 0 {
		return res[0][1], nil
	}
	return "", errors.New("No Response")
}
