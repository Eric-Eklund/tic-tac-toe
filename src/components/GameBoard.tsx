import type {GameBoard} from "../types/shared.types.tsx";

type GameBoardProps = {
    onSelectSquare: (rowIndex: number, cellIndex: number) => void;
    board: GameBoard;
    disabled?: boolean;
};

export default function GameBoardComp({ onSelectSquare, board, disabled = false }: GameBoardProps) {
    return <ol id="game-board">
        {board.map((row, rowIndex) => (
            <li key={rowIndex}>
                <ol className="board-row">
                    {row.map((playerSymbol, cellIndex) => (
                        <li key={cellIndex} className="board-cell">
                            <button 
                                onClick={() => onSelectSquare(rowIndex, cellIndex)}
                                disabled={disabled || playerSymbol !== null}
                            >
                                {playerSymbol || ""}
                            </button>
                        </li>
                    ))}
                </ol>
            </li>
        ))}
    </ol>
}