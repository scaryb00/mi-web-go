package main

import ( //librerias
    "fmt" //imprimir texto en consola
    "net/http" //herramientas para crear el servidor web
)

func main() {
    fs := http.FileServer(http.Dir("static")) //FileServer convierte static en un servidor de archivos
    http.Handle("/static/", http.StripPrefix("/static/", fs)) 
	
	// Página principal: devuelve el HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
	})

	fmt.Println("Servidor funcionando en http://localhost:8080")
    error := http.ListenAndServe(":8080", nil) //instrucción que enciende el servidor utilizando el puerto 8080, guardamos en una variable error

	if error != nil {  //si "error" no es nil, entonces ha habido un error y se imprime por pantalla
		fmt.Printf("Error al iniciar el servidor: %v\n", error)
	}
	
	
}