

// fetchAndUpdateData - получение данных
export  function fetchData(path) {
    return fetch(path)
        .then(response => {
            if (!response.ok) {
                alert('Произошла ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
            }
            return response.json();
        })
}