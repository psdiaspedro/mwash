package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Data struct {
	Dia    string
	Mes    string
	Ano    string
	Datadb string
}

func (data *Data) VerificaData(dataString string) (Data, error) {
	quantidadeHifen := strings.Count(dataString, "-")
	dataFormatada := data.formataData(dataString)
	dataRaw := strings.Trim(dataString, "-")
	dataRaw = strings.TrimSpace(dataRaw)

	switch quantidadeHifen {
	case 0:
		d, erro := data.parseAno(dataFormatada, dataRaw)
		if erro != nil {
			return d, erro
		}
		return d, nil
	case 1:
		d, erro := data.parseAnoMes(dataFormatada, dataRaw)
		if erro != nil {
			return d, erro
		}
		return d, nil
	case 2:
		d, erro := data.parseAnoMesDia(dataFormatada, dataRaw)
		if erro != nil {
			return d, erro
		}
		return d, nil
	default:
		var d Data
		d.Dia = ""
		d.Mes = ""
		d.Ano = ""
		d.Datadb = ""
		return d, errors.New("formato de data inválido, formato esperado: AAAA-MM-DD")
	}
}

func (data *Data) formataData(dataString string) string {
	dataFormatada := strings.Trim(dataString, "-")
	dataFormatada = strings.TrimSpace(dataString)
	dataFormatada = strings.ReplaceAll(dataString, "-", "/")
	return dataFormatada
}

func (data *Data) parseAno(dataFormatada string, dataRaw string) (Data, error) {
	var d Data

	dataParsed, erro := time.Parse("2006", dataFormatada)
	if erro != nil {
		return d, errors.New("formato de data inválido, formato esperado: AAAA-MM-DD")
	}
	d.Dia = ""
	d.Mes = ""
	d.Ano = fmt.Sprintf("%d", dataParsed.Year())
	d.Datadb = dataRaw
	return d, nil
}

func (data *Data) parseAnoMes(dataFormatada string, dataRaw string) (Data, error) {
	var d Data

	dataParsed, erro := time.Parse("2006/01", dataFormatada)
	if erro != nil {
		return d, errors.New("formato de data inválido, formato esperado: AAAA-MM-DD")
	}
	d.Dia = ""
	d.Mes = fmt.Sprintf("%d", dataParsed.Month())
	d.Ano = fmt.Sprintf("%d", dataParsed.Year())
	d.Datadb = dataRaw
	return d, nil
}

func (data *Data) parseAnoMesDia(dataFormatada string, dataRaw string) (Data, error) {
	var d Data

	dataParsed, erro := time.Parse("2006/01/02", dataFormatada)
	if erro != nil {
		return d, errors.New("formato de data inválido, formato esperado: AAAA-MM-DD")
	}
	d.Dia = fmt.Sprintf("%d", dataParsed.Day())
	d.Mes = fmt.Sprintf("%d", dataParsed.Month())
	d.Ano = fmt.Sprintf("%d", dataParsed.Year())
	d.Datadb = dataRaw
	return d, nil
}

func (data *Data) GerarQueryString(d Data) (string, []any) {
	query := "select a.*, p.cliente_id, p.cidade, p.bairro, p.CEP, p.logadouro, p.numero, p.complemento, u.nome, u.email, u.contato from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id INNER JOIN usuarios u on u.id = p.cliente_id where "

	var valores []any

	if d.Dia == "" && d.Mes == "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? order by dia_agendamento asc"
		valores = append(valores, d.Ano)
	} else if d.Dia == "" && d.Mes != "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? and extract(month from dia_agendamento) = ? order by dia_agendamento asc"
		valores = append(valores, d.Ano)
		valores = append(valores, d.Mes)
	} else if d.Dia != "" && d.Mes != "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? and extract(month from dia_agendamento) = ? and extract(day from dia_agendamento) = ? order by checkout asc"
		valores = append(valores, d.Ano)
		valores = append(valores, d.Mes)
		valores = append(valores, d.Dia)
	}

	return query, valores
}

func (data *Data) GerarQueryStringUsuarioId(d Data, usuarioId uint64) (string, []any) {
	query := "select a.* from agendamentos a INNER JOIN propriedades p ON p.id = a.propriedade_id INNER JOIN usuarios u ON u.id = p.cliente_id where u.id = ? and "

	var valores []any
	valores = append(valores, usuarioId)

	if d.Dia == "" && d.Mes == "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? order by dia_agendamento asc"
		valores = append(valores, d.Ano)
	} else if d.Dia == "" && d.Mes != "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? and extract(month from dia_agendamento) = ? order by dia_agendamento asc"
		valores = append(valores, d.Ano)
		valores = append(valores, d.Mes)
	} else if d.Dia != "" && d.Mes != "" && d.Ano != "" {
		query += "extract(year from dia_agendamento) = ? and extract(month from dia_agendamento) = ? and extract(day from dia_agendamento) = ? order by dia_agendamento asc"
		valores = append(valores, d.Ano)
		valores = append(valores, d.Mes)
		valores = append(valores, d.Dia)
	}

	return query, valores
}
