# IIC2523_T1_2

**Grupo:**
- Javiera Inostroza
- Oscar Cárcamo

## Respuesta Conceptos

**1. Si dos tareas se ejecutan en paralelo, ¿cuál de las siguientes afirmaciones es correcta?**
- Están utilizando hardware diferente, pero se ejecutan al mismo tiempo.

**2. ¿Qué dice el cuello de botella de von Neumann sobre las arquitecturas de computador?**
- El tiempo de acceso a la memoria es un cuello de botella del rendimiento

**3. ¿Qué dice la ley de Moore?**
- La potencia del procesador se duplica cada 2 años.

**4. Una de las ventajas de la ejecución concurrente en un único procesador es que puede ocultar la latencia. ¿Qué significa esto?**
- Una tarea puede ejecutarse mientras otra tarea está esperando algo.


**5. ¿Qué es la exclusión mutua?**
- Cuando varias goroutinas no pueden ejecutar bloques de código simultáneamente.

**6. ¿Qué aseveraciones son ciertas acerca de los deadlocks?**
- Está causado por una cadena de dependencia circular entre subrutinas.
- Puede ser causado por la espera en los canales.

**7. ¿Cuál es el método de sync.mutex que debe ser llamado al principio de una región crítica?**
- Lock().

**8. ¿Qué línea de código podría usarse para definir un loop que lea iterativamente de un canal ch1?**
- for i := range ch1

**9. ¿Qué hace el keyword select?**
- Ejecuta un conjunto de sentencias case.

**10. ¿Qué significa la cláusula default dentro de un select?**
- La cláusula por defecto se ejecuta si todas las cláusulas case están bloqueadas.



## Respuesta P1 - Explicación No Determinismo en concurrencia.

Lo que está ocurriendo con el programa es que se están creando diversas rutinas que se ejecutan de forma paralela al hilo principal de ejecución donde cada una de estas recibe la dirección de memoria de la misma variable, es decir, están haciendo uso de una memoria compartida de forma simultanea. Entonces al extraer el valor de dicha dirección de memoria, como todo está ocurriendo en paralelo, se da el caso de que las rutinas intentan cambiar el valor al mismo tiempo, sin esperar a que uno termine primero, por lo que el ultimo en hacer los cambios no se alcanza a informar del ultimo valor de la variable y sobreescribe los cambios de las otras rutinas.
Entonces, cada vez que se ejecute el programa se podrán ver resultados distintos que dependeran del azar con las que una rutina termine antes que otra.

## Respuesta P2 - Explicación estructura del código.

**Supuesto: Se deben de ingresar exactamente 16 digitos por consola al ejecutar el script ´subrutinas.go´**

Para ordenar el arreglo donde se almacenan los 16 inputs, se crean 4 subrutinas, donde cada una de estas se encarga de ordenar mediante ´mergeSort´ un subarreglo de largo 4. 

Luego, se espera a que se reciban mediante canales los subarreglos ordenados y se fusionan en dos arreglos que se pasan a dos subrutinas respectivamente para que se vuelvan a ordenar.

Ahora, con los dos arreglos ya ordenados se vuelven a fusionar y se mandan a una unica rutina que termina de ordenar el arreglo completo, haciendo ´print´ del resultado final.

Se puede notar que es una simulacion del algoritmo de ´merge sort´ pero haciendo uso de las ´go routines´.
