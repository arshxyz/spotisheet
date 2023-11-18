package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/arshxyz/spotisheet/template"
	"github.com/pkg/browser"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func BuildUrl(client_id, client_secret string) {
	url := fmt.Sprintf("https://accounts.spotify.com/en/authorize?response_type=code&client_id=%s&scope=user-read-playback-state user-modify-playback-state&redirect_uri=http://localhost:8888/callback&state=lol", client_id)
	browser.OpenURL(url)
}

func startServer(done chan bool, ci, cs string) {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		ci_cs := base64.StdEncoding.EncodeToString([]byte(ci + ":" + cs))
		key := fmt.Sprintf("%s:%s", r.URL.Query().Get("code"), ci_cs)
		reshtml := fmt.Sprintf(template.Html, key)
		fmt.Fprintf(w, reshtml)
		done <- true
	})
	go func() {
		http.ListenAndServe(":8888", nil)
		<-done
	}()
}

func main() {
	done := make(chan bool)
	ci := StringPrompt("Enter your Client ID: ")
	cs := StringPrompt("Enter your Client Secret:")
	go startServer(done, ci, cs)
	BuildUrl(ci, cs)
	<-done
	// Can't be bothered to do srv.Shutdown for something as simple as this
	// Assuming the server takes less than 5 seconds to serve a requst
	// ugh
	time.Sleep(5 * time.Second)
}
