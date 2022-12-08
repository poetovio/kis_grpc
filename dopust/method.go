package dopust

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	UnimplementedDopustServiceServer
}

func (s *Server) CreateZaposlen(ctx context.Context, in *CreateZaposlenRequest) (*Zaposlen, error) {
	log.Printf("Prejeto -> %v", in.GetIme())

	readBytes, err := ioutil.ReadFile("zaposleni.json")
	var zaposlen_list *ZaposlenList = &ZaposlenList{}
	var user_id = int64(rand.Intn(1000))

	kreiranUporabnik := &Zaposlen{IdZaposlenega: user_id, Ime: in.GetIme(), Priimek: in.GetPriimek(),
		Spol: in.GetSpol(), DatumRojstva: in.GetDatumRojstva(), Podjetje: in.GetPodjetje()}

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error -> File not found. Creating new file\n", "zaposleni.json")
			zaposlen_list.Zaposleni = append(zaposlen_list.Zaposleni, kreiranUporabnik)
			jsonBytes, err := protojson.Marshal(zaposlen_list)
			if err != nil {
				log.Fatalf("Napaka pri Marshalanju: %v", err)
			}

			if err := ioutil.WriteFile("zaposleni.json", jsonBytes, 0664); err != nil {
				log.Fatalf("Napaka pri pisanju v file -> %v", err)
			}
			return kreiranUporabnik, nil
		} else {
			log.Fatalf("Napaka ppri branju datoteke -> %v", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, zaposlen_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	zaposlen_list.Zaposleni = append(zaposlen_list.Zaposleni, kreiranUporabnik)
	jsonBytes, err := protojson.Marshal(zaposlen_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("zaposleni.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return kreiranUporabnik, nil
}

func (s *Server) GetDopustRequest(ctx context.Context, GetDopustRequest *GetDopustRequest) (*Dopust, error) {
	log.Printf("Prejeli smo dopust od client-a: %d", GetDopustRequest.IdDopust)

	return &Dopust{IdDopust: 1, IdZaposlenega: 2, DatumZacetka: "2022-03-03", DatumKonca: "2022-03-05"}, nil
}
