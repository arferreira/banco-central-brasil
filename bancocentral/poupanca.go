package bancocentral

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Poupanca struct {
	queryURL string
	acesso   *AcessarBancoCentral
	req      []byte
	err      error
}

func NewPoupanca() *Poupanca {
	return &Poupanca{
		queryURL: "https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/poupanca",
		acesso:   NewAcessarBancoCentral("https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/poupanca"),
		req:      nil,
		err:      nil,
	}
}

// Helper function to clean content
func cleanContent(content string) string {
	fix := map[string]string{
		"&lt;": "<",
		"&gt;": ">",
	}
	for key, value := range fix {
		content = strings.ReplaceAll(content, key, value)
	}
	content = strings.ReplaceAll(content, "\r\n", "")
	content = strings.ReplaceAll(content, "/>    <content", "/> <content")
	return content
}

func (p *Poupanca) GetPoupancaTax() (float64, error) {
	p.req, p.err = p.acesso.GetURL() // Corrected method name here
	if p.err != nil {
		return 0, p.err
	}

	poupanca := cleanContent(string(p.req))
	regex := regexp.MustCompile(`<div id=value>(\d*[\.,]?\d+)</div>`)
	tax := regex.FindStringSubmatch(poupanca)
	if len(tax) < 2 {
		return 0, errors.New("atributo não encontrado")
	}

	taxFloat, err := strconv.ParseFloat(strings.ReplaceAll(tax[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return taxFloat, nil
}
