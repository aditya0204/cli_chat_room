package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	//"github.com/google/uuid"
)

var connections []net.Conn

var chatrooms =make(map[string][]net.Conn)

var clientChatroomMap = make(map[net.Conn]string)

func readMsg(conn net.Conn)(string,error){

	for {
		msgToRead ,err :=  bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// Close the connection when you're done with it.
			    //	removeConn(conn)
				conn.Close()
				return "", err
			}
			log.Println(err)
			return "",err
		}
		if err!=nil{
			fmt.Println(err)
		}

		splittedMsg := strings.Split(msgToRead[:len(msgToRead)-2]," ")
		if len(splittedMsg)>0{
			//	fmt.Println("--->",splittedMsg[0],"<-----",len(splittedMsg[0]))
			if splittedMsg[0]=="1" {
				fmt.Fprintf(conn, "\n Enter name to chatroom you want to join \t for e.g. if room is abc then type join-abc :  ")

			}
			if splittedMsg[0]=="2"{
				fmt.Fprintf(conn,"\n Type name for your chatroom : \t for e.g. if you want to create room name abc then type create-abc : ")
			}

			if strings.HasPrefix(splittedMsg[0],"create-"){
				room := strings.Split(splittedMsg[0],"-")
				if len(room)==2{
					fmt.Println("User want to create room ",room[1])
					chatrooms[room[1]]=append(chatrooms[room[1]],conn)
					clientChatroomMap[conn]=room[1]
					fmt.Fprintf(conn,"\n *******Chat room --->%s<----******created \n \n type join-%s to join this chatroom for other users\n",room[1],room[1])

				}
			}

			if strings.HasPrefix(splittedMsg[0],"join-"){
				room := strings.Split(splittedMsg[0],"-")
				if len(room)==2{
					fmt.Println("User want to create room ",room[1])
					if chatrooms[room[1]]==nil{
						fmt.Fprintf(conn,"Chat room does not exist\n")
					}else{
						chatrooms[room[1]]=append(chatrooms[room[1]],conn)
						clientChatroomMap[conn]=room[1]
						fmt.Fprintf(conn,"************Welcome to the chat room********\n")
					}
				}
			}

		}

		fmt.Println(chatrooms)

		if chatroom, ok := clientChatroomMap[conn];ok {
			     fmt.Println("publishing to chatroom ",chatroom)
                 emitToChatroom(conn,msgToRead,chatroom)
		}


		return msgToRead,nil
	}


}

func emitToChatroom(conn net.Conn, msg ,room string){

	allPeerConnections := chatrooms[room]

	fmt.Println("sending to total peers ",len(allPeerConnections))
	for i:= range allPeerConnections{


		tc := allPeerConnections[i]
		if tc == conn{

		}else{
			tc.Write([]byte(msg))
		}


	}


}

func emit(conn net.Conn, msg string){

	for i:= range connections{


		tc := connections[i]
		if tc == conn{

		}else{
			tc.Write([]byte(msg))
		}


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

	l, err:=net.Listen("tcp",":5000")
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Server is running at tcp5000")

	for {
		conn,err:=l.Accept()
		if err!=nil{
			fmt.Println(err)
			return
		}


		conn.Write([]byte("\n\n OPTIONS :-  \n type 1 to enter chatroom \n type 2 to create chatroom "))


		connections = append(connections,conn)
		go handler(conn)
	}



}
func handler(conn net.Conn){
	for {
		msg ,err:=readMsg(conn)
		if err!=nil{
			if err==io.EOF{
				fmt.Println("Closing the server")
			}
			return
		}


		fmt.Printf("Message Received: %s", msg)
		//emit(conn, msg)
	}
}
