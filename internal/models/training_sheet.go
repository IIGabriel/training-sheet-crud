package models

type TrainingSheet struct {
	ID            uint   `json:"id"`
	FuncionarioID uint   `json:"funcionario_id"`
	Modalidade    string `json:"modalidade"`
	Exercicios    string `json:"exercicios"`
	Ativo         bool   `json:"ativo"`
}
