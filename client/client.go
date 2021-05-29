package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func readMsg(conn net.Conn){

	for {
		msgToRead ,err :=  bufio.NewReader(conn).ReadString('\n')
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("MSG received is ",msgToRead)
	}


}

func writemsg(conn net.Conn){
	for {
		msgToSend ,err:= bufio.NewReader(os.Stdin).ReadString('\n')
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Fprintf(conn,msgToSend)
		}
	}

}

func main() {
	conn,err:= net.Dial("tcp",":5000")
	if err!=nil{
		fmt.Println(err)
		return
	}

		fmt.Println("Type the msg...")
   defer conn.Close()
	go readMsg(conn)
   writemsg(conn)
}