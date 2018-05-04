package main

import "hust_wireless"

func main(){
	u:= new(hust_wireless.User)
	u.Init("HUSTWIRELESS","123456")
	status, str:= u.Login()
	println(status, str)
	//u.Logout()
}
