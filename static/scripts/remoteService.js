import {setRootDirecory, rootDirecory} from "./buttonsScript.js"

// fetchData - получение данных
export  function fetchData(path, sort) {
    return fetch(`/fs?dst=${path}&sort=${sort}`)
        .then(response => {
            if (!response.ok) {
                alert('Произошла ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
                return []
            }
            return response.json();
        }).then(responseData => {
            if(responseData.status===200){
                if(rootDirecory===undefined){
                    setRootDirecory(responseData.root_dir)
                }
                return responseData.data;
            }else{
                alert(`Ошибка запроса: ${responseData.text_error}`)
                return responseData.data
            }
        })
}