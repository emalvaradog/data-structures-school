#include "practica1.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void crearTDA(Container* tda, int node_id, char* node_name, int node_key, 
  int block_id, char* content, int block_key){
  // Node asignation
  tda->node.node_id = node_id;
  tda->node.node_name = node_name;
  tda->node.node_key = node_key;

  // Block asignation
  tda->block.block_id = block_id;
  tda->block.content = content;
  tda->block.block_key = block_key;
};

void imprimirTDA(Container tda) {
  printf(" ---- NODE: ---- \n");
  printf("ID: %d\n", tda.node.node_id);
  printf("Name: %s\n", tda.node.node_name);
  printf("Key: %d\n", tda.node.node_key);
  printf("\n");
  printf("---- Block: ---- \n");
  printf("ID: %d\n", tda.block.block_id);
  printf("Content: %s\n", tda.block.content);
  printf("Key: %d\n", tda.block.block_key);
  printf("\n");
};

void copiarTDA(Container original, Container* copia){
  copia->block.block_id = original.block.block_id;
  copia->block.content = original.block.content;
  copia->block.block_key = original.block.block_key;

  copia->node.node_id =  original.node.node_id;
  copia->node.node_name = original.node.node_name;
  copia->node.node_key = original.node.node_key;
  printf("\n");
}


int start() {
  Container mainTDA;
  Container newTDA;
  printf("Crear tda \n");
  crearTDA(&mainTDA, 5, "Node", 10, 3, "Content", 2);
  printf("Imprimir TDA \n");
  imprimirTDA(mainTDA); 
  printf("Copiar tda \n");
  copiarTDA(mainTDA, &newTDA); 
  printf("Imprimir nuevo TDA \n");
  imprimirTDA(newTDA);
  return 0;
}