package strategy

import "strings"

type ReglaCienciasInvestigacion struct{}

func (r ReglaCienciasInvestigacion) Nombre() string {
	return "R1 - Bonificación Ciencias Investigación"
}

func (r ReglaCienciasInvestigacion) Prioridad() int {
	return 1
}

func (r ReglaCienciasInvestigacion) Aplica(ctx *Contexto) bool {
	return ctx.Docente.Departamento == "Ciencias" &&
		ctx.Competencia.Tipo == "investigacion"
}

func (r ReglaCienciasInvestigacion) Calcular(ctx *Contexto) float64 {
	return ctx.Evaluacion.PuntajeFinal * 0.10
}

type ReglaDirectorGestion struct{}

func (r ReglaDirectorGestion) Nombre() string {
	return "R2 - Reducción Directores Gestión"
}

func (r ReglaDirectorGestion) Prioridad() int {
	return 1
}

func (r ReglaDirectorGestion) Aplica(ctx *Contexto) bool {
	return ctx.Docente.Cargo == "Director" &&
		ctx.Competencia.Tipo == "gestion"
}

func (r ReglaDirectorGestion) Calcular(ctx *Contexto) float64 {
	return -(ctx.Evaluacion.PuntajeFinal * 0.10)
}

type ReglaAntiguedad struct{}

func (r ReglaAntiguedad) Nombre() string {
	return "R3 - Bonificación Antigüedad"
}

func (r ReglaAntiguedad) Prioridad() int {
	return 2
}

func (r ReglaAntiguedad) Aplica(ctx *Contexto) bool {

	tieneR1 := false

	for _, regla := range ctx.ReglasAplicadas {
		if strings.Contains(regla, "R1") {
			tieneR1 = true
			break
		}
	}

	return ctx.Docente.Antiguedad > 10 && tieneR1
}

func (r ReglaAntiguedad) Calcular(ctx *Contexto) float64 {
	return ctx.Evaluacion.PuntajeFinal * 0.05
}

type ReglaCoordinadorAcademica struct{}

func (r ReglaCoordinadorAcademica) Nombre() string {
	return "R4 - Bonificación Coordinadores Académica"
}

func (r ReglaCoordinadorAcademica) Prioridad() int {
	return 1
}

func (r ReglaCoordinadorAcademica) Aplica(ctx *Contexto) bool {
	return ctx.Docente.Cargo == "Coordinador" &&
		ctx.Competencia.Tipo == "academica"
}

func (r ReglaCoordinadorAcademica) Calcular(ctx *Contexto) float64 {
	return ctx.Evaluacion.PuntajeFinal * 0.05
}

type ReglaAdministracionGestion struct{}

func (r ReglaAdministracionGestion) Nombre() string {
	return "R5 - Bonificación Administración Gestión"
}

func (r ReglaAdministracionGestion) Prioridad() int {
	return 1
}

func (r ReglaAdministracionGestion) Aplica(ctx *Contexto) bool {
	return ctx.Docente.Departamento == "Administración" &&
		ctx.Competencia.Tipo == "gestion"
}

func (r ReglaAdministracionGestion) Calcular(ctx *Contexto) float64 {
	return ctx.Evaluacion.PuntajeFinal * 0.15
}