package strategy

type EstrategiaPonderacion interface {
	Calcular(ctx Contexto, puntaje float64) float64
	Aplica(ctx Contexto) bool
	Nombre() string
	Prioridad() int
}