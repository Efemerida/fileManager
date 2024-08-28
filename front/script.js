    currentDir = "/home/artem"  //путь к текущей директории
    currentSort = "asc"         // флаг текущей сортировки

    let showCurrentDirectory    // элемент отображающий путь к текущей директории

    // загрузка DOM дерева
    document.addEventListener('DOMContentLoaded', function() {
        fetchAndUpdateData();
        initScript()
    });


// initScript - опрделение обработки нажатий на кнопки и отображение текущей директории
    function initScript(){
        //кнопка сортировки
        sortButton = document.getElementById("sortButton")
        sortButton.addEventListener('click', changeSort)

        //кнопка назад
        let backButton = document.getElementById("backButton")
        backButton.addEventListener('click', returnToPreviousDirectory)

        //получение и сохранения элемента отображения текущей директории
        showCurrentDirectory = document.getElementById("currentDirectory")
    }


    //changeSortFlag - изменение флага сортировки и обновление данных
    function changeSort(){

        if(currentSort == 'asc'){
            currentSort = 'desc'
            this.textContent = "Сортировка по размеру (убывание)"
        }
        else{
            currentSort = 'asc'
            this.textContent = "Сортировка по размеру (возрастание)"
        }
        fetchAndUpdateData()
    }


    //returnToPreviousDirectory - возвращение на предыдущую директорию в пути обновление данных
    function returnToPreviousDirectory(){
        if(currentDir=='/home'){
            alert("вы достигли корневой директории")
            return
        }
        arr = currentDir.split('/')
        arr.pop()
        currentDir = arr.join("/")
        fetchAndUpdateData()
    }

     // directoryTraversal - переход на дирректорию вглубь и обновление данных
     function directoryTraversal(path){
        currentDir = `${currentDir}/${path}`
        fetchAndUpdateData()
    }



    // fetchAndUpdateData - получение и обновление данных
    function fetchAndUpdateData() {
        fetch(`http://localhost:8080/fs?dst=${currentDir}&sort=${currentSort}`)
            .then(response => {
                if (!response.ok) {
                    console.error('Ошибка:', error);
                    alert('Произошла ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
                }
                return response.json();
            })
            .then(responseData => {

                if(responseData.status==200){
                    updateData(responseData.data);
                    document.getElementById("currentDirectory").textContent = currentDir
                }else{
                    alert(`Ошибка запроса: ${responseData.text_error}`)
                }
            })
    }


    //updateData - обновление данных на странице
    function updateData(files) {

        //обновление отображения текущего пути
        showCurrentDirectory.textContent = currentDir

        // поиск и очищение таблицы
        const tbody = document.getElementById('file-table').getElementsByTagName('tbody')[0];
        tbody.innerHTML = ''; 


        files.forEach(file => {

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

            //добалвение строки в таблицу
            tbody.appendChild(row);
        });


        
    }


