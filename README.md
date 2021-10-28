# Modificando módulos con Go
Se descargo el codigo del framework [echo](https://github.com/labstack/echo) para modificarlo

## Comandos

### Editar el archivo mod de un remoto a local, reemplazamos la ruta
```go mod edit -replace github.com/labstack/echo/v4=./echo```

### Verificar que todos los modulos esten bien
``` go mod verify  ```

### Volver al estado inicial antes del replace
``` go mod edit -dropreplace github.com/labstack/echo/v4 ```

### Crear el Vendor
Con este comando lo que hacemos es crear el vendor, que vendria siendo el
node_modules en javascript para tener el codigo de tercers y siga funcionando 
nuestro codigo

``` go mod vendor  ```

### go.sum
Es un registro de los paquetes que se usan en el proyecto, es como un package-lock.json
verifica la version del paquete y se mantiene encriptado, esto es para verificar que el 
codigo este con la mejor salud posible

### Limpiar los paquetes que ya no se usen en el proyecto
``` go mod tid  ```