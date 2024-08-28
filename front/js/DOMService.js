import {fetchData} from "./remoteService.js"

const asc = "asc"
const desc = "desc"

let rootDirecory            //корневая директория
let currentDir = "/home"  //путь к текущей директории
let currentSort = asc         // флаг текущей сортировки

let showCurrentDirectory    // элемент отображающий путь к текущей директории

let placeholder             //переменная контейнера плейсхолдера
let containerContent        //переменная контейнера с данными


//setShowCurrentDirectory - функция для задания переменной showCurrentDirectory 
export function setShowCurrentDirectory(newShowCurrentDirectory){
    showCurrentDirectory = newShowCurrentDirectory
}

//setPlaceholder - функция для задания переменной placeholder 
export function setPlaceholder(newPlaceholder){
    placeholder = newPlaceholder
}

//setContainerContent - функция для задания переменной containerContent 
export function setContainerContent(newContainerContent){
    containerContent = newContainerContent
}



//changeSortFlag - изменение флага сортировки и обновление данных
export function changeSort(){
    if(currentSort == asc){
        currentSort = desc
        this.textContent = "Сортировка по размеру (убывание)"
    }
    else{
        currentSort = asc
        this.textContent = "Сортировка по размеру (возрастание)"
    }
    getAndUpdateData();
}



//returnToPreviousDirectory - возвращение на предыдущую директорию в пути обновление данных
export function returnToPreviousDirectory(){
    if(currentDir==rootDirecory){
        alert("вы достигли корневой директории")
        return
    }
    let arr = currentDir.split('/')
    arr.pop()
    currentDir = arr.join("/")
    getAndUpdateData();
}

 // directoryTraversal - переход на дирректорию вглубь и обновление данных
 function directoryTraversal(path){
    currentDir = `${currentDir}/${path}`
    getAndUpdateData();
}

//getDataAndParseData - получение данные и их представление
function getDataAndParseData(){
    return fetchData(`http://localhost:8080/fs?dst=${currentDir}&sort=${currentSort}`).then(responseData => {
        if(responseData.status==200){
            rootDirecory = responseData.root_dir
            return responseData.data;
        }else{
            alert(`Ошибка запроса: ${responseData.text_error}`)
        }
    })
}

//updateData - обнавление данных
function updateData(files){
     //обновление отображения текущего пути
     showCurrentDirectory.textContent = currentDir

     // поиск и очищение таблицы
     const tbody = document.getElementById('file-table').getElementsByTagName('tbody')[0];
     tbody.innerHTML = ''; 


     files.forEach(file => {
         tbody.appendChild(createRowFromFile(file))
        
     });
}


//getAndUpdateData - получение и обновление данных на странице
export function getAndUpdateData() {

    containerContent.style.display = "none";
    placeholder.style.display = "block"
    
    getDataAndParseData().then(
        files =>{
            updateData(files)
            containerContent.style.display = "block";
            placeholder.style.display = "none"
        } 
    )
}

//createRowFromFile - создание строки таблицы из данных
function createRowFromFile(file){
                //создание элементов таблицы
                const row = document.createElement('tr');
                const fileName = document.createElement('td');
                const fileSize = document.createElement('td');
                const fileType = document.createElement('td');

                //присовение элементам строки значений
                fileName.textContent = file.file_name;
                fileSize.textContent = `${file.file_size.toFixed(2)} ${file.file_size_type}`;
                fileType.textContent = file.file_type;

                //добалвение элементов в строку таблицы
                row.appendChild(fileName);
                row.appendChild(fileSize);
                row.appendChild(fileType);

                //если это директория - сделать ее кликабельной и добавить ей стиль
                if(file.file_type=="Директория"){
                    row.classList.add("directStyle")
                    row.addEventListener('click', function(){
                        directoryTraversal(file.file_name)
                    })
                }


                //добавлние классов
                fileName.classList.add("rowTableStyle")
                fileSize.classList.add("rowTableStyle")
                fileType.classList.add("rowTableStyle")
                row.classList.add("rowTableStyle")

                //возвращение строки в таблицу
                return  row
}