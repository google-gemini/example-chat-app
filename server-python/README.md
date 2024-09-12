
# Flask Gemini Chat API

This application provides a web API using Flask to interact with the Google Generative AI Gemini API, designed to get you started building with Gemini's capabilities. 

## Getting Started

### Prerequisites

Ensure you have Python 3.11 or greater installed on your machine. You can download it from [python.org](https://www.python.org/downloads/).

### Installation

Create a new virtual environment:

 - macOS:
   ```bash
   $ python -m venv venv
   $ . venv/bin/activate
   ```

 - Windows:
   ```cmd
   > python -m venv venv
   > .\venv\Scripts\activate
   ```

 - Linux:
    ```bash
    $ python -m venv venv
    $ source venv/bin/activate
    ```

Install the required Python packages:
```bash
pip install -r requirements.txt
```

### Configuration

Make a copy of the example environment variables file by copying the `.env.example` file to `.env`:

```bash
cp .env.example .env
```

Add your [API key](https://ai.google.dev/gemini-api/docs/api-key) to the newly created `.env` file or as an environment variable.

### Running the Application

Run the application with the following command:

```bash
python app.py
```

The server will start on `localhost:9000`.
