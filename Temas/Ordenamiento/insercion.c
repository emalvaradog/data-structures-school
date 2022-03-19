#include <stdio.h>

typedef struct {
  int dato;
}Dato;

typedef struct {
  Dato generico;    
}TDA;

void insercion(TDA arreglo[], int tam){
  TDA pivote;
  int j;
  for(int i = 0; i < tam; i++){
    pivote = arreglo[i];
    j = i-1;
    while(j >=0 && arreglo[j].generico.dato > pivote.generico.dato){
      arreglo[j+1] = arreglo[j];
      j--;
    }
    arreglo[j+1] = pivote;
  }
}

int main() {
  TDA arreglo[5] = {
    {5},
    {8},
    {1},
    {10},
    {3}
  };

  insercion(arreglo, 5);

  for(int i = 0; i<5; i++){
    printf("%d, ", arreglo[i].generico.dato);
  }
}