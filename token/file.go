package token

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const (
	TokenFilePostfix = "_TOKEN_FILE"
)

func FromFile(envPrefix string) (*oauth2.Token, error) {
	filename := getFilename(envPrefix)
	tokenFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer tokenFile.Close()

	tok := &oauth2.Token{}
	jd := json.NewDecoder(tokenFile)
	if err = jd.Decode(tok); err != nil {
		return nil, err
	}

	return tok, nil
}

func ToFile(envPrefix string, oauthClient *http.Client) {
	filename := getFilename(envPrefix)

	// todo: сделать сравнение текущего токена с тем, что лежит в файле во
	//       избежание бесполезной/неудачной записи
	tokenFile, err := os.Create(filename)
	if err != nil {
		return
	}
	defer tokenFile.Close()

	tokenSource, ok := oauthClient.Transport.(*oauth2.Transport)
	if !ok {
		return
	}

	tok, err := tokenSource.Source.Token()
	if err != nil {
		return
	}

	je := json.NewEncoder(tokenFile)
	if err = je.Encode(tok); err != nil {
		return
	}
}

func getFilename(prefix string) string {
	return os.Getenv(prefix + TokenFilePostfix)
}
