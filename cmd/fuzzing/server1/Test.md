# Test


Curl con datos binarios:

Generar datos aleatorios: Utilizaremos el comando head de Linux para generar un archivo binario de un tamaño específico.
Enviar los datos: Usaremos curl con la opción -d para enviar el archivo como datos POST.

# Generar un archivo de 1000 bytes de datos aleatorios
head -c 1000 /dev/urandom > large_file.bin

# Enviar el archivo como una solicitud POST
curl -k -d @large_file.bin localhost:8080/handler
