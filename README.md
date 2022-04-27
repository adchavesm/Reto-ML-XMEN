# Examen Mercadolibre

Magneto necesita buscar Mutantes, esta herramienta desarrollada en GO y desplegada con docker en AWS, te ayudará!  


## ¡Consulta si eres mutante!

El ejercicio quedó resuelto y desplegado a través de AWS. Por lo que si estás leyendo esto, puedes interactuar con este.
Este servicio se encuentra respondiendo en la direccion http://35.171.169.217, por el puerto: 8080.

### El servicio cuenta con dos recursos expuestos:

- Solo debes contar con un generador de solicitudes HTTP como postman, o tu navegador para las peticiones GET. O también con CURL.
- A continuación dejo los CURLs de cada recurso para que puedan ser importados.

- Si quieres consultar si el Adn de un Humano, pertenece a un mutante: 

```
curl --location --request POST 'http://35.171.169.217:8080/mutant' \
--header 'Content-Type: application/json' \
--data-raw '{
"dna": [
"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}'
```

- Si quieres consultar las estadísticas de humanos y mutantes verificados.

```
curl --location --request GET 'http://35.171.169.217:8080/stats'
```

## Instrucciones de ejecución

#### Si quieres ejecutar esto en tu máquina local, debes seguir los siguientes pasos:

- Tener instalado docker y docker compose en tu computadora.
- Clonar este repositorio en un directorio de tu computadora.
- Una vez clonado, entras al repositorio clonado en local y en tu terminal ejecutas: $ docker-compose up
- Debido a que tanto la aplicación backend y la base de datos están en contenedores, este comando le indicará a docker que corra los contenedores descritos en los archivos docker.
- Solo tendrás que esperar y ya tendrás el servicio disponible en el puerto 8080.

## Consideraciones

- El algoritmo para identificar mutantes, funciona para cualquier matriz de tamaño N.
- Se asume que el conjunto de Mutantes pertenece al conjunto de humanos. Por lo tanto todos los mutantes son humanos, pero no todos los humanos son mutantes.

## Tecnologias utilizadas

- Como lenguaje de programación se utilizó GO y su framework Gin para gestionar el servidor.
- Se utilizó una base de datos MySQL para la persistencia de información.
- Se creó un contenedor de docker tanto para la app, como para la base de datos.
- Se desplegó en una máquina EC2 de AWS, a través de docker-compose y el repositorio de github.
