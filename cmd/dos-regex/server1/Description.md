# Description

## API Vulnerable a DoS con Regex

Esta API simula una vulnerabilidad de Denegación de Servicio (DoS) al procesar correos electrónicos con una expresión regular ineficiente. Un atacante podría enviar un correo electrónico extremadamente largo y malformado, haciendo que la expresión regular consuma muchos recursos y, potencialmente, haga que el servidor no responda.

### Ejemplo de uso


    curl -X POST -H "Content-Type: application/json" -d '{"email": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa@malicious-domain-with-very-very-very-long-name.com"}' http://localhost:8080/update_email


### Cómo Evaluar el Comportamiento Para evaluar si realmente se está produciendo un DoS, podrías intentar lo siguiente:

Envía Múltiples Solicitudes Concurrentes: Intenta hacer múltiples solicitudes de manera concurrente para ver si el servidor comienza a ralentizarse o deja de responder. Puedes hacer esto usando una herramienta como ab (Apache Benchmark) o wrk.
    
    apt install apache2-utils

    ab -n 1000 -c 10 -p post_data.json -T application/json http://localhost:8080/update_email

    c: conexiones concurrentes; n: cantidad de solicitudes

    ps -p 1442998 -o %cpu,%mem,cmd

    top -p 1442998



No se cae depende del host.


Se levanta un contenedor con la API limitada.


    docker run -it -d --name dos-regex-api -h dos-regex diegoall1990/dos-regex-api:0.0.1

    docker run -it -d --name dos-regex-api -h dos-regex --cpus=".5" --memory="100m" diegoall1990/dos-regex-api:0.0.1

    docker stats dos-regex-api

    docker run -it -d --name dos-regex-api -h dos-regex --cpus=".2" --memory="50m" diegoall1990/dos-regex-api:0.0.1

    docker exec -it dos-regex-api /bin/sh

    docker exec -it dos-regex-api netstat -tuln
    docker exec -it dos-regex-api lsof -i:8080


    docker run -it -d --name dos-regex-api -h dos-regex --cpus=".1" --memory="8m" -p 8080:8080 diegoall1990/dos-regex-api:0.0.1


ab -n 1000000 -c 1000 -p post_data.json -T application/json http://localhost:8080/update_email

Benchmarking localhost (be patient)
apr_socket_recv: Connection refused (111)
Total of 4454 requests completed