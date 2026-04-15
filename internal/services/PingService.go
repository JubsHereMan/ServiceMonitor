package services

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)


func parseUrl(s string)(*http.Response, string, error){
	u,err := url.Parse(s)
	if err != nil {
		return nil, "", fmt.Errorf("URL inválida: %w", err)
	}
	
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if u.Scheme != "" {
			resp, err := client.Get(s)
			if err != nil {
				return nil, u.Scheme, err
			}
			return resp, u.Scheme, nil
		}

	httpsURL := "https://" + s
	resp, err := client.Get(httpsURL)
		if err == nil {
			return resp, "https", nil
		}

	httpURL := "http://" + s
	resp, err = client.Get(httpURL)
		if err == nil {
			return resp, "http", nil
		}

	return nil, "", fmt.Errorf("não foi possível acessar a URL com HTTP ou HTTPS")
}

func GetStatusResponse(s string) bool {
	resp,_,err:=parseUrl(s)
		if err !=nil{
			fmt.Println("Erro ao acessar a URL:", err)
		}
	
	defer resp.Body.Close()
	status:= resp.StatusCode
	var result bool
	
		if status >= 200 && status < 400{
			result = true
		}else{
			result = false
		}
	return result
}