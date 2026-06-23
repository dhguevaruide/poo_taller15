package strategy

type ReglaCienciasInvestigacion struct{}

func (r ReglaCienciasInvestigacion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Ciencias" &&
		ctx.Competencia.Tipo == "investigación"
}

func (r ReglaCienciasInvestigacion) Calcular(ctx Contexto, p float64) float64 {
	return p * 1.10
}

func (r ReglaCienciasInvestigacion) Nombre() string {
	return "R1 Bonificación Ciencias Investigación"
}

func (r ReglaCienciasInvestigacion) Prioridad() int {
	return 1
}

type ReglaDirectorGestion struct{}

func (r ReglaDirectorGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Director" &&
		ctx.Competencia.Tipo == "gestión"
}

func (r ReglaDirectorGestion) Calcular(ctx Contexto, p float64) float64 {
	return p * 0.90
}

func (r ReglaDirectorGestion) Nombre() string {
	return "R2 Reducción Directores Gestión"
}

func (r ReglaDirectorGestion) Prioridad() int {
	return 1
}

type ReglaAntiguedad struct{}

func (r ReglaAntiguedad) Aplica(ctx Contexto) bool {

	r1Aplicada :=
		ctx.ReglasAplicadas["R1 Bonificación Ciencias Investigación"]

	return ctx.Docente.AniosAntiguedad > 10 &&
		r1Aplicada
}

func (r ReglaAntiguedad) Calcular(ctx Contexto, p float64) float64 {
	return p * 1.05
}

func (r ReglaAntiguedad) Nombre() string {
	return "R3 Bonificación Antigüedad"
}

func (r ReglaAntiguedad) Prioridad() int {
	return 2
}

type ReglaCoordinadorAcademica struct{}

func (r ReglaCoordinadorAcademica) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Coordinador" &&
		ctx.Competencia.Tipo == "académica"
}

func (r ReglaCoordinadorAcademica) Calcular(ctx Contexto, p float64) float64 {
	return p * 1.05
}

func (r ReglaCoordinadorAcademica) Nombre() string {
	return "R4 Bonificación Coordinadores Académica"
}

func (r ReglaCoordinadorAcademica) Prioridad() int {
	return 1
}

type ReglaAdministracionGestion struct{}

func (r ReglaAdministracionGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Administración" &&
		ctx.Competencia.Tipo == "gestión"
}

func (r ReglaAdministracionGestion) Calcular(ctx Contexto, p float64) float64 {
	return p * 1.15
}

func (r ReglaAdministracionGestion) Nombre() string {
	return "R5 Bonificación Administración Gestión"
}

func (r ReglaAdministracionGestion) Prioridad() int {
	return 1
}