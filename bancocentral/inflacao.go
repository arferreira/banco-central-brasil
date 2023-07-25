package bancocentral

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Inflacao struct {
	queryURL string
	acesso   *AcessarBancoCentral
	req      []byte
}

func NewInflacao() *Inflacao {
	i := &Inflacao{
		queryURL: "https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/inflacao",
		acesso:   NewAcessarBancoCentral("https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/inflacao"),
		req:      nil,
	}

	i.req, _ = i.acesso.GetURL() // Ignore error for simplicity, as per previous code

	return i
}

func (i *Inflacao) GetMetaTax() (float64, error) {
	inflacao := cleanContent(string(i.req))
	tax := regexp.MustCompile(`<div id=label>Meta</div><div id=rate><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(inflacao)
	if tax == nil {
		return 0, errors.New("atributo não encontrado")
	}

	metaTax, err := strconv.ParseFloat(strings.ReplaceAll(tax[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return metaTax, nil
}

func (i *Inflacao) GetAcumuladaTax() (float64, error) {
	inflacao := cleanContent(string(i.req))
	tax := regexp.MustCompile(`<div id=label>Acumulada</div><div id=rate><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(inflacao)
	if tax == nil {
		return 0, errors.New("atributo não encontrado")
	}

	acumuladaTax, err := strconv.ParseFloat(strings.ReplaceAll(tax[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return acumuladaTax, nil
}
