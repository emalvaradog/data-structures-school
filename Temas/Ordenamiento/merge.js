const array = [
  { contenedor: { dato: 5 } },
  { contenedor: { dato: 4 } },
  { contenedor: { dato: 0 } },
  { contenedor: { dato: 2 } },
  { contenedor: { dato: 10 } },
  { contenedor: { dato: 8 } },
];

function sort(left, right) {
  let arr = [];
  while (left.length && right.length) {
    if (left[0].contenedor.dato < right[0].contenedor.dato) {
      arr.push(left.shift());
    } else {
      arr.push(right.shift());
    }
  }

  return [...arr, ...left, ...right];
}

function mergeSort(array) {
  const half = array.length / 2;

  if (array.length < 2) {
    return array;
  }

  const left = array.splice(0, half);

  return sort(mergeSort(left), mergeSort(array));
}

console.log(mergeSort(array));
