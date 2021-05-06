package server

import (
	
	"log"
	
	
	"testing"
	"net"
	
)

 func BenchmarkSerever(b *testing.B) {
	
 	host := "0.0.0.0"
 	port := "9999"

 	srv := NewServer(net.JoinHostPort(host, port))

 	
 	
	
 	for i := 0; i < b.N; i++ {

 	errN := srv.Start()

 	conn, err := net.Dial("tcp", net.JoinHostPort(host, port)) 
     if err != nil { 
         log.Println(err) 
        
     } 
     defer conn.Close() 

	if errN != nil {
 		log.Print(errN)
 	}
}
 }