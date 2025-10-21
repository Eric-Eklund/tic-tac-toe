import axios from "axios"
import type {GameBoard, Player} from "../types/shared.types.tsx"


const api = axios.create({
    baseURL: "http://localhost:8080/api",
    timeout: 5000,
    headers: {"Content-Type": "application/json"},
});

export async function getInitialPlayers(): Promise<Player[]> {
    const response = await api.get<Player[]>("/players");
    return response.data;
}

export async function getInitialGameBoard(): Promise<GameBoard> {
    const response = await api.get<GameBoard>("/gameboard");
    return response.data
}