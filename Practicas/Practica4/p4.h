#include <stdio.h>

typedef struct{
    char* pais;
    char* nombre;
    float altura;
    int telefono;
}Dato;

typedef struct{
    Dato datos;
}Persona;

void imprimirArreglo(Persona arr[], int tam);
void ordenarAltura(Persona arr[], int tam);
void ordenarNombre(Persona arr[], int tam);
void ordenarTelefono(Persona arr[], int tam);
void ordenarPais(Persona arr[], int tam);
void principal();
