package website

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Prismic ...
type Prismic struct {
	Results []struct {
		Data interface{} `json:"data"`
	} `json:"results"`
}

// PrismicRef ...
type PrismicRef struct {
	Refs []struct {
		ID  string `json:"id"`
		Ref string `json:"ref"`
	} `json:"refs"`
}

func getPrismicMasterRef() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(
		`%s`,
		os.Getenv("PRISMIC_URL"),
	), nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		ref := PrismicRef{}
		err = json.Unmarshal(r, &ref)
		if err != nil {
			return "", err
		}

		for _, refs := range ref.Refs {
			if refs.ID == "master" {
				return refs.Ref, nil
			}
		}
	}

	return "", fmt.Errorf("no master found")
}

func getPrismic(page string) (Prismic, error) {
	ref, err := getPrismicMasterRef()
	if err != nil {
		return Prismic{}, err
	}
	query := strings.Replace("%5B%5Bat(document.type,%22REPLACETHIS%22)%5D%5D", "REPLACETHIS", page, -1)
	orderings := strings.Replace("%5BREPLACETHIS%5D", "REPLACETHIS", "document.first_publication_date", -1)
	sendURL := fmt.Sprintf(
		`%s/documents/search?ref=%s&q=%s&orderings=%s`,
		os.Getenv("PRISMIC_URL"),
		ref,
		query,
		orderings,
	)

	fmt.Println(sendURL)

	req, err := http.NewRequest("GET", sendURL, nil)
	if err != nil {
		return Prismic{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Prismic{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Prismic{}, err
		}

		pr := Prismic{}
		err = json.Unmarshal(r, &pr)
		if err != nil {
			return pr, err
		}

		return pr, nil
	}

	return Prismic{}, nil
}
