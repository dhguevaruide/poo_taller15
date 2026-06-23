package engine

import (
	"poo_taller15/model"
	"poo_taller15/strategy"
	"sort"
)

type CalculadorReglas struct {
	Estrategias []strategy.EstrategiaPonderacion
}

func NuevoCalculador() *CalculadorReglas {

	estrategias := []strategy.EstrategiaPonderacion{
		strategy.ReglaCienciasInvestigacion{},
		strategy.ReglaDirectorGestion{},
		strategy.ReglaCoordinadorAcademica{},
		strategy.ReglaAdministracionGestion{},
		strategy.ReglaAntiguedad{},
	}

	sort.SliceStable(estrategias,
		func(i, j int) bool {
			return estrategias[i].Prioridad() <
				estrategias[j].Prioridad()
		})

	return &CalculadorReglas{
		Estrategias: estrategias,
	}
}

func (c *CalculadorReglas) CalcularPuntajeConReglas(
	docente model.Docente,
	competencia model.Competencia,
	evaluacion model.Evaluacion,
) (float64, []string) {

	ctx := strategy.Contexto{
		Docente:     docente,
		Competencia: competencia,
		Evaluacion:  evaluacion,
	}

	ctx.Evaluacion.PuntajeFinal =
		ctx.Evaluacion.PuntajeBase

	for _, regla := range c.Estrategias {

		if regla.Aplica(&ctx) {

			ajuste := regla.Calcular(&ctx)

			ctx.Evaluacion.PuntajeFinal += ajuste

			ctx.ReglasAplicadas =
				append(
					ctx.ReglasAplicadas,
					regla.Nombre(),
				)
		}
	}

	return ctx.Evaluacion.PuntajeFinal,
		ctx.ReglasAplicadas
}