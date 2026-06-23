# Sistema de Evaluación de Competencias con Strategy

## Justificación de Diseño

### 1. ¿Por qué Strategy en lugar de otros patrones?

Se eligió el patrón Strategy porque permite encapsular cada regla de evaluación como una clase independiente. Esto facilita agregar, eliminar o modificar reglas sin afectar el sistema principal. A diferencia de Chain of Responsibility, aquí no se detiene la ejecución, sino que se pueden aplicar múltiples reglas sobre el mismo contexto.

---

### 2. Manejo de prioridades

Las reglas tienen un método `Prioridad()` que permite ordenarlas antes de ejecutarlas. Esto garantiza un orden controlado de aplicación, evitando inconsistencias en el cálculo del puntaje.

---

### 3. Composición vs Agregación

Se utiliza composición en el `Contexto` porque el objeto depende directamente de `Docente`, `Competencia` y `Evaluacion` para funcionar. Sin estos elementos, el sistema no puede calcular el puntaje, por lo que su ciclo de vida está fuertemente ligado.

---

### 4. Extensibilidad para historial

El sistema puede extenderse agregando un campo `Historial []Evaluacion` en el Contexto. Esto permitiría crear nuevas reglas basadas en evaluaciones anteriores del docente, como promedio histórico o desempeño acumulado.

---

### 5. Análisis de performance

El sistema tiene una complejidad O(n * r), donde n es el número de evaluaciones y r el número de reglas. Aunque puede escalar a miles de evaluaciones, el impacto se mantiene controlado porque cada regla es una operación simple de comparación y cálculo.

---

## Ejecución

Para ejecutar el proyecto:

```bash
go run main.go