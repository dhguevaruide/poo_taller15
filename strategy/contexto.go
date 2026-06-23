package strategy

import "poo_taller15/model"

type Contexto struct {
	Docente     model.Docente
	Competencia model.Competencia
	Evaluacion  model.Evaluacion

	PuntajeActual float64

	ReglasAplicadas map[string]bool
}