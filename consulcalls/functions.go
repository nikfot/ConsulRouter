package consulcalls
<<<<<<< HEAD

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"nikfot/ConsulRouter/common"

	"github.com/wunderlist/moxy"
)

func getConsulURL() []string {
	url := []string{}
	envVars := common.GetEnvVars()
	serverName := "CONSUL-SERVER-"
	serial := 1
	key := serverName + string(serial)
	for range envVars {
		if common.IsEnvExist(key) {
			url = append(url, envVars[key])
		}
		serial += serial
		key = serverName + string(serial)
	}
	return url
}

func parseRequestBody(request *http.Request) reqDecoder {
	decoder := requestBodyDecoder(request)

	var payload reqDecoder
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	return payload
}

// Get a json decoder for a given requests body
func requestBodyDecoder(request *http.Request) *json.Decoder {
	// Read body to buffer
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

//serveConsulProxy serves as a reverse proxy to Consul APIs
func serveConsulProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	fmt.Println("Before Proxy serve")
	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
	//proxy.Transport = &transport{http.DefaultTransport}
	fmt.Println("THIS IS A RESPONSE", res)
	//req.Body.Close()
	fmt.Println(req.Body, "///////", req.Response)
}

func moxyConsul(hosts []string, res http.ResponseWriter, req *http.Request) {
	filters := []moxy.FilterFunc{}
	proxy := moxy.NewReverseProxy(hosts, filters)
	proxy.ServeHTTP(res, req)

}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
}

func forwardRequest(url string, req *http.Request) (string, error) {
	request := &ForwardRequest{}
	request.Request = req
	response, err := getConsulForward(*request, url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return response, nil
}

var getConsulForward = func(r ForwardRequest, target string) (string, error) {
	//client := &http.Client{}
	r.Request.RequestURI = ""
	fmt.Println(target, "//", r.Request.URL.Path[1:])
	u, _ := url.Parse(target + "/" + r.Request.URL.Path[1:])
	fmt.Println("Forwarding to this url: ", u)
	url := (target + "/" + r.Request.URL.Path[1:])
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		fmt.Println("Error", fmt.Sprintf("%v\n", err))
	}
	defer resp.Body.Close()
	f, _ := ioutil.ReadAll(resp.Body)
	body := string(f)
	return body, nil

}

//IsHealthy checks if consul agents are reachable
func IsHealthy() bool {
	urls := getConsulURL()
	var healthy bool
	for _, url := range urls {
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("Consul Server: ", url, " is not reachable", err.Error())
			healthy = false
			break
		} else {
			healthy = true
		}

	}
	return healthy
}
=======
>>>>>>> Initial Commit
