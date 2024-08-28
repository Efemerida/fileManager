import {changeSort, setShowCurrentDirectory, returnToPreviousDirectory, getAndUpdateData,
    setContainerContent,setPlaceholder
} from "./DOMService.js"






// загрузка DOM дерева
document.addEventListener('DOMContentLoaded', function() {
        initScript()
});


// initScript - опрделение обработки нажатий на кнопки и отображение текущей директории
export function initScript(){

    //кнопка сортировки
    let sortButton = document.getElementById("sort-button")
    sortButton.addEventListener('click', changeSort)

    //кнопка назад
    let backButton = document.getElementById("back-button")
    backButton.addEventListener('click', returnToPreviousDirectory)

    //получение и сохранения элемента отображения текущей директории
    setShowCurrentDirectory(document.getElementById("current-directory"))

    //получение контейнера, где отображаются данные, плейсхолдера, кнопки сортировки плейсхолдера
    setContainerContent(document.getElementById('container-data'))
    setPlaceholder(document.getElementById('placeholder'))

    //получение данных при загрузке страницы
    getAndUpdateData();
}



