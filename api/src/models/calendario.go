package models

type Calendario struct {
	AgendamentoID  uint64 `json:"agendamentoid,omitempty"`
	PropriedadeID  uint64 `json:"propriedadeId,omitempty"`
	DiaAgendamento string `json:"diaAgendamento,omitempty"`
	Checkin        string `json:"checkin,omitempty"`
	Checkout       string `json:"checkout,omitempty"`
	Obs            string `json:"obs,omitempty"`
	ProprietarioID uint64 `json:"proprietarioid,omitempty"`
	Cidade         string `json:"cidade,omitempty"`
	Bairro         string `json:"bairro,omitempty"`
	CEP            string `json:"cep,omitempty"`
	Logadouro      string `json:"logadouro,omitempty"`
	Numero         string `json:"numero,omitempty"`
	Complemento    string `json:"complemento,omitempty"`
	Nome           string `json:"nome,omitempty"`
	Email          string `json:"email,omitempty"`
}
