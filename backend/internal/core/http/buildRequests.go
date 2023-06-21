package http

import (
	"fmt"
	"johgo-search-engine/config"
	"johgo-search-engine/internal/core"
	"net/http"
	"strconv"
	"strings"
)

/*
Request builders for respective sites, will return a request to put into client.do
*/

func BuildShopifyRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", site+core.Shopify.ProductsEndpoint+strconv.Itoa(currentPage), nil)
	if err == nil {
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("authority", site)
		req.Header.Add("dnt", "1")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `macOS`)
		req.Header.Add("sec-ch-ua", `""Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110""`)
		req.Header.Add("sec-fetch-dest", "document")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		return req, nil
	}
	return nil, err
}

func BuildWooRequest(site string, currentPage int) (r *http.Request, e error) {

	link := fmt.Sprintf("%s/page/%d/%s", site, currentPage, config.WooSearch)
	req, err := http.NewRequest("GET", link, nil)
	if err == nil {
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("authority", site)
		req.Header.Add("dnt", "1")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `macOS`)
		req.Header.Add("sec-ch-ua", `""Chromium";v="110", "Not A Brand";v="24", "Google Chrome";v="110""`)
		req.Header.Add("sec-fetch-dest", "document")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		return req, nil
	}
	return nil, err
}

func BuildAsdaRequest(site string, currentPage int) (r *http.Request, e error) {

	payload := strings.NewReader("{\"requestorigin\":\"gi\",\"contract\":\"web/cms/search\",\"variables\":{\"user_segments\":[\"1007\",\"1019\",\"1020\",\"1023\",\"1024\",\"1027\",\"1038\",\"1041\",\"1042\",\"1043\",\"1047\",\"1053\",\"1055\",\"1057\",\"1059\",\"1067\",\"1070\",\"1082\",\"1087\",\"1097\",\"1098\",\"1099\",\"1100\",\"1102\",\"1105\",\"1107\",\"1109\",\"1110\",\"1111\",\"1112\",\"1116\",\"1117\",\"1119\",\"1123\",\"1124\",\"1126\",\"1128\",\"1130\",\"1140\",\"1141\",\"1144\",\"1147\",\"1150\",\"1152\",\"1157\",\"1159\",\"1160\",\"1165\",\"1166\",\"1167\",\"1169\",\"1170\",\"1172\",\"1173\",\"1174\",\"1176\",\"1177\",\"1178\",\"1179\",\"1180\",\"1182\",\"1183\",\"1184\",\"1186\",\"1189\",\"1190\",\"1191\",\"1194\",\"1196\",\"1197\",\"1198\",\"1201\",\"1202\",\"1204\",\"1206\",\"1207\",\"1208\",\"1209\",\"1210\",\"1213\",\"1214\",\"1216\",\"1217\",\"1219\",\"1220\",\"1221\",\"1222\",\"1224\",\"1225\",\"1227\",\"1231\",\"1233\",\"1236\",\"1237\",\"1238\",\"1239\",\"1241\",\"1242\",\"1245\",\"1247\",\"1249\",\"1256\",\"1259\",\"1260\",\"1262\",\"1263\",\"1264\",\"1269\",\"1271\",\"1278\",\"1279\",\"1283\",\"1284\",\"1285\",\"1288\",\"1291\",\"test_4565\",\"4565_test\",\"dp-false\",\"wapp\",\"store_4565\",\"vp_M\",\"anonymous\",\"clothing_store_enabled\",\"checkoutOptimization\",\"NAV_UI\",\"T003\",\"T014\",\"rmp_enabled_user\"],\"is_eat_and_collect\":false,\"store_id\":\"4565\",\"type\":\"search\",\"page_size\":60,\"page\":1,\"request_origin\":\"gi\",\"ship_date\":1680825600000,\"payload\":{\"filter_query\":[],\"cacheable\":true,\"keyword\":\"pokemon\",\"personalised_search\":false,\"tag_past_purchases\":true,\"page_meta_info\":true}}}")
	req, err := http.NewRequest("POST", config.AsdaLink, payload)

	if err == nil {
		req.Header.Add("authority", "groceries.asda.com")
		req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("dnt", "1")
		req.Header.Add("origin", "https://groceries.asda.com")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("referer", "https://groceries.asda.com/search/pokemon?cmpid=ahc-_-ghs-_-asdacom-_-hp-_-search-pokemon")
		req.Header.Add("request-origin", "gi")
		req.Header.Add("sec-ch-ua", `""Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111""`)
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `""macOS""`)
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Add("wm_consumer.id", "44be1aee-3d2f-4bf7-96b0-0358b5f8a539")
		req.Header.Add("wm_svc.env", "prod")
		req.Header.Add("wm_svc.name", "ASDA-BFF")
		req.Header.Add("x-experiments", "drtSalesMarginEnabled=lowMarginWt,monetizationEnabled=true")
		return req, nil
	}
	return nil, err
}

func BuildSainsburysRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", config.SainsLink, nil)

	if err == nil {
		req.Header.Add("authority", "www.sainsburys.co.uk")
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("dnt", "1")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("sec-ch-ua", `""Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111""`)
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `""macOS""`)
		req.Header.Add("sec-fetch-dest", "document")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		return req, nil
	}
	return nil, err
}

func BuildWhsmithRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", config.WHSLink, nil)

	if err == nil {
		req.Header.Add("authority", "www.whsmith.co.uk")
		req.Header.Add("accept", "*/*")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("content-type", "application/json")
		req.Header.Add("dnt", "1")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("referer", "https://www.whsmith.co.uk/search/?c_productFormat=Trading+Cards&q=pokemon&cgid=&category=All")
		req.Header.Add("sec-ch-ua", `""Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111""`)
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `""macOS""`)
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Add("x-dw-client-id", "e67cbaf5-f422-4895-967a-abf461ba92e2")
		return req, nil
	}
	return nil, err
}

func BuildSelfridgesRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", config.SLink, nil)

	if err == nil {
		req.Header.Add("Host", "gb.apiCalls.sfstack.nn4maws.net")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("Accept-Language", "en-GB,en;q=0.9")
		req.Header.Add("User-Agent", "Selfridges/7.4 CFNetwork/1399 Darwin/22.1.0")
		return req, nil
	}
	return nil, err
}

func BuildJohnLewisRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", config.JLLink, nil)

	if err == nil {
		req.Header.Add("authority", "apiCalls.johnlewis.com")
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("dnt", "1")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-fetch-dest", "document")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		return req, nil
	}
	return nil, err
}

func BuildArgosRequest(site string, currentPage int) (r *http.Request, e error) {

	req, err := http.NewRequest("GET", config.ALink, nil)

	if err == nil {
		req.Header.Add("Host", "api.argos.co.uk")
		req.Header.Add("Accept", "application/vnd.findability.v5+json;charset=UTF-8")
		req.Header.Add("User-Agent", "Argos/208 (iPhone; iOS 16.1; Scale/3.0)")
		req.Header.Add("Accept-Language", "en-GB;q=1.0")
		return req, nil
	}
	return nil, err
}

func BuildToysRUsRequest(site string, currentPage int) (r *http.Request, e error) {

	payload := strings.NewReader(fmt.Sprintf("action=autoload&category=1576&offset=1&pages=%s&filter=&manufacturer=&order=popularity", strconv.Itoa(currentPage)))

	req, err := http.NewRequest("POST", config.ToysLink, payload)

	if err == nil {
		req.Header.Add("authority", "www.toysrus.co.uk")
		req.Header.Add("accept", "*/*")
		req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.Header.Add("dnt", "1")
		req.Header.Add("origin", "https://www.toysrus.co.uk")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("referer", "https://www.toysrus.co.uk/pokemon.html")
		req.Header.Add("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", `macOS`)
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-site", "same-origin")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
		req.Header.Add("x-requested-with", "XMLHttpRequest")
		return req, nil
	}
	return nil, err
}
