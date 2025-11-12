# Main.gd - Attach to root Node2D
extends Node2D

# Backend connection
var backend_url = "http://localhost:8080"
var ws_url = "ws://localhost:8080/ws"
var http_request: HTTPRequest
var websocket: WebSocketPeer

# Game state
var match_id = ""
var current_board = []
var players = []
var current_turn = 0
var game_over = false
var winner = null

# UI nodes
var grid_container: GridContainer
var status_label: Label
var rematch_button: Button
var player1_label: Label
var player2_label: Label

# Cell buttons (3x3 grid)
var cell_buttons = []

func _ready():
	print("=== Game Starting ===")
	setup_ui()
	setup_http()
	setup_websocket()
	print("=== Starting new match ===")
	start_new_match()

func setup_ui():
	# Create status label at top
	status_label = Label.new()
	status_label.position = Vector2(50, 20)
	status_label.add_theme_font_size_override("font_size", 24)
	add_child(status_label)
	
	# Create player labels
	player1_label = Label.new()
	player1_label.position = Vector2(50, 60)
	player1_label.add_theme_font_size_override("font_size", 18)
	player1_label.add_theme_color_override("font_color", Color(1, 0.6, 0.2))  # Orange
	add_child(player1_label)
	
	player2_label = Label.new()
	player2_label.position = Vector2(250, 60)
	player2_label.add_theme_font_size_override("font_size", 18)
	player2_label.add_theme_color_override("font_color", Color(0.2, 0.6, 1))  # Blue
	add_child(player2_label)
	
	# Create 3x3 grid
	grid_container = GridContainer.new()
	grid_container.columns = 3
	grid_container.position = Vector2(100, 120)
	grid_container.add_theme_constant_override("h_separation", 10)
	grid_container.add_theme_constant_override("v_separation", 10)
	add_child(grid_container)
	
	# Create 9 cell buttons
	for i in range(9):
		var button = Button.new()
		button.custom_minimum_size = Vector2(100, 100)
		button.add_theme_font_size_override("font_size", 48)
		button.pressed.connect(_on_cell_pressed.bind(i))
		grid_container.add_child(button)
		cell_buttons.append(button)
	
	# Rematch button
	rematch_button = Button.new()
	rematch_button.text = "Rematch!"
	rematch_button.position = Vector2(150, 480)
	rematch_button.custom_minimum_size = Vector2(150, 50)
	rematch_button.pressed.connect(_on_rematch_pressed)
	rematch_button.visible = false
	add_child(rematch_button)
	
	update_status("Connecting to server...")

func setup_http():
	print("Setting up HTTP request node")
	http_request = HTTPRequest.new()
	add_child(http_request)
	http_request.request_completed.connect(_on_http_request_completed)
	print("HTTP request node ready")

func setup_websocket():
	websocket = WebSocketPeer.new()

func start_new_match():
	var error = http_request.request(backend_url + "/api/new-match")
	if error != OK:
		update_status("Failed to connect to server")

func _on_http_request_completed(result, response_code, headers, body):
	print("HTTP Response - Result: ", result, " Code: ", response_code)
	
	if response_code == 200:
		var json = JSON.new()
		var body_string = body.get_string_from_utf8()
		print("Response body: ", body_string)
		
		var parse_result = json.parse(body_string)
		if parse_result == OK:
			var data = json.data
			match_id = data.get("id", "")
			players = data.get("players", [])
			current_turn = 0
			
			print("Match ID: ", match_id)
			print("Players: ", players)
			
			# Parse the 2D board into a flat array
			var game_board = data.get("game_board", {})
			var board_2d = game_board.get("board", [])
			
			current_board = []
			for row in board_2d:
				for cell in row:
					current_board.append(cell)
			
			# Ensure board has 9 elements
			if current_board.size() != 9:
				current_board = ["", "", "", "", "", "", "", "", ""]
			
			print("Board loaded with ", current_board.size(), " cells")
			update_ui()
			connect_websocket()
		else:
			update_status("Failed to parse server response")
			print("JSON parse error: ", parse_result)
	else:
		update_status("Server error: " + str(response_code))
		print("Server returned error code: ", response_code)

