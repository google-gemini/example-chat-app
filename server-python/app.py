# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from flask import (
    Flask,
    request,
    Response,
    stream_with_context
)
from flask_cors import CORS
import google.generativeai as genai
from dotenv import load_dotenv
import os

# Load environment variables from a .env file located in the same directory.
load_dotenv()

# Initialize a Flask application. Flask is used to create and manage the web server.
app = Flask(__name__)

# Apply CORS to the Flask app which allows it to accept requests from all domains.
# This is especially useful during development and testing.
CORS(app)

# WARNING: Do not share code with you API key hard coded in it.
# Configure the Google Generative AI's Google API key obtained
# from the environment variable. This key authenticates requests to the Gemini API.
genai.configure(api_key=os.getenv("GOOGLE_API_KEY"))

# Initialize the generative model with the specified model name.
# This model will be used to process user inputs and generate responses.
model = genai.GenerativeModel(
    model_name="gemini-1.5-flash"
)

@app.route('/chat', methods=['POST'])
def chat():
    """Processes user input and returns AI-generated responses.

    This function handles POST requests to the '/chat' endpoint. It expects a JSON payload
    containing a user message and an optional conversation history. It returns the AI's
    response as a JSON object.

    Args:
        None (uses Flask `request` object to access POST data)

    Returns:
        A JSON object with a key "text" that contains the AI-generated response.
    """
    # Parse the incoming JSON data into variables.
    data = request.json
    msg = data.get('chat', '')
    chat_history = data.get('history', [])

    # Start a chat session with the model using the provided history.
    chat_session = model.start_chat(history=chat_history)

    # Send the latest user input to the model and get the response.
    response = chat_session.send_message(msg)

    return {"text": response.text}

@app.route("/stream", methods=["POST"])
def stream():
    """Streams AI responses for real-time chat interactions.

    This function initiates a streaming session with the Gemini AI model,
    continuously sending user inputs and streaming back the responses. It handles
    POST requests to the '/stream' endpoint with a JSON payload similar to the
    '/chat' endpoint.

    Args:
        None (uses Flask `request` object to access POST data)

    Returns:
        A Flask `Response` object that streams the AI-generated responses.
    """
    def generate():
        data = request.json
        msg = data.get('chat', '')
        chat_history = data.get('history', [])

        chat_session = model.start_chat(history=chat_history)
        response = chat_session.send_message(msg, stream=True)

        for chunk in response:
            yield f"{chunk.text}"

    return Response(stream_with_context(generate()), mimetype="text/event-stream")

# Configure the server to run on port 9000.
if __name__ == '__main__':
    app.run(port=os.getenv("PORT"))
