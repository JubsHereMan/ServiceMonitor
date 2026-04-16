package services

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"monitoramento/internal/model"
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

func GetStatusResponse(s string) *model.Response{

	resp,_,err:= parseUrl(s)
		if err !=nil{
			fmt.Println(err)
			return &model.Response{
				Service: s,
				Status:  "NÃO POSSIVEL ACESSAR SITE",
				Code: 0,
			}
			
		}
	
	defer resp.Body.Close()
	status:= resp.StatusCode


	
		if status >= 200 && status < 400{
				return &model.Response{
				Service: s,
				Status: "ONLINE",
				Code: 1,
			
			}
		}else{
			fmt.Println(s,status)
				return &model.Response{
				Service: s,
				Status: "OFFLINE",
				Code: 0,
			}
				
		}
	}



func ResponsesStatus()[]*model.Response{
	data, err := os.ReadFile("internal/data/sites.txt")
	if err != nil{
		fmt.Println("erro aconteceu ao acessar o arquivo:", err)
	}
	lines := strings.Split(string(data),"\n")

	var results []*model.Response

	for _, linha := range lines {
		linha = strings.TrimSpace(linha)
		if linha == "" {
			continue
		}
		result := GetStatusResponse(linha)
		results = append(results, result)
	}
	return results
}

var lastResults []*model.Response

func UpdateStatus(){
	lastResults= ResponsesStatus()

}

func GetLastResults() []*model.Response {
	if lastResults == nil {
		return []*model.Response{
			{
				Service: "system",
				Status:  "loading",
				Code:    0,
			},
		}
	}
	return lastResults
}
