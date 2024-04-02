Esto es un compilador de go hecho para el proyecto de Compiladores

Integrantes:

***Anthony Jimenez**
***Kevin Varela**
***Fabian Rojas**



Pasos para iniciar


Servidor y compilador
  1)Configurar correctamente el parser y el scanner con las rutas apropiadas
  2)Ejecutar el main.go


Interfaz Web

Dentro de la carpeta donde se encuentra la interfaz web, escribir los siguientes comandos en la terminal
  1) npm create astro@latest
  2) npm run dev
  3) copiar la direccion que aparece en la consola (para más información visitar la documentación de Astro https://astro.build/


Funcionamiento

Una vez realizados los pasos anteriores, se debe mostrar en la terminal del servidor que se está ejecutando correctamente
lo mismo con la terminal de la interfaz web, que debe mostrar una dirrecion localhost donde se encuentra alojada la página

La pagina cuenta con 2 bloques de texto, en el de la izquierda se debe escribir el código, una vez terminado de escribir
se debe presionar en el botón "compile" y enviará un mensaje al servidor, el parser procesará el código y se retornará
una respuesta que se mostrará en el bloque de texto de la derecha, si se encuentran errores, los mostrará enumerados y si no,
mostrará que no hay errores

Se adjunta un enlace de YouTube donde se muestra el funcionamiento del proyecto
