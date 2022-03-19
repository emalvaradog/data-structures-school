const Arreglo = [
  { contenedor: { dato: 5 } },
  { contenedor: { dato: 7 } },
  { contenedor: { dato: 10 } },
  { contenedor: { dato: 2 } },
  { contenedor: { dato: 3 } },
  { contenedor: { dato: 1 } },
];

function insercion(arreglo) {
  let pivote;
  let j;
  for (let i = 0; i < arreglo.length; i++) {
    pivote = arreglo[i];
    j = i - 1;
    while (j >= 0 && arreglo[j].contenedor.dato > pivote.contenedor.dato) {
      arreglo[j + 1] = arreglo[j];
      j--;
    }
    arreglo[j + 1] = pivote;
  }
  return arreglo;
}

console.log(insercion(Arreglo));
