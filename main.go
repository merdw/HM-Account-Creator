package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strings"
)

var email, password string
var yil int = 2001
var ay int = 1
var gun int = 1
var sayac int = 1
var durum bool
var buffer string

func main() {
	fmt.Print("Kullanici adi giriniz:")
	fmt.Scanln(&email)
	password = "Aqwer12*"

	for ay != 13 {
		gun = 1
		for gun != 32 {

			CreateMember(email, password, gun, ay)

			if durum == false {
				fmt.Println("false dondu aga")
				file, err := os.OpenFile("sorunnotcreated.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
				file.WriteString(fmt.Sprintf("%s%d@gmail.com:%s     %d/%d/%d SORUN: %s\n", email, sayac, password, gun, ay, yil, buffer))
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("sorun yazildi!")
				}
			} else {
				fmt.Println("true donmus aga")
				file, err := os.OpenFile("created.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
				file.WriteString(fmt.Sprintf("%s%d@gmail.com:%s     %d/%d/%d\n", email, sayac, password, gun, ay, yil))
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("yazildi!")
				}
			}
			sayac++
			gun++
		}
		ay++
	}
	fmt.Println(sayac)
}
func CreateMember(email string, password string, gun int, ay int) {

	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://app2.hm.com/tr_tr/register/newcustomer")
	req.Header.SetMethod("POST")
	req.Header.Add("x-devid", "A738E345-11F8-4868-AD31-50A85AC32B1E")
	req.Header.Add("x-app-devicetype", "iPhone 14")
	req.Header.Add("x-app-version", "22.37.0")
	req.Header.Add("x-os-version", "16.0")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	req.Header.Add("user-agent", "targetapp_ios_21")
	req.SetBody([]byte(fmt.Sprintf("{\"lastName\":\"\",\"prefix\":\"\",\"firstName\":\"\",\"birthDate\":\"%d-%d-%d\",\"hmNewsSubscription\":\"true\",\"day\":\"%d\",\"postalCode\":\"\",\"year\":\"%d\",\"month\":\"%d\",\"pwd\":\"%s\",\"email\":\"%s%d@gmail.com\",\"gender\":\"\",\"cellPhone\":\"\"}", yil, ay, gun, gun, yil, ay, password, email, sayac)))

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer = string(resp.Body())
	println(buffer)
	if strings.Contains(string(buffer), "hmClubJoin\":true") {
		durum = true
		return
	}
	durum = false
}
