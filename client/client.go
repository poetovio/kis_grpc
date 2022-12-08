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

	r, err := client.CreateZaposlen(ctx, &dopust.CreateZaposlenRequest{Ime: "Kralj", Priimek: "Matjaz",

		Spol: "Moski", DatumRojstva: "1948-03-01", Podjetje: "Mercator d.o.o."})
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

	r2, err := client.CreateDopust(ctx, &dopust.CreateDopustRequest{IdZaposlenega: 2, DatumZacetka: "2021-06-07",
		DatumKonca: "2021-06-30", VrstaDopusta: "Letni", StevilkaDopusta: 1})
	if err != nil {
		log.Fatalf("Napaka pri kreiranju dopusta -> %v", err)
	}

	log.Printf(`Podatki o dopustu ->
	ID -> %d
	Datum zacetka -> %s
	Datum konca -> %s
	Vrsta dopusta -> %s
	Stevilka dopusta -> %d. dopust
	`, r2.GetIdDopust(), r2.GetDatumZacetka(), r2.GetDatumKonca(), r2.GetVrstaDopusta(), r2.GetStevilkaDopusta())

}
