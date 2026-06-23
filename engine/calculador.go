package engine

import (
	"poo_taller15/model"
	"poo_taller15/strategy"
)

func CalcularPuntajeConReglas(
	eval model.Evaluacion,
	doc model.Docente,
	comp model.Competencia,
	reglas []strategy.EstrategiaPonderacion,
) (float64, []string) {

	ctx := strategy.Contexto{
		Docente:        doc,
		Competencia:    comp,
		Evaluacion:     eval,
		PuntajeActual:  eval.PuntajeBase,
		ReglasAplicadas: make(map[string]bool),
	}

	puntaje := eval.PuntajeBase
	aplicadas := []string{}

	for _, regla := range reglas {

		ctx.PuntajeActual = puntaje

		if regla.Aplica(ctx) {

			cambio := regla.Calcular(ctx)
			puntaje += cambio

			aplicadas = append(aplicadas, regla.Nombre())
			ctx.ReglasAplicadas[regla.Nombre()] = true
		}
	}

	return puntaje, aplicadas
}