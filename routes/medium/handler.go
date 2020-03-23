package medium

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/conclave-dev/celoist-backend/util"
)

const mediumFeedBase = "https://medium.com/feed/"

func getCelo(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("https://api.rss2json.com/v1/api.json?rss_url=%s%s", mediumFeedBase, "celohq"))
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.RespondWithError(err, r, w)
	}

	util.RespondWithData(body, w)
}
