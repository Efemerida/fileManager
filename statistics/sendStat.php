<?php

// Берем данные аунтификации из соответсвющего файла
include 'auth.php';


try{

    // Проверяем, был ли отправлен POST-запрос
    if ($_SERVER['REQUEST_METHOD'] === 'POST') {


        //Получаем данные из POST-запроса
        header('Content-Type: application/json'); 
        $jsonData = file_get_contents('php://input');
        $data = json_decode($jsonData, true); 

        $root = $data['root'];
        $size = $data['size'];
        $elapsedTime = $data['elapsed_time'];

        // Подключаемся к базе данных
        $conn = new mysqli($servername, $username, $password, $dbname);

        // Проверяем подключение
        if ($conn->connect_error) {
            throw new Exception("Ошибка подключения:  . $conn->connect_error .");
        }

        // Подготавливаем запрос INSERT
        $sql = "INSERT INTO stats (root, size, elapsed_time, date_created) VALUES (?, ?, ?, ?)";
        $stmt = $conn->prepare($sql);

        // Связываем параметры с запросом
        $currentDateTime = date('Y-m-d H:i:s');
        $stmt->bind_param("sids", $root, $size, $elapsedTime, $currentDateTime);

        // Выполняем запрос
        if ($stmt->execute()) {
            echo json_encode(['status' => 200, 'message' => 'Данные успешно добавлены в базу данных!']);
        } else {
            throw new Exception("Ошибка при добавлении данных:  . $stmt->error .");
        }

    }else {
        throw new Exception('Неверный метод запроса');
    }
}catch(Exception $exeption){
    echo json_encode(['status' => 500, 'message' => $exeption -> getMessage()]);
} finally {

    // Закрываем подготовленный запрос
    if (!empty($stmt)) $stmt->close();

    // Закрываем подключение
    if (!empty($conn)) $conn->close();
}
    

