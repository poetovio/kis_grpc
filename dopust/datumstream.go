package dopust

import (
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

type StreamServer struct {
	UnimplementedStreamDopustServer
}

func (s *StreamServer) DopustStream(input *GetDopustiRequest, stream StreamDopust_DopustStreamServer) error {

	readBytes, err := ioutil.ReadFile("dopusti.json")
	if err != nil {
		log.Fatalf("Napaka pri branju datoteke -> %v", err)
	}

	var dopust_list *DopustList = &DopustList{}

	if err := protojson.Unmarshal(readBytes, dopust_list); err != nil {
		log.Fatalf("Napaka pri parasnju jsona -> %v", err)
	}

	for _, dopust := range dopust_list.Dopusti {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			time.Sleep(1 * time.Second)

			err := stream.SendMsg(&DatumStreamResponse{
				Id:              dopust.GetIdDopust(),
				IdZaposlenega:   dopust.GetIdZaposlenega(),
				DatumZacetka:    dopust.GetDatumZacetka(),
				DatumKonca:      dopust.GetDatumKonca(),
				VrstaDopusta:    dopust.GetVrstaDopusta(),
				StevilkaDopusta: dopust.GetStevilkaDopusta(),
			})
			if err != nil {
				log.Fatalf("Napaka pri posiljanju datuma preko streama -> %v", err)
			}
		}
	}

	return nil
}
