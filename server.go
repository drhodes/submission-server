// Copyright <2021> (see CONTRIBUTERS file)
// for license, see LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	auth "github.com/abbot/go-http-auth"
	"golang.org/x/crypto/bcrypt"
)

// This function is passed up to the authenticator library for
// authentication purposes.
func secretPass(user, realm string) string {
	if user == os.Getenv("STAFF_SUBMITTER_USERID") {
		secret := []byte(os.Getenv("STAFF_SUBMITTER_PASSWD"))
		hashedPassword, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
		if err == nil {
			return string(hashedPassword)
		} else {
			log.Println(Err(err, "something happened while hashing the staff password"))
		}
	} else if user == "student" {
		secret := []byte("student")
		hashedPassword, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
		if err == nil {
			return string(hashedPassword)
		} else {
			log.Println(Err(err, "something happened while hashing the student password"))
		}
	}
	return ""
}

// this might not work! Make sure to check that the ip of the request
// is actually the ip of the host and not a proxy.

// another thing: check to make sure the addresses located in
// pod-tracker-store directory are ip4 address that look unique
// and seem to be from student pods.

// more info here: https://golangcode.com/get-the-request-ip-addr/
func getIP(r *auth.AuthenticatedRequest) string {
	addr := r.RemoteAddr
	xs := strings.Split(addr, ":")
	return strings.TrimSpace(xs[0])
}

// request comes in from
func podStartingHandler(w http.ResponseWriter, req *auth.AuthenticatedRequest) {
	// make sure the authed user is staff.
	if req.Username == "staff" {
		// userid will look like: jupyter-35dd7e9124c8847ec5-030ef
		edx_anon_id := req.FormValue("edx-anon-id")

		// make sure this userid is not just the edx anon_id, but full
		// jupyterhubized user id, which can be generated with
		// util.go/generate_jupyterhub_userid

		if edx_anon_id == "" {
			fmt.Fprintf(w, `{"ok": "false", "error": "edx-anon-id missing from form data"}`)
		} else {
			//recordPodStarting(edx_anon_id, req.
			ip := getIP(req)
			fmt.Printf("Got ip: %s\n", ip)
			pt := NewPodTracker("./pod-tracker-store")
			err := pt.StoreRecord(edx_anon_id, ip)
			logif(err)

			ok, err := pt.CheckRecord(edx_anon_id, ip)
			logif(err)
			if !ok {
				log.Println("Check record not working")
			}
		}

	} else {
		log.Println("student tried to enter pod starting mode", req)
	}
}

// request comes in from
func submitAnswerHandler(w http.ResponseWriter, req *auth.AuthenticatedRequest) {
	// make sure the authed user is student
	if req.Username != "student" {
		// fail with error back to magic.
		log.Println("student tried to enter pod starting mode", req)
	}

	// userid will look like: jupyter-35dd7e9124c8847ec5-030ef
	edxAnonId := req.FormValue("edx-anon-id")
	labName := req.FormValue("labname")

	if edxAnonId == "" {
		fmt.Fprintf(w, `{"ok": "false", "error": "edx-anon-id missing from form data"}`)
		return
	}

	ip := getIP(req)
	pt := NewPodTracker("pod-tracker-store")

	// check to make sure the ip address in the pod tracker
	// store is the same as the one in the request.
	matches, err := pt.CheckRecord(edxAnonId, ip)
	logif(err)

	if matches {
		answersJson := req.FormValue("lab-answers")
		err := sendAnswers(edxAnonId, labName, answersJson)
		if err != nil {
			log.Println(err)
			// log this and give the user some indication that
			// the answer was not sent.

			// TODO from within magic check the response code
			// from the request that was sent
			// w.Write(oh no!)
		}
	} else {
		// todo: figure out legit logging.
		log.Println("someone may be trying to cheat. TODO msg staff somehow")
	}
}

func serve() {
	host := ":3000"
	authenticator := auth.NewBasicAuthenticator(host, secretPass)
	http.HandleFunc("/pod-starting", authenticator.Wrap(podStartingHandler))
	http.HandleFunc("/submit-answers", authenticator.Wrap(submitAnswerHandler))
	log.Fatal(http.ListenAndServe(host, nil))
}
