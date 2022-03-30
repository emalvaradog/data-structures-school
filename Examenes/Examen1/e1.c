#include "e1.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>


void crearComputadora(Computadora** arr, int pos){
  fflush(stdin);
  scanf("%s", arr[pos-1]->marca);
}


void iniciar() {
  int tam = 1;
  Computadora** arr = (Computadora**)malloc(sizeof(Computadora*) * tam);

  if(arr == NULL){
    printf("ERROR");
  }

  for(int i = 0; i<tam; i++){
    arr[i] = (Computadora*)malloc(sizeof(Computadora)*1);
    if(arr[i] == NULL){
      printf("Error");
    }
  }
  
  crearComputadora(arr, tam);

  // arr[0]->marca = "A";
  // arr[0]->procesador.modelo = "B";
  printf("%s \n", arr[0]->marca);
  // printf("%s \n", arr[0]->procesador.modelo);

};
