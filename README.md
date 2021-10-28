# Concurrencia y Paralelismo
[Programación concurrente](https://ferestrepoca.github.io/paradigmas-de-programacion/progconcurrente/concurrente_teoria/index.html)

## Concurrencia
Hay Concurrencia entre varios procesos cuando existen al mismo tiempo. Un ejemplo de esto es que el programa manda a ejecutar 4 tareas a un procesador y éste le asigna un pequeño tiempo a cada tarea 
si no se completa la tarea 1 en la fracción de tiempo que se le asigno sigue con la segunda tarea 
y asi sucesivamente hasta llegar con la ultima tarea, al terminar el procesador verifica si alguna de las tareas anteriores no se completo vuelve otra vez y le asigna un pequeño tiempo para hacerlo y
vuelve el ciclo
### Ejemplo
###### Primera Vuela

```Nucleo 1``` -> ``` Tarea 1 (1ms) -> Tarea 2 (1ms)(Se Completo) -> Tera 3 (1ms) (Se Completo) -> Tarea 4 (1ms) ```

###### Segunda Vuela

```Nucleo 1``` ->``` Tarea 1 (1ms) (Se Completo) -> Tarea 4 (1ms) ```

###### Tercera Vuela

```Nucleo 1``` -> ``` Tarea 4 (1ms) (Se Completo) ```

Esto sucede tan rápido que se piensa que se esta ejecutando al mismo tiempo.

## Paralelismo
Ejecuta dos tareas al mismo tiempo, para que se de tiene que haber varios núclos.
En este caso cada tarea se va hacia un nucleo del procedador y se ejecutan al mismo tiempo.
(procesador con dos nucleos)
### Ejemplo
###### Primera Vuela
```Nucleo 1``` -> ``` Tarea 1 (1ms) -> Tarea 3 (1ms) ```

```Nucleo 2``` -> ``` Tarea 2 (1ms) -> Tarea 4 (1ms) ```


## Diferencia de Go
La diferencia entre Go con otros lenguajes concurrente es que por ejemplo
Java tiene su tarea y crea un hilo y se lo manda al procesador ya el programam de Java 
no hace nada hasta que el procesador le devuelva el ok de que termino 

EL que maneja la concurrencia en si es el procesador el SO, no Java con sus hilos
Otro Ejemplo es que si hay un programam en Java y el le manda 6 hilos a un procesador de 2 nucleos
no quiere decir que en cada nucleo se le vaya asignar un hilo ya que el procesador puede estar 
manejando otros programas, servicios, etc. EL SO manejara de que forma ejecuta esos hilos puede 
que solo mande un hilo al nucleo 1 y mantega en espera los otros hilos en el mismo nucleo o los 
ira repartiendo, el punto es que el SO es el que se encargara de ejecutar los hilos ya sea en 
paralelo usando los dos nucleos o poniendolos en concurrencia en un solo nucleo, o la combinacion
del paralelismo con concurrencia para hacerlo aun mas rapido.

Ahora en Java cada hilo es una funcion si tu quieres ejecutar 10 funciones en paralelo debes mandar 
10 hilos. Con Go, las gorutine tu puedes tener cientos de funciones en un hilo y cuando ese hilo
llega al procesador (nucleo) (ten en cuenta que el SO es el que maneja cuando va a ejecutar ese hilo
ya que pueden haber procesos delante) y le asigna 1ms (Ejemplo refencial) go tiene el control de ese 1ms
o sea Go hace lo que quiera en ese 1ms el toma el control y ejecuta todas las funciones que puede
con la gorutine.

Por eso Go es un lenguaje poderoso, esta es la pontencia de Go ya que puede manejar de forma eficiente el uso de los multiples nucleos.

Por ejemplo un programa en java necesita 8 nucleos y 12 GB de ram un programa en Go solo puede necesitar
2 nucleos y 4 GB de ram para el mismo programa