package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/midorigreen/test-kube/protoc"
	"google.golang.org/grpc"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	// dummy
	backHost := "localhost"
	backPort := "8080"

	// WithInsecure options set no secure transport
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", backHost, backPort), grpc.WithInsecure())
	if err != nil {
		responseError(w, err, "Failed")
		return
	}

	client := pb.NewPingClient(conn)
	// background is empty context
	ctx := context.Background()
	res, err := client.Ok(ctx, &pb.OkRequest{Quetion: "ping"})
	if err != nil {
		responseError(w, err, "Failed")
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		responseError(w, err, "json marshal failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(b))
}

func responseError(w http.ResponseWriter, err error, mes string) {
	log.Fatalln(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "Failed")
}
