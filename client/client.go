package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"grpcserver/dopust"
)

func main() {
	var connection *grpc.ClientConn

	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer connection.Close()

	client := dopust.NewDopustServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.CreateZaposlen(ctx, &dopust.CreateZaposlenRequest{Ime: "Peter", Priimek: "Poles",

		Spol: "Moski", DatumRojstva: "1978-02-06", Podjetje: "Podjetje d.o.o."})
	if err != nil {
		log.Fatalf("Napaka pri kreiranju zaposlenega -> %v", err)
	}

	log.Printf(`Podatki zaposlenega -> 
	ID -> %d
	Ime -> %s
	Priimek -> %s
	Spol -> %s
	Datum rojstva -> %s
	Podjetje -> %s`, r.GetIdZaposlenega(), r.GetIme(), r.GetPriimek(), r.GetSpol(), r.GetDatumRojstva(), r.GetPodjetje())

}
