#include <stdio.h>

typedef struct dato{
    int valor;
}Dato;

typedef struct contenedor{
    Dato contenedor;
}Contenedor;

void ordenamientoSeleccion(Contenedor array[], int tam){
    for(int i = 0; i < tam; i++){
        for(int j = i+1; j < tam; j++){
            if(array[j].contenedor.valor < array[i].contenedor.valor){
                Contenedor cont = array[i];
                array[i] = array[j];
                array[j] = cont;
            }
        }
    }
}

int main() {
    Contenedor arrCont[7];
    arrCont[0].contenedor.valor = 5;
    arrCont[1].contenedor.valor = 3;
    arrCont[2].contenedor.valor = 9;
    arrCont[3].contenedor.valor = 0;
    arrCont[4].contenedor.valor = 6;
    arrCont[5].contenedor.valor = 8;
    arrCont[6].contenedor.valor = 2;

    ordenamientoSeleccion(arrCont, 7);
    
    for(int i = 0; i < 7; i++){
        printf(%d, arrCont[i].contenedor.valor);
    }
        
    return 0;
}
