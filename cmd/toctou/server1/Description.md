# Description


Time Of Check: El servidor comprueba si el archivo existe con os.Stat().
Time Of Use: Posteriormente, el servidor intenta leer y eliminar el archivo.
Entre estas dos operaciones, un atacante podría aprovechar la ventana de tiempo para modificar o reemplazar el archivo, explotando así la vulnerabilidad TOCTOU. Por ejemplo, después de la verificación y antes de que el archivo sea leído o eliminado, el archivo podría ser reemplazado, lo que permitiría la ejecución de código no deseado o la eliminación de un archivo incorrecto.

Mitigación
Una forma de mitigar esta vulnerabilidad es utilizar bloqueos adecuados para asegurar que la verificación y el uso ocurran de manera atómica, o evitar este tipo de comprobaciones separadas y realizar la operación de forma que sea segura contra race conditions.