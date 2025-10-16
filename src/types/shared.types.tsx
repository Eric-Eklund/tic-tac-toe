export type Player = {
    id: number;
    name: string;
    symbol: string;
};

export type Players = [Player, Player]; // Exactly two players

export type GameTurn = {
    cell: { row: number; col: number };
    player: Player;
};

export type GameBoard = (string | null)[][];
