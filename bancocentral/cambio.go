package bancocentral

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// TODO: Implement tests
type Cambio struct {
	queryURL string
	acesso   *AcessarBancoCentral
	req      []byte
	cambio   string
	err      error
}

func NewCambio() *Cambio {
	c := &Cambio{
		queryURL: "https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/cambio",
		acesso:   NewAcessarBancoCentral("https://conteudo.bcb.gov.br/api/feed/pt-br/PAINEL_INDICADORES/cambio"),
		req:      nil,
		cambio:   "",
		err:      nil,
	}

	// Realize a solicitação HTTP e capture a resposta e o erro
	c.req, c.err = c.acesso.GetURL()
	if c.err != nil {
		// Se houver um erro, retorne a instância de `Cambio` com o erro
		return c
	}

	// Se a solicitação for bem-sucedida, atualize o campo `cambio` com o conteúdo limpo
	c.cambio = cleanContent(string(c.req))

	return c
}

func (c *Cambio) GetDolarCompraPtax() (float64, error) {
	data := regexp.MustCompile(`<entry>.*?<id>painel_indicadores_INDICADOR_CAMBIO_DOLAR_PTAX</id>(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	compra := regexp.MustCompile(`<div id=rate><div id=label>Compra</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if compra == nil {
		return 0, errors.New("atributo não encontrado")
	}

	dolarCompraPtax, err := strconv.ParseFloat(strings.ReplaceAll(compra[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return dolarCompraPtax, nil
}

func (c *Cambio) GetDolarVendaPtax() (float64, error) {
	data := regexp.MustCompile(`<entry>.*?<id>painel_indicadores_INDICADOR_CAMBIO_DOLAR_PTAX</id>(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo INDICADOR_CAMBIO_DOLAR_PTAX não encontrado")
	}

	venda := regexp.MustCompile(`<div id=rate><div id=label>Venda</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if venda == nil {
		return 0, errors.New("atributo Venda não encontrado")
	}

	dolarVendaPtax, err := strconv.ParseFloat(strings.ReplaceAll(venda[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return dolarVendaPtax, nil
}

func (c *Cambio) GetDataDolarPtax() (string, error) {
	data := regexp.MustCompile(`<entry>.*?<id>painel_indicadores_INDICADOR_CAMBIO_DOLAR_PTAX</id>(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return "", errors.New("atributo INDICADOR_CAMBIO_DOLAR_PTAX não encontrado")
	}

	log.Println("data: ", data)

	rate := regexp.MustCompile(`<div id=rate>.*?<div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if rate == nil {
		return "", errors.New("valor de PTAX não encontrado")
	}

	log.Println("rate", rate)

	dataDolar := regexp.MustCompile(`<div id=data>([\d/]+\s[\d:]+)`).FindStringSubmatch(data[1])
	if dataDolar == nil {
		return "", errors.New("data de PTAX não encontrada")
	}

	ptaxValue := strings.ReplaceAll(rate[1], ",", ".")
	return fmt.Sprintf("%s %s", ptaxValue, dataDolar[1]), nil
}

func (c *Cambio) GetDolarCompra() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_DOLAR(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	compra := regexp.MustCompile(`<div id=rate><div id=label>Compra</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if compra == nil {
		return 0, errors.New("atributo não encontrado")
	}

	dolarCompra, err := strconv.ParseFloat(strings.ReplaceAll(compra[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return dolarCompra, nil
}

func (c *Cambio) GetDolarVenda() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_DOLAR(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	venda := regexp.MustCompile(`<div id=rate><div id=label>Venda</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if venda == nil {
		return 0, errors.New("atributo não encontrado")
	}

	dolarVenda, err := strconv.ParseFloat(strings.ReplaceAll(venda[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return dolarVenda, nil
}

func (c *Cambio) GetDataDolar() (string, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_DOLAR(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return "", errors.New("atributo não encontrado")
	}

	search := regexp.MustCompile(`<div id=data>[a-zA-Z\s]*([\d/]+\s[\d:]+)</div>`).FindStringSubmatch(data[1])
	if search == nil {
		return "", errors.New("atributo não encontrado")
	}

	return search[1], nil
}

func (c *Cambio) GetEuroCompraPtax() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO_PTAX(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	compra := regexp.MustCompile(`<div id=rate><div id=label>Compra</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if compra == nil {
		return 0, errors.New("atributo não encontrado")
	}

	euroCompraPtax, err := strconv.ParseFloat(strings.ReplaceAll(compra[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return euroCompraPtax, nil
}

func (c *Cambio) GetEuroVendaPtax() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO_PTAX(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	venda := regexp.MustCompile(`<div id=rate><div id=label>Venda</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if venda == nil {
		return 0, errors.New("atributo não encontrado")
	}

	euroVendaPtax, err := strconv.ParseFloat(strings.ReplaceAll(venda[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return euroVendaPtax, nil
}

func (c *Cambio) GetDataEuroPtax() (string, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO_PTAX(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return "", errors.New("atributo não encontrado")
	}

	search := regexp.MustCompile(`<div id=data>[a-zA-Z\s]*([\d/]+\s[\d:]+)</div>`).FindStringSubmatch(data[1])
	if search == nil {
		return "", errors.New("atributo não encontrado")
	}

	return search[1], nil
}

func (c *Cambio) GetEuroCompra() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	compra := regexp.MustCompile(`<div id=rate><div id=label>Compra</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if compra == nil {
		return 0, errors.New("atributo não encontrado")
	}

	euroCompra, err := strconv.ParseFloat(strings.ReplaceAll(compra[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return euroCompra, nil
}

func (c *Cambio) GetEuroVenda() (float64, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return 0, errors.New("atributo não encontrado")
	}

	venda := regexp.MustCompile(`<div id=rate><div id=label>Venda</div><div id=value>(\d*[\.,]?\d+)</div>`).FindStringSubmatch(data[1])
	if venda == nil {
		return 0, errors.New("atributo não encontrado")
	}

	euroVenda, err := strconv.ParseFloat(strings.ReplaceAll(venda[1], ",", "."), 64)
	if err != nil {
		return 0, err
	}

	return euroVenda, nil
}

func (c *Cambio) GetDataEuro() (string, error) {
	data := regexp.MustCompile(`INDICADOR_CAMBIO_EURO(.*?)</entry>`).FindStringSubmatch(c.cambio)
	if data == nil {
		return "", errors.New("atributo não encontrado")
	}

	search := regexp.MustCompile(`<div id=data>[a-zA-Z\s]*([\d/]+\s[\d:]+)</div>`).FindStringSubmatch(data[1])
	if search == nil {
		return "", errors.New("atributo não encontrado")
	}

	return search[1], nil
}
