#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct {
  char* modelo;
  char* marca;
  char* velocidad;
}Procesador;

typedef struct {
  char* cantidadRAM;
  char* cantidadHDD;
}Extras;

typedef struct {
  Procesador procesador;
  Extras extras;
  char* marca;
}Computadora;

void crearComputadora(Computadora** arr, int pos);

void iniciar();

