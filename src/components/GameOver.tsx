import type {Player} from "../types/shared.types.tsx";

export default function GameOver({winner, onRestart}: {winner: Player | null, onRestart: () => void}) {
    return <div id="game-over">
        <h2>Game Over!</h2>
        {winner ? (
            <>
                <p>Congratulations!</p>
                <p>The winner is {winner.name}!</p>
            </>
        ) : (
            <p>It's a draw!</p>
        )}
        <p>
            <button onClick={onRestart}>Rematch!</button>
        </p>
    </div>
}
