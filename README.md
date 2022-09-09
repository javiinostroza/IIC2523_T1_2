# IIC2523_T1_2

**Grupo:**
- Javiera Inostroza
- Oscar Cárcamo

## Respuesta P1 - Explicación No Determinismo en concurrencia.

Lo que está ocurriendo con el programa es que se están creando diversas rutinas que se ejecutan de forma paralela al hilo principal de ejecución donde cada una de estas recibe la dirección de memoria de la misma variable, es decir, están haciendo uso de una memoria compartida de forma simultanea. Entonces al extraer el valor de dicha dirección de memoria, como todo está ocurriendo en paralelo, se da el caso de que las rutinas intentan cambiar el valor al mismo tiempo, sin esperar a que uno termine primero, por lo que el ultimo en hacer los cambios no se alcanza a informar del ultimo valor de la variable y sobreescribe los cambios de las otras rutinas.
Entonces, cada vez que se ejecute el programa se podrán ver resultados distintos que dependeran del azar con las que una rutina termine antes que otra.

## Respuesta P2 - Explicación estructura del código.

**Supuesto: Se deben de ingresar exactamente 16 digitos por consola al ejecutar el script ´subrutinas.go´**

Para ordenar el arreglo donde se almacenan los 16 inputs, se crean 4 subrutinas, donde cada una de estas se encarga de ordenar mediante ´mergeSort´ un subarreglo de largo 4. 

Luego, se espera a que se reciban mediante canales los subarreglos ordenados y se fusionan en dos arreglos que se pasan a dos subrutinas respectivamente para que se vuelvan a ordenar.

Ahora, con los dos arreglos ya ordenados se vuelven a fusionar y se mandan a una unica rutina que termina de ordenar el arreglo completo, haciendo ´print´ del resultado final.

Se puede notar que es una simulacion del algoritmo de ´merge sort´ pero haciendo uso de las ´go routines´.
