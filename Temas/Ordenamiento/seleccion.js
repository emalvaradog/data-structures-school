const array = [
  { contenedor: { dato: 5 } },
  { contenedor: { dato: 4 } },
  { contenedor: { dato: 0 } },
  { contenedor: { dato: 2 } },
  { contenedor: { dato: 10 } },
  { contenedor: { dato: 8 } },
];

function ordenamientoSeleccion(array) {
  for (let i = 0; i < array.length; i++) {
    for (let j = i + 1; j < array.length; j++) {
      if (array[j].contenedor.dato < array[i].contenedor.dato) {
        let a = array[i];
        array[i] = array[j];
        array[j] = a;
      }
    }
  }
  return array;
}

console.log(ordenamientoSeleccion(array));
