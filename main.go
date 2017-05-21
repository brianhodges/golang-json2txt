package main
import(
    "fmt"
    "os"
    "bufio"
    "encoding/json"
    "strconv"
    "./person"
)

const nl = "\n"

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    var people []person.Person
    configFile, err := os.Open("./data.json")
    check(err)

    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&people); err != nil {
        fmt.Println(err.Error())
    }

    //Write to .txt File
    f, err := os.Create("./log.txt")
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    
    for _, p := range people {
        _, err = fmt.Fprintf(w, "ID: " + strconv.Itoa(p.ID) + nl)
        check(err)
        _, err = fmt.Fprintf(w, "Full Name: " + p.Full_Name() + nl)
        check(err)
        _, err = fmt.Fprintf(w, "Email: " + p.Email + nl)
        check(err)
        _, err = fmt.Fprintf(w, "BTC Address: " + p.BTC_Address + nl)
        check(err)
        _, err = fmt.Fprintf(w, "IP Address: " + p.IP_Address + nl)
        check(err)
        _, err = fmt.Fprintf(w, nl)
        check(err)
    }
    w.Flush()
}