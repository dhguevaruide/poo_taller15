# Sistema de Evaluación de Competencias con Strategy

## Justificación de Diseño

### 1. ¿Por qué Strategy en lugar de otros patrones?

Se eligió Strategy porque cada regla de evaluación representa un algoritmo independiente que puede ser agregado, eliminado o modificado sin afectar las demás reglas. Chain of Responsibility está orientado a pasar solicitudes entre objetos hasta encontrar quién la procese, mientras que Visitor requiere modificar la estructura de los elementos visitados cuando se agregan nuevas operaciones. Strategy ofrece mayor flexibilidad para reglas dinámicas.

### 2. Manejo de prioridades

Las estrategias se ordenan utilizando sort.SliceStable(). Cuando dos reglas poseen la misma prioridad, se conserva el orden de inserción original, garantizando resultados deterministas.

### 3. Composición vs Agregación

El Contexto posee una relación de composición con Docente, Competencia y Evaluacion, ya que dichos objetos forman parte del proceso de evaluación y el contexto no tiene sentido sin ellos. Por ello su ciclo de vida está ligado al contexto durante el cálculo.

### 4. Extensibilidad para historial

Se podría agregar un nuevo componente RepositorioHistorial dentro del Contexto. Las estrategias consultarían este repositorio mediante interfaces, permitiendo crear reglas basadas en evaluaciones anteriores sin modificar las estrategias existentes.

### 5. Análisis de performance

La complejidad temporal sería O(n × r), donde n representa las evaluaciones y r las reglas. Para 10.000 evaluaciones y 50 reglas se realizarían aproximadamente 500.000 verificaciones, una carga perfectamente manejable para Go. Además, el uso de interfaces permite paralelizar el procesamiento mediante goroutines si el volumen aumenta.
