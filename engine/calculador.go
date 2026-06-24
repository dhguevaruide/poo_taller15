package engine

import (
	"poo_taller15/model"
	"poo_taller15/strategy"
	"sort"
)

type Calculador struct {
	estrategias []strategy.EstrategiaPonderacion
}

func NewCalculador() *Calculador {
	return &Calculador{
		estrategias: []strategy.EstrategiaPonderacion{
			strategy.ReglaCienciasInvestigacion{},
			strategy.ReglaDirectoresGestion{},
			strategy.ReglaAntiguedad{},
			strategy.ReglaCoordinadoresAcademica{},
			strategy.ReglaAdministracionGestion{},
		},
	}
}

func (c *Calculador) CalcularPuntajeConReglas(
	e model.Evaluacion,
	docente model.Docente,
	competencia model.Competencia,
	antiguedad int,
) (float64, []string) {

	ctx := strategy.Contexto{
		Docente:         docente,
		Competencia:     competencia,
		Evaluacion:      e,
		ReglasAplicadas: []string{},
		Antiguedad:      antiguedad,
	}

	sort.Slice(c.estrategias, func(i, j int) bool {
		return c.estrategias[i].Prioridad() < c.estrategias[j].Prioridad()
	})

	puntaje := e.PuntajeBase

	for _, regla := range c.estrategias {
		if regla.Aplica(ctx) {
			puntaje += regla.Calcular(ctx)
			ctx.ReglasAplicadas = append(ctx.ReglasAplicadas, regla.Nombre())
		}
	}

	return puntaje, ctx.ReglasAplicadas
}
