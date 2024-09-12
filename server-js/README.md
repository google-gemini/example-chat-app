# Gemini chat app with Node.js

## Installation

### Installation of npm and node
To run this app you need `node` and `npm` installed in your system.
Check if the packages are already installed.
```
node -v
npm -v
```
If these packages are not installed, follow the instructions for your particular os from the [official npm website](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm#using-a-node-version-manager-to-install-nodejs-and-npm).

### Install the package dependencies

You can quickly install the required packages using the `package.json` file. 
#### Install from `package.json`
1. Navigate to the app directory, `chat-js` (i.e. where package.json is located).
2. Run `npm install`. This will install all the required packages mentioned in `packages.json`.

## Configuration
1. Navigate to the app directory, `server-js`
2. Copy the `.env.example` file to `.env`:
```
cp .env.example .env
```
3. If you don't have a Gemini API key ready, you can create a key with one click in [Google AI Studio](https://aistudio.google.com/app/apikey).
4. Add the Gemini API key to the variable `GOOGLE_API_KEY` in the `.env` file.
   ```
   GOOGLE_API_KEY=<your_api_key>
   ```

## Run the app
To run the node.js chat app, use the following command.

```node --env-file=.env app.js```

`--env-file=.env` tells node.js where the .env file lies.

By default, the app will run on port 9000.

If you want to specify your port, edit the `PORT` key in your `.env` file.
`PORT=xxxx`
