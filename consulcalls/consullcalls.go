package consulcalls
<<<<<<< HEAD

import (
	"fmt"
	"net/http"
)

//RouteToConsul Receivesconsul URIs and organizes calls to Consul APIs
func RouteToConsul(req *http.Request) (res string) {
	url := getConsulURL()
	length := len(url)
	resp := make([]string, length)
	var err error
	for i := range url {
		resp[i], err = forwardRequest(url[i], req)
		if err != nil {
			fmt.Println(err)
		}
	}
	response := []string{}
	for _, v := range resp {
		if v == "[]" {
			continue
		}
		response = append(response, v)
	}
	var w string
	if len(response) > 1 {
		for k := range response {
			if k == 0 {
				last := len(response[k]) - 1
				w = resp[k][:last] + ","
			} else if k == (len(resp) - 1) {
				trimFirst := resp[k][1:]
				w = w + trimFirst
			} else {
				last := len(response[k]) - 1
				trimLast := resp[k][:last]
				trimFirst := trimLast[1:]
				w = w + trimFirst + ","
			}
		}
	} else {
		w = response[0]
	}

	return w
}
=======
>>>>>>> Initial Commit
