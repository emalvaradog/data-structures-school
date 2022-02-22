#include <stdio.h>
#include <string.h>

typedef struct{
  int node_id;
  char* node_name;
  int node_key;
}Node;

typedef struct{
  int block_id;
  char* content;
  int block_key;
}Block;

typedef struct {
  Node node;
  Block block;
}Container;

void crearTDA(Container* tda, int node_id, char* node_name, int node_key, 
  int block_id, char* content, int block_key);
void imprimirTDA(Container tda);
void copiarTDA(Container original, Container* copia);
int start();
