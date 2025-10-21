import type {GameTurn, Player} from "../types/shared.types.tsx";

type LogProps = {
    turns: GameTurn[];
    players: Player[];
};

export default function Log({turns}: LogProps) {
    return <ol id="log">
        {turns.map((turn, index) => (
            <li key={`${turn.cell.row}-${turn.cell.col}-${index}`}>
                {turn.player.name} selected row {turn.cell.row + 1}, column {turn.cell.col + 1}
            </li>
        ))}
    </ol>
}
