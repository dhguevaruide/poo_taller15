# TALLER: Patrón Strategy - Sistema de Evaluación de Competencias

Este repositorio contiene la implementación de un sistema dinámico de reglas para la evaluación docente utilizando el patrón de diseño **Strategy** en Go.

---

## Justificación de Diseño (Respuestas al Cuestionario)

### 1. ¿Por qué elegiste el patrón Strategy en lugar de Chain of Responsibility o Visitor?
* **Frente a Chain of Responsibility (CoR):** CoR está diseñado para que una petición sea procesada por un eslabón de la cadena y se detenga o pase al siguiente si este no pudo manejarla. En nuestro sistema, **múltiples reglas deben aplicar de forma acumulativa e independiente** sobre la misma evaluación. El patrón Strategy nos permite evaluar todas las estrategias en un orden específico sin que una interrumpa a las demás.
* **Frente a Visitor:** Visitor se utiliza para añadir operaciones a una estructura de objetos compleja y heterogénea sin modificar las clases de dicha estructura. En este problema, la estructura de datos es simple (`Docente`, `Competencia`), y el núcleo del requerimiento consiste en intercambiar dinámicamente los algoritmos de cálculo, que es el propósito exacto de **Strategy**.

### 2. ¿Cómo manejaste el caso donde dos reglas tienen la misma prioridad?
Se manejó implementando un ordenamiento estable en el `Calculador` mediante la función `sort.Slice` de Go. Si dos estrategias concretas devuelven exactamente el mismo valor numérico en su método `Prioridad()`, el motor de evaluación mantiene el **orden natural de registro** en el que fueron añadidas al slice de reglas. Esto garantiza un comportamiento determinista y predecible.

### 3. ¿Por qué usaste composición en lugar de agregación para ciertas relaciones?
En el struct `strategy.Contexto`, las entidades `Docente`, `Competencia` y `Evaluacion` se acoplan mediante **composición por valor** (copias de los objetos dentro del contexto). 
* **Ejemplo:** El ciclo de vida de este objeto `Contexto` es efímero; nace al iniciar el cálculo de una evaluación y muere al terminarlo. Al pasar las estructuras por valor, nos aseguramos de que el motor de reglas trabaje con una "fotografía" inmutable de los datos en ese instante de tiempo, evitando efectos secundarios o mutaciones concurrentes accidentales de la data original.

### 4. ¿Cómo extenderías el sistema para soportar reglas que dependen del historial del docente?
Para soportar reglas históricas (por ejemplo, *"si el docente mantuvo una calificación excelente los últimos 3 años"*), el diseño se extendería de la siguiente manera:
1. Se añadiría un campo al struct `Contexto`: `HistorialEvaluaciones []model.Evaluacion`.
2. Al inicializar el contexto en el `Calculador`, se consultaría la base de datos para cargar este histórico.
3. Se crearía una nueva estrategia (ej. `ReglaHistorialExcelente`) que implemente `EstrategiaPonderacion`. En su método `Aplica(ctx)`, leería el slice de históricos para retornar `true` o `false`. **Ninguna interfaz existente ni regla previa tendría que modificarse**, cumpliendo con el principio Open/Closed.

### 5. ¿Qué impacto tiene tu diseño en la performance si hay 10,000 evaluaciones con 50 reglas cada una?
El impacto del algoritmo actual es de complejidad lineal $O(N \times M)$, donde $N$ es el número de evaluaciones (10,000) y $M$ el número de reglas (50). Esto se traduce en 500,000 iteraciones simples con evaluaciones booleanas en memoria, lo cual toma pocos milisegundos.
Dado que cada proceso de evaluación es completamente independiente y su `Contexto` es único y no comparte memoria mutable, la solución es **altamente paralelizable**. En un entorno real, se podría optimizar utilizando un pool de **Goroutines** para procesar las 10,000 evaluaciones en paralelo, aprovechando el modelo de concurrencia nativo de Go sin riesgo de condiciones de carrera (*race conditions*).