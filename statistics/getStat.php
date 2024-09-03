<?php
// Параметры подключения к базе данных
include 'auth.php';

// Создание подключения
$conn = new mysqli($servername, $username, $password, $dbname);

// Проверка подключения
if ($conn->connect_error) {
    echo("Ошибка подключения: " . $conn->connect_error);
}

// Запрос данных из таблицы stats
$sql = "SELECT id, root, size, elapsed_time, date_created FROM stats";
$result = $conn->query($sql);

// Генерация HTML
    echo "<!DOCTYPE html>
    <html lang='ru'>
    <head>
        <meta charset='UTF-8'>
        <meta name='viewport' content='width=device-width, initial-scale=1.0'>
        <title>Статистика</title>
        <style>
            table {
                width: 100%;
                border-collapse: collapse;
            }
            table, th, td {
                border: 1px solid black;
                padding: 8px;
                text-align: left;
            }
            th {
                background-color: #f2f2f2;
            }
        </style>
    </head>
    <body>
        <h1>Статистика</h1>
        <table>
            <tr>
                <th>ID</th>
                <th>Директория</th>
                <th>Размер (байты)</th>
                <th>Время выполнения</th>
                <th>Дата создания</th>
            </tr>";

    // Вывод данных построчно
    while($row = $result->fetch_assoc()) {
        echo "<tr>
                <td>" . $row["id"] . "</td>
                <td>" . ($row["root"]) . "</td>
                <td>" . $row["size"] . "</td>
                <td>" . $row["elapsed_time"] . "</td>
                <td>" . $row["date_created"] . "</td>
              </tr>";
    }

    echo "</table>
    </body>
    </html>";

// Закрытие подключения
$conn->close();
?>