func connect_websocket():
	# Close existing connection if any
	if websocket.get_ready_state() != WebSocketPeer.STATE_CLOSED:
		print("Closing existing WebSocket connection")
		websocket.close()
	
	var ws_url_to_use = ws_url
	
	# If running on web, construct WebSocket URL from current location
	if OS.has_feature("web"):
		# Use wss:// if page is https://, otherwise ws://
		var protocol = "ws://"
		ws_url_to_use = protocol + JavaScriptBridge.eval("window.location.host", true) + "/ws"
		print("Web mode - WebSocket URL: ", ws_url_to_use)
	
	var error = websocket.connect_to_url(ws_url_to_use)
	if error != OK:
		update_status("WebSocket connection failed")
		print("WebSocket error: ", error)
		return
	update_status("Connected! Game started")

func _process(_delta):
	if websocket.get_ready_state() == WebSocketPeer.STATE_OPEN:
		websocket.poll()
		
		while websocket.get_available_packet_count():
			var packet = websocket.get_packet()
			var message = packet.get_string_from_utf8()
			handle_websocket_message(message)

func handle_websocket_message(message: String):
	var json = JSON.new()
	var parse_result = json.parse(message)
	if parse_result == OK:
		var data = json.data
		var msg_type = data.get("type", "")
		
		if msg_type == "welcome":
			print("WebSocket connected: ", data.get("message", ""))
		# Add more message handlers as your backend adds them

func _on_cell_pressed(index: int):
	if game_over:
		return
	
	# Make sure board is loaded
	if current_board.size() != 9:
		update_status("Waiting for game to load...")
		return
		
	# Check if cell is empty
	if current_board[index] != "":
		return
	
	# Make move
	var player_symbol = players[current_turn].get("symbol", "X")
	current_board[index] = player_symbol
	
	# Update locally
	cell_buttons[index].text = player_symbol
	cell_buttons[index].disabled = true
	
	# Switch turns
	current_turn = 1 - current_turn
	
	# Check win condition
	check_game_over()
	
	# TODO: Send move to backend via WebSocket when you implement it
	# var move_data = {"type": "move", "matchId": match_id, "position": index}
	# websocket.send_text(JSON.stringify(move_data))

func check_game_over():
	# Check rows
	for row in range(3):
		var start = row * 3
		if current_board[start] != "" and \
		   current_board[start] == current_board[start + 1] and \
		   current_board[start] == current_board[start + 2]:
			end_game(current_board[start])
			return
	
	# Check columns
	for col in range(3):
		if current_board[col] != "" and \
		   current_board[col] == current_board[col + 3] and \
		   current_board[col] == current_board[col + 6]:
			end_game(current_board[col])
			return
	
	# Check diagonals
	if current_board[0] != "" and \
	   current_board[0] == current_board[4] and \
	   current_board[0] == current_board[8]:
		end_game(current_board[0])
		return
		
	if current_board[2] != "" and \
	   current_board[2] == current_board[4] and \
	   current_board[2] == current_board[6]:
		end_game(current_board[2])
		return
	
	# Check draw
	var is_draw = true
	for cell in current_board:
		if cell == "":
			is_draw = false
			break
	
	if is_draw:
		end_game(null)

func end_game(winning_symbol):
	game_over = true
	
	if winning_symbol:
		var winner_name = ""
		for player in players:
			if player.get("symbol", "") == winning_symbol:
				winner_name = player.get("name", "Player")
				break
		update_status("🏆 " + winner_name + " wins!")
	else:
		update_status("It's a draw!")
	
	rematch_button.visible = true

func _on_rematch_pressed():
	# Reset game state
	game_over = false
	winner = null
	rematch_button.visible = false
	
	# Reset board
	for button in cell_buttons:
		button.text = ""
		button.disabled = false
	
	# Start new match
	start_new_match()

func update_ui():
	# Update board
	for i in range(9):
		if i < current_board.size():
			var symbol = current_board[i]
			cell_buttons[i].text = symbol
			cell_buttons[i].disabled = symbol != ""
	
	# Update player labels
	if players.size() >= 2:
		player1_label.text = players[0].get("name", "Player 1") + " (X)"
		player2_label.text = players[1].get("name", "Player 2") + " (O)"
	else:
		# Fallback if backend doesn't send players
		player1_label.text = "Player 1 (X)"
		player2_label.text = "Player 2 (O)"
	
	# Update turn status
	if not game_over:
		if players.size() > current_turn:
			var current_player = players[current_turn].get("name", "Player")
			update_status(current_player + "'s turn")
		else:
			update_status("Player " + str(current_turn + 1) + "'s turn")

func update_status(text: String):
	status_label.text = text
