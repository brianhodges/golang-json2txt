package person

type Person struct {
    ID int `json:"id"`
    First_Name string `json:"first_name"`
    Last_Name  string `json:"last_name"`
    Email string `json:"email"`
    BTC_Address string `json:"btc_address"`
    IP_Address string `json:"ip_address"`
}

func (p Person) Full_Name() string {
    return p.First_Name + " " + p.Last_Name
}