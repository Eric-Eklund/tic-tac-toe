import './index.css'
import Player from "./components/Player.tsx";

function App() {
    return  <main>
                <div id="game-container">
                    <ol id="players">
                        <Player name="Player 1" symbol="X"/>
                        <Player name="Player 2" symbol="O"/>
                    </ol>
                    Game Board
                </div>
                LOG
        </main>
}

export default App
