package main

import(
	"fmt"
	"log"
)

type server struct{

	hostname string
	ip string
	os string
	version int
}
func main(){

	s1:= server{

		hostname: "hv1",
		ip: "10.4.208.30",
		os: "Windows",
		version: 2025,
	}

	fmt.Println(s1)
	log.Println(s1)




}

func (s server) String() string {
return fmt.Sprintf("hostname is %s. IP is %s", s.hostname, s.os)
}

