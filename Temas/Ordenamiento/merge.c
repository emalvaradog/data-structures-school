typedef struct {
  int valor;
}Generico;

typedef struct Dato{
  Generico datoG;
}Dato;

void sort(int arreglo1[], int n1, int arreglo2[], int n2, int arreglo3[]){
  int a=0, b=0, c=0;
 
  while (a<n1 && b<n2) {
    if (arreglo1[a]<arreglo2[b]) {
      arreglo3[c] = arreglo1[a];
      a++;
    } else {
      arreglo3[c] = arreglo2[b];
      b++;
    }
    c++;
  }
  while (a<n1) {
    arreglo3[c] = arreglo1[a];
    a++;
    c++;
  }

  while (b<n2) {
    arreglo3[c] = arreglo2[b];
    b++;
    c++;
  }
}
 
void merge(int arreglo[], int n){
  int *array1, *array2, n1, n2,m,k;
  if (n>1){
    if (n%2 == 0){
      n1=n2=(int) n / 2;
    } else {
      n1=(int) n / 2;
      n2=n1+1;
    }

    array1=(int *) malloc(sizeof(int)*n1);
    array2=(int *) malloc(sizeof(int)*n2);
    if(array1 && array2==NULL){
      perror("Se acabo la memoria");
      exit(-1);
    }

    for(m=0;m<n1;m++){
      array1[m]=arreglo[m];
    }

    for(k=0;k<n2;m++,k++){
      array2[k]=arreglo[m];
    }
    merge(array1, n1);
    merge(array2, n2);
    sort(array1, n1, array2, n2, arreglo);
    free(array1);
    free(array2);
  }    
}