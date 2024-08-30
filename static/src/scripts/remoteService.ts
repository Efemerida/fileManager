import { rootDirecory, setRootDirecory } from "./buttonsScript";
import {FileData} from "./tableScript"


// ResponseData - структура ответных данных
interface ResponseData {
    status: number;         //статус
    root_dir?: string;      // корневая директория
    data: any;              // данные
    text_error?: string;    // текст ошибки
}

// fetchData - получение данных
export function fetchData(path: string, sort: string): Promise<FileData[]>  {
    return fetch(`/fs?dst=${path}&sort=${sort}`)
        .then((response: Response) => {
            if (!response.ok) {
                alert('Произошла ошибка при загрузке данных. Пожалуйста, попробуйте позже.');
                return [];
            }
            return response.json()
        })
        .then((responseData : ResponseData) => {
            if (responseData.status === 200) {
                if (rootDirecory === undefined) {
                    setRootDirecory(responseData.root_dir!);
                }
                return responseData.data;
            } else {
                alert(`Ошибка запроса: ${responseData.text_error}`);
                return responseData.data; // Обратите внимание, что может вернуться неполный массив данных
            }
        });
}