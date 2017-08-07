package main

import (
"net/http"
"io/ioutil"
"encoding/json"
"reflect"
"fmt"
	"log"
)

type balance struct {
	Address  string `json:"address"`
	AccountName            string `json:"accountname"`
	Balance          float64    `json:"balance"`
}
func get_holders() []balance{
	url:="http://127.0.0.1:6869/assets/4eooJSqSpHCzHp6Ao26maWBCR57gjWegbKbPDRhZ7YRt/distribution"
	accounts:=[]string{"ITE Investment trust Security Account No 1", "ITE Investment trust Security Account No 2","ITE Investment trust Security Account No 3", "ITE Investment trust Security Account No 4", "J Scott-Thompson atf International Investment Trust", "J Scott-Thompson atf APIF Trust", "J Scott-Thompson atf International Trade Exchange Trust","J Scott-Thompson atf Reserve Account No 2", "J Scott-Thompson Reserve Account No 1"}
	info := map[string]string{
		"3P7Jh6tfwAJkCfzpo6dwvW7vvKgrbpMRfKm": accounts[0],
		"3PLND7xhE6X4MEMyx8HVPSResS4nZ8XiVq3": accounts[1],
		"3PBAoDi79itvw7MZaJzNpTSWZ4RU9CwdCFY": accounts[2],
		"3PBztGMqYeyHz8f5dtoa6YbWMZbsA3z23VD": accounts[4],
		"3PDPw7eL3YGGzxCLvzC1xfc7o5b8BjLFfEG": accounts[3],
		"3PBvH5kc1ET77e9J1tqvpBdTBKqN7sFQKVc": accounts[5],
		"3P51NYYoSqBjfJCoUH5bYncH5SL1j7Uunun": accounts[6],
		"3PMyFHxg8byNfzdNwNVBS3asrqHL6TVz5aM": accounts[7],
		"3PQjg9r5aTscnsVuGD4AqP3kXoGHqXLox3o": accounts[8],
	}
	var pair map[string] interface{};
	res, err:=http.Get(url)
	if err!=nil{
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	json.Unmarshal([]byte(body),&pair)
	keys := reflect.ValueOf(pair).MapKeys()
	keys1 := reflect.ValueOf(info).MapKeys()
	out:=make([]balance,len(keys))
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
		out[i].Address=keys[i].String();
		for j := 0; j < len(keys1); j++ {
			if keys1[j].String()==strkeys[i]{
				out[i].AccountName=info[keys1[j].String()];
			}
		}
		out[i].Balance=pair[strkeys[i]].(float64)/100000000
	}
	//fmt.Printf("result %v", out)
	//fmt.Printf("result: %T", reflect.ValueOf(pair).MapIndex(reflect.ValueOf("3PQjg9r5aTscnsVuGD4AqP3kXoGHqXLox3o")))
	return out
}
func handler(w http.ResponseWriter, r *http.Request) {
	var resp []balance;
	resp=get_holders()
	res, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(res);
}
func main(){

	http.HandleFunc("/", handler) // set router
	err := http.ListenAndServe("192.168.1.115:8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
