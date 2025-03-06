package handler

import (
	"fmt"
	"net/http"
)

// HomeHandler provides a welcome message and API usage guide
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Welcome to the Country Information Service!\n"+
		"Available endpoints:\n"+
		"- /countryinfo/v1/info/{country_code} (Get general country info)\n"+
		"- /countryinfo/v1/population/{country_code}?limit=startYear-endYear (Get population data)\n"+
		"- /countryinfo/v1/status (Check API status)")
}
