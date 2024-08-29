import { getAndUpdateData } from "./tableScript";

export let rootDirecory: string;  // корневая директория
export let currentDir: string;     // путь к текущей директории

export const asc: string = "asc";
const desc: string = "desc";
let currentSort: string = asc; // флаг текущей сортировки

// setRootDirecory - установка корневой директории и переотправка запроса с корневой
export function setRootDirecory(path: string): void {
    rootDirecory = path;
    getAndUpdateData(rootDirecory, currentSort);
}

// setCurrentDir - установка текущей директории, в момент когда определяется корневая
export function setCurrentDir(): void {
    currentDir = rootDirecory;
}

// changeSort - изменение флага сортировки и обновление данных
export function changeSort(this: HTMLButtonElement): void {
    if (currentSort === asc) {
        currentSort = desc;
        this.textContent = "Сортировка по размеру (убывание)";
    } else {
        currentSort = asc;
        this.textContent = "Сортировка по размеру (возрастание)";
    }
    getAndUpdateData(currentDir, currentSort);
}

// goIntoDirectory - переход на директорию вглубь и обновление данных
export function goIntoDirectory(path: string): void {
    if (currentDir === "/") {
        currentDir = "";
    }
    currentDir = `${currentDir}/${path}`;
    getAndUpdateData(currentDir, currentSort);
}

// returnToPreviousDirectory - возвращение на предыдущую директорию в пути и обновление данных
export function returnToPreviousDirectory(): void {
    if (currentDir === rootDirecory) {
        alert("Вы достигли корневой директории");
        return;
    }
    let arr = currentDir.split('/');
    arr.pop();
    currentDir = arr.join("/");
    if (currentDir === "") {
        currentDir = "/";
    }
    getAndUpdateData(currentDir, currentSort);
}
