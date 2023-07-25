package bancocentral

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Selic struct {
	queryURL string
	acesso   *AcessarBancoCentral
	req      []byte
	err      error
}

func NewSelic() *Selic {
	return &Selic{
		queryURL: "https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/juros",
		acesso:   NewAcessarBancoCentral("https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/juros"),
		req:      nil,
		err:      nil,
	}
}

func (s *Selic) GetSelicMeta() (float64, error) {
	s.req, s.err = s.acesso.GetURL()
	if s.err != nil {
		return 0, s.err
	}

	selic := cleanContent(string(s.req))
	regex := regexp.MustCompile(`<div id=ratevalue>(\d*[\.,]?\d+)</div>`)
	selicMeta := regex.FindStringSubmatch(selic)
	if len(selicMeta) < 2 {
		return 0, errors.New("atributo não encontrado")
	}

	selicMetaFloat, err := strconv.ParseFloat(strings.ReplaceAll(selicMeta[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return selicMetaFloat, nil
}

func (s *Selic) GetDataSelicMeta() (string, error) {
	s.req, s.err = s.acesso.GetURL()
	if s.err != nil {
		return "", s.err
	}

	selic := cleanContent(string(s.req))
	regex := regexp.MustCompile(`<div=ratedate>([\d/]+)</div>`)
	search := regex.FindStringSubmatch(selic)
	if len(search) < 2 {
		return "", errors.New("atributo não encontrado")
	}

	return search[1], nil
}

func (s *Selic) GetSelicReal() (float64, error) {
	s.req, s.err = s.acesso.GetURL()
	if s.err != nil {
		return 0, s.err
	}

	selic := cleanContent(string(s.req))
	regex := regexp.MustCompile(`<div id=dailyratevalue>(\d*[\.,]?\d+)</div>`)
	selicReal := regex.FindStringSubmatch(selic)
	if len(selicReal) < 2 {
		return 0, errors.New("atributo não encontrado")
	}

	selicRealFloat, err := strconv.ParseFloat(strings.ReplaceAll(selicReal[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return selicRealFloat, nil
}

func (s *Selic) GetDataSelicReal() (string, error) {
	s.req, s.err = s.acesso.GetURL()
	if s.err != nil {
		return "", s.err
	}

	selic := cleanContent(string(s.req))
	regex := regexp.MustCompile(`<div id=dailyratedate>([\d/]+)</div>`)
	search := regex.FindStringSubmatch(selic)
	if len(search) < 2 {
		return "", errors.New("atributo não encontrado")
	}

	return search[1], nil
}
