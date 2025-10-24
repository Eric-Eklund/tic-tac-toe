import axios from "axios"
import type {NewGameResponse} from "../types/shared.types.tsx"


const api = axios.create({
    baseURL: "http://localhost:8080/api",
    timeout: 5000,
    headers: {"Content-Type": "application/json"},
});

export async function startNewMatch(): Promise<NewGameResponse> {
    const response = await api.get<NewGameResponse>("/new-match");
    return response.data;
}
