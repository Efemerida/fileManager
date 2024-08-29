import {
    setShowCurrentDirectory,
    getAndUpdateData,
    setContainerContent,
    setPlaceholder
} from "./tableScript";

import {
    changeSort,
    returnToPreviousDirectory,
    asc
} from "./buttonsScript";

import './../styles/style.css';


// загрузка DOM дерева
document.addEventListener('DOMContentLoaded', function() {
    initScript();
});

// initScript - определение обработки нажатий на кнопки и отображение текущей директории
function initScript(): void {
    // кнопка сортировки
    const sortButton = document.getElementById("sort-button") as HTMLButtonElement;
    if (sortButton) {
        sortButton.addEventListener('click', changeSort);
    }

    // кнопка назад
    const backButton = document.getElementById("back-button") as HTMLButtonElement;
    if (backButton) {
        backButton.addEventListener('click', returnToPreviousDirectory);
    }

    // получение и сохранение элемента отображения текущей директории
    const currentDirectoryElement = document.getElementById("current-directory") as HTMLElement;
    setShowCurrentDirectory(currentDirectoryElement);

    // получение контейнера, где отображаются данные, плейсхолдера, кнопки сортировки плейсхолдера
    const containerDataElement = document.getElementById('container-data') as HTMLElement;
    setContainerContent(containerDataElement);
    
    const placeholderElement = document.getElementById('placeholder') as HTMLElement;
    setPlaceholder(placeholderElement);

    // получение данных при загрузке страницы
    getAndUpdateData("/home", asc);
}
