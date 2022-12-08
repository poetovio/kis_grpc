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

/*
type StreamServer struct {
	UnimplementedStreamDopustServer
}
*/

func (s *Server) GetZaposlen(ctx context.Context, input *GetEnZaposlenRequest) (*Zaposlen, error) {
	readBytes, err := ioutil.ReadFile("zaposleni.json")
	if err != nil {
		log.Fatalf("Napaka pri branju datoteke -> %v", err)
	}

	var zaposlen_list *ZaposlenList = &ZaposlenList{}

	if err := protojson.Unmarshal(readBytes, zaposlen_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for _, zaposlen := range zaposlen_list.Zaposleni {
		if zaposlen.GetIdZaposlenega() == input.IdZaposlen {
			kreiranZaposlen := &Zaposlen{IdZaposlenega: zaposlen.GetIdZaposlenega(), Ime: zaposlen.GetIme(), Priimek: zaposlen.GetPriimek(),
				Spol: zaposlen.GetSpol(), DatumRojstva: zaposlen.GetDatumRojstva(), Podjetje: zaposlen.GetPodjetje()}

			return kreiranZaposlen, nil
		}
	}

	return nil, nil
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
			log.Fatalf("Napaka pri branju datoteke -> %v", err)
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

func (s *Server) CreateDopust(ctx context.Context, in *CreateDopustRequest) (*Dopust, error) {
	log.Printf("Prejeto -> %v", in.GetStevilkaDopusta())

	readBytes, err := ioutil.ReadFile("dopusti.json")
	var dopust_list *DopustList = &DopustList{}
	var dopust_id = int64(rand.Intn(1000))

	kreiranDopust := &Dopust{IdDopust: dopust_id, IdZaposlenega: in.GetIdZaposlenega(),
		DatumZacetka: in.GetDatumZacetka(), DatumKonca: in.GetDatumKonca(),
		VrstaDopusta: in.GetVrstaDopusta(), StevilkaDopusta: in.GetStevilkaDopusta()}

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error -> File not found. Creating new file\n", "dopusti.json")
			dopust_list.Dopusti = append(dopust_list.Dopusti, kreiranDopust)
			jsonBytes, err := protojson.Marshal(dopust_list)
			if err != nil {
				log.Fatalf("Napaka pri Marshalanju: %v", err)
			}

			if err := ioutil.WriteFile("dopusti.json", jsonBytes, 0664); err != nil {
				log.Fatalf("Napaka pri pisanju v file -> %v", err)
			}
			return kreiranDopust, nil
		} else {
			log.Fatalf("Napaka pri branju datoteke -> %v", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, dopust_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	dopust_list.Dopusti = append(dopust_list.Dopusti, kreiranDopust)
	jsonBytes, err := protojson.Marshal(dopust_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("dopusti.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return kreiranDopust, nil
}

/*
func (s *Server) DopustStream(ctx context.Context, input *GetDopustiRequest, serv StreamDopust_DopustStreamServer) error {
	var wg sync.WaitGroup

	readBytes, err := ioutil.ReadFile("dopusti.json")
	if err != nil {
		log.Fatalf("Napaka pri branju datoteke -> %v", err)
	}

	var dopust_list *DopustList = &DopustList{}

	if err := protojson.Unmarshal(readBytes, dopust_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for id, dopust := range dopust_list.Dopusti {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			time.Sleep(time.Duration(count) * time.Millisecond)

			if err := serv.Send(dopust.DatumZacetka); err != nil {
				log.Printf("Napaka pri izvajanju streama! -> %v", err)
			}
			log.Printf("Zakljucen request #%d", id)

		}(int64(id * 200))
	}

	wg.Wait()

	return nil
}
*/

func (s *Server) DeleteDopust(ctx context.Context, input *GetDopustRequest) (*DeleteDopustResponse, error) {
	readBytes, err := ioutil.ReadFile("dopusti.json")
	if err != nil {
		log.Fatalf("Napaka pri branju datoteke -> %v", err)
	}

	var dopust_list *DopustList = &DopustList{}

	if err := protojson.Unmarshal(readBytes, dopust_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for id, dopust := range dopust_list.Dopusti {
		if dopust.GetIdDopust() == input.GetIdDopust() {
			copy(dopust_list.Dopusti[id:], dopust_list.Dopusti[id+1:])
			dopust_list.Dopusti[len(dopust_list.Dopusti)-1] = nil
			dopust_list.Dopusti = dopust_list.Dopusti[:len(dopust_list.Dopusti)-1]
		}
	}

	jsonBytes, err := protojson.Marshal(dopust_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("dopusti.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return &DeleteDopustResponse{Success: true}, nil
}

func (s *Server) DeleteZaposlen(ctx context.Context, input *GetZaposlenRequest) (*DeleteZaposlenResponse, error) {
	readBytes, err := ioutil.ReadFile("zaposleni.json")
	if err != nil {
		log.Fatalf("Napaka pri branju datoteke -> %v", err)
	}

	var zaposlen_list *ZaposlenList = &ZaposlenList{}

	if err := protojson.Unmarshal(readBytes, zaposlen_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for id, zaposlen := range zaposlen_list.Zaposleni {
		if zaposlen.GetIdZaposlenega() == input.GetIdZaposlen() {
			copy(zaposlen_list.Zaposleni[id:], zaposlen_list.Zaposleni[id+1:])
			zaposlen_list.Zaposleni[len(zaposlen_list.Zaposleni)-1] = nil
			zaposlen_list.Zaposleni = zaposlen_list.Zaposleni[:len(zaposlen_list.Zaposleni)-1]
		}
	}

	jsonBytes, err := protojson.Marshal(zaposlen_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("zaposleni.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return &DeleteZaposlenResponse{Success: true}, nil
}

func (s *Server) UpdateDopust(ctx context.Context, input *UpdateDopustRequest) (*Dopust, error) {
	readBytes, err := ioutil.ReadFile("dopusti.json")
	var dopust_list *DopustList = &DopustList{}
	var updejtanId int64 = input.GetId()

	kreiranDopust := &Dopust{IdDopust: updejtanId, IdZaposlenega: input.GetIdZaposlenega(),
		DatumZacetka: input.GetDatumZacetka(), DatumKonca: input.GetDatumKonca(),
		VrstaDopusta: input.GetVrstaDopusta(), StevilkaDopusta: input.GetStevilkaDopusta()}

	if err := protojson.Unmarshal(readBytes, dopust_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for id, dopust := range dopust_list.Dopusti {
		if dopust.GetIdDopust() == updejtanId {
			dopust_list.Dopusti[id] = kreiranDopust
		}
	}

	jsonBytes, err := protojson.Marshal(dopust_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("dopusti.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return kreiranDopust, nil
}

func (s *Server) UpdateZaposlen(ctx context.Context, input *UpdateZaposlenRequest) (*Zaposlen, error) {
	readBytes, err := ioutil.ReadFile("zaposleni.json")
	var zaposlen_list *ZaposlenList = &ZaposlenList{}
	var updejtanId int64 = input.GetIdZaposlenega()

	kreiranUporabnik := &Zaposlen{IdZaposlenega: updejtanId, Ime: input.GetIme(), Priimek: input.GetPriimek(),
		Spol: input.GetSpol(), DatumRojstva: input.GetDatumRojstva(), Podjetje: input.GetPodjetje()}

	if err := protojson.Unmarshal(readBytes, zaposlen_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for id, zaposlen := range zaposlen_list.Zaposleni {
		if zaposlen.GetIdZaposlenega() == updejtanId {
			zaposlen_list.Zaposleni[id] = kreiranUporabnik
		}
	}

	jsonBytes, err := protojson.Marshal(zaposlen_list)

	if err != nil {
		log.Fatalf("JSON marshilanje failed -> %v", err)
	}

	if err := ioutil.WriteFile("zaposleni.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Napaka pri pisanju v datoteko -> %v", err)
	}

	return kreiranUporabnik, nil
}
