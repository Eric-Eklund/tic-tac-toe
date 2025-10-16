import {useState} from "react";
import * as React from "react";
import type {Player} from "../types/shared.types.tsx";

type PlayerProps = {
    player: Player;
    isActive: boolean;
    onNameChange: (playerId: number, newName: string) => void;
};

export default function PlayerComp({player, isActive, onNameChange}: PlayerProps) {
    const [playerName, setPlayerName] = useState(player.name);
    const [isEditing, setIsEditing] = useState(false);

    function handleEditClick() {
        if (isEditing) {
            // Save the name
            onNameChange(player.id, playerName);
        }
        setIsEditing(isEditing => !isEditing);
    }

    function handlePlayerNameChange(event: React.ChangeEvent<HTMLInputElement>) {
        setPlayerName(event.target.value);
    }

    let editablePlayerName = <span className="player-name">{playerName}</span>

    if (isEditing) {
        editablePlayerName = <input type="text" required value={playerName} onChange={handlePlayerNameChange}/>
    }

    return (
        <li className={isActive ? "active" : undefined}>
            <span className="player">
                { editablePlayerName }
                <span className="player-symbol">{player.symbol}</span>
            </span>
            <button onClick={handleEditClick}>{isEditing ? "Save" : "Edit"}</button>
        </li>
    )}