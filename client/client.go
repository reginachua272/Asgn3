package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	gRPC "github.com/PatrickMatthiesen/DSYS-gRPC-template/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Same principle as in client. Flags allows for user specific arguments/values
var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("server", "5400", "Tcp server")

var state = 0
var server gRPC.TemplateClient  //the server
var ServerConn *grpc.ClientConn //the server connection
var t = 0

type Client struct {
	id         int
	portNumber int
	name string
}

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//log to file instead of console
	//setLog()

	// // Create a client
	// client := &Client{
	// 	id:         1,
	// 	portNumber: 5400,
	// 	name: clientName,
	// }

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer()
	defer ServerConn.Close()
	// Wait for the client (user) to ask for the time
	// go waitForTimeRequest(client)
	for {
		if state == 1 {
			break
		}
	}

}

// connect to server
func ConnectToServer() {

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	// timeContext, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	//dial the server, with the flag "server", to get a connection to it
	log.Printf("client %s: Attempts to dial on port %s\n", *clientsName, *serverPort)
	// Here we just go the port locally, but if you want to connect to an other device, then you would just add the 
	//Ip of the other device. (use ipconfig in the terminal or call GetOutboundIP() which can be found in server)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = gRPC.NewTemplateClient(conn)
	//Now you can just say server.<endpoint name> to call the endpoint
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
	clientName := os.Args[1]
	// Create a client
	client := &Client{
		id:         1,
		portNumber: 5400,
		name: clientName,
	}
	go waitForTimeRequest(client)
	ctx := context.Background()
	go chat(ctx,clientName)

}

	

func incrementVal(val int64) {
	//create amount type
	amount := &gRPC.Amount{
		ClientName: *clientsName,
		Value:      val, //cast from int to int32
	}

	//Make gRPC call to server with amount, and recieve acknowlegdement back.
	ack, err := server.Increment(context.Background(), amount)
	if err != nil {
		log.Printf("Client %s: no response from the server, attempting to reconnect", *clientsName)
		log.Println(err)
	}

	// check if the server has handled the request correctly
	if ack.NewValue >= val {
		fmt.Printf("Success, the new value is now %d\n", ack.NewValue)
	} else {
		// something could be added here to handle the error
		// but hopefully this will never be reached
		fmt.Println("Oh no something went wrong :(")
	}
}

func sayHi() {
	// get a stream to the server
	stream, err := server.SayHi(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	// send some messages to the server
	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "Hi"})
	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "How are you?"})
	stream.Send(&gRPC.Greeding{ClientName: *clientsName, Message: "I'm fine, thanks."})
	

	// close the stream
	farewell, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("server says: ", farewell)
}

// Function which returns a true boolean if the connection to the server is ready, and false if it's not.
func conReady(s gRPC.TemplateClient) bool {
	return ServerConn.GetState().String() == "READY"
}

// sets the logger to use a log.txt file instead of the console
func setLog() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}



func waitForTimeRequest(client *Client) {
	

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	stream, err := server.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		input := scanner.Text()
		mes := input
		if input == "exit" {
			t += 1
			mes := "Participant " + client.name + " left Chitty-Chat at Lamport time " + strconv.Itoa(t)
			if err := stream.SendMsg(&gRPC.ChatRequest{Message: mes,Time: int64(t)}); err != nil {
				log.Fatal(err)
			}
			ServerConn.Close()
			state = 1
			break
		}
		//send for broadcast
		t+=1
		if err := stream.SendMsg(&gRPC.ChatRequest{Message: mes,Time: int64(t)}); err != nil {
			log.Fatal(err)
		}

		log.Printf("sent: %s", mes)
		
	}
}


func chat(ctx context.Context,clientName string){
	stream, err := server.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}
	t += 1
	mes := "Participant " + clientName + " joined Chitty-Chat at Lamport time " + strconv.Itoa(t)

	if err := stream.SendMsg(&gRPC.ChatRequest{Message: mes,Time: int64(t)}); err != nil {
		log.Fatal(err)
	}
	log.Printf("sent: %s", mes)

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		// msg := resp.Message
		// length := len(resp.Message)
		// value,_ := strconv.Atoi(msg[length-2:length-1])
		if int(resp.Time) > t {
			t = int(resp.Time)
		}
		log.Printf("recv: %s,%d", resp.Message,int(resp.Time))
	}
	
}