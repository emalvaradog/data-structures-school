#include <stdio.h>
#include <stdlib.h>

typedef struct libro{
  char* titulo;
  char* tipoFuente;
  float alto;
  float ancho;
}Libro;

void ordSelec(int tam, long arr[]);
void ordIns(int tam, long arr[]);

void ordAlto(int tam, Libro bd[], int type);
void ordTitulo(int tam, Libro bd[], int type);
void ordAncho(int tam, Libro bd[], int type);
void ordFuente(int tam, Libro bd[], int type);
void imprimirTDA(int tam, Libro bd[]);

void inicio();