package strategy

type ReglaCienciasInvestigacion struct{}

func (r ReglaCienciasInvestigacion) Nombre() string {
	return "R1 - Bonificación Ciencias Investigación"
}
func (r ReglaCienciasInvestigacion) Prioridad() int { return 1 }
func (r ReglaCienciasInvestigacion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Ciencias" && ctx.Competencia.Tipo == "investigacion"
}
func (r ReglaCienciasInvestigacion) Calcular(ctx Contexto) float64 {
	return ctx.Evaluacion.PuntajeBase * 0.10
}

type ReglaDirectoresGestion struct{}

func (r ReglaDirectoresGestion) Nombre() string {
	return "R2 - Reducción Directores Gestión"
}
func (r ReglaDirectoresGestion) Prioridad() int { return 2 }
func (r ReglaDirectoresGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Director" && ctx.Competencia.Tipo == "gestion"
}
func (r ReglaDirectoresGestion) Calcular(ctx Contexto) float64 {
	return ctx.Evaluacion.PuntajeBase * -0.10
}

type ReglaAntiguedad struct{}

func (r ReglaAntiguedad) Nombre() string {
	return "R3 - Bonificación Antigüedad (>10 años)"
}
func (r ReglaAntiguedad) Prioridad() int { return 3 }
func (r ReglaAntiguedad) Aplica(ctx Contexto) bool {
	for _, regla := range ctx.ReglasAplicadas {
		if regla == "R1 - Bonificación Ciencias Investigación" {
			return ctx.Antiguedad > 10
		}
	}
	return false
}
func (r ReglaAntiguedad) Calcular(ctx Contexto) float64 {
	return ctx.Evaluacion.PuntajeBase * 0.05
}

type ReglaCoordinadoresAcademica struct{}

func (r ReglaCoordinadoresAcademica) Nombre() string {
	return "R4 - Bonificación Coordinadores Académica"
}
func (r ReglaCoordinadoresAcademica) Prioridad() int { return 4 }
func (r ReglaCoordinadoresAcademica) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Coordinador" && ctx.Competencia.Tipo == "academica"
}
func (r ReglaCoordinadoresAcademica) Calcular(ctx Contexto) float64 {
	return ctx.Evaluacion.PuntajeBase * 0.05
}

type ReglaAdministracionGestion struct{}

func (r ReglaAdministracionGestion) Nombre() string {
	return "R5 - Bonificación Administración Gestión"
}
func (r ReglaAdministracionGestion) Prioridad() int { return 5 }
func (r ReglaAdministracionGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Administración" && ctx.Competencia.Tipo == "gestion"
}
func (r ReglaAdministracionGestion) Calcular(ctx Contexto) float64 {
	return ctx.Evaluacion.PuntajeBase * 0.15
}
