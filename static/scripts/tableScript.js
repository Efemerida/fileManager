import {fetchData} from "./remoteService.js"
import {goIntoDirectory, currentDir, setCurrentDir} from "./buttonsScript.js"



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



//updateData - обнавление данных
function updateData(files){

     // поиск и очищение таблицы
     const tbody = document.getElementById('file-table').getElementsByTagName('tbody')[0];
     tbody.innerHTML = ''; 


     files.forEach(file => {
         tbody.appendChild(createRowFromFile(file))
        
     });
}


//removePlaceholder - скрытие плейсхолдера
export function removePlaceholder(){
    containerContent.style.display = "block";
    placeholder.style.display = "none"
}

//getAndUpdateData - получение и обновление данных на странице
export function getAndUpdateData(path, sort) {

    containerContent.style.display = "none";
    placeholder.style.display = "block"

    //обновление отображения текущего пути
    showCurrentDirectory.textContent = path

    fetchData(path, sort).then(
        files =>{

            //если корневая директория еще не была получена
            if(currentDir===undefined){
                setCurrentDir()
                return
            }

            updateData(files)
            removePlaceholder()
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
                if(file.file_type==="Директория"){
                    row.classList.add("direct_style")
                    row.addEventListener('click', function(){
                        goIntoDirectory(file.file_name)
                    })
                }


                //добавлние классов
                fileName.classList.add("row_table_style")
                fileSize.classList.add("row_table_style")
                fileType.classList.add("row_table_style")
                row.classList.add("row_table_style")

                //возвращение строки в таблицу
                return  row
}