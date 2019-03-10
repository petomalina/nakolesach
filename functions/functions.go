package functions

import (
	"context"
	"encoding/json"
	"log"
	"nakolesach/models"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const (
	collectionRides = "rides"
)

var (
	fsClient *firestore.Client
	fsRides  *firestore.CollectionRef
)

func init() {
	conf := &firebase.Config{ProjectID: "nakolesach-sk"}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	fsClient = client
	fsRides = client.Collection(collectionRides)
}

func Ride(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		var d models.Ride
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			w.WriteHeader(400)
			return
		}

		doc, _, err := fsRides.Add(context.Background(), &d)
		if err != nil {
			log.Printf("fsRides.Add: %v", err)
			w.WriteHeader(503)
			return
		}

		d.ID = doc.ID

		if err := json.NewEncoder(w).Encode(&d); err != nil {
			w.WriteHeader(503)
			log.Printf("json.Encode: %v", err)
			return
		}

	} else if r.Method == "GET" {
		id := r.URL.Path[1:]
		doc, err := fsRides.Doc(id).Get(context.Background())
		if err != nil {
			log.Printf("fsRide.Doc().Get(): %v", err)
			w.WriteHeader(503)
			return
		}

		d := models.Ride{}
		err = doc.DataTo(&d)
		if err != nil {
			log.Printf("doc.DataTo(): %v", err)
			w.WriteHeader(503)
			return
		}

		d.ID = id

		err = json.NewEncoder(w).Encode(&d)
		if err != nil {
			log.Printf("json.Encode(): %v", err)
			w.WriteHeader(503)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// } else if r.Method == "PUT" {
	// 	id := strings.TrimPrefix(r.URL.Path, "/Ride/")
	// 	var d models.Ride
	// 	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
	// 		w.WriteHeader(400)
	// 		return
	// 	}
	// 	fsRides.Doc(id).Update(ctx, updates, preconds)
	// 	firestore.Update{}
	// 	doc, err := fsRides.Doc(id).Get(context.Background())
	// }

	// if d.ID == "" {
	// 	fmt.Fprint(w, "Hello, World!")
	// 	return
	// }
	// fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.ID))
}

func ListRides(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	docs, err := fsRides.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("fsRides.GetAll: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := make([]models.Ride, len(docs))
	for i, doc := range docs {
		err = doc.DataTo(&result[i])
		if err != nil {
			log.Printf("doc.DataTo: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		result[i].ID = doc.Ref.ID
	}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Printf("json.Encode(): %v", err)
		w.WriteHeader(503)
	}
}
