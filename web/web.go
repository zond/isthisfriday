package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		unix, err := strconv.ParseInt(r.FormValue("unixtime"), 10, 64)
		if err != nil {
			fmt.Fprintln(w, err)
			w.WriteHeader(400)
			return
		}
		t := time.Unix(unix, 0)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"friday": t.Weekday() == time.Friday,
		}); err != nil {
			fmt.Fprintln(w, err)
			w.WriteHeader(500)
			return
		}
		return
	})
}
