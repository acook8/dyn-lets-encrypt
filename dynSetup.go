package main

import (
    "encoding/json"
    "log"
    "os"
    "time"

    "github.com/monaco-io/request"
)

// Token struct
type Token struct {
    Status string `json:"status"`
    Data   struct {
        Token   string `json:"token"`
        Version string `json:"version"`
    } `json:"data"`
    JobID int64 `json:"job_id"`
    Msgs  []struct {
        INFO   string      `json:"INFO"`
        SOURCE string      `json:"SOURCE"`
        ERRCD  interface{} `json:"ERR_CD"`
        LVL    string      `json:"LVL"`
    } `json:"msgs"`
}

// function to generate a token
func getToken() string {
    client := request.Client{
        URL:    "https://api.dynect.net/REST/Session/",
        Method: "POST",
        Body: []byte(`{
            "customer_name": "CUSTOMER_NAME",
            "user_name": "USER_NAME",
            "password": "PASSWORD"
        }`),
    }
    resp, err := client.Do()

    if err != nil {
        log.Println(err)
    }

    var tokenData Token
    json.Unmarshal([]byte(resp.Data), &tokenData)
    var token string
    token = tokenData.Data.Token

    return token
}

func deleteTxt(token string, targetDomain string, recordName string) {

    deleteURL := "https://api.dynect.net/REST/TXTRecord/" + targetDomain + "/" + recordName
    client := request.Client{
        URL:    deleteURL,
        Method: "DELETE",
        Header: map[string]string{"Auth-Token": token, "Content-Type": "application/json"},
        // ContentType: request.ApplicationJSON,
    }
    resp, err := client.Do()

    if err != nil {
        log.Println(resp.Code, string(resp.Data), err)
    }
}

func updateZone(token string, targetDomain string, recordName string) {
    updateZoneURL := "https://api.dynect.net/REST/Zone/" + targetDomain
    client := request.Client{
        URL:    updateZoneURL,
        Method: "PUT",
        Header: map[string]string{"Auth-Token": token, "Content-Type": "application/json"},
        Body:   []byte(`{"publish": true}`),
    }

    resp, err := client.Do()

    if err != nil {
        log.Println(resp.Code, string(resp.Data), err)
    }
}

func addTxt(token string, targetDomain string, recordName string, recordValue string) {
    addURL := "https://api.dynect.net/REST/TXTRecord/" + targetDomain + "/" + recordName
    recordValue = `"` + recordValue + `"`
    body := `{"rdata":{"txtdata": ` + recordValue + `},"ttl":"0"}`
    client := request.Client{
        URL:    addURL,
        Method: "POST",
        Header: map[string]string{"Auth-Token": token, "Content-Type": "application/json"},
        Body:   []byte(body),
    }
    resp, err := client.Do()

    if err != nil {
        log.Println(resp.Code, string(resp.Data), err)
    }
}

func main() {

    token := getToken()
    targetDomain := os.Args[1]
    recordName := os.Args[2]
    recordValue := os.Args[3]
    runType := os.Args[4]

    // Certify the Web passes *.example.com, but we need the zone
    targetDomain = targetDomain[2:]

    if runType == "setup" {
        //make sure there are no existing _acme-challege records
        deleteTxt(token, targetDomain, recordName)
        updateZone(token, targetDomain, recordName)

        //create new record
        addTxt(token, targetDomain, recordName, recordValue)
        updateZone(token, targetDomain, recordName)

        //sleep for ten seconds so dns changes can be propegated
        time.Sleep(10 * time.Second)
    } else if runType == "cleanup" {
        //cleanup the records
        deleteTxt(token, targetDomain, recordName)
        updateZone(token, targetDomain, recordName)
    } else {
        log.Println("There was an error, please check the arguments")
    }

}
