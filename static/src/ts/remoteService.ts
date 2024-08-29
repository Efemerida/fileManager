import { rootDirecory, setRootDirecory } from "./buttonsScript";
import {FileData} from "./tableScript"


// Определение интерфейса для структуры ответных данных
interface ResponseData {
    status: number;
    root_dir?: string; // Може быть undefined
    data: any; // Замените any на конкретный тип, если знаете структуру данных
    text_error?: string; // Може быть undefined
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
