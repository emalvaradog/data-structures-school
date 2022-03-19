#include <stdio.h>

typedef struct dato{
  int valor;
}Dato;

typedef struct Contendedor{
  Dato contenedor;
}Contenedor;

void ordenarBurbuja(Contenedor *a, int len){
  int i=0;
  int ordenado = 0;
  while ( !ordenado ){
    ordenado=1;
    for(i = 0; i<len-1; i++){
      if(a[i].contenedor.valor>a[i+1].contenedor.valor){
        Contenedor aux = a[i];
        a[i] = a[i+1];
        a[i+1] = aux;
        ordenado = 0;
      }
    }
  }
}

int main() {  
  int m,i;
  Contenedor a[cantidad];
  srand (time(NULL));

  for( m=0; m<cantidad; m++){
    a[m].contenedor.valor = rand()%(459564+1);
    scanf("%d", &arrCont[m].contenedor.valor);
  }

  arrCont[0].contenedor.valor = 5;
  arrCont[1].contenedor.valor = 4;        
  arrCont[2].contenedor.valor = 9;
  arrCont[3].contenedor.valor = 0;
  arrCont[4].contenedor.valor = 1;

  ordenarBurbuja(a, cantidad);
    
  for( i = 0; i < cantidad; i++){
    if(a[i].contenedor.valor%2==0){
    printf("%d,", a[i].contenedor.valor);
    }
  }

  return 0;
}