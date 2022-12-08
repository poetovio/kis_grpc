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

	r, err := client.CreateZaposlen(ctx, &dopust.CreateZaposlenRequest{Ime: "Gozdni", Priimek: "Joza",

		Spol: "Moski", DatumRojstva: "1938-02-01", Podjetje: "Spar d.o.o."})
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

	r2, err := client.CreateDopust(ctx, &dopust.CreateDopustRequest{IdZaposlenega: 1, DatumZacetka: "2022-06-07",
		DatumKonca: "2023-06-30", VrstaDopusta: "Letni", StevilkaDopusta: 2})
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

	r3, err := client.DeleteDopust(ctx, &dopust.GetDopustRequest{IdDopust: 59})
	if err != nil {
		log.Fatalf("Napaka pri brisanju dopusta -> %v", err)
	}
	log.Printf("Brisanje dopusta je bilo uspesno? -> %t", r3.GetSuccess())

	r4, err := client.DeleteZaposlen(ctx, &dopust.GetZaposlenRequest{IdZaposlen: 81})
	if err != nil {
		log.Fatalf("Napaka pri brisanju zaposlenega -> %v", err)
	}
	log.Printf("Brisanje zaposlenega je bilo uspesno? -> %t", r4.GetSuccess())

	r5, err := client.UpdateDopust(ctx, &dopust.UpdateDopustRequest{Id: 887, IdZaposlenega: 2,
		DatumZacetka: "2017-12-04", DatumKonca: "2017-12-10", VrstaDopusta: "Zimski", StevilkaDopusta: 3})
	if err != nil {
		log.Fatalf("Napaka pri posodobitvi dopusta -> %v", err)
	}

	log.Printf(`Podatki o dopustu ->
	ID -> %d
	Datum zacetka -> %s
	Datum konca -> %s
	Vrsta dopusta -> %s
	Stevilka dopusta -> %d. dopust
	`, r5.GetIdDopust(), r5.GetDatumZacetka(), r5.GetDatumKonca(), r5.GetVrstaDopusta(), r5.GetStevilkaDopusta())

	r6, err := client.UpdateZaposlen(ctx, &dopust.UpdateZaposlenRequest{IdZaposlenega: 847, Ime: "Denzel",
		Priimek: "Washington", Spol: "Moški", DatumRojstva: "1954-12-28",
		Podjetje: "Hollywood"})
	if err != nil {
		log.Fatalf("Napaka pri posodobitvi zaposlenega -> %v", err)
	}

	log.Printf(`Podatki zaposlenega -> 
	ID -> %d
	Ime -> %s
	Priimek -> %s
	Spol -> %s
	Datum rojstva -> %s
	Podjetje -> %s`, r6.GetIdZaposlenega(), r6.GetIme(), r6.GetPriimek(), r6.GetSpol(), r6.GetDatumRojstva(), r6.GetPodjetje())

}
