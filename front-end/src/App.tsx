import './index.css'
import PlayerComp from "./components/Player.tsx";
import GameBoardComp from "./components/GameBoard.tsx";
import {useState, useEffect} from "react";
import Log from "./components/Log.tsx";
import type {GameTurn, Player, Players, GameBoard} from "./types/shared.types.tsx";
import {WINNING_COMBINATIONS} from "./assets/winning-combinations.ts";
import GameOver from "./components/GameOver.tsx";
import {getInitialGameBoard, getInitialPlayers} from "./services/backend.tsx";

let INITIAL_GAME_BOARD: GameBoard = {
    board: [
        ["", "", ""],
        ["", "", ""],
        ["", "", ""]
    ]
}

function deriveActivePlayer(turns: GameTurn[], players: Player[]): Player {
    if (turns.length % 2 === 0) {
        return players[0]; // Player 1 (X)
    } else {
        return players[1]; // Player 2 (O)
    }
}

function deriveWinner(gameBoard: GameBoard, players: Player[]) {
    let winner: Player | null = null;
    for (const winningCombination of WINNING_COMBINATIONS) {
        const firstCell = gameBoard.board[winningCombination[0].row][winningCombination[0].col];
        const secondCell = gameBoard.board[winningCombination[1].row][winningCombination[1].col];
        const thirdCell = gameBoard.board[winningCombination[2].row][winningCombination[2].col];

        if (firstCell &&
            firstCell === secondCell &&
            secondCell === thirdCell) {

            winner = players.find(p => p.symbol === firstCell) || null;
            break;
        }
    }

    return winner;
}

function deriveGameBoard(gameTurns: GameTurn[]) {
    const gameBoard: GameBoard = {
        board: INITIAL_GAME_BOARD.board.map(row => [...row])
    };
    for (const turn of gameTurns) {
        const {cell, player} = turn;
        const {row, col} = cell;
        gameBoard.board[row][col] = player.symbol;
    }
    return gameBoard;
}

function App() {
    const [players, setPlayers] = useState<Player[]>([]);
    const [gameTurns, setGameTurns] = useState<GameTurn[]>([]);

    useEffect(() => {
        (async () => {
            try {
                const data = await getInitialPlayers();
                setPlayers(data);
                const gameBoard = await getInitialGameBoard();
                INITIAL_GAME_BOARD = gameBoard;
                console.log(gameBoard);
            } catch (err) {
                console.error(err);
            }
        })();
    }, []);

    const activePlayer = deriveActivePlayer(gameTurns, players);
    const gameBoard = deriveGameBoard(gameTurns);
    const winner = deriveWinner(gameBoard, players);
    const hasDraw = gameTurns.length === 9 && !winner;

    function handleSelectSquare(rowIndex: number, cellIndex: number) {
        setGameTurns((prevTurns) => {
            const currentPlayer = deriveActivePlayer(prevTurns, players);
            
            return [
                {cell: {row: rowIndex, col: cellIndex}, player: currentPlayer},
                ...prevTurns
            ];
        });
    }

    function handlePlayerNameChange(playerId: number, newName: string) {
        setPlayers(prevPlayers => 
            prevPlayers.map(player => 
                player.id === playerId 
                    ? { ...player, name: newName }
                    : player
            ) as Players
        );
    }

    function handleRematch() {
        setGameTurns([])
    }

    return <main>
        <div id="game-container">
            <ol id="players" className="highlight-player">
                {players.map(player => (
                    <PlayerComp 
                        key={player.id}
                        player={player}
                        isActive={activePlayer.id === player.id}
                        onNameChange={handlePlayerNameChange}
                    />
                ))}
            </ol>
            {(winner || hasDraw) && (
                <GameOver 
                    winner={winner} 
                    onRestart={handleRematch}
                />
            )}
            <GameBoardComp
                onSelectSquare={handleSelectSquare}
                board={gameBoard}
                disabled={!!winner || hasDraw}
            />
        </div>
        <Log turns={gameTurns} players={players}/>
    </main>
}

export default App
