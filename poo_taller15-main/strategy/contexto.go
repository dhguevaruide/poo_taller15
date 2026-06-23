package strategy

import "poo_taller15-main/model"

type Contexto struct {
	Docente     model.Docente
	Competencia model.Competencia
	Evaluacion  model.Evaluacion

	ReglasAplicadas map[string]bool
}