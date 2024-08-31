# Description


Este código implementa una API en Golang con un endpoint /vulnerable que recibe un parámetro input a través de la URL. La vulnerabilidad se introduce al copiar sin restricciones el contenido de input a un buffer de tamaño fijo (buffer := make([]byte, 1024)), lo cual podría causar un buffer overflow si input es más grande que el buffer.


## Cómo Exploitarla con Fuzzing


    sudo apt-get update
    sudo apt-get install afl++

    CC=afl-clang-fast go build -o vulnerable_api

    echo "test" > input.txt
    mkdir inputs
    mv input.txt inputs/

    afl-fuzz -i inputs -o findings -- ./vulnerable_api


    root@pho3nix:/home/diegoall/FALCO/falco-rules-examples/cmd/fuzzing/server2# CC=afl-clang-fast go build -o vulnerable_api
    # falco-rules-examples/cmd/fuzzing/server2
    net(.text._cgo_97ab22c4dc7b_C2func_getaddrinfo): relocation target __afl_area_ptr not defined
    net(.text._cgo_97ab22c4dc7b_Cfunc_freeaddrinfo): relocation target __afl_area_ptr not defined
    net(.text._cgo_97ab22c4dc7b_Cfunc_gai_strerror): relocation target __afl_area_ptr not defined
    runtime/cgo(.text._cgo_try_pthread_create): relocation target __afl_area_ptr not defined
    runtime/cgo(.text._cgo_set_stacklo): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_bindm): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_notify_runtime_init_done): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_init): relocation target __afl_area_ptr not defined
    runtime/cgo(.text._cgo_sys_thread_start): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.threadentry): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_mmap): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_munmap): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_setenv): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_unsetenv): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_sigaction): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_getstackbound): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_callers): relocation target __afl_area_ptr not defined
    runtime/cgo(.text.x_cgo_thread_start): relocation target __afl_area_ptr not defined

AFL++ no es compatible directamente con el compilador de Go. AFL++ está diseñado principalmente para trabajar con programas escritos en lenguajes como C, C++, y otros lenguajes que se compilan a código nativo. El proceso de compilación de Go es diferente porque utiliza un runtime y un sistema de compilación que no son compatibles con los mecanismos de instrumentación de AFL++.

¿Qué pasó?

Cuando intentaste compilar tu aplicación Go con afl-clang-fast (el compilador de AFL++), el compilador de Go intentó enlazar con su runtime y las bibliotecas estándar, pero los símbolos específicos requeridos por AFL++ (__afl_area_ptr, por ejemplo) no están definidos en las bibliotecas de Go. Esto se debe a que AFL++ espera trabajar con código en lenguajes más cercanos al sistema operativo que permiten instrumentación a nivel de código máquina, mientras que Go maneja muchas de estas cosas internamente de manera diferente.