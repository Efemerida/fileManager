<?php
// Параметры подключения к базе данных
include 'auth.php';

try{

    // Создание подключения
    $conn = new mysqli($servername, $username, $password, $dbname);

    // Проверка подключения
    if ($conn->connect_error) {
        throw new Exception("Ошибка подключения: " . $conn->connect_error);
    }

    // Запрос данных из таблицы stats
    $sql = "SELECT id, root, size, elapsed_time, date_created FROM stats";
    $stmt = $conn->prepare($sql);
    $stmt->execute();
    $result = $stmt->get_result();

    // Сохраняем данные в массив
    $rows = $result->fetch_all(MYSQLI_ASSOC);

    // Разбиение массива на 2 с размером и временем выполнения
    $sizes = array();
    $elapsedTimes = array();
    foreach ($rows as $row) {
        $sizes[] = $row['size'];
        $elapsedTimes[] = $row['elapsed_time'];
    }

    // Создание массива для передачи в график
    $dataPoints = [];
    for ($i = 0; $i < count($sizes); $i++) {
        $dataPoints[] = [
            'x' => $sizes[$i],
            'y' => $elapsedTimes[$i]
        ];
    }

    // Сортировка по размеру 
    usort($dataPoints, function($a, $b) {
        return $a['x'] - $b['x'];
    });

    
    $dataPoints = json_encode($dataPoints);

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

        echo "</table>";

        //Печать графика
        echo "<h2>График зависимости размера от времени запроса</h2>

        <canvas id='myChart' width='400' height='200'></canvas>
        <script>
        const ctx = document.getElementById('myChart').getContext('2d');
        const dataPoints = $dataPoints;

        const myChart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: ['Размер', 'Время запроса'],
                datasets: [{
                    label: 'Зависимость времени выполнения от размера директории',
                    data: dataPoints,
                    borderColor: 'rgba(75, 192, 192, 1)',
                    borderWidth: 1,
                    fill: false
                }]
            },
            options: {
                scales: {
                    x: {
                        type: 'logarithmic',
                        title: {
                            display: true,
                            text: 'Размер (байты)'
                        }
                    },
                    y: {
                        type: 'logarithmic',
                        title: {
                            display: true,
                            text: 'Время (секунды)'
                        }
                    }
                }
            }
        });
    </script>

        </body>
        </html>";
}
catch(Exception $exeption){
    echo json_encode(['status' => 500, 'message' => $exeption -> getMessage()]);
}
finally{
    // Закрытие подключения
    $conn->close();
}
?>
