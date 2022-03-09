package main

import (
	"encoding/json"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	v1 "k8s.io/api/core/v1"
)

func main() {
	http.HandleFunc("/", acceptPod)
	log.Fatal(
		http.ListenAndServeTLS(":8443", "./server.crt", "./server.key", nil))
}

func acceptPod(w http.ResponseWriter, r *http.Request) {
	review := admissionv1.AdmissionReview{}
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.Write([]byte("could not read the admission review"))
		w.WriteHeader(400)
		return
	}

	pod := v1.Pod{}
	err = json.Unmarshal(review.Request.Object.Raw, (&pod))
	if err != nil {
		w.Write([]byte("could not read the object inside the admission review"))
		w.WriteHeader(400)
		return
	}
	review.Response = createResponse(review.Request)
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	encoder.Encode(review)
}

func podFromRequest(request *admissionv1.AdmissionRequest) (v1.Pod, error) {

	pod := v1.Pod{}
	err := json.Unmarshal(request.Object.Raw, (&pod))
	if err != nil {
		return v1.Pod{}, err
	}
	return pod, nil
}

func createResponse(request *admissionv1.AdmissionRequest) *admissionv1.AdmissionResponse {
	// no need to check the pod because we accept everything
	//pod, err := podFromRequest(request)
	//	if err != nil {
	//		log.Println(err)
	//	}
	response := admissionv1.AdmissionResponse{
		UID:      request.UID,
		Allowed:  true,
		Warnings: []string{"This is a super warning"},
	}

	return &response
}
