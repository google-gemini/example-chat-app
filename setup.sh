#!/bin/bash

set -e

exit_message() {
    echo "Valid backend option not provided."
    echo "Select 'BACKEND=nodejs' for Node.js"
    echo "Select 'BACKEND=python' for Python/Flask"
    echo "Select 'BACKEND=all' for both Python/Flask and Node.js"
    echo "Example:"
    echo "BACKEND=nodejs ./setup.sh"
}

if [ -z ${BACKEND+x} ]; then
    exit_message
    exit 1
fi

nvm_node_install() { 
    printf "\nChecking nvm installation\n";
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

    printf "\nChecking node installation\n"
    if command -v node &> /dev/null; then
        printf "\nnode installation found\n"
    else
        printf "\nnode installation not found. Installing node\n"
        nvm install 22.6.0
    fi
}

react_setup() {
    printf "\nSetting up frontend(ReactJS + Vite)\n"
    cd client-react/
    npm install 
}

case $BACKEND in
    "nodejs")
        nvm_node_install
        react_setup
        echo "----------------------------"
        printf "\nSetting up NodeJs\n"
        cd ../server-js/
        npm install
        cd ..
        ;;
    "python")
        nvm_node_install
        react_setup
        echo "----------------------------"
        printf "\nSetting up Python/Flask\n"
        cd ../server-python
        if [ -d ".venv" ]; then
            printf "\nVirtual environment venv already exists.\n"
        else
            python3 -m venv .venv  
        fi
        source .venv/bin/activate
        pip install -r requirements.txt
        cd ..
        ;;
    "all")
        nvm_node_install
        react_setup
        echo "----------------------------"
        printf "\nSetting up NodeJs\n"
        cd ../server-js/
        npm install
        echo "----------------------------"
        printf "\nSetting up Python/Flask\n"
        cd ../server-python
        if [ -d ".venv" ]; then
            printf "\nVirtual environment venv already exists.\n"
        else
            python3 -m venv .venv  
        fi
        source .venv/bin/activate
        pip install -r requirements.txt
        cd ..
        ;;
    *)
        exit_message
        exit 1
        ;;
esac
