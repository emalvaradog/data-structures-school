#include "p2.h"
#include <stdio.h>
#include <stdlib.h>

Libro bd[] = {
  [0] = {"Anecdota de un dia", "Tiahoma Black", 8.5, 3.18},
  [1] = {"Burros Blancos", "Agency", 7.5, 0.74},
  [2] = {"Cien años y uno mas", "Abscissa", 6.5, 0.3},
  [3] = {"Diamantes y dados", "HP Simplified", 5.5, 0.14},
  [4] = {"EL Avestrus", "Arial", 4.5, 0.58},
  [5] = {"El Cuento", "Calibri", 3.5, 0.46},
  [6] = {"Elefantes", "Impact", 2.5, 1.46},
  [7] = {"Fiestas", "Alanis  Hand", 1.5, 1.9},
  [8] = {"Gatoman", "Jasmine UPC", 0.5, 2.34},
  [9] = {"Hipopotamos en su habitat", "Latha", 0.5, 2.78},
  [10] = {"Italia, un nuevo mundo", "Kaiti", 1.5, 3.22},
  [11] = {"Juanacatlan", "Nyala", 2.5, 3.66},
  [12] = {"Kids, where are them?", "Verdana", 3.5, 4.1},
  [13] = {"La niña", "Tahoma", 4.5, 4.54},
  [14] = {"Metafisica", "Modern", 5.5, 4.98},
  [15] = {"Nuevos horizontes", "Univers", 6.5, 5.42},
  [16] = {"Osos Polares", "Yu Gotic", 7.5, 5.86},
  [17] = {"Pato Hood", "Ravie", 8.5, 6.3},
  [18] = {"Poblaciones vulnerables", "OCRB", 9.5, 6.75},
  [19] = {"Quo duo", "Selawik", 10.5, 7.18},
  [20] = {"Robin Hood", "Perpetua", 11.5, 7.62},
  [21] = {"Significado de la vida", "Times", 12.5, 8.06},
  [22] = {"Soberanos", "Webdings", 13.5, 8.5},
  [23] = {"Trasatlantico", "Sweer Pea", 14.5, 8.94},
  [24] = {"Transportes en el futuro", "Bell MT", 15.5, 9.38},
  [25] = {"Trotamundos", "Ebrima", 16.5, 9.82},
  [26] = {"Uvas para el vino", "Mangal", 17.5, 10.26},
  [27] = {"Verduras y frutas", "Cambria", 18.5, 10.7},
  [28] = {"Viajando ando", "Dotum", 19.5, 11.14},
  [29] = {"Walter Cross, ¿Quien fue?", "Papyrus", 20.5, 11.58},
  [30] = {"Xcaret, un lugar magico", "Arial Black", 21.5, 12.02},
  [31] = {"Yardas", "Georgia", 22.5, 12.46},
  [32] = {"Zapato Viejo", "French Script",23.5, 12.9},
};
  

void ordSelec(int tam, long arr[]){
  for(int i = 0; i < tam; i++){
    for(int j = i+1; j < tam; j++){
      if(arr[j]%2==0 && arr[i]%2 == 0){
        if(arr[j] < arr[i]){
          int cont = arr[i];
          arr[i] = arr[j];
          arr[j] = cont;
        }
      }
    }
  }
}

void ordIns(int tam, long arr[]){
  int pivote;
  int j;
  for(int i = 0; i<tam; i++){
    if(arr[i]%2 !=0){
      pivote = arr[i];
      j = i - 1;
      while(j<= 0 && arr[j] < pivote) {
        arr[j+1] = arr[j];
        j--;
      }
      arr[j + 1] = pivote;
    }
  }
}


void ordAlto(int tam, Libro bd[], int type){
  if(type == 1){
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].alto < bd[i].alto){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  } else {
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].alto > bd[i].alto){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  }  
}

void ordTitulo(int tam, Libro bd[], int type){
  if(type == 1){
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].titulo[0]< bd[i].titulo[0]){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  } else {
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].titulo[0]> bd[i].titulo[0]){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  }
}

void ordAncho(int tam, Libro bd[], int type){
  if(type == 1){
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].ancho< bd[i].ancho){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  } else {
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].ancho > bd[i].ancho){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  }
}

void ordFuente(int tam, Libro bd[], int type){
  if(type == 1){
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].tipoFuente[0]< bd[i].tipoFuente[0]){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  } else {
    for(int i = 0; i < tam; i++){
      for(int j = i+1; j < tam; j++){
        if(bd[j].tipoFuente[0]> bd[i].tipoFuente[0]){
          Libro lib = bd[i];
          bd[i] = bd[j];
          bd[j] = lib;
        }
      }
    }
  }
  
}

void imprimirTDA(int tam, Libro bd[]){
  for(int i=0; i<tam; i++){
    printf("%s | %s | %f | %f \n",bd[i].titulo, bd[i].tipoFuente, bd[i].alto, bd[i].ancho);
  }
}

void inicio(){
  long arr[10000];

  if(arr == NULL){
    printf("No fue posible alocar memoria");
    exit(0);
  }

  for(int i = 0; i < 10000; i++){
    arr[i] = (long) rand();
  }

  ordSelec(10000, arr);
  ordIns(10000, arr);
  for(int i = 0; i < 10000; i++){
    printf("%ld \n", arr[i]);
  }

  int opc = 0;
  while(opc != 5){
    printf("Ordenar libros por: \n");
    printf("1- Alto \n");
    printf("2- Titulo \n");
    printf("3- Ancho \n");
    printf("4- Fuente \n");
    printf("5- Salir \n");
    scanf("%d", &opc);
    int type = 0;
    switch(opc) {
      case 1: 
        printf("Ordenar libros por alto \n");
        printf("1. Ascendente \n");
        printf("2. Descendente \n");
        scanf("%d", &type);
        ordAlto(32, bd, type);
        imprimirTDA(32, bd);
        break;
      
      case 2: 
        printf("Ordenar libros por titulo \n");
        printf("1. Ascendente \n");
        printf("2. Descendente \n");
        scanf("%d", &type);
        ordTitulo(32, bd, type);
        imprimirTDA(32, bd);
        break;
      
      case 3: 
        printf("Ordenar libros por ancho \n");
        printf("1. Ascendente \n");
        printf("2. Descendente \n");
        scanf("%d", &type);
        ordAncho(32, bd, type);
        imprimirTDA(32, bd);
        break;
      
      case 4: 
        printf("Ordenar libros por fuente \n");
        printf("1. Ascendente \n");
        printf("2. Descendente \n");
        scanf("%d", &type);
        ordFuente(32, bd, type);
        imprimirTDA(32, bd);
        break;

      default:
       break; 
    }
  }
}
