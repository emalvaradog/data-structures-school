const Arreglo = [
  { contenedor: { dato: 5 } },
  { contenedor: { dato: 7 } },
  { contenedor: { dato: 10 } },
  { contenedor: { dato: 2 } },
  { contenedor: { dato: 3 } },
  { contenedor: { dato: 1 } },
];

function burbuja(array) {
  let ordenado = false;
  while (!ordenado) {
    ordenado = true;
    for (let i = 0; i < array.length - 1; i++) {
      if (array[i].contenedor.dato > array[i + 1].contenedor.dato) {
        let aux = array[i];
        array[i] = array[i + 1];
        array[i + 1] = aux;
        ordenado = false;
      }
    }
  }
  return array;
}

console.log(burbuja(Arreglo));
