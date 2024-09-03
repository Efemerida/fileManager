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
$stmt = $conn->prepare($sql);
$stmt->execute();
$result = $stmt->get_result();

// Сохраняем данные в массив
$rows = $result->fetch_all(MYSQLI_ASSOC);


// Извлекаем столбцы size и elapsed_time в отдельные массивы
$sizes = array();
$elapsedTimes = array();
foreach ($rows as $row) {
    $sizes[] = $row['size'];
    $elapsedTimes[] = $row['elapsed_time'];
}


// Генерация HTML
    echo "<!DOCTYPE html>
    <html lang='ru'>
    <head>
        <meta charset='UTF-8'>
        <meta name='viewport' content='width=device-width, initial-scale=1.0'>
        <script src='https://cdn.jsdelivr.net/npm/chart.js@3.9.1/dist/chart.min.js'></script>
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
    foreach ($rows as $row) {
        echo "<tr>
                <td>" . $row["id"] . "</td>
                <td>" . $row["root"] . "</td>
                <td>" . $row["size"] . "</td>
                <td>" . $row["elapsed_time"] . "</td>
                <td>" . $row["date_created"] . "</td>
              </tr>";
    }

    echo "</table>

<h2>График зависимости размера от времени запроса</h2>

<canvas id='myChart' width='400' height='200'></canvas>
    <script>
    const ctx = document.getElementById('myChart').getContext('2d');
    const dataPoints = <?php echo $dataPoints; ?>;

    const myChart = new Chart(ctx, {
        type: 'line',
        data: {
                    labels: labels,
                    datasets: [{
                        label: 'Время выполнения (мс)',
                        data: $elapsedTimes,
                        backgroundColor: 'rgba(255, 99, 132, 0.2)',
                        borderColor: 'rgba(255, 99, 132, 1)',
                        borderWidth: 1
                    }, {
                        label: 'Размер данных (КБ)',
                        data: $sizes,
                        backgroundColor: 'rgba(54, 162, 235, 0.2)',
                        borderColor: 'rgba(54, 162, 235, 1)',
                        borderWidth: 1
                    }]
                },
        options: {
            scales: {
                y: {
                    beginAtZero: true,
                    title: {
                        display: true,
                        text: 'Размер (байты) / Время (секунды)'
                    }
                }
            }
        }
    });
</script>

    </body>
    </html>";

// Закрытие подключения
$conn->close();
?>
