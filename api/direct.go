package api

import (
	"errors"
	"fmt"
	"net/http"
)

func HandlePwdRequest(w http.ResponseWriter, r *http.Request) {
	
	// Get ready to handle errors
	var err error
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With,Content-Type, Accept")
	if (r.Method == http.MethodOptions) {
		w.WriteHeader(http.StatusOK);
		return
	}
	
	// Handle requests
	switch r.Method {
	case http.MethodPost:
		err = Get(w, r)
	case http.MethodPut:
		err = Put(w, r)
	case http.MethodPatch:
		err = Patch(w, r)
	case http.MethodDelete:
		err = Delete(w, r)
	default:
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
		w.Write([]byte(err.Error()))
	}
}

func HandleCmpRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With,Content-Type, Accept")
	if (r.Method == http.MethodOptions) {
		w.WriteHeader(http.StatusOK);
		return
	}

	switch r.Method {
	case http.MethodPatch:
		err = PatchCmp(w, r)
	default:
		err = errors.ErrUnsupported
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
		w.Write([]byte(err.Error()))
	}
}

func HandleAllRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With,Content-Type, Accept")
	if (r.Method == http.MethodOptions) {
		w.WriteHeader(http.StatusOK);
		return
	}

	switch r.Method {
	case http.MethodPost:
		err = GetAll(w, r)
	default:
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusBadRequest)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
		w.Write([]byte(err.Error()))
	}
}

func HandleTransfRequest(w http.ResponseWriter, r *http.Request) {
     fmt.Println("handle transf.")

	var err error
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With,Content-Type, Accept")
	if (r.Method == http.MethodOptions) {
		w.WriteHeader(http.StatusOK);
		return
	}

	switch r.Method {
		case http.MethodPost:
			err = PostTransf(w, r)
		default:
			err = errors.ErrUnsupported
			w.WriteHeader(http.StatusBadRequest)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
		w.Write([]byte(err.Error()))
	}
}

func HandleRecvRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With,Content-Type, Accept")

	if (r.Method == http.MethodOptions) {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
		case http.MethodPost:
			err = RecvRequest(w, r)
		default:
			err = errors.ErrUnsupported
			w.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
		w.Write([]byte(err.Error()))
	}
}
