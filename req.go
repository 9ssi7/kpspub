package kpspub

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func parseResponse(response *http.Response) (bool, error) {
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	return strings.Contains(string(data), "true"), nil
}

func makeRequest(ctx context.Context, xml string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", ApiUrl, strings.NewReader(xml))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Host", "tckimlik.nvi.gov.tr")
	req.Header.Add("Content-Length", strconv.Itoa(len(xml)))
	req.Header.Add("SOAPAction", "http://tckimlik.nvi.gov.tr/WS/KisiVeCuzdanDogrula")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	return parseResponse(response)
}
