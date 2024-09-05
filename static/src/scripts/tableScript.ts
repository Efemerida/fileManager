import { fetchData } from "./remoteService";
import { goIntoDirectory} from "./buttonsScript";

// Определение типов для файлов
export interface FileData {
    file_name: string;              //название файла
    file_size: number;              //размер файла
    file_size_type: string;         //тип размера
    file_type: string;              //тип файла
}

let showCurrentDirectory: HTMLElement; // элемент отображающий путь к текущей директории
let placeholder: HTMLElement; // переменная контейнера плейсхолдера
let containerContent: HTMLElement; // переменная контейнера с данными

// setShowCurrentDirectory - функция для задания переменной showCurrentDirectory 
export function setShowCurrentDirectory(newShowCurrentDirectory: HTMLElement): void {
    showCurrentDirectory = newShowCurrentDirectory;
}

// setPlaceholder - функция для задания переменной placeholder 
export function setPlaceholder(newPlaceholder: HTMLElement): void {
    placeholder = newPlaceholder;
}

// setContainerContent - функция для задания переменной containerContent 
export function setContainerContent(newContainerContent: HTMLElement): void {
    containerContent = newContainerContent;
}

// updateData - обновление данных
function updateData(files: FileData[]): void {
    // поиск и очищение таблицы
    const tbody = document.getElementById('file-table')!.getElementsByTagName('tbody')[0];
    tbody.innerHTML = '';

    //добавление данных
    files.forEach(file => {
        tbody.appendChild(createRowFromFile(file));
    });
}

// removePlaceholder - скрытие плейсхолдера
export function removePlaceholder(): void {
    containerContent.style.display = "block";
    placeholder.style.display = "none";
}

// getAndUpdateData - получение и обновление данных на странице
export function getAndUpdateData(path: string, sort: string): void {
    
    //отображение плейсхолдера
    containerContent.style.display = "none";
    placeholder.style.display = "block";

    // обновление отображения текущего пути
    showCurrentDirectory.textContent = path;

    //загрузка и обновление данных
    fetchData(path, sort).then(files => {
        updateData(files);
        removePlaceholder();
    });
}

// createRowFromFile - создание строки таблицы из данных
function createRowFromFile(file: FileData): HTMLTableRowElement {

    // создание элементов таблицы
    const row = document.createElement('tr');
    const fileName = document.createElement('td');
    const fileSize = document.createElement('td');
    const fileType = document.createElement('td');

    // присвоение элементам строки значений
    fileName.textContent = file.file_name;
    fileSize.textContent = `${file.file_size.toFixed(2)} ${file.file_size_type}`;
    fileType.textContent = file.file_type;

    // добавление элементов в строку таблицы
    row.appendChild(fileName);
    row.appendChild(fileSize);
    row.appendChild(fileType);

    // если это директория - сделать ее кликабельной и добавить ей стиль
    if (file.file_type === "Директория") {
        row.classList.add("direct_style");
        row.addEventListener('click', function() {
            goIntoDirectory(file.file_name);
        });
    }

    // добавление классов
    fileName.classList.add("row_table_style");
    fileSize.classList.add("row_table_style");
    fileType.classList.add("row_table_style");
    row.classList.add("row_table_style");

    // возвращение строки в таблицу
    return row;
}