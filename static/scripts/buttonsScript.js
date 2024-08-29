import{getAndUpdateData} from "./tableScript.js"

export let rootDirecory             //корневая директория
export let currentDir               //путь к текущей директории

export const asc = "asc"
const desc = "desc"
let currentSort = asc               // флаг текущей сортировки


//setRootDirecory - установка корневой директори и переотправка запроса с корневой
export function setRootDirecory(path){
    rootDirecory = path
    getAndUpdateData(rootDirecory, currentSort)
}

//setCurrentDir - установка текущей директории, в момент когда определяется корневая
export function setCurrentDir(){
    currentDir = rootDirecory
}

//changeSort - изменение флага сортировки и обновление данных
export function changeSort(){
    if(currentSort === asc){
        currentSort = desc
        this.textContent = "Сортировка по размеру (убывание)"
    }
    else{
        currentSort = asc
        this.textContent = "Сортировка по размеру (возрастание)"
    }
    getAndUpdateData(currentDir, currentSort);
}


 // goIntoDirectory - переход на дирректорию вглубь и обновление данных
export function goIntoDirectory(path){
    if(currentDir==="/"){
        currentDir = ""
    }
    currentDir = `${currentDir}/${path}`
    getAndUpdateData(currentDir, currentSort);
}


//returnToPreviousDirectory - возвращение на предыдущую директорию в пути и обновление данных
export function returnToPreviousDirectory(){
    if(currentDir === rootDirecory){
        alert("вы достигли корневой директории")
        return
    }
    let arr = currentDir.split('/')
    arr.pop()
    currentDir = arr.join("/")
    if(currentDir === ""){
        currentDir = "/"
    }
    getAndUpdateData(currentDir, currentSort);
}