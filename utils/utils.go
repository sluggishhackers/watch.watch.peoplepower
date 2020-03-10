package utils

import "net/url"

func QueryStringValueFromURL(URL *url.URL, key string) string {
	q, _ := url.ParseQuery(URL.RawQuery)
	return q[key][0]
}

func QueryStringValueFromRawURL(URL string, key string) string {
	parsedURL, _ := url.Parse(URL)
	q, _ := url.ParseQuery(parsedURL.RawQuery)
	return q[key][0]
}
