package common

import (
	"net/http"
	"os"
)

//TruncateString trancates any given string to the letter count given
func TruncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

//IsEnvExist checks if environmental variable exists
func IsEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

//GetEnvVars gets environmental variables
func GetEnvVars() map[string]string {
	vars := make(map[string]string)
	//	a_condtion_url := os.Getenv("A_CONDITION_URL")
	vars["PORT"] = os.Getenv("PORT")
	vars["HOST"] = os.Getenv("HOST")
	serverName := "CONSUL-SERVER-"
	serial := 1
	key := serverName + string(serial)
	for IsEnvExist(key) {
		vars[key] = os.Getenv(key)
		serial += serial
		key = serverName + string(serial)
	}
	return vars

}

//GetIP reads the IP from which request was made
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
