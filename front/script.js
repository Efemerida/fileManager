currentDir = "/home/artem"
currentSort = "asc"

document.addEventListener('DOMContentLoaded', function() {
    fetchData(); 
});

document.getElementById('fetch-data-btn').addEventListener('click', fetchData);

function fetchData() {
    fetch('http://localhost:8080/fs?dst='+currentDir+'&sort='+currentSort)
        .then(response => {
            if (!response.ok) {
                throw new Error('Сетевая ошибка: ' + response.status);
            }
            return response.json();
        })
        .then(data => {
            console.log(data.data);
            populateTable(data.data);
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
        });
}

function populateTable(files) {
    const tbody = document.getElementById('file-table').getElementsByTagName('tbody')[0];
    tbody.innerHTML = ''; // Очистить предыдущее содержимое таблицы

    files.forEach(file => {
        const row = document.createElement('tr');
        const nameCell = document.createElement('td');
        const typeCell = document.createElement('td');
        const sizeCell = document.createElement('td');

        nameCell.textContent = file.file_name;
        typeCell.textContent = file.file_size;
        sizeCell.textContent = file.file_type;

        row.appendChild(nameCell);
        row.appendChild(typeCell);
        row.appendChild(sizeCell);
        tbody.appendChild(row);
    });
}