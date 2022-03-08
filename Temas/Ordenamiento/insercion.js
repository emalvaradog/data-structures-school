const generico = { dato };
const dato = { generico };

function ordenar(arreglo) {
  let pivote;
  let j;
  for (let i = 0; i < array.length; i++) {
    pivote = arreglo[i];
    j = i - 1;
    while (j >= 0 && arreglo[j].generico.dato > pivote.generico.dato) {
      arreglo[j + 1] = arreglo[j];
      j--;
    }
    arreglo[j + 1] = pivote;
  }
}
