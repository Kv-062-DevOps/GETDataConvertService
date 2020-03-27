package handlers

import (
        "encoding/json"
        "errors"
        "gopkg.in/yaml.v3"
        "io/ioutil"
        "net/http"
        "os"
        "fmt"
        "strconv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Println(os.Getenv("ENDPOINT"))
        req, err := http.NewRequest("GET", os.Getenv("ENDPOINT"),r.Body)
        if err != nil {
                http.Error(w, "Troubles with forming request to DB service", 500)
                return
        }
        req.Header = r.Header

        client := &http.Client{}
        resp, err := client.Do(req)

        if err != nil {
                http.Error(w, "Troubles with receiving data from DB service", 500)
                return
        }
        defer resp.Body.Close()

        var response []Employee

        data, _ := ioutil.ReadAll(resp.Body)

        e := yaml.Unmarshal(data, &response)
        if e != nil {
                var emp Employee
                err := yaml.Unmarshal(data, &emp)
                if err != nil {
                        http.Error(w, "Error unmarshalling YAML OR No valid data returned by DB", 500)
                        return
                }
                emp.Salary, err = calculateSalary(&emp)
                if err != nil {
                        http.Error(w, "Error unmarshalling YAML OR No valid data returned by DB", 500)
                        return
                }
                emp.Salary, err = calculateSalary(&emp)
                if err != nil {
                        http.Error(w, "Troubles in calculating salary", 500)
                        return
                }
                er := json.NewEncoder(w).Encode(emp)
                if er != nil {
                        http.Error(w, "Error encoding data to JSON", 500)
                        return
                }
                w.WriteHeader(200)
                return
        }

        for i := 0; i < len(response); i++ {
                emp := &response[i]
                emp.Salary, err = calculateSalary(emp)
                if err != nil {
                        http.Error(w, "Troubles in calculating salary", 500)
                        return
                 }
        }

        er := json.NewEncoder(w).Encode(response)
        if er != nil {
                http.Error(w, "Error encoding data to JSON", 500)
                return
        }
}

func calculateSalary(emp* Employee) (string, error) {
        exp, err := strconv.Atoi(emp.Experience)
        if err != nil {
                return "", errors.New("Trouble converting data")
        }
        def, err := strconv.Atoi(emp.DefaultSalary)
        if err != nil {
                return "", errors.New("Trouble converting data")
        }
        temp := def * exp
        return strconv.Itoa(temp), nil
}



