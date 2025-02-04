from flask import Flask, request, jsonify
from dotenv import load_dotenv
import os
import json
# from debug.debug import debug

# debug()

#########################################################
# Get the current script directory /prome-alert-gateway_dev/mockup_accesstoken_py
current_dir = os.path.dirname(os.path.abspath(__file__))

# Move up one levels and then go into '"/prome-alert-gateway_dev/.env"'
expected_env_path = os.path.abspath(os.path.join(current_dir, "..", ".env"))

print("Expected Path:", expected_env_path)
#########################################################


# Load environment variables from .env
load_dotenv(expected_env_path)

app = Flask(__name__)

# Load user store from .env
store_json = os.getenv("STORE")

# Convert string to dictionary (JSON)
if store_json:
    store = json.loads(store_json)
else:
    store = {}


# Route: POST /api/v2/authen - ตรวจสอบ username + password และคืนค่า access_token
@app.post("/api/v2/authen") #http://localhost:5000/api/v2/authen
def gen_token():
    data = request.json
    username=data.get('username')                 
    password=data.get('password')

    if not username or not password:
        return jsonify({"error": "Missing username or password"}), 400

    user = store.get(username)
    # print(user)

    if not user or user["password"] != password:
        return jsonify({"error": "Invalid username or password"}), 401

    # return jsonify({"username": username, "access_token": user["access_token"]})
    return user["access_token"]

# Start Flask server
if __name__ == "__main__":
    app.run(debug=True, port=5000)


