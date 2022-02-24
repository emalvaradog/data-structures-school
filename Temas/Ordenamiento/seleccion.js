const array = [{obj: {dato: 5}},{obj: {dato: 4}},{obj: {dato: 0}},{obj: {dato: 2}},{obj: {dato: 10}},{obj: {dato: 8}},]

function ordenamientoSeleccion(array, tam){
    for(let i = 0; i<tam; i++){
        for(let j = i+1; j<tam; j++){
            if(array[j].obj.dato < array[i].obj.dato){
                let a = array[i];
                array[i] = array[j];
                array[j] = a;
            }
        }
    }
    console.log(array)
}

ordenamientoSeleccion(array, array.length);
